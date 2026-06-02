// Package chiropractor is the first FlintCraft demo vertical: a single-page
// marketing site for the fictional Alpine Spine Chiropractic & Wellness in
// Helena, Montana. It is the template later verticals (law, dentist, …) copy:
// all copy lives here as typed structs, the sections compose in page.templ,
// and the shared internal/ packages carry the reusable plumbing.
//
// Everything is fictional. No real practice, doctor, or review is depicted.
package chiropractor

import "github.com/firefly-software-mt/standard-template/internal/structdata"

// ── Domain structs ──────────────────────────────────────────────────────────

// BusinessInfo is the practice's core identity and contact details.
type BusinessInfo struct {
	Name       string // "Alpine Spine"
	LegalName  string // "Alpine Spine Chiropractic & Wellness"
	Tagline    string
	Street     string
	Locality   string // city
	Region     string // state code
	PostalCode string
	PhoneLabel string // "(406) 555-0142"
	PhoneTel   string // "+14065550142"
	Email      string
	Latitude   float64
	Longitude  float64
	MapAltText string
}

// Hours is one row of the location hours table.
type Hours struct {
	Days   string
	Time   string
	Closed bool
}

// Stat is one figure in the trust band.
type Stat struct {
	Number string // leading number, e.g. "4.9"
	Unit   string // trailing accent, e.g. "★", "+", "k+"
	Label  string
}

// Path is one side of the dual-path hero (new vs. existing patient).
type Path struct {
	Variant  string // "new" | "exist" — drives the .path--{variant} class
	Eyebrow  string
	Title    string
	Desc     string
	Primary  CTA
	Ghost    CTA
	GhostTel bool // render the ghost CTA as a tel: link with a phone icon
}

// CTA is a labelled link target (anchor or tel:).
type CTA struct {
	Label string
	Href  string
}

// ExpectStep is one numbered "what to expect" step.
type ExpectStep struct {
	Title string
	Body  string
}

// SpecialtyIcon identifies which inline SVG a specialty card renders.
type SpecialtyIcon string

const (
	IconSports   SpecialtyIcon = "sports"
	IconFamily   SpecialtyIcon = "family"
	IconAuto     SpecialtyIcon = "auto"
	IconWellness SpecialtyIcon = "wellness"
)

// Specialty is one card in the specialties grid.
type Specialty struct {
	Icon     SpecialtyIcon
	Title    string
	Desc     string
	LinkText string
}

// Doctor is one provider card. Photo (a webp under static/img/) is the headshot;
// when empty, the Initials monogram on the Variant (a/b/c) gradient is shown
// instead. Bio is the always-visible blurb; LongBio expands inline via the
// card's "Read bio" disclosure.
type Doctor struct {
	Name        string
	Initials    string
	Variant     string // "a" | "b" | "c" — gradient when no Photo
	Credentials string
	Bio         string
	LongBio     string // expanded bio revealed by the card's disclosure toggle
	Photo       string // filename under static/img/, e.g. "doctor-whitlock.webp"
	PhotoAlt    string
}

// Review is one testimonial. Rating is 1–5; it renders as cards and as
// schema.org Review nodes from this same data.
type Review struct {
	Quote    string
	Name     string // "Dana R."
	Initials string // "DR"
	Location string
	Rating   int
}

// ReasonOption is one choice in the intake "reason for visit" select.
type ReasonOption string

// Content is the whole site's content, assembled once.
type Content struct {
	Business     BusinessInfo
	HeroTitle    []string // lines; last line renders in the italic accent
	HeroSub      string
	HeroBadge    string
	HeroImage    string // base name under static/img/ ("hero" → hero-600/hero-1120.webp); "" = gradient
	HeroImageAlt string
	CalmImage    string // reassurance image base ("reassure" → reassure-600/reassure-1000.webp); "" = gradient
	CalmImageAlt string
	Paths        []Path
	Stats        []Stat
	ExpectSteps  []ExpectStep
	Specialties  []Specialty
	Doctors      []Doctor
	TreatmentCol [][]string // treatments grouped into columns
	Reviews      []Review
	Reasons      []ReasonOption
	HoursRows    []Hours
	HoursSummary string
}

// ── Populated content ─────────────────────────────────────────────────────────

