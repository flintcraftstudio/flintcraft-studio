package view

import "time"

// SiteName is the display name used in templates.
const SiteName = "FlintCraft Studio"

// Tagline is the brand tagline displayed in the nav.
const Tagline = "Still shaped by hand."

// Contact info
const (
	Phone = "(406) 871-9875"
	Email = "hello@flintcraftstudio.com"
)

// Tracking IDs and Turnstile site key, set once at startup from config.
var (
	PixelID          string
	GtagID           string
	TurnstileSiteKey string
	// BaseURL is the site's public origin (e.g. "https://flintcraftstudio.com"),
	// set once at startup. Used to build absolute canonical / Open Graph URLs.
	BaseURL string
)

// Year returns the current year for copyright notices.
func Year() int {
	return time.Now().Year()
}

// ReferralLabel turns a sanitized ?ref= tag into a friendly phrase for the
// contact form ("You're coming from …"). Returns "" for no/unknown referral so
// the banner stays hidden — but the raw tag is still carried in the hidden
// field and the email either way. Add a case per vertical/surface as demos ship.
func ReferralLabel(ref string) string {
	switch ref {
	case "":
		return ""
	case "chiropractor-demo":
		return "the Alpine Spine chiropractor demo"
	case "chiropractor-landing":
		return "our chiropractor website page"
	case "industries-hub":
		return "our industries page"
	default:
		return "one of our industry examples"
	}
}
