package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSanitizeRef(t *testing.T) {
	cases := map[string]string{
		"chiropractor-demo":         "chiropractor-demo",
		"  Chiropractor-Demo ":      "chiropractor-demo",
		"<script>alert(1)</script>": "scriptalert1script",
		"":                          "",
		"a/b?c=d&e=f":               "abcdef",
	}
	for in, want := range cases {
		if got := sanitizeRef(in); got != want {
			t.Errorf("sanitizeRef(%q) = %q, want %q", in, got, want)
		}
	}
	// Length cap.
	if got := sanitizeRef(strings.Repeat("x", 80)); len(got) != 40 {
		t.Errorf("sanitizeRef did not cap length: got %d", len(got))
	}
}

func TestContactGETCapturesRef(t *testing.T) {
	r := httptest.NewRequest("GET", "/contact?ref=chiropractor-demo", nil)
	w := httptest.NewRecorder()
	Contact().ServeHTTP(w, r)
	body := w.Body.String()

	if !strings.Contains(body, `name="ref" value="chiropractor-demo"`) {
		t.Error("contact page missing hidden ref field")
	}
	if !strings.Contains(body, "the Alpine Spine chiropractor demo") {
		t.Error("contact page missing referral banner label")
	}
}

func TestContactGETSanitizesMaliciousRef(t *testing.T) {
	r := httptest.NewRequest("GET", "/contact?ref="+url.QueryEscape("<script>x</script>"), nil)
	w := httptest.NewRecorder()
	Contact().ServeHTTP(w, r)
	body := w.Body.String()
	if strings.Contains(body, "<script>x</script>") {
		t.Error("unsanitized ref reflected into the page")
	}
}

func TestContactSubmitPreservesRefOnError(t *testing.T) {
	form := url.Values{
		"email":   {"a@b.co"},
		"message": {"hi"},
		"ref":     {"chiropractor-demo"},
		// name omitted -> validation error, form re-renders
	}
	r := httptest.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ContactSubmit(nil, "").ServeHTTP(w, r)

	if got := w.Result().StatusCode; got != http.StatusOK {
		t.Fatalf("status = %d", got)
	}
	if !strings.Contains(w.Body.String(), `name="ref" value="chiropractor-demo"`) {
		t.Error("ref not preserved through validation error")
	}
}

func TestSanitizeTimeline(t *testing.T) {
	cases := map[string]string{
		"Within a month":     "Within a month",
		" 1–3 months ":       "1–3 months",
		"Just exploring":     "Just exploring",
		"":                   "",
		"yesterday":          "",
		"<script>x</script>": "",
		"within a month":     "", // exact match only
	}
	for in, want := range cases {
		if got := sanitizeTimeline(in); got != want {
			t.Errorf("sanitizeTimeline(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestContactSubmitPreservesTimelineOnError(t *testing.T) {
	form := url.Values{
		"email":    {"a@b.co"},
		"message":  {"hi"},
		"timeline": {"1–3 months"},
		// name omitted -> validation error, form re-renders
	}
	r := httptest.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ContactSubmit(nil, "").ServeHTTP(w, r)

	if !strings.Contains(w.Body.String(), `value="1–3 months" selected`) {
		t.Error("timeline selection not preserved through validation error")
	}
}

func TestContactSubmitSuccessMessage(t *testing.T) {
	form := url.Values{
		"name":    {"Pat"},
		"email":   {"a@b.co"},
		"message": {"hi"},
	}
	r := httptest.NewRequest("POST", "/contact", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ContactSubmit(nil, "").ServeHTTP(w, r)

	body := w.Body.String()
	for _, s := range []string{"reply by end of the next business day", `href="tel:+14068719875"`} {
		if !strings.Contains(body, s) {
			t.Errorf("success response missing %q", s)
		}
	}
}
