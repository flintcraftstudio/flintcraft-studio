package view

// DefaultOGImage is the site-wide social preview: the arrowhead mark on
// obsidian, 1200x630.
const DefaultOGImage = "/static/images/og-default.jpg"

// PageMeta is everything the document <head> needs for one page.
type PageMeta struct {
	Title             string // short title; " | FlintCraft Studio" is appended unless FullTitle is set
	FullTitle         string // optional: the complete <title>, no suffix added
	Description       string // meta description (search)
	SocialDescription string // optional shorter OG/Twitter description; defaults to Description
	Image             string // optional OG/Twitter image path; defaults to DefaultOGImage
	ActivePath        string // nav highlight
	CanonicalPath     string // optional; defaults to ActivePath
}

func (m PageMeta) documentTitle() string {
	if m.FullTitle != "" {
		return m.FullTitle
	}
	return m.Title + " | " + SiteName
}

func (m PageMeta) canonical() string {
	if m.CanonicalPath != "" {
		return m.CanonicalPath
	}
	return m.ActivePath
}

func (m PageMeta) socialDescription() string {
	if m.SocialDescription != "" {
		return m.SocialDescription
	}
	return m.Description
}

func (m PageMeta) image() string {
	if m.Image != "" {
		return m.Image
	}
	return DefaultOGImage
}
