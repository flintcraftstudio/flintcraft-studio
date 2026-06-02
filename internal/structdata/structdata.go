// Package structdata renders schema.org JSON-LD for the demo sites.
//
// It is shared across verticals: each demo builds a LocalBusiness value from
// its own content structs and renders it in the page <head>. Emitting the
// reviews and rating from the same data that renders the on-page cards is a
// deliberate selling point — the structured data lives in the site's own
// markup, not in a third-party plugin.
package structdata

import (
	"context"
	"encoding/json"
	"io"

	"github.com/a-h/templ"
)

// PostalAddress is a schema.org PostalAddress node.
type PostalAddress struct {
	Type            string `json:"@type"` // "PostalAddress"
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
	PostalCode      string `json:"postalCode"`
	AddressCountry  string `json:"addressCountry"`
}

// GeoCoordinates is a schema.org GeoCoordinates node.
type GeoCoordinates struct {
	Type      string  `json:"@type"` // "GeoCoordinates"
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// OpeningHours is a schema.org OpeningHoursSpecification node.
// Days are full schema.org URLs or short day names (e.g. "Monday").
type OpeningHours struct {
	Type      string   `json:"@type"` // "OpeningHoursSpecification"
	DayOfWeek []string `json:"dayOfWeek"`
	Opens     string   `json:"opens"`  // "07:30"
	Closes    string   `json:"closes"` // "18:00"
}

// Rating is a schema.org Rating node used inside a Review.
type Rating struct {
	Type        string `json:"@type"` // "Rating"
	RatingValue string `json:"ratingValue"`
	BestRating  string `json:"bestRating,omitempty"`
}

// AggregateRating is a schema.org AggregateRating node.
type AggregateRating struct {
	Type        string `json:"@type"` // "AggregateRating"
	RatingValue string `json:"ratingValue"`
	ReviewCount string `json:"reviewCount"`
	BestRating  string `json:"bestRating,omitempty"`
}

// Author is a schema.org Person node (review author).
type Author struct {
	Type string `json:"@type"` // "Person"
	Name string `json:"name"`
}

// Review is a schema.org Review node.
type Review struct {
	Type          string `json:"@type"` // "Review"
	Author        Author `json:"author"`
	ReviewRating  Rating `json:"reviewRating"`
	ReviewBody    string `json:"reviewBody"`
	DatePublished string `json:"datePublished,omitempty"`
}

// LocalBusiness is a schema.org LocalBusiness (or a more specific subtype such
// as MedicalClinic / ChiropracticBusiness, set via Type). It carries the full
// set of fields the demos emit; omitempty keeps unused fields out of the JSON.
type LocalBusiness struct {
	Context         string           `json:"@context"` // "https://schema.org"
	Type            string           `json:"@type"`    // e.g. "ChiropracticBusiness"
	Name            string           `json:"name"`
	Description     string           `json:"description,omitempty"`
	URL             string           `json:"url,omitempty"`
	Telephone       string           `json:"telephone,omitempty"`
	Email           string           `json:"email,omitempty"`
	PriceRange      string           `json:"priceRange,omitempty"`
	Image           string           `json:"image,omitempty"`
	Address         *PostalAddress   `json:"address,omitempty"`
	Geo             *GeoCoordinates  `json:"geo,omitempty"`
	OpeningHours    []OpeningHours   `json:"openingHoursSpecification,omitempty"`
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty"`
	Review          []Review         `json:"review,omitempty"`
}

// Provider is the organization offering a Service.
type Provider struct {
	Type      string `json:"@type"` // "Organization" / "LocalBusiness"
	Name      string `json:"name"`
	URL       string `json:"url,omitempty"`
	Telephone string `json:"telephone,omitempty"`
}

// Service is a schema.org Service node — used by the industry landing pages to
// describe an offering (e.g. "Chiropractor Website Design") and its area served.
type Service struct {
	Context     string    `json:"@context"` // "https://schema.org"
	Type        string    `json:"@type"`    // "Service"
	ServiceType string    `json:"serviceType"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	AreaServed  string    `json:"areaServed,omitempty"`
	Provider    *Provider `json:"provider,omitempty"`
}

// Answer is the accepted answer to a FAQ Question.
type Answer struct {
	Type string `json:"@type"` // "Answer"
	Text string `json:"text"`
}

// Question is one FAQ entry.
type Question struct {
	Type           string `json:"@type"` // "Question"
	Name           string `json:"name"`
	AcceptedAnswer Answer `json:"acceptedAnswer"`
}

// FAQPage is a schema.org FAQPage node built from a page's Q&A list.
type FAQPage struct {
	Context    string     `json:"@context"` // "https://schema.org"
	Type       string     `json:"@type"`    // "FAQPage"
	MainEntity []Question `json:"mainEntity"`
}

// Script returns a templ.Component that renders v as a
// <script type="application/ld+json"> tag in the page head.
//
// encoding/json escapes "<", ">" and "&" to their \u00xx forms by default, so
// a "</script>" sequence in any string value cannot break out of the tag.
func Script(v any) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if _, err := io.WriteString(w, `<script type="application/ld+json">`); err != nil {
			return err
		}
		if _, err := w.Write(data); err != nil {
			return err
		}
		_, err = io.WriteString(w, `</script>`)
		return err
	})
}
