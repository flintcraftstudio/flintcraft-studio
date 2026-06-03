# Granite Peak Trial Lawyers — Law / personal-injury demo

A single-page (long-scroll) marketing site for a **fictional** personal-injury
trial firm, **Granite Peak Trial Lawyers** in Helena, Montana. It is the
**second** FlintCraft demo vertical (after the chiropractor) and a
copy-reskin-recontent of that template — it reuses the shared `internal/`
plumbing and only adds its own content, palette, and PI-specific sections.

> Everything here is fictional — firm name, attorneys, results, reviews,
> address, phone. **Attorney advertising is regulated:** every results/outcome
> claim is modest and carries the "prior results do not guarantee a similar
> outcome" disclaimer, and the footer carries the full attorney-advertising
> notice. All of it is clearly-marked placeholder text the firm's real counsel
> would review before launch.

## Positioning

Lead with the **advocacy / "we go to trial" story**, not big settlement figures.
The firm is the credible anti-mill: real trial lawyers, personal attention, no
fee unless we win. Results are shown **modestly** (the story is the fight, not
the jackpot) and always under a disclaimer.

## Route

- Page: **`GET /demos/law/`** (bare `/demos/law` 301s to the slash form)
- Intake: **`POST /demos/law/intake`** (free case review)
- Static: **`/demos/law/static/`** (CSS, fonts, images)

Mounted from `cmd/server/main.go` via `law.Register(mux, "/demos/law", cfg.BaseURL)`.

## Files

```
content.go      all copy + typed structs (FirmInfo, Attorney, PracticeArea,
                Step, Result, Review, Stat, MoneyPoint, LocRow) + LegalService JSON-LD
icons.go        Lucide icon inner-SVG map (ISC license), keyed by icon name
page.templ      layout shell (own <head>: Spectral/Libre Franklin, JSON-LD),
                utility strip, sticky nav + mobile drawer, footer (legal notices),
                the icon() + photoPanel() helpers, reveal script
sections.templ  hero + stats, contingency-fee block, practice index, how-it-works
                timeline, attorneys, results (+ disclaimer), testimonials, location
intake.templ    the htmx free-case-review form + field helpers + success partial
routes.go       Register() + intake wiring + static handler + theme
helpers.go      small render helpers + intake type aliases
render_test.go  renders the page/partials and asserts key content + structured data
static/         fonts/ (self-hosted Spectral + Libre Franklin woff2),
                img/ (logos), css/site.css (built, gitignored)
```

## Shared packages it depends on

- **`internal/intake`** — parses/validates the form, handles spam (honeypot),
  logs submissions, and swaps the right markup (htmx partial vs. full-page
  fallback). This demo injects its own templ via `intake.Renderer`. The shared
  `Submission` gained two optional fields here (`When`, `Contact`) for the law
  form's "when did it happen" / "best way to reach you" — generic enough to
  share; the chiropractor leaves them empty.
- **`internal/structdata`** — renders the `LegalService` JSON-LD (address, geo,
  hours, `aggregateRating`, and `Review` nodes) from the same review data that
  renders the cards.
- **`internal/theme`** — emits the demo's stylesheet `<link>` + font preloads.
- **`internal/ui`** — skip link and the FlintCraft demo-framing bar.

PI-specific structures (results cards + disclaimers, contingency-fee messaging,
case-type paths, the attorney-advertising footer) live **only** here in
`demos/law/` — they were not retrofitted onto the chiropractor or forced into a
shared abstraction.

## Running

Uses the repo's existing tooling — no new top-level commands:

```bash
mage Dev            # templ generate + build all stylesheets + run server
# then open http://localhost:8080/demos/law/
```

Build steps individually:

```bash
mage BuildLawCSS    # compile this demo's stylesheet -> static/css/site.css  (NEW target)
mage GenerateTempl  # regenerate *_templ.go
mage Build          # full production build (now also runs BuildLawCSS)
```

Tests:

```bash
go test ./demos/law/ ./internal/intake/
```

## Intake form (free case review)

- Progressive enhancement: the form is a plain `POST` that validates and
  re-renders server-side **without JS**; htmx enhances it to swap just the card.
- Server-side validation (shared `internal/intake`): required name, valid phone
  (10+ digits), valid-looking email, and "what happened" selected. Errors swap
  back into the form with values preserved; success swaps in a warm confirmation
  that sets a callback expectation and offers the 24/7 phone.
- Honeypot (`company`) field drops bots silently. A per-IP rate limit belongs in
  `internal/intake` before any real launch (see the `// TODO: wire to Postmark`).