// site returns the fully populated Alpine Spine content. Kept as a function
// (not a package var) so the data is built per render and stays easy to read.
func site() Content {
	biz := BusinessInfo{
		Name:       "Alpine Spine",
		LegalName:  "Alpine Spine Chiropractic & Wellness",
		Tagline:    "Move better. Feel better. Live better.",
		Street:     "1240 Eagle Bend Dr",
		Locality:   "Helena",
		Region:     "MT",
		PostalCode: "59601",
		PhoneLabel: "(406) 555-0142",
		PhoneTel:   "+14065550142",
		Email:      "hello@alpinespinehelena.com",
		Latitude:   46.5927,
		Longitude:  -112.0205,
		MapAltText: "Map showing Alpine Spine at 1240 Eagle Bend Drive, Helena, Montana",
	}

	return Content{
		Business:     biz,
		HeroTitle:    []string{"Move better. Feel better.", "Live better."},
		HeroSub:      "Gentle, individualized chiropractic care in Helena — built around your body, your goals, and your pace.",
		HeroBadge:    "from 600+ Helena neighbors",
		HeroImage:    "hero",
		HeroImageAlt: "Two hikers on an alpine ridge at golden hour, mountains glowing in the Montana light",
		CalmImage:    "reassure",
		CalmImageAlt: "A smiling team member at the bright, welcoming Alpine Spine front desk",
		Paths: []Path{
			{
				Variant: "new", Eyebrow: "New here?", Title: "Your first visit",
				Desc:    "Never been to a chiropractor? You're in good hands — we'll walk you through everything.",
				Primary: CTA{"Book your first visit", "#book"},
				Ghost:   CTA{"What to expect", "#expect"},
			},
			{
				Variant: "exist", Eyebrow: "Existing patient", Title: "Welcome back",
				Desc:     "Pick up where you left off — your chart's ready when you are.",
				Primary:  CTA{"Book appointment", "#book"},
				Ghost:    CTA{"Call the office", "tel:" + biz.PhoneTel},
				GhostTel: true,
			},
		},
		Stats: []Stat{
			{"4.9", "★", "Average rating"},
			{"600", "+", "Five-star reviews"},
			{"15", "k+", "Patients cared for"},
			{"20", "+", "Treatment options"},
		},
		ExpectSteps: []ExpectStep{
			{"We listen", "A real conversation about what's going on — what hurts, and what you want to get back to."},
			{"We assess", "A thorough, gentle exam to see how you move and where you're holding tension. No surprises."},
			{"We explain", "Your options, in plain language, before anything happens — so you're always in the driver's seat."},
			{"We start", "Care at your pace, and only what you're comfortable with. Never more force than you need."},
		},
		Specialties: []Specialty{
			{IconSports, "Sports & Performance", "Keep training, recover faster, and move the way your sport demands.", "Get back in it"},
			{IconFamily, "Family & Pediatric", "Gentle care for every age, from newborns to grandparents.", "Care for your family"},
			{IconAuto, "Auto & Work Injury", "Clear documentation and focused recovery after a crash or on-the-job injury.", "Start your recovery"},
			{IconWellness, "General & Wellness", "Relief from the everyday aches of desk work, sleep, and life in Montana.", "Feel better daily"},
		},
		Doctors: []Doctor{
			{
				Name: "Dr. Sarah Whitlock", Initials: "SW", Variant: "a", Credentials: "DC, CCSP · Founder",
				Bio:      "Sports-focused and 14 years in practice — a former collegiate runner who's warm, thorough, and always explains the \"why.\"",
				LongBio:  "Sarah founded Alpine Spine in 2014 after a decade treating endurance athletes. A Certified Chiropractic Sports Physician and former collegiate runner, she's as comfortable with a marathoner's IT band as she is with a desk-worker's stiff neck. Patients leave knowing exactly what's going on and why — she won't touch you until you understand the plan.",
				Photo:    "doctor-whitlock.webp",
				PhotoAlt: "Dr. Sarah Whitlock, founder of Alpine Spine, smiling",
			},
			{
				Name: "Dr. Marcus Reyes", Initials: "MR", Variant: "b", Credentials: "DC · Family & Pediatric",
				Bio:      "Gentle, low-force techniques for the whole family. A dad of three who coaches Little League on the weekends.",
				LongBio:  "Marcus focuses on family and pediatric care, with gentle, low-force techniques suited to newborns, kids, and grandparents alike. A dad of three who coaches Little League, he's unhurried with nervous first-timers and especially good with the wiggly ones. Prenatal and postpartum care is a particular focus.",
				Photo:    "doctor-reyes.webp",
				PhotoAlt: "Dr. Marcus Reyes, family and pediatric chiropractor",
			},
			{
				Name: "Dr. Hannah Briggs", Initials: "HB", Variant: "c", Credentials: "DC, ART · Soft-tissue & Auto",
				Bio:      "Soft-tissue and auto-injury specialist. Calm, detail-oriented, and known for never rushing an exam.",
				LongBio:  "Hannah is our soft-tissue and auto-injury specialist, certified in Active Release Technique. Calm and meticulous, she's known for thorough documentation that holds up with insurers and attorneys after a crash or on-the-job injury — and for never rushing an exam, no matter how full the schedule.",
				Photo:    "doctor-briggs.webp",
				PhotoAlt: "Dr. Hannah Briggs, soft-tissue and auto-injury specialist",
			},
		},
		TreatmentCol: [][]string{
			{"Spinal Adjustments", "Extremity Adjustments", "Active Release (ART)", "Graston / IASTM", "Dry Needling", "Cupping"},
			{"Kinesiology Taping", "Electric Muscle Stim", "Therapeutic Ultrasound", "Shockwave Therapy", "Spinal Decompression", "Corrective Exercise"},
			{"Postural Rehab", "Prenatal Care", "Pediatric Care", "Sports Massage", "Custom Orthotics", "Nutritional Guidance"},
		},
		Reviews: []Review{
			{"I came in barely able to turn my neck. Dr. Briggs actually listened first and never rushed me. Three weeks later I'm back on my bike.",
				"Dana R.", "DR", "Helena, MT", 5},
			{"First-timer, honestly nervous about being \"cracked.\" Dr. Reyes walked me through every single step. Calmest clinic I've ever been in.",
				"Tyler K.", "TK", "Helena, MT", 5},
			{"They worked around my pregnancy with so much care. I left every visit feeling lighter and genuinely looked after. Wish I'd come sooner.",
				"Maria S.", "MS", "Helena, MT", 5},
		},
		Reasons: []ReasonOption{
			"Back or neck pain", "Sports injury", "Auto or work injury",
			"Prenatal or pediatric", "Routine adjustment / wellness", "Something else",
		},
		HoursRows: []Hours{
			{"Monday – Thursday", "7:30a – 6:00p", false},
			{"Friday", "7:30a – 3:00p", false},
			{"Saturday", "8:00a – 12:00p", false},
			{"Sunday", "Closed", true},
		},
		HoursSummary: "Mon–Thu 7:30a–6p · Sat 8a–12p",
	}
}

