package law

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/firefly-software-mt/standard-template/internal/intake"
)

// testState returns a representative pageState for rendering tests.
func testState() pageState {
	prefix := "/demos/law"
	return pageState{
		prefix:  prefix,
		baseURL: "https://flintcraftstudio.com" + prefix,
		theme:   graniteTheme(prefix),
	}
}

func TestPageRenders(t *testing.T) {
	var b strings.Builder
	if err := Page(site(), testState()).Render(context.Background(), &b); err != nil {
		t.Fatalf("render page: %v", err)
	}
	html := b.String()

	// Dump for manual inspection.
	_ = os.WriteFile("/tmp/law-page.html", []byte(html), 0o644)

	must := []string{
		"Granite Peak",                                 // brand
		"Granite Peak Trial Lawyers, PLLC",             // legal name in title
		"(406) 555-0188",                               // phone label renders
		"fight the insurance company alone",            // hero headline
		"unless we win",                                // money-fear accent + promise
		`application/ld+json`,                          // JSON-LD present
		`"@type":"LegalService"`,                       // structured data type
		`"aggregateRating"`,                            // rating in JSON-LD
		`"@type":"Review"`,                             // review nodes in JSON-LD
		"Sarah Whitcomb",                               // attorney
		"Auto Accidents",                               // practice area
		"Prior results do not guarantee",               // results disclaimer (visible)
		"Attorney advertising.",                        // footer legal notice
		"does not create an attorney-client",           // intake + footer disclaimer
		`hx-post="/demos/law/intake"`,                  // htmx wiring
		`action="/demos/law/intake"`,                   // no-JS fallback
		`name="company"`,                               // honeypot
		"/demos/law/static/css/site.css",               // themed stylesheet
		"/demos/law/static/img/logo-mark.png",          // nav logo
		`class="skip-link"`,                            // a11y skip link
		"Helena, MT",                                   // review location
		`/demos/law/static/img/attorney-whitcomb.webp`, // real headshot wired in
		`/demos/law/static/img/hero-1120.webp`,         // hero photo wired in
		`loading="eager"`,                              // hero is eager (LCP)
		`rel="preload"`,                                // hero preloaded for LCP
		`fetchpriority="high"`,                         // hero prioritized
		`loading="lazy"`,                               // below-fold headshots lazy-loaded
		`width="640" height="576"`,                     // headshot dimensions (no layout shift)
		`class="map-pin"`,                              // CSS-drawn location map (no raster)
		`role="img" aria-label=`,                       // hero/map remain accessible image regions
		`name="robots" content="noindex`,               // demo kept out of search
		`class="demo-bar"`,                             // FlintCraft framing bar
		`href="/contact?ref=law-demo"`,                 // bar links back with referral tag
		`class="hero2-stats"`,                          // trust stats strip
		`class="pindex"`,                               // editorial practice index
		`class="timeline"`,                             // how-it-works timeline
		"reveal-armed",                                 // reveal script present (JS-armed)
		"No fee unless we win",                         // contingency promise
		`id="fcToast"`,                                 // celebration toast present
		"Talk to FlintCraft",                           // toast CTA back to marketing
		`href="/contact?ref=law-demo"`,                 // toast carries the demo referral tag
		"window.fcCelebrate",                           // celebration script wired
		".form-success",                                // celebration keyed to the success partial
		`class="mcall"`,                                // mobile persistent call bar
		`class="btn-sending"`,                          // intake submit pending state
		`role="dialog"`,                                // mobile drawer is a labelled dialog
		`aria-modal="true"`,                            // …that traps to the modal
		`x-on:keydown.escape.window="closeMenu()"`,     // Esc closes the drawer
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
	errs := intake.Errors{
		"name":   "Please tell us your name",
		"reason": "Tell us what happened",
	}
	sub := intake.Submission{Phone: "406", Email: "nope"}
	var b strings.Builder
	if err := IntakeCard(site(), "/demos/law", sub, errs).Render(context.Background(), &b); err != nil {
		t.Fatalf("render intake card: %v", err)
	}
	html := b.String()
	for _, s := range []string{`class="field err"`, "Please tell us your name", "Tell us what happened", `value="406"`} {
		if !strings.Contains(html, s) {
			t.Errorf("intake card missing %q", s)
		}
	}
}

func TestIntakeSuccessIsAnnouncedRegion(t *testing.T) {
	var b strings.Builder
	if err := IntakeSuccess(site(), "Jordan").Render(context.Background(), &b); err != nil {
		t.Fatalf("render success: %v", err)
	}
	html := b.String()
	// The confirmation must be a focusable, announced region (celebrateScript
	// focuses it after the htmx swap so screen readers hear it).
	for _, s := range []string{`role="status"`, `tabindex="-1"`} {
		if !strings.Contains(html, s) {
			t.Errorf("success partial missing %q", s)
		}
	}
}

func TestPageNoJSIntakeState(t *testing.T) {
	// Error state: the full page should embed the form with the error.
	st := testState()
	st.intake = &intakeState{
		sub:  intake.Submission{Phone: "12"},
		errs: intake.Errors{"phone": "We need a number to reach you"},
	}
	var b strings.Builder
	if err := Page(site(), st).Render(context.Background(), &b); err != nil {
		t.Fatalf("render: %v", err)
	}
	if html := b.String(); !strings.Contains(html, "We need a number to reach you") || !strings.Contains(html, `class="field err"`) {
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
	if err := IntakeSuccess(site(), "Jordan").Render(context.Background(), &b); err != nil {
		t.Fatalf("render success: %v", err)
	}
	if !strings.Contains(b.String(), "Thanks, Jordan") {
		t.Errorf("success partial did not greet by name")
	}
}
