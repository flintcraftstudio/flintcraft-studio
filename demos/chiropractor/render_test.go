package chiropractor

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/firefly-software-mt/standard-template/internal/intake"
)

// testState returns a representative pageState for rendering tests.
func testState() pageState {
	prefix := "/demos/chiropractor"
	return pageState{
		prefix:  prefix,
		baseURL: "https://flintcraftstudio.com" + prefix,
		theme:   alpineTheme(prefix),
	}
}

func TestPageRenders(t *testing.T) {
	var b strings.Builder
	if err := Page(site(), testState()).Render(context.Background(), &b); err != nil {
		t.Fatalf("render page: %v", err)
	}
	html := b.String()

	// Dump for manual inspection.
	_ = os.WriteFile("/tmp/chiro-page.html", []byte(html), 0o644)

	must := []string{
		"Alpine Spine", // brand
		"Alpine Spine Chiropractic &amp; Wellness", // legal name in title (escaped)
		"(406) 555-0142",                          // phone label actually renders
		"Move better. Feel better.",               // hero headline
		"Live better.",                            // italic accent line
		`application/ld+json`,                     // JSON-LD present
		`"@type":"ChiropracticBusiness"`,          // structured data type
		`"aggregateRating"`,                       // rating in JSON-LD
		"Spinal Adjustments",                      // treatments
		"Dr. Sarah Whitlock",                      // doctor
		`hx-post="/demos/chiropractor/intake"`,    // htmx wiring
		`action="/demos/chiropractor/intake"`,     // no-JS fallback
		`name="company"`,                          // honeypot
		"/demos/chiropractor/static/css/site.css", // themed stylesheet
		`class="skip-link"`,                       // a11y skip link
		"Helena, MT",                              // review location
		`/static/img/doctor-whitlock.webp`,        // real headshot
		`loading="lazy"`,                          // below-fold images lazy-loaded
		`width="640" height="576"`,                // headshot dimensions (no layout shift)
		`/static/img/hero-1120.webp`,              // real hero photo
		`/static/img/reassure-1000.webp`,          // reassurance photo
		`loading="eager"`,                         // hero is eager (LCP)
		`rel="preload" as="image"`,                // hero preloaded for LCP
		`fetchpriority="high"`,                    // hero prioritized
		`name="robots" content="noindex`,          // demo kept out of search
		`class="demo-bar"`,                        // FlintCraft framing bar
		`href="/contact?ref=chiropractor-demo"`,   // bar links back to marketing with referral tag
		`class="spec-row"`,                        // specialties are whole-row links now
		`class="link-arrow doc-card__toggle"`,     // doctor bio is an accessible disclosure
		"founded Alpine Spine in 2014",            // expanded doctor bio present
		`class="btn-sending"`,                     // intake submit pending state
		`class="expect__grid" data-reveal`,        // steps marked for staggered reveal
		"reveal-armed",                            // reveal script present (JS-armed)
		`id="fcToast"`,                            // celebration toast present
		"Talk to FlintCraft",                      // toast CTA back to marketing
		"window.fcCelebrate",                      // celebration script wired
	}
	for _, s := range must {
		if !strings.Contains(html, s) {
			t.Errorf("page HTML missing %q", s)
		}
	}

	// There must be exactly one <h1>.
	if n := strings.Count(html, "<h1"); n != 1 {
		t.Errorf("expected exactly one <h1>, got %d", n)
	}
}

func TestIntakeCardShowsErrors(t *testing.T) {
	errs := intake.Errors{"name": "Please tell us your name.", "email": "Add an email for your confirmation."}
	sub := intake.Submission{Phone: "406", Email: "nope"}
	var b strings.Builder
	if err := IntakeCard(site(), "/demos/chiropractor", sub, errs).Render(context.Background(), &b); err != nil {
		t.Fatalf("render intake card: %v", err)
	}
	html := b.String()
	for _, s := range []string{"field field--error", "Please tell us your name.", "Add an email for your confirmation.", `value="406"`} {
		if !strings.Contains(html, s) {
			t.Errorf("intake card missing %q", s)
		}
	}
}

func TestPageNoJSIntakeState(t *testing.T) {
	// Error state: the full page should embed the form with the error.
	st := testState()
	st.intake = &intakeState{
		sub:  intake.Submission{Phone: "12"},
		errs: intake.Errors{"phone": "A number we can reach you at, please."},
	}
	var b strings.Builder
	if err := Page(site(), st).Render(context.Background(), &b); err != nil {
		t.Fatalf("render: %v", err)
	}
	if html := b.String(); !strings.Contains(html, "A number we can reach you at, please.") || !strings.Contains(html, "field field--error") {
		t.Error("no-JS error page missing inline form error")
	}

	// Success state: the full page should embed the confirmation.
	st2 := testState()
	st2.intake = &intakeState{sub: intake.Submission{Name: "Pat Lee"}, success: true}
	var b2 strings.Builder
	if err := Page(site(), st2).Render(context.Background(), &b2); err != nil {
		t.Fatalf("render: %v", err)
	}
	if !strings.Contains(b2.String(), "Thanks, Pat") {
		t.Error("no-JS success page missing confirmation")
	}
}

func TestIntakeSuccessGreetsByName(t *testing.T) {
	var b strings.Builder
	if err := IntakeSuccess("Jordan").Render(context.Background(), &b); err != nil {
		t.Fatalf("render success: %v", err)
	}
	if !strings.Contains(b.String(), "Thanks, Jordan") {
		t.Errorf("success partial did not greet by name")
	}
}
