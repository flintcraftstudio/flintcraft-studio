package handler

import (
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

// Industries handles GET /industries and renders the verticals hub.
func Industries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.IndustriesHubPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}

// IndustryLanding renders a single industry landing page. The data is bound at
// registration time (one route per vertical), mirroring how the demos register.
func IndustryLanding(d view.IndustryLanding) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.IndustryLandingPage(d).Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}
