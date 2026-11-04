package view

import (
	"strings"

	"github.com/firefly-software-mt/standard-template/internal/structdata"
)

// FAQParagraph is one paragraph of an FAQ answer. Lead, if set, renders bold
// ahead of Text (e.g. "Hosting:"); either may be empty.
type FAQParagraph struct {
	Lead string
	Text string
}

// FAQEntry is one question on the Services page. The same data renders the
// on-page FAQ and the FAQPage JSON-LD, so the two can't drift apart.
type FAQEntry struct {
	Q string
	A []FAQParagraph
}

// plainAnswer flattens an answer to plain text for structured data.
func (f FAQEntry) plainAnswer() string {
	parts := make([]string, 0, len(f.A))
	for _, p := range f.A {
		parts = append(parts, strings.TrimSpace(p.Lead+" "+p.Text))
	}
	return strings.Join(parts, " ")
}

// faqLD builds a schema.org FAQPage node from FAQ entries.
func faqLD(entries []FAQEntry) structdata.FAQPage {
	qs := make([]structdata.Question, 0, len(entries))
	for _, f := range entries {
		qs = append(qs, structdata.Question{
			Type:           "Question",
			Name:           f.Q,
			AcceptedAnswer: structdata.Answer{Type: "Answer", Text: f.plainAnswer()},
		})
	}
	return structdata.FAQPage{Context: "https://schema.org", Type: "FAQPage", MainEntity: qs}
}

// studioLD is the site-wide ProfessionalService node for FlintCraft Studio.
// No street address is published on the site, so only locality and region are set.
func studioLD() structdata.LocalBusiness {
	return structdata.LocalBusiness{
		Context:     "https://schema.org",
		Type:        "ProfessionalService",
		Name:        SiteName,
		Description: "Custom websites for Montana businesses: research, copywriting, design, build, hosting and care.",
		URL:         BaseURL,
		Telephone:   Phone,
		Email:       Email,
		Image:       BaseURL + DefaultOGImage,
		Logo:        BaseURL + "/static/images/logo.png",
		Address: &structdata.PostalAddress{
			Type:            "PostalAddress",
			AddressLocality: "Helena",
			AddressRegion:   "MT",
			AddressCountry:  "US",
		},
		AreaServed: []string{"Helena, MT", "Montana", "United States"},
	}
}

// ServicesFAQ is the Services page FAQ, in display order.
var ServicesFAQ = []FAQEntry{
	{"What does the first conversation look like?", []FAQParagraph{
		{Text: "No pitch, no deck. We look at your current setup before we talk — site, search presence, how you're showing up for the customers you want. When we get on the call we walk through what we see: what's working, what isn't, and what we'd do differently. If it makes sense to work together, we'll tell you. If it doesn't, we'll tell you that too."},
	}},
	{"Do I own the site?", []FAQParagraph{
		{Lead: "Yes. As soon as the build is paid."},
		{Text: "Your domain, your content, your photos and your customer data are yours from day one, no matter what. The website itself (the design and the code) becomes yours the moment your build fee is paid in full. Pay up front and that's day one."},
		{Text: "If you ever move the site somewhere else, we transfer the files and source code to the host of your choice within 14 days. Nothing is held back."},
	}},
	{"How long does the build take?", []FAQParagraph{
		{Text: "Most sites launch in four to six weeks. Timeline depends on how quickly you can review and approve — we don't sit on things. The faster you can react, the faster you're live."},
	}},
	{"What counts as a monthly change?", []FAQParagraph{
		{Text: "Copy updates, new photos, team changes, updated hours, added service descriptions, seasonal content. On unlimited-edit plans that means text updates, photo swaps, service and price changes, hours, minor layout adjustments, and up to [four] new portfolio entries a month. New pages, structural redesigns and third-party integrations are quoted separately and priced honestly before we touch them."},
	}},
	{"What about paid ads — is that included?", []FAQParagraph{
		{Text: "No, but it's available. Google Ads and Facebook campaigns are scoped and managed separately for clients where paid visibility makes sense. We'll tell you in the first conversation whether we think it's worth it for your market."},
	}},
	{"Do you work with businesses outside Montana?", []FAQParagraph{
		{Text: "Yes. No geographic limit. We're based in Helena and Montana brands are close to our heart, but most of what we do works just as well over a call and a shared screen. If you're in Billings or Birmingham, the process is the same."},
	}},
	{"What happens if I want to cancel?", []FAQParagraph{
		{Lead: "Hosting:", Text: "thirty days' written notice, any time. We hand over your site files."},
		{Lead: "If you're on a payment plan:", Text: "the remaining build balance is due when hosting ends. Then the site is fully yours."},
		{Lead: "If you cancel during the build:", Text: "we bill the time worked at our hourly rate, apply what you've already paid, and refund anything left over. If we're the ones who back out, you get every build payment back."},
	}},
}
