package chiropractor

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/templ"

	"github.com/firefly-software-mt/standard-template/internal/intake"
	"github.com/firefly-software-mt/standard-template/internal/theme"
)

// intakeLogFile is where intake submissions are appended for the demo (no real
// email is sent). Path is relative to the server's working directory and is
// gitignored. See the TODO in internal/intake for wiring a real mailer.
const intakeLogFile = "demos/chiropractor/intake-submissions.log"

// Register mounts the Alpine Spine demo under prefix (e.g. "/demos/chiropractor")
// on mux. baseURL is the site's public origin, used only for the JSON-LD url.
//
// This is the single wiring point each vertical exposes — adding the next demo
// is one more Register call in the marketing site's router. Switching from a
// path prefix to host-based routing only changes the prefix passed here.
func Register(mux *http.ServeMux, prefix, baseURL string) {
	prefix = strings.TrimRight(prefix, "/")

	st := pageState{
		prefix:  prefix,
		baseURL: strings.TrimRight(baseURL, "/") + prefix,
		theme:   alpineTheme(prefix),
	}

	// Long-scroll homepage at the prefix root.
	mux.Handle("GET "+prefix+"/{$}", page(st))
	// Bare prefix (no trailing slash) → redirect to the canonical slash form.
	mux.Handle("GET "+prefix, http.RedirectHandler(prefix+"/", http.StatusMovedPermanently))

	// Intake form: shared handler, this demo's markup + log file.
	mux.Handle("POST "+prefix+"/intake", intake.Handler(renderer(st), intakeLogFile))

	// Per-demo static assets (CSS, fonts, images) under a matching prefix.
	fs := http.FileServer(http.Dir("demos/chiropractor/static"))
	mux.Handle("GET "+prefix+"/static/", http.StripPrefix(prefix+"/static/", fs))
}

// page renders the full homepage with no intake state.
func page(st pageState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, Page(site(), st))
	}
}

// renderer wires this demo's templ markup into the shared intake handler so the
// validation/logging/spam logic stays in internal/intake.
func renderer(st pageState) intake.Renderer {
	c := site()
	return intake.Renderer{
		Form: func(s intake.Submission, errs intake.Errors) templ.Component {
			return IntakeCard(c, st.prefix, s, errs)
		},
		Success: func(firstName string) templ.Component {
			return IntakeSuccess(firstName)
		},
		Page: func(s intake.Submission, errs intake.Errors, success bool) templ.Component {
			ps := st
			ps.intake = &intakeState{sub: s, errs: errs, success: success}
			return Page(c, ps)
		},
	}
}

func render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.Render(r.Context(), w); err != nil {
		slog.Error("chiropractor render error", "err", err)
	}
}

// alpineTheme is this vertical's visual shell: its compiled stylesheet and the
// two above-the-fold fonts worth preloading.
func alpineTheme(prefix string) theme.Theme {
	return theme.Theme{
		StylesheetHref: prefix + "/static/css/site.css",
		FontPreloads: []theme.FontPreload{
			{Href: prefix + "/static/fonts/dmsans-latin.woff2"},
			{Href: prefix + "/static/fonts/newsreader-normal-latin.woff2"},
		},
	}
}
