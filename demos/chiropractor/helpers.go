package chiropractor

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

// telURL builds a sanitized tel: URL for the practice phone number.
func telURL(c Content) templ.SafeURL {
	return templ.SafeURL("tel:" + c.Business.PhoneTel)
}

// ratingLabel renders an accessible label for a star rating, e.g. "5 out of 5 stars".
func ratingLabel(n int) string {
	return itoa(n) + " out of 5 stars"
}

// fieldClass returns the wrapper class for an intake field, adding the error
// modifier when that field failed validation.
func fieldClass(errs intakeErrors, key string) string {
	if _, bad := errs[key]; bad {
		return "field field--error"
	}
	return "field"
}

// fieldHelp returns the field's error message when present, else its help copy.
func fieldHelp(errs intakeErrors, key, fallback string) string {
	if msg, bad := errs[key]; bad {
		return msg
	}
	return fallback
}
