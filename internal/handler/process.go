package handler

import (
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

// Process handles GET /process and renders the process page.
func Process() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.ProcessPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}
