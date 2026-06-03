package law

import (
	"github.com/a-h/templ"

	"github.com/firefly-software-mt/standard-template/internal/intake"
)

// Local aliases keep the templates terse and decoupled from the shared package
// name while still using the shared types underneath.
type (
	submission   = intake.Submission
	intakeErrors = intake.Errors
)

// telURL builds a sanitized tel: URL for the firm phone number.
func telURL(c Content) templ.SafeURL {
	return templ.SafeURL("tel:" + c.Firm.PhoneTel)
}

// ratingLabel renders an accessible label for a star rating, e.g. "5 out of 5 stars".
func ratingLabel(n int) string {
	return itoa(n) + " out of 5 stars"
}

// iconInner returns the inner SVG markup for a Lucide icon name (see icons.go).
// Unknown names render nothing rather than breaking the page.
func iconInner(name string) string {
	return lucideIcons[name]
}

// fieldClass returns the wrapper class for an intake field, adding the error
// modifier when that field failed validation.
func fieldClass(errs intakeErrors, key string) string {
	if _, bad := errs[key]; bad {
		return "field err"
	}
	return "field"
}

// fieldErr returns the field's error message when present, else "".
func fieldErr(errs intakeErrors, key string) string {
	if msg, bad := errs[key]; bad {
		return msg
	}
	return ""
}

// firstName returns the first whitespace token of a name, for the warm
// confirmation copy ("Thanks, Jordan —"); falls back to "there".
func firstName(name string) string {
	for i := 0; i < len(name); i++ {
		if name[i] == ' ' || name[i] == '\t' {
			if i > 0 {
				return name[:i]
			}
		}
	}
	if name == "" {
		return "there"
	}
	return name
}
