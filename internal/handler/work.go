package handler

import (
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

// Work handles GET /work and renders the portfolio listing page.
func Work() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.WorkPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}

// WorkProject handles GET /work/{slug} and renders an individual project page.
func WorkProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		project, ok := view.Projects[slug]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			if err := view.NotFoundPage().Render(r.Context(), w); err != nil {
				slog.Error("render error", "err", err)
			}
			return
		}
		if err := view.ProjectPage(project).Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}
