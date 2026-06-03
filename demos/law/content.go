// Package law is the second FlintCraft demo vertical: a single-page marketing
// site for the fictional Granite Peak Trial Lawyers, a personal-injury trial
// firm in Helena, Montana. It follows the chiropractor demo's shape — all copy
// lives here as typed structs, the sections compose in page.templ, and the
// shared internal/ packages carry the reusable plumbing (intake, structdata,
// theme, ui).
//
// Everything is fictional. No real firm, attorney, result, or review is
// depicted. Attorney advertising is regulated: every results/outcome claim is
// modest and carries the "prior results do not guarantee a similar outcome"
// disclaimer, and the footer carries the full attorney-advertising notice —
// all clearly-marked placeholder text the firm's real counsel would review.
package law

import "github.com/firefly-software-mt/standard-template/internal/structdata"

// ── Domain structs ──────────────────────────────────────────────────────────

// FirmInfo is the firm's core identity and contact details.
type FirmInfo struct {
	Name       string // "Granite Peak" — short brand
	Suffix     string // "Trial Lawyers" — lockup line under the brand
	LegalName  string // "Granite Peak Trial Lawyers, PLLC"
	Tagline    string
	Closing    string // italic footer closing line
	Street     string
	Locality   string // city
	Region     string // state code
	PostalCode string
	PhoneLabel string // "(406) 555-0188"
	PhoneTel   string // "+14065550188"
	Email      string
	Latitude   float64
	Longitude  float64
	MapAltText string
}

// Stat is one figure in the hero trust strip.
type Stat struct {
	Number string // "40+", "1,200+", "$0", "24/7"
	Label  string
}

// MoneyPoint is one reassurance line in the contingency-fee block.
type MoneyPoint struct {
	Icon string // lucide name
	Text string
}

// PracticeArea is one row in the editorial practice-areas index.
type PracticeArea struct {
	Icon  string // lucide name
	Title string
	Desc  string
}

// Step is one node in the "how it works after a crash" timeline.
type Step struct {
	Title string
	Body  string
}

// Attorney is one card in the team grid. Photo (a webp/jpg under static/img/)
// is the headshot; when empty, a labelled placeholder panel is shown instead.
type Attorney struct {
	Name        string
	Role        string
	Credentials string // bar number + focus line
	Bio         string
	Photo       string // filename under static/img/; "" = placeholder panel
	PhotoAlt    string
}

// Result is one modest, credible outcome card. The story leads; the figure
// stays quiet. Every results display sits under the section disclaimer.
type Result struct {
	Type     string // "Auto accident · Settled"
	Story    string
	FigLabel string // "Recovered" / "Jury verdict"
	FigValue string // modest, rounded — "$500,000"
}

// Review is one testimonial. Rating is 1–5; it renders as cards and as
// schema.org Review nodes from this same data.
type Review struct {
	Quote    string
	Name     string // "Marcus T."
	Initials string // "MT"
	Location string // "Helena, MT"
	Rating   int
}

// Option is one <option> in an intake select (value + visible label).
type Option struct {
	Value string
	Label string
}

// LocRow is one line in the location/availability list.
type LocRow struct {
	Icon    string // lucide name
	Heading string
	// Body is rendered with a small amount of inline markup (a link), so it is
	// modelled as plain text plus an optional link appended after it.
	Body     string
	LinkText string // optional trailing link label
	LinkHref string // optional trailing link href
	LinkTel  bool   // render the link as the firm tel: number
}

// Content is the whole site's content, assembled once.
type Content struct {
	Firm           FirmInfo
	HeroEyebrow    string
	HeroTitle      string
	HeroSub        string
	HeroImage      string // base name under static/img/ ("hero" → hero-600/hero-1120.webp); "" = placeholder panel
	HeroImageAlt   string // alt/label for the hero media
	Stats          []Stat
	MoneyTitle     []string // two parts; the second renders in the brass italic accent
	MoneyBody      string
	MoneyPoints    []MoneyPoint
	PracticeAreas  []PracticeArea
	Steps          []Step
	Attorneys      []Attorney
	Results        []Result
	ResultsDisc    string // section-level disclaimer
	Reviews        []Review
	Incidents      []Option // "what happened" select
	ContactMethods []Option // "best way to reach you" select
	LocRows        []LocRow
	LegalNotice    string // footer attorney-advertising / disclaimer block
}

// ── Populated content ─────────────────────────────────────────────────────────

