package handler

import (
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

// Terms handles GET /terms and renders the terms of use page.
func Terms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.TermsPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}
