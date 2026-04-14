package handler

import (
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

// Services handles GET /services and renders the services page.
func Services() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.ServicesPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}
