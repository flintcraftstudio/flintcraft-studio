// Package intake holds the new-patient intake form logic shared by every demo
// vertical (chiropractor, law, dentist, …). Parsing, validation, spam handling
// and submission logging live here; each vertical injects its own templ markup
// for the form, error, and success states via Renderer.
//
// No real email is sent for the demos — submissions are appended to a local
// (gitignored) JSON-lines file. See the TODO in Handler where a real send would
// be wired.
package intake

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/a-h/templ"
)

// Submission is one parsed intake form post. Name/Email/Phone/Reason/Message
// are shared across every vertical; PatientType, When and Contact are optional
// extras some verticals collect (the law demo asks "when did it happen" and
// "best way to reach you"). Unused fields stay empty and validation ignores them.
type Submission struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Reason      string `json:"reason"`
	PatientType string `json:"patientType,omitempty"` // chiropractor: "new" | "existing"
	When        string `json:"when,omitempty"`        // law: when the incident happened
	Contact     string `json:"contact,omitempty"`     // law: preferred contact method
	Message     string `json:"message"`
}

// Errors maps a field name (or "form" for a form-level error) to a message.
type Errors map[string]string

// honeypotField is a hidden input bots tend to fill; humans never see it.
const honeypotField = "company"

var (
	emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	digitRe = regexp.MustCompile(`\D`)
)

// Parse reads the intake fields from a (already-parsed) request form.
func Parse(r *http.Request) Submission {
	return Submission{
		Name:        strings.TrimSpace(r.FormValue("name")),
		Email:       strings.TrimSpace(r.FormValue("email")),
		Phone:       strings.TrimSpace(r.FormValue("phone")),
		Reason:      strings.TrimSpace(r.FormValue("reason")),
		PatientType: strings.TrimSpace(r.FormValue("ptype")),
		When:        strings.TrimSpace(r.FormValue("when")),
		Contact:     strings.TrimSpace(r.FormValue("contact")),
		Message:     strings.TrimSpace(r.FormValue("message")),
	}
}

// FirstName returns the first whitespace-delimited token of the name, for the
// warm confirmation copy ("Thanks, Jordan —").
func (s Submission) FirstName() string {
	if i := strings.IndexAny(s.Name, " \t"); i > 0 {
		return s.Name[:i]
	}
	return s.Name
}

// Validate checks the submission server-side and returns field-level errors.
// Mirrors the client-side checks so the no-JS path behaves identically.
func (s Submission) Validate() Errors {
	errs := Errors{}
	if s.Name == "" {
		errs["name"] = "Please tell us your name."
	}
	if digits := digitRe.ReplaceAllString(s.Phone, ""); len(digits) < 10 {
		errs["phone"] = "A number we can reach you at, please."
	}
	if !emailRe.MatchString(s.Email) {
		errs["email"] = "Add an email for your confirmation."
	}
	if s.Reason == "" {
		errs["reason"] = "Pick the closest reason — we'll dig in together."
	}
	return errs
}

// isSpam reports whether the honeypot field was filled (bot signal).
func isSpam(r *http.Request) bool {
	return strings.TrimSpace(r.FormValue(honeypotField)) != ""
}

// Renderer supplies the per-vertical markup the shared handler renders.
//   - Form re-renders the intake form with entered values + errors (htmx swap).
//   - Success renders the confirmation partial (htmx swap).
//   - Page re-renders the whole page reflecting form state, for the no-JS
//     full-page fallback (progressive enhancement).
type Renderer struct {
	Form    func(s Submission, errs Errors) templ.Component
	Success func(firstName string) templ.Component
	Page    func(s Submission, errs Errors, success bool) templ.Component
}

// Handler returns an http.HandlerFunc for POST {prefix}/intake.
//
// It validates the submission and, on htmx requests (HX-Request header),
// swaps in the form (with errors) or the success partial. Without JS it
// re-renders the full page so the form still works as a plain POST.
// Successful submissions are appended to logPath as JSON lines.
func Handler(rnd Renderer, logPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		htmx := r.Header.Get("HX-Request") != ""
		sub := Parse(r)

		// Honeypot: pretend success, persist nothing.
		if isSpam(r) {
			slog.Warn("intake honeypot triggered", "remote", r.RemoteAddr)
			render(w, r, htmx, rnd.Success(sub.FirstName()), rnd.Page(sub, nil, true))
			return
		}

		if errs := sub.Validate(); len(errs) > 0 {
			render(w, r, htmx, rnd.Form(sub, errs), rnd.Page(sub, errs, false))
			return
		}

		if err := appendLog(logPath, sub); err != nil {
			slog.Error("intake log write failed", "err", err, "path", logPath)
		}
		// TODO: wire to Postmark (mail.Client) for real verticals; demos log only.
		// A basic rate limit (per-IP token bucket) belongs here too before launch.
		slog.Info("intake submission", "name", sub.Name, "reason", sub.Reason, "type", sub.PatientType)

		render(w, r, htmx, rnd.Success(sub.FirstName()), rnd.Page(sub, nil, true))
	}
}

// render writes the htmx partial when the request came from htmx, otherwise the
// full-page fallback.
func render(w http.ResponseWriter, r *http.Request, htmx bool, partial, full templ.Component) {
	c := full
	if htmx {
		c = partial
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.Render(r.Context(), w); err != nil {
		slog.Error("intake render error", "err", err)
	}
}

// appendLog appends one submission as a JSON line to path (created if absent).
func appendLog(path string, s Submission) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return writeJSONLine(f, s)
}

func writeJSONLine(w io.Writer, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	_, err = io.WriteString(w, "\n")
	return err
}
