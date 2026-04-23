package handler

import (
	"encoding/xml"
	"log/slog"
	"net/http"

	"github.com/firefly-software-mt/standard-template/internal/view"
)

type sitemapURL struct {
	Loc        string `xml:"loc"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

type urlSet struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

// Sitemap handles GET /sitemap.xml and returns an XML sitemap.
func Sitemap(baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urls := []sitemapURL{
			{Loc: baseURL + "/", ChangeFreq: "monthly", Priority: "1.0"},
			{Loc: baseURL + "/services", ChangeFreq: "monthly", Priority: "0.8"},
			{Loc: baseURL + "/work", ChangeFreq: "monthly", Priority: "0.8"},
			{Loc: baseURL + "/process", ChangeFreq: "monthly", Priority: "0.7"},
			{Loc: baseURL + "/about", ChangeFreq: "monthly", Priority: "0.7"},
			{Loc: baseURL + "/contact", ChangeFreq: "yearly", Priority: "0.6"},
			{Loc: baseURL + "/privacy", ChangeFreq: "yearly", Priority: "0.3"},
			{Loc: baseURL + "/terms", ChangeFreq: "yearly", Priority: "0.3"},
		}

		for _, slug := range view.ProjectOrder {
			urls = append(urls, sitemapURL{
				Loc:        baseURL + "/work/" + slug,
				ChangeFreq: "monthly",
				Priority:   "0.7",
			})
		}

		sitemap := urlSet{
			XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
			URLs:  urls,
		}

		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(xml.Header))
		if err := xml.NewEncoder(w).Encode(sitemap); err != nil {
			slog.Error("sitemap encode error", "err", err)
		}
	}
}
