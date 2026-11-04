package view

import (
	"strings"
	"testing"
)

func TestHomeMetaAndStudioLD(t *testing.T) {
	withBaseURL(t)
	html := renderView(t, HomePage())
	for _, s := range []string{
		"<title>Custom Websites for Montana Businesses | FlintCraft Studio, Helena</title>",
		`"@type":"ProfessionalService"`,
		`"addressLocality":"Helena"`,
		`og:image" content="https://flintcraftstudio.com/static/images/og-default.jpg"`,
	} {
		if !strings.Contains(html, s) {
			t.Errorf("home missing %q", s)
		}
	}
}

func TestServicesFAQStructuredData(t *testing.T) {
	withBaseURL(t)
	html := renderView(t, ServicesPage())
	for _, s := range []string{
		`"@type":"FAQPage"`,
		`"name":"Do I own the site?"`,
		"Yes. As soon as the build is paid. Your domain", // flattened answer in JSON-LD
		"Shaped with intent.",                            // tagline in meta description
	} {
		if !strings.Contains(html, s) {
			t.Errorf("services missing %q", s)
		}
	}
	if n := strings.Count(html, "<details"); n != len(ServicesFAQ) {
		t.Errorf("rendered %d FAQ items, want %d", n, len(ServicesFAQ))
	}
}

func TestProjectCanonicalAndImage(t *testing.T) {
	withBaseURL(t)
	p := Projects[ProjectOrder[0]]
	html := renderView(t, ProjectPage(p))
	for _, s := range []string{
		`rel="canonical" href="https://flintcraftstudio.com/work/` + p.Slug + `"`,
		`og:image" content="https://flintcraftstudio.com` + p.Screenshot + `"`,
	} {
		if !strings.Contains(html, s) {
			t.Errorf("project page missing %q", s)
		}
	}
}

// withBaseURL sets the production origin for the duration of a test, since
// canonical/OG tags only render when BaseURL is set.
func withBaseURL(t *testing.T) {
	t.Helper()
	old := BaseURL
	BaseURL = "https://flintcraftstudio.com"
	t.Cleanup(func() { BaseURL = old })
}
