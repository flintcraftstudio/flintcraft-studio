package view

import "time"

// SiteName is the display name used in templates.
const SiteName = "FlintCraft Studio"

// Tagline is the brand tagline displayed in the nav.
const Tagline = "Still shaped by hand."

// Contact info
const (
	Phone     = "(406) 871-9875"
	Email     = "hello@flintcraftstudio.com"
	ReviewURL = "https://g.page/r/Cc5efB_VXb5vEBM/review"
)

// Tracking IDs and Turnstile site key, set once at startup from config.
var (
	PixelID          string
	GtagID           string
	TurnstileSiteKey string
)

// Year returns the current year for copyright notices.
func Year() int {
	return time.Now().Year()
}