// jsonLD builds the schema.org ChiropracticBusiness node for the page <head>,
// including aggregateRating and Review nodes generated from the same review
// data that renders the on-page cards. baseURL is the site's public origin.
func (c Content) jsonLD(baseURL string) structdata.LocalBusiness {
	b := c.Business
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
		Type:        "ChiropracticBusiness",
		Name:        b.LegalName,
		Description: c.HeroSub,
		URL:         baseURL,
		Telephone:   b.PhoneLabel,
		Email:       b.Email,
		PriceRange:  "$$",
		Address: &structdata.PostalAddress{
			Type:            "PostalAddress",
			StreetAddress:   b.Street,
			AddressLocality: b.Locality,
			AddressRegion:   b.Region,
			PostalCode:      b.PostalCode,
			AddressCountry:  "US",
		},
		Geo: &structdata.GeoCoordinates{
			Type: "GeoCoordinates", Latitude: b.Latitude, Longitude: b.Longitude,
		},
		OpeningHours: []structdata.OpeningHours{
			{Type: "OpeningHoursSpecification", DayOfWeek: []string{"Monday", "Tuesday", "Wednesday", "Thursday"}, Opens: "07:30", Closes: "18:00"},
			{Type: "OpeningHoursSpecification", DayOfWeek: []string{"Friday"}, Opens: "07:30", Closes: "15:00"},
			{Type: "OpeningHoursSpecification", DayOfWeek: []string{"Saturday"}, Opens: "08:00", Closes: "12:00"},
		},
		AggregateRating: &structdata.AggregateRating{
			Type: "AggregateRating", RatingValue: "4.9", ReviewCount: "600", BestRating: "5",
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
