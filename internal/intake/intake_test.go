package intake

import (
	"context"
	"io"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

// marker returns a templ component that writes a fixed tag, so tests can tell
// which render path the handler took.
func marker(tag string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		_, err := io.WriteString(w, tag)
		return err
	})
}

func testRenderer() Renderer {
	return Renderer{
		Form:    func(s Submission, errs Errors) templ.Component { return marker("[FORM]") },
		Success: func(first string) templ.Component { return marker("[SUCCESS:" + first + "]") },
		Page:    func(s Submission, errs Errors, ok bool) templ.Component { return marker("[PAGE]") },
	}
}

func post(form url.Values, htmx bool) *httptest.ResponseRecorder {
	r := httptest.NewRequest("POST", "/intake", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if htmx {
		r.Header.Set("HX-Request", "true")
	}
	w := httptest.NewRecorder()
	logPath := filepath.Join(os.TempDir(), "intake-test.log")
	Handler(testRenderer(), logPath).ServeHTTP(w, r)
	return w
}

func valid() url.Values {
	return url.Values{
		"name":   {"Jordan Hill"},
		"phone":  {"(406) 555-0142"},
		"email":  {"jordan@example.com"},
		"reason": {"Back or neck pain"},
		"ptype":  {"new"},
	}
}

func TestHandlerSuccessHTMX(t *testing.T) {
	w := post(valid(), true)
	if got := w.Body.String(); !strings.Contains(got, "[SUCCESS:Jordan]") {
		t.Errorf("expected success partial greeting by first name, got %q", got)
	}
}

func TestHandlerValidationHTMX(t *testing.T) {
	f := valid()
	f.Set("email", "not-an-email")
	f.Del("name")
	w := post(f, true)
	if got := w.Body.String(); got != "[FORM]" {
		t.Errorf("expected form partial on validation error, got %q", got)
	}
}

func TestHandlerNoJSFallback(t *testing.T) {
	// No HX-Request header → full-page render on both error and success.
	if got := post(valid(), false).Body.String(); got != "[PAGE]" {
		t.Errorf("expected full page on no-JS success, got %q", got)
	}
	f := valid()
	f.Del("name")
	if got := post(f, false).Body.String(); got != "[PAGE]" {
		t.Errorf("expected full page on no-JS error, got %q", got)
	}
}

func TestHandlerHoneypot(t *testing.T) {
	f := valid()
	f.Set(honeypotField, "spammybot")
	w := post(f, true)
	if got := w.Body.String(); !strings.HasPrefix(got, "[SUCCESS:") {
		t.Errorf("honeypot should fake success, got %q", got)
	}
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name string
		sub  Submission
		bad  []string
	}{
		{"all good", Submission{Name: "A B", Phone: "4065550142", Email: "a@b.co", Reason: "x"}, nil},
		{"missing all", Submission{}, []string{"name", "phone", "email", "reason"}},
		{"short phone", Submission{Name: "A", Phone: "12345", Email: "a@b.co", Reason: "x"}, []string{"phone"}},
		{"bad email", Submission{Name: "A", Phone: "4065550142", Email: "nope", Reason: "x"}, []string{"email"}},
	}
	for _, c := range cases {
		errs := c.sub.Validate()
		if len(errs) != len(c.bad) {
			t.Errorf("%s: got errors %v, want fields %v", c.name, errs, c.bad)
		}
		for _, f := range c.bad {
			if _, ok := errs[f]; !ok {
				t.Errorf("%s: expected error on %q", c.name, f)
			}
		}
	}
}

func TestFirstName(t *testing.T) {
	if got := (Submission{Name: "Jordan Hill"}).FirstName(); got != "Jordan" {
		t.Errorf("FirstName = %q, want Jordan", got)
	}
	if got := (Submission{Name: "Cher"}).FirstName(); got != "Cher" {
		t.Errorf("FirstName = %q, want Cher", got)
	}
}