// site returns the fully populated Granite Peak content. Kept as a function
// (not a package var) so the data is built per render and stays easy to read.
func site() Content {
	firm := FirmInfo{
		Name:       "Granite Peak",
		Suffix:     "Trial Lawyers",
		LegalName:  "Granite Peak Trial Lawyers, PLLC",
		Tagline:    "Real trial lawyers. Ready to fight for you.",
		Closing:    "Montana's people, Montana's courts — we're ready.",
		Street:     "318 Fuller Ave",
		Locality:   "Helena",
		Region:     "MT",
		PostalCode: "59601",
		PhoneLabel: "(406) 555-0188",
		PhoneTel:   "+14065550188",
		Email:      "intake@granitepeaktriallaw.com",
		Latitude:   46.5884,
		Longitude:  -112.0245,
		MapAltText: "Map showing Granite Peak Trial Lawyers at 318 Fuller Avenue, Helena, Montana",
	}

	return Content{
		Firm:         firm,
		HeroEyebrow:  "Helena, Montana · Personal-injury trial lawyers",
		HeroTitle:    "Hurt in a crash? You shouldn't have to fight the insurance company alone.",
		HeroSub:      "Granite Peak is a Montana trial firm. We handle the insurance company, build your case, and we're ready to take it to court — no fee unless we win.",
		HeroImage:    "hero",
		HeroImageAlt: "A Granite Peak attorney shaking hands with a client across the desk in the firm's warm, book-lined Helena office",
		Stats: []Stat{
			{"40+", "Years combined trial experience"},
			{"1,200+", "Montana injury cases handled"},
			{"$0", "Up-front cost to you"},
			{"24/7", "We answer the phone"},
		},
		MoneyTitle: []string{"You don't pay", "unless we win."},
		MoneyBody:  "The first conversation is free, and you'll never get a surprise bill from us. We work on contingency — our fee comes out of what we recover for you, not your pocket. If we don't win, you owe us nothing.",
		MoneyPoints: []MoneyPoint{
			{"wallet", "The first conversation is free — and you'll never get a surprise bill from us."},
			{"badge-percent", "We work on contingency. Our fee comes out of what we recover, not your pocket."},
			{"shield-check", "If we don't win, you owe us nothing. No recovery, no fee."},
		},
		PracticeAreas: []PracticeArea{
			{"car-front", "Auto Accidents", "Cars, motorcycles, pedestrians — when someone else's mistake left you hurt."},
			{"truck", "Truck & Commercial Vehicle", "Bigger vehicles, bigger insurers, harder fights. We know how to take them on."},
			{"hard-hat", "Workplace Injury", "Hurt on the job? Know your rights before you sign anything."},
			{"heart-handshake", "Wrongful Death", "Compassionate, determined representation when your family has lost someone."},
		},
		Steps: []Step{
			{"Free case review", "Tell us what happened. No cost, no obligation."},
			{"We take it from here", "We deal with the insurance company and the paperwork."},
			{"We build your case", "Thorough, documented, and ready for court if it comes to that."},
			{"You focus on healing", "We fight for what you're owed."},
		},
		Attorneys: []Attorney{
			{
				Name: "Sarah Whitcomb", Role: "Founding Partner · Trial Attorney",
				Credentials: "Montana Bar #4821 · 18 yrs trying cases",
				Bio:         "A former insurance-defense lawyer who now uses that playbook for the people. Direct, warm, and relentless in trial.",
				Photo:       "attorney-whitcomb.webp", PhotoAlt: "Headshot of Sarah Whitcomb, founding partner",
			},
			{
				Name: "Daniel Reyes", Role: "Partner · Trial Attorney",
				Credentials: "Montana Bar #6093 · Truck & commercial vehicle",
				Bio:         "Handles the firm's trucking and commercial-vehicle cases. Calm and meticulous — reads every page of the policy. Coaches youth wrestling.",
				Photo:       "attorney-reyes.webp", PhotoAlt: "Headshot of Daniel Reyes, partner",
			},
			{
				Name: "Megan Holloway", Role: "Associate Attorney",
				Credentials: "Montana Bar #7140 · Workplace & wrongful death",
				Bio:         "Workplace injury and wrongful death. Known for keeping clients genuinely informed at every step. Grew up in Helena.",
				Photo:       "attorney-holloway.webp", PhotoAlt: "Headshot of Megan Holloway, associate attorney",
			},
		},
		Results: []Result{
			{"Auto accident · Settled",
				"A highway collision the insurer first denied. We documented every injury and recovered the full policy limits for the family.",
				"Recovered", "$500,000"},
			{"Workplace injury · Jury verdict",
				"A safety guard removed to save time, and an insurer who offered almost nothing. We took it to a Helena jury.",
				"Jury verdict", "$1.2M"},
			{"Truck accident · Settled",
				"An 18-wheeler ran a light and the trucking company's insurer dug in. We resolved it after depositions, before trial.",
				"Recovered", "$750,000"},
		},
		ResultsDisc: "Prior results do not guarantee a similar outcome. Every case is different and must be evaluated on its own facts.",
		Reviews: []Review{
			{"The adjuster kept calling and I felt buried. Granite Peak took it all over the same day, kept me updated every step, and got a fair result. I never felt like just a case number.",
				"Marcus T.", "MT", "Helena, MT", 5},
			{"After my husband's accident I was overwhelmed and scared about money. They were honest, patient, and never rushed us — and they handled the insurance company so we could grieve.",
				"Donna R.", "DR", "Helena, MT", 5},
			{"I almost signed what the insurer first offered. So glad I called first. They explained everything in plain language and recovered far more than I expected.",
				"Curtis B.", "CB", "Helena, MT", 5},
		},
		Incidents: []Option{
			{"", "Select what happened"},
			{"auto", "Auto accident"},
			{"truck", "Truck or commercial vehicle"},
			{"workplace", "Workplace injury"},
			{"wrongful-death", "Wrongful death"},
			{"other", "Something else"},
		},
		ContactMethods: []Option{
			{"phone", "A phone call"},
			{"text", "A text message"},
			{"email", "An email"},
		},
		LocRows: []LocRow{
			{Icon: "map-pin", Heading: "Our office",
				Body:     firm.Street + ", " + firm.Locality + ", " + firm.Region + " " + firm.PostalCode,
				LinkText: "Get directions",
				LinkHref: "https://maps.google.com/?q=" + "318+Fuller+Ave,+Helena,+MT+59601"},
			{Icon: "clock", Heading: "Hours",
				Body: "Office: Mon–Fri 8:00a–5:30p · Phones answered 24/7"},
			{Icon: "phone", Heading: "Call us",
				Body: "— a real person answers, any time", LinkText: firm.PhoneLabel, LinkTel: true},
			{Icon: "accessibility", Heading: "Getting here",
				Body: "Free parking · Wheelchair accessible"},
		},
		LegalNotice: "Attorney advertising. This website is for informational purposes only and is not legal advice. Contacting Granite Peak Trial Lawyers, or submitting the case-review form, does not create an attorney-client relationship. Prior results do not guarantee a similar outcome; every case is different and must be evaluated on its own facts. The attorneys of Granite Peak Trial Lawyers are licensed to practice law in the State of Montana.",
	}
}