- **No real email is sent.** Submissions are appended to
  `demos/law/intake-submissions.log` (gitignored).
- Confidentiality microcopy near submit: *Free, confidential, no obligation* —
  plus the *does not create an attorney-client relationship* / attorney-advertising
  notice.
- On success, a confetti burst + a FlintCraft "nudge" toast appear (demo chrome,
  same as the chiropractor — keyed to the `.form-success` partial, brass/evergreen
  palette, skipped under `prefers-reduced-motion`). The toast and the top demo bar
  link to `/contact?ref=law-demo`; the contact form reads that `ref`, shows a
  "You're coming from the Granite Peak Trial Lawyers demo" banner, and carries the
  tag through to the lead email so demo-driven leads are attributed. The
  `/law-firm-website-design` landing CTAs use `ref=law-landing`. (Friendly labels
  live in `internal/view/shared.go`'s `ReferralLabel`.)

## Compliance / legal placeholders

This is a demo. Treat every claim as placeholder a real firm's counsel must
review:

- Results are modest and rounded, framed around the fight. Each results display
  carries: *"Prior results do not guarantee a similar outcome. Every case is
  different…"* (visible in the section and in the footer).
- The footer carries the attorney-advertising notice, the "informational
  purposes only / not legal advice" line, the "no attorney-client relationship"
  line, and the Montana-licensure jurisdiction note.

## Photography

The **hero** and the **three attorney headshots** are real photos (webp in
`static/img/`). The hero is eager-loaded with a responsive `srcset`, a matching
`<link rel="preload" as="image">` for LCP, and `fetchpriority="high"`; headshots
are lazy. The **location map** is a pure-CSS stylized map (warm stone field,
brass highway, sage river/park, a brass pin with the mountain mark) — the same
technique as the chiropractor demo, reskinned to the Granite Peak palette, so it
adds zero raster weight.

| Slot | Current | Replace with |
| --- | --- | --- |
| Hero media | **real photo** `hero-600/1120.webp` (3:2, `object-fit: cover`) | the firm's own warm office / consultation shot or a Helena locale, same widths; **no** crash scene |
| 3 attorney headshots | **real photos** `attorney-*.webp` (4:3) | the firm's own headshots, same crop — consistent, warm, **not** arms-crossed power poses; `loading="lazy"` |
| Location map | `.loc-map` CSS map | a static map of 318 Fuller Ave, Helena, or a warm office exterior; `loading="lazy"` |

Re-crop a replacement hero (centered 3:2) with:

```bash
magick INPUT.jpg -gravity center -crop 3:2 +repage -resize 1120x -strip -quality 80 static/img/hero-1120.webp
magick INPUT.jpg -gravity center -crop 3:2 +repage -resize 600x  -strip -quality 80 static/img/hero-600.webp
```

Clear `HeroImage` in `content.go` to fall back to the CSS placeholder panel.

Each attorney in `content.go` has a `Photo` (file under `static/img/`) and
`PhotoAlt`; clearing `Photo` falls back to the placeholder panel. Re-crop a
replacement to the card aspect (4:3) with:

```bash
magick INPUT.jpg -auto-orient -resize 640x576^ -gravity center \
  -extent 640x576 -strip -quality 80 static/img/attorney-<name>.webp
```

### Image credits

- `logo-mark.png` / `logo-mark-ondark.png` — the Granite Peak geometric mountain
  mark (light and on-dark variants), provided for this demo. The brass peak is
  the single accent; the dark mountain carries a faint evergreen tint.
- `attorney-whitcomb.webp`, `attorney-reyes.webp`, `attorney-holloway.webp` —
  shared from the chiropractor demo (real Unsplash-licensed portraits, see
  `demos/chiropractor/README.md` for photographer credits). They stand in for
  the fictional attorneys in this demo only; swap them for the firm's real team
  before any production use.
- `hero-600/1120.webp` — Adobe Stock (licensed), a warm book-lined office
  consultation/handshake. A demo placeholder for the fictional firm; replace
  with the real firm's own photography before production. `demo-preview.webp`
  (the screenshot shown on the `/law-firm-website-design` landing page) is
  regenerated from the live hero.

## Fonts

Self-hosted in `static/fonts/` (no Google Fonts request, no render-blocking):
**Spectral** (display/headings — refined transitional serif, law-review gravitas)
and **Libre Franklin** (body/UI — grounded institutional sans). woff2, latin +
latin-ext subsets, the weights the design uses (Spectral 400/500/600 + italic;
Libre Franklin 400/500/600). The two above-the-fold faces are `preload`ed via
the theme.
