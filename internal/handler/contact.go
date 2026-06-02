package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/firefly-software-mt/standard-template/internal/mail"
	"github.com/firefly-software-mt/standard-template/internal/view"
)

// refSanitize strips a referral tag down to a safe slug so an arbitrary ?ref=
// query value can't inject anything into the page or the outgoing email.
var refSanitize = regexp.MustCompile(`[^a-z0-9-]`)

// sanitizeRef normalizes a referral source to lowercase [a-z0-9-], capped at 40
// chars. Empty if nothing usable. Demo/landing CTAs pass values like
// "chiropractor-demo" or "chiropractor-landing".
func sanitizeRef(s string) string {
	s = refSanitize.ReplaceAllString(strings.ToLower(strings.TrimSpace(s)), "")
	if len(s) > 40 {
		s = s[:40]
	}
	return s
}

// Contact handles GET /contact and renders the contact form. A ?ref= query
// (set by industry demo/landing CTAs) is captured so we know the lead came in
// through a demo — it rides through the form in a hidden field.
func Contact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var values map[string]string
		if ref := sanitizeRef(r.URL.Query().Get("ref")); ref != "" {
			values = map[string]string{"ref": ref}
		}
		if err := view.ContactPage(nil, values, false).Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}

// ContactSubmit handles POST /contact, validates input, and sends a message.
func ContactSubmit(mailer *mail.Client, turnstileSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		values := map[string]string{
			"name":     strings.TrimSpace(r.FormValue("name")),
			"business": strings.TrimSpace(r.FormValue("business")),
			"email":    strings.TrimSpace(r.FormValue("email")),
			"phone":    strings.TrimSpace(r.FormValue("phone")),
			"website":  strings.TrimSpace(r.FormValue("website")),
			"message":  strings.TrimSpace(r.FormValue("message")),
			"ref":      sanitizeRef(r.FormValue("ref")),
		}

		errors := validate(values)

		if len(errors) > 0 {
			if err := view.ContactForm(errors, values, false).Render(r.Context(), w); err != nil {
				slog.Error("render error", "err", err)
			}
			return
		}

		// Verify Turnstile token
		if turnstileSecret != "" {
			token := r.FormValue("cf-turnstile-response")
			if !verifyTurnstile(turnstileSecret, token, r.RemoteAddr) {
				errors = map[string]string{"form": "Verification failed. Please try again."}
				if err := view.ContactForm(errors, values, false).Render(r.Context(), w); err != nil {
					slog.Error("render error", "err", err)
				}
				return
			}
		}

		// Log every submission (incl. referral source) so demo-driven leads are
		// recorded even when no mailer is configured.
		slog.Info("contact submission", "name", values["name"], "email", values["email"], "ref", values["ref"])

		if mailer != nil {
			body := values["message"]
			if values["business"] != "" {
				body = fmt.Sprintf("Business: %s\n\n%s", values["business"], body)
			}
			if values["phone"] != "" {
				body = fmt.Sprintf("%s\n\nPhone: %s", body, values["phone"])
			}
			if values["website"] != "" {
				body = fmt.Sprintf("%s\nWebsite: %s", body, values["website"])
			}
			subject := fmt.Sprintf("Contact form: %s", values["name"])
			if values["ref"] != "" {
				body = fmt.Sprintf("⟶ Referred via: %s\n\n%s", values["ref"], body)
				subject = fmt.Sprintf("Contact form: %s [via %s]", values["name"], values["ref"])
			}
			msg := mail.Message{
				Name:    values["name"],
				Email:   values["email"],
				Subject: subject,
				Body:    body,
			}
			if err := mailer.Send(msg); err != nil {
				slog.Error("postmark send error", "err", err)
				errors = map[string]string{"form": "Failed to send message. Please try again."}
				if err := view.ContactForm(errors, values, false).Render(r.Context(), w); err != nil {
					slog.Error("render error", "err", err)
				}
				return
			}
		}

		if err := view.ContactForm(nil, nil, true).Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}

// validate checks contact form values and returns a map of field errors.
func validate(values map[string]string) map[string]string {
	errors := make(map[string]string)

	if values["name"] == "" {
		errors["name"] = "Name is required."
	}
	if values["email"] == "" {
		errors["email"] = "Email is required."
	} else if !strings.Contains(values["email"], "@") {
		errors["email"] = "Enter a valid email address."
	}
	if values["message"] == "" {
		errors["message"] = "Message is required."
	}

	return errors
}

// verifyTurnstile checks a Turnstile token against the Cloudflare API.
func verifyTurnstile(secret, token, remoteIP string) bool {
	resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify", url.Values{
		"secret":   {secret},
		"response": {token},
		"remoteip": {remoteIP},
	})
	if err != nil {
		slog.Error("turnstile verify request failed", "err", err)
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		slog.Error("turnstile verify decode failed", "err", err)
		return false
	}

	if !result.Success {
		slog.Warn("turnstile verification failed")
	}
	return result.Success
}
