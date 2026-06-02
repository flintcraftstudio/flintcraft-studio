package view

import "github.com/firefly-software-mt/standard-template/internal/structdata"

// IndustryLanding is the content for one keyword-targeted industry landing page
// (e.g. "Chiropractor Website Design in Montana"). These pages are the SEO entry
// points for verticals FlintCraft is pursuing; each links to a live demo as
// proof. Adding a vertical is a new entry here plus its demo — no new templates.
type IndustryLanding struct {
	Slug            string // flat keyword URL, no leading slash, e.g. "chiropractor-website-design"
	Profession      string // "chiropractic", used in prose
	Keyword         string // primary keyword phrase, e.g. "chiropractor website design"
	H1              string
	Lede            string
	MetaTitle       string // Base appends " | FlintCraft Studio"
	MetaDescription string
	Intro           []string
	KeyPoints       []IndustryPoint
	Cities          []string
	FAQ             []IndustryFAQ
	DemoURL         string // live demo, opened in a new tab
	DemoLabel       string // short label for the demo, e.g. "Alpine Spine"
	DemoCaption     string // one line under the preview
	PreviewImage    string // framed preview visual
	PreviewAlt      string
}

// IndustryPoint is one "what your site gets" item.
type IndustryPoint struct {
	Title string
	Body  string
}

// IndustryFAQ is one question/answer pair (rendered on-page and as FAQ JSON-LD).
type IndustryFAQ struct {
	Q string
	A string
}

// CanonicalURL is the absolute URL of this landing page.
func (d IndustryLanding) CanonicalURL() string { return BaseURL + "/" + d.Slug }

// serviceLD builds the schema.org Service node for this landing page.
func (d IndustryLanding) serviceLD() structdata.Service {
	return structdata.Service{
		Context:     "https://schema.org",
		Type:        "Service",
		ServiceType: d.Keyword,
		Name:        d.MetaTitle,
		Description: d.MetaDescription,
		URL:         d.CanonicalURL(),
		AreaServed:  "Montana",
		Provider: &structdata.Provider{
			Type:      "LocalBusiness",
			Name:      SiteName,
			URL:       BaseURL,
			Telephone: Phone,
		},
	}
}

// faqLD builds the schema.org FAQPage node from the page's Q&A list.
func (d IndustryLanding) faqLD() structdata.FAQPage {
	qs := make([]structdata.Question, 0, len(d.FAQ))
	for _, f := range d.FAQ {
		qs = append(qs, structdata.Question{
			Type:           "Question",
			Name:           f.Q,
			AcceptedAnswer: structdata.Answer{Type: "Answer", Text: f.A},
		})
	}
	return structdata.FAQPage{Context: "https://schema.org", Type: "FAQPage", MainEntity: qs}
}

// IndustryOrder is the display/registration order for industry landing pages.
var IndustryOrder = []string{
	"chiropractor-website-design",
}

// IndustryLandings is the canonical set of industry landing pages, keyed by slug.
var IndustryLandings = map[string]IndustryLanding{
	"chiropractor-website-design": {
		Slug:            "chiropractor-website-design",
		Profession:      "chiropractic",
		Keyword:         "chiropractor website design",
		H1:              "Chiropractor Website Design in Montana",
		Lede:            "Custom, fast, easy-to-find websites for Montana chiropractors — built to bring in new patients, not just sit there. No WordPress, no templates, no plugin upkeep.",
		MetaTitle:       "Montana Chiropractor Website Design",
		MetaDescription: "Custom chiropractor website design for Montana practices. Online new-patient intake, real review structured data, near-instant load. See a live demo, then start a project.",
		Intro: []string{
			"Most chiropractic websites are a tired WordPress theme with a booking plugin bolted on — slow, hard to update, and a little broken on mobile. Patients notice. So does Google.",
			"FlintCraft builds chiropractor websites from scratch: a site that loads in well under a second, turns first-time visitors into booked new patients, and stays sharp without you touching a plugin. Below is a real, working example of what that looks like.",
		},
		KeyPoints: []IndustryPoint{
			{"New-patient intake that actually converts", "A calm, dual-path layout for new and returning patients, with an intake form that validates instantly and reassures nervous first-timers — the moment most practice sites lose the booking."},
			{"Your reviews, in your own site's data", "Star ratings and reviews are written into the site's own structured data, so Google can show them in search — not rented from a third-party review plugin you don't control."},
			{"Found by the patients near you", "Built-in local SEO and Google Business Profile setup, with copy written for the Montana towns you actually serve."},
			{"Fast and accessible by default", "Near-perfect Lighthouse scores, sub-second loads, and full keyboard/screen-reader accessibility — built in, not patched on. Nothing to update on Tuesday night."},
		},
		Cities: []string{"Helena", "Bozeman", "Missoula", "Billings", "Great Falls", "Kalispell"},
		FAQ: []IndustryFAQ{
			{"How much does a chiropractor website cost?", "Every project is custom, so pricing depends on scope — but it's a simple, predictable monthly relationship that includes hosting, maintenance, SEO, and updates, not a surprise invoice every time something needs to change. The first conversation is free; we'll give you an honest read before you commit to anything."},
			{"Do you only work with chiropractors in Helena?", "No — we're based in Helena, Montana and work with chiropractic and wellness practices across the state, including Bozeman, Missoula, Billings, Great Falls, and Kalispell. The work is remote-friendly and the relationship is hands-on either way."},
			{"Can patients book or request appointments online?", "Yes. The demo includes a new-patient intake form with instant validation and a warm confirmation. It can connect to your scheduling or simply notify your front desk — your call."},
			{"What's wrong with a WordPress or Wix site?", "Nothing, until it's slow, breaks on an update, looks like a template, and the structured data that gets you into Google lives in a plugin you don't own. We build the whole thing from scratch so it's fast, durable, and yours."},
		},
		DemoURL:      "/demos/chiropractor/",
		DemoLabel:    "Alpine Spine Chiropractic & Wellness",
		DemoCaption:  "A complete sample site we built for a (fictional) Helena chiropractor. Click through the whole thing — the booking form works.",
		PreviewImage: "/demos/chiropractor/static/img/hero-1120.webp",
		PreviewAlt:   "Preview of the Alpine Spine chiropractor demo website",
	},
}
