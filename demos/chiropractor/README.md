# Alpine Spine — Chiropractor demo

A single-page (long-scroll) marketing site for a **fictional** chiropractic
practice, **Alpine Spine Chiropractic & Wellness** in Helena, Montana. It is the
first of the FlintCraft demo series (see the repo README's "Demo Sites") and the
template later verticals copy.

> Everything here is fictional — practice name, doctors, reviews, address, phone.
> The Helena/Montana flavor is intentional; no real practice is depicted.

## Route

- Page: **`GET /demos/chiropractor/`** (bare `/demos/chiropractor` 301s to the slash form)
- Intake: **`POST /demos/chiropractor/intake`**
- Static: **`/demos/chiropractor/static/`** (CSS, fonts)

Mounted from `cmd/server/main.go` via `chiropractor.Register(mux, "/demos/chiropractor", cfg.BaseURL)`.

## Files

```
content.go      all copy + typed structs (BusinessInfo, Doctor, Specialty, …)
page.templ      layout shell (own <head>: Newsreader/DM Sans, JSON-LD), nav, footer
sections.templ  hero, trust band, reassurance, what-to-expect, specialties,
                team, treatments, reviews, location + icon/star helpers
intake.templ    the htmx intake form + error/success partials
routes.go       Register() + intake wiring + static handler + theme
helpers.go      small render helpers + intake type aliases
render_test.go  renders the page/partials and asserts key content
static/         fonts/ (self-hosted DM Sans + Newsreader woff2), css/site.css (built)
```

## Shared packages it depends on

- **`internal/intake`** — parses/validates the form, handles spam (honeypot),
  logs submissions, and swaps the right markup (htmx partial vs. full-page
  fallback). This demo injects its own templ via `intake.Renderer`.
- **`internal/structdata`** — renders the `ChiropracticBusiness` JSON-LD
  (address, geo, hours, `aggregateRating`, and `Review` nodes) from the same
  review data that renders the cards.
- **`internal/theme`** — emits the demo's stylesheet `<link>` + font preloads.
- **`internal/ui`** — skip link and the reused phone/arrowhead icons.

## Running

Uses the repo's existing tooling — no new top-level commands:

```bash
mage Dev            # templ generate + build both stylesheets + run server
# then open http://localhost:8080/demos/chiropractor/
```

Build steps individually:

```bash
mage BuildChiroCSS  # compile this demo's stylesheet -> static/css/site.css  (NEW target)
mage GenerateTempl  # regenerate *_templ.go
mage Build          # full production build (now also runs BuildChiroCSS)
```

Tests:

```bash
go test ./demos/chiropractor/ ./internal/intake/
```

## Intake form

- Progressive enhancement: the form is a plain `POST` that validates and
  re-renders server-side **without JS**; htmx enhances it to swap just the card.
- Server-side validation mirrors the client expectations (name, valid email,
  10+ digit phone, reason selected). Errors swap back into the form with values
  preserved; success swaps in a warm confirmation.
- **No real email is sent.** Submissions are appended to
  `demos/chiropractor/intake-submissions.log` (gitignored). The
  `// TODO: wire to Postmark` marker in `internal/intake/intake.go` is where a
  real send (and a per-IP rate limit) would go.

## Photography

The **hero and the three doctor headshots are real Unsplash photos** (webp in
`static/img/`). The hero is eager-loaded with a responsive `srcset`, a matching
`<link rel=preload as=image>` for LCP, and `fetchpriority=high`; headshots are
lazy. Everything else is still a warm CSS treatment — no raster requests:

| Slot | Current | Replace with |
| --- | --- | --- |
| Hero | **real photo** `hero-600/1120.webp` (4:4.3) | client photo, same crop + widths |
| Reassurance | **real photo** `reassure-600/1000.webp` (4:3.5) | client photo, same crop + widths |
| Atmospheric band | `.ph-band` gradient strip | wide Montana-light banner — `loading="lazy"` |
| 3 doctor headshots | **real photos** (see below) | client's own headshots, same 4:3.6 crop |
| Location map | `.loc__map` CSS map | static map image or embedded iframe |

`HeroImage`/`HeroImageAlt` and each doctor's `Photo`/`PhotoAlt` live in
`content.go`; clear the field to fall back to the CSS gradient/monogram. Re-crop
the hero with:

```bash
magick INPUT.jpg -auto-orient -resize 600x645^   -gravity center -extent 600x645   -strip -quality 80 static/img/hero-600.webp
magick INPUT.jpg -auto-orient -resize 1120x1204^ -gravity center -extent 1120x1204 -strip -quality 78 static/img/hero-1120.webp
```

Each doctor in `content.go` has a `Photo` (file under `static/img/`) and
`PhotoAlt`; leaving `Photo` empty falls back to the monogram-on-gradient avatar.
To re-crop a replacement to the card aspect:

```bash
magick INPUT.jpg -auto-orient -resize 640x576^ -gravity center \
  -extent 640x576 -strip -quality 80 static/img/doctor-<name>.webp
```

### Image credits (Unsplash License — free to use)

| File | Maps to | Photographer | Source |
| --- | --- | --- | --- |
| `hero-600/1120.webp` | Hero (alpine ridge) | Adria Sanchez Roque | https://unsplash.com/photos/wbYDAnrVnMg |
| `reassure-600/1000.webp` | Reassurance (front desk) | True Agency | https://unsplash.com/photos/JwP90y9wgr4 |
| `doctor-whitlock.webp` | Dr. Sarah Whitlock | Clay Elliot | https://unsplash.com/photos/mpDV4xaFP8c |
| `doctor-briggs.webp` | Dr. Hannah Briggs | Christina @ wocintechchat.com | https://unsplash.com/photos/SJvDxw0azqw |
| `doctor-reyes.webp` | Dr. Marcus Reyes | Jurica Koletic | https://unsplash.com/photos/7YVZYZeITc8 |

These stand in for the fictional doctors in this demo only; swap them for the
client's real team before any production use.

## Fonts

Self-hosted in `static/fonts/` (no Google Fonts request, no render-blocking):
DM Sans (body/UI) and Newsreader (display/headings), each a variable woff2 with
latin + latin-ext subsets. The two above-the-fold faces are `preload`ed via the
theme.
