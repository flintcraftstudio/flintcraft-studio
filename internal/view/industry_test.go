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
		"Chiropractor Website Design in Montana",                  // H1
		"Montana Chiropractor Website Design | FlintCraft Studio", // <title> (Base appends suffix)
		`rel="canonical" href="https://flintcraftstudio.com/chiropractor-website-design"`,
		`property="og:title"`,
		`"@type":"Service"`, // service JSON-LD
		`"serviceType":"chiropractor website design"`,
		`"@type":"FAQPage"`,           // FAQ JSON-LD
		`href="/demos/chiropractor/"`, // links to live demo
		`target="_blank"`,             // opens in new tab
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

func TestIndustriesHubRenders(t *testing.T) {
	html := renderView(t, IndustriesHubPage())
	for _, s := range []string{
		"Built for your",
		`href="/chiropractor-website-design"`, // live vertical links out
		"Coming soon",                         // dentist / law placeholders
		"Dentists",
		"Law firms",
	} {
		if !strings.Contains(html, s) {
			t.Errorf("hub missing %q", s)
		}
	}
}
