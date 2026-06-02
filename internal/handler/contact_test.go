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