// jsonLD builds the schema.org LegalService node for the page <head>, including
// aggregateRating and Review nodes generated from the same review data that
// renders the on-page cards. baseURL is the site's public origin. Emitting the
// reviews from the site's own structs (not a plugin) is part of the pitch.
func (c Content) jsonLD(baseURL string) structdata.LocalBusiness {
	f := c.Firm
	reviews := make([]structdata.Review, 0, len(c.Reviews))
	for _, r := range c.Reviews {
		reviews = append(reviews, structdata.Review{
			Type:         "Review",
			Author:       structdata.Author{Type: "Person", Name: r.Name},
			ReviewRating: structdata.Rating{Type: "Rating", RatingValue: itoa(r.Rating), BestRating: "5"},
			ReviewBody:   r.Quote,
		})
	}

	return structdata.LocalBusiness{
		Context:     "https://schema.org",
		Type:        "LegalService",
		Name:        f.LegalName,
		Description: c.HeroSub,
		URL:         baseURL,
		Telephone:   f.PhoneLabel,
		Email:       f.Email,
		PriceRange:  "No fee unless we win",
		Address: &structdata.PostalAddress{
			Type:            "PostalAddress",
			StreetAddress:   f.Street,
			AddressLocality: f.Locality,
			AddressRegion:   f.Region,
			PostalCode:      f.PostalCode,
			AddressCountry:  "US",
		},
		Geo: &structdata.GeoCoordinates{
			Type: "GeoCoordinates", Latitude: f.Latitude, Longitude: f.Longitude,
		},
		// Office hours for in-person; phones are answered 24/7 but the office
		// hours are what schema.org openingHours describes.
		OpeningHours: []structdata.OpeningHours{
			{Type: "OpeningHoursSpecification", DayOfWeek: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}, Opens: "08:00", Closes: "17:30"},
		},
		AggregateRating: &structdata.AggregateRating{
			Type: "AggregateRating", RatingValue: "5.0", ReviewCount: itoa(len(c.Reviews)), BestRating: "5",
		},
		Review: reviews,
	}
}

// itoa is a tiny int→string helper to avoid pulling strconv into content.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
