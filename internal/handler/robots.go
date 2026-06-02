package handler

import (
	"fmt"
	"net/http"
)

// Robots handles GET /robots.txt. It allows all crawling and points at the
// sitemap. The fictional demo sites under /demos/ set their own noindex meta,
// so they're excluded from search without needing a Disallow here.
func Robots(baseURL string) http.HandlerFunc {
	body := "User-agent: *\nAllow: /\n"
	if baseURL != "" {
		body += fmt.Sprintf("\nSitemap: %s/sitemap.xml\n", baseURL)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(body))
	}
}
