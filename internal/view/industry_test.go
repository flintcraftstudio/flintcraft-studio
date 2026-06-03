package view

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func renderView(t *testing.T, c templ.Component) string {
	t.Helper()
	var b strings.Builder
	if err := c.Render(context.Background(), &b); err != nil {
		t.Fatalf("render: %v", err)
	}
	return b.String()
}

func TestChiropractorLandingRenders(t *testing.T) {
	old := BaseURL
	BaseURL = "https://flintcraftstudio.com"
	defer func() { BaseURL = old }()

	d := IndustryLandings["chiropractor-website-design"]
	var b strings.Builder
	if err := IndustryLandingPage(d).Render(context.Background(), &b); err != nil {
		t.Fatalf("render: %v", err)
	}
	html := b.String()

	must := []string{
		"books patients while you adjust them",                    // humanized H1
		"chiropractor website design",                             // keyword retained in body (lede)
		"Montana Chiropractor Website Design | FlintCraft Studio", // <title> (Base appends suffix)
		`rel="canonical" href="https://flintcraftstudio.com/chiropractor-website-design"`,
		`property="og:title"`,
		`"@type":"Service"`, // service JSON-LD
		`"serviceType":"chiropractor website design"`,
		`"@type":"FAQPage"`,           // FAQ JSON-LD
		`href="/demos/chiropractor/"`, // links to live demo
		`target="_blank"`,             // opens in new tab
		`href="#see-it"`,              // hero secondary CTA scrolls to proof, not a 3rd demo link
		"WordPress",                   // comparison
		"Helena",                      // local SEO cities
	}
	for _, s := range must {
		if !strings.Contains(html, s) {
			t.Errorf("landing page missing %q", s)
		}
	}
	// The landing page must be indexable (no robots noindex).
	if strings.Contains(html, "noindex") {
		t.Error("landing page should not be noindex")
	}
	if n := strings.Count(html, "<h1"); n != 1 {
		t.Errorf("expected exactly one <h1>, got %d", n)
	}
}

func TestLawFirmLandingRenders(t *testing.T) {
	old := BaseURL
	BaseURL = "https://flintcraftstudio.com"
	defer func() { BaseURL = old }()

	d := IndustryLandings["law-firm-website-design"]
	var b strings.Builder
	if err := IndustryLandingPage(d).Render(context.Background(), &b); err != nil {
		t.Fatalf("render: %v", err)
	}
	html := b.String()

	must := []string{
		"earns the call when someone needs you most",          // humanized H1
		"law firm website design",                             // keyword retained in body
		"Montana Law Firm Website Design | FlintCraft Studio", // <title>
		`rel="canonical" href="https://flintcraftstudio.com/law-firm-website-design"`,
		`"@type":"Service"`,
		`"serviceType":"law firm website design"`,
		`"@type":"FAQPage"`,
		`href="/demos/law/"`,         // links to the live law demo
		`target="_blank"`,            // opens in new tab
		"Granite Peak Trial Lawyers", // demo label
		"Helena",                     // local SEO cities
	}
	for _, s := range must {
		if !strings.Contains(html, s) {
			t.Errorf("law landing missing %q", s)
		}
	}
	if strings.Contains(html, "noindex") {
		t.Error("landing page should not be noindex")
	}
	if n := strings.Count(html, "<h1"); n != 1 {
		t.Errorf("expected exactly one <h1>, got %d", n)
	}
}

func TestIndustriesHubRenders(t *testing.T) {
	html := renderView(t, IndustriesHubPage())
	for _, s := range []string{
		"Built for your",
		`href="/chiropractor-website-design"`, // live vertical links out
		`href="/law-firm-website-design"`,     // law vertical now live
		"Coming soon",                         // dentist placeholder remains
		"Dentists",
		"Law firms",
	} {
		if !strings.Contains(html, s) {
			t.Errorf("hub missing %q", s)
		}
	}
}
