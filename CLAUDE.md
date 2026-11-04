# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

FlintCraft Studio website — a rebrand of the existing Firefly Software site (`../fs-website`). Built with Go stdlib `net/http`, templ (type-safe HTML templating), Tailwind CSS (standalone CLI), HTMX, and Alpine.js.

One Go binary serves two things:
1. **The marketing site** (`internal/view`, `internal/handler`) — pricing, portfolio, process, about, contact, plus keyword-targeted industry landing pages. Reference `../fs-website` for copy and page structure. Full design system is in `./design-system.md`.
2. **Demo sites** (`demos/<vertical>/`) — fictional example sites for Montana professions (chiropractor "Alpine Spine", law "Granite Peak"), used as sales proof and linked from the industry landing pages.

## Commands

```bash
# Development
mage Dev               # templ generate + all CSS (site + every demo), then go run on :8080
                       # No watcher: re-run after editing .templ files or Tailwind classes

# One-time setup
mage InstallTailwind   # Download Tailwind v3.4.17 standalone CLI to ./tailwind/tailwindcss

# Production build
mage Build             # BuildCSS + BuildChiroCSS + BuildLawCSS + BuildGo -> ./bin/server

# Individual build steps
mage BuildCSS          # Marketing site CSS -> web/static/css/site.css
mage BuildChiroCSS     # Demo CSS -> demos/chiropractor/static/css/site.css
mage BuildLawCSS       # Demo CSS -> demos/law/static/css/site.css
mage GenerateTempl     # templ generate
mage BuildGo           # GenerateTempl + go build -> ./bin/server

# Tests (run `templ generate` first — *_templ.go files are gitignored)
go test ./...
go test ./demos/law -run TestPageRenders     # single test

# Linting
go vet ./...
golangci-lint run

# Docker
docker compose up      # builds locally, reads .env, serves on :8080
```

## Architecture

**Routing**: Go 1.22+ pattern-based routing via `http.NewServeMux()` in `cmd/server/main.go`. Routes use method+pattern syntax (e.g., `"GET /{$}"`, `"POST /contact"`). `GET /` at the end is the 404 catch-all.

**Handler pattern**: Handlers in `internal/handler/` return `http.HandlerFunc` closures, keeping dependencies (mailer, config) in closure scope.

**Templating**: `.templ` files compile to gitignored `*_templ.go` via `templ generate`. Marketing pages compose into `Base(title, description, activePath, body)` in `internal/view/layout.templ`. Package-level vars in `view/shared.go` (BaseURL, GtagID, PixelID, TurnstileSiteKey) are set from config in `main.go`.

**Content lives in Go data, not templates**:
- Portfolio: `internal/view/portfolio.go` — `Projects` map + `ProjectOrder`. `/work/{slug}` and the sitemap read from these.
- Industry landing pages: `internal/view/industries.go` — `IndustryLandings` + `IndustryOrder`. Each entry is auto-registered as a flat keyword route (e.g. `/chiropractor-website-design`) and added to the sitemap. A new landing page is a new data entry; no new template needed.

**Static files are served from disk, not embedded**: `http.Dir("web/static")` at `/static/`, and each demo serves `demos/<vertical>/static` under its own prefix. Paths are relative to CWD, so the server must run from the repo root. The Dockerfile copies `web/` and `demos/` into the image for the same reason. (This differs from the usual embed.FS default.)

**Styling**: Tailwind standalone CLI (not npm). Marketing config is `tailwind/tailwind.config.js` with content scoped to `./internal/view/**/*.templ`; FlintCraft palette uses the `fc-*` prefix. Each demo has its own config + input in `tailwind/demos/<vertical>.{config.js,css}` with its own palette and fonts. All compiled `site.css` outputs are gitignored.

**Middleware**: `internal/middleware/logging.go` wraps the mux with structured JSON logging via `slog` and generates request IDs.

**Graceful degradation**: All integrations (Postmark email, Cloudflare Turnstile, Google Analytics, Facebook Pixel) are optional. If env vars are missing, features disable with log warnings rather than failing. `BASE_URL` (default `https://flintcraftstudio.com`) drives canonical/OG URLs, the sitemap, and demo JSON-LD.

### Demo sites

Each vertical is a self-contained package exposing `Register(mux, prefix, baseURL)`, called from `main.go`. Keep shared code shared and per-demo code isolated:
- **Shared** (`internal/`): `intake` (form parse/validate/honeypot/log, htmx partial vs. no-JS full-page fallback), `structdata` (schema.org JSON-LD), `theme` (per-demo `<head>` stylesheet + font preloads), `ui` (skip link, icons).
- **Per-demo** (`demos/<vertical>/`): `content.go` (all copy as typed structs), `page.templ`/`sections.templ`/`intake.templ`, `routes.go`, `static/`, `render_test.go`. The demo injects its markup into the shared intake handler via `intake.Renderer`.
- Demo intake doesn't send email. It appends to a gitignored `demos/<vertical>/intake-submissions.log`. Demos are noindex and excluded from the sitemap.
- Profession-specific structures (e.g. the law demo's results disclaimers, contingency-fee copy, attorney-advertising footer) stay in that demo, not in shared packages.

**Adding a vertical** touches several places. Use `demos/law/` as the worked example, and see the README's "Demo Sites" section and each `demos/<vertical>/README.md`:
1. Copy a demo package and rewrite content/sections.
2. Add `tailwind/demos/<name>.{config.js,css}` and a `Build<Name>CSS` mage target, wired into **both** `Build` and `Dev`.
3. Add a matching `tailwindcss` line to the **Dockerfile** CSS step (it doesn't use mage).
4. Add the `Register` call in `cmd/server/main.go`.
5. Add an `IndustryLandings` entry (and `IndustryOrder` slug) linking to the demo.

## Design System

Full spec in `./design-system.md`. Key points:

**Color tokens** (8 total plus white, closed palette — no gradients, shadows, or additional colors; every text pairing passes WCAG AA):
- `fc-obsidian` (#17181A) — dark bg / primary text on light
- `fc-flint` (#3B3E43) — dark surface (cards, nav) / body text on light
- `fc-slate` (#666B72) — muted text on light
- `fc-slate-light` (#A7ADB5) — muted text on dark (pairs with slate)
- `fc-ash` (#D5D7DA) — body text on dark / borders and dividers on light (0.5px weight)
- `fc-chalk` (#F5F5F3) — light page bg / primary text on dark
- `white` — raised light surfaces (cards, sections) on chalk
- `fc-ember-tint` (#FCEEDF) — soft accent background on light
- `fc-ember` (#EE9440) — accent on dark backgrounds; filled ember takes obsidian text
- `fc-ember-deep` (#A85A14) — accent on light backgrounds; filled ember-deep takes chalk text (never use ember on light)

**Typography**: Cormorant Garamond (`font-display`, headlines) + DM Sans (`font-body`, body/UI). Three DM Sans weights only: 300, 400, 500. Cormorant italic reserved for tagline, hero, pull quotes.

**Brand voice**: "Shaped with intent." (tagline), "Carved for purpose." (hero). Direct, understated. CTA verbs: "Start a project", "See our work".

**Layout**: Light/dark section alternation. Max width 1100px (`max-w-content`). Tight radii (3px interactive, 6px containers). No decorative backgrounds.

**Logos**: `logo.svg`/`logo.png` (clean silhouette, nav/small), `logo_wordmark.svg`/`logo_wordmark.png` (full arrowhead, hero/large). Favicons in `./favicon/` (served copies in `web/static/favicon/`).

The `FlintCraft Studio Design System/` bundle (CSS tokens, HTML previews, JSX UI kit) and `project-highlight/` are design references and are not served. Root-level standalone `.html` files, `before-after/`, and `helena-laundry/` (a separate Go module that `./...` ignores) are scratch/reference material.

## Deployment

Pushing to `main` triggers `.github/workflows/deploy.yml`, which builds the Dockerfile and pushes to GHCR. The SSH deploy step is currently a placeholder (`echo "deploy"`). `deploy/docker-compose.yml` mirrors `/opt/flintcraft/docker-compose.yml` on the VPS (reads `.env`, publishes on `127.0.0.1:8101`). Keep the two in sync.

**templ version must match `go.mod`** (currently `v0.3.1020`). A newer CLI generates code that calls runtime functions the pinned module lacks, and the build fails with `undefined: templruntime.…`. The Dockerfile pins the CLI. Locally, use `go run github.com/a-h/templ/cmd/templ@v0.3.1020 generate` if the installed CLI is newer. When upgrading templ, bump `go.mod` and the Dockerfile together.

## Key Conventions

- No external Go frameworks — stdlib `net/http` only
- No npm/Node.js dependency — Tailwind runs as a standalone binary via Mage
- Mage (Go-based task runner) instead of Make
- Environment config loaded in `internal/config/config.go`; all vars optional with defaults
- `.env` file parsed by hand in `main.go` (no third-party dotenv library); real env vars take precedence
- Module path: `github.com/firefly-software-mt/standard-template`

## Custom Skills

- `/two-variation-site` — Builds a split-panel site presenting two brand directions for client comparison
- `/qc` — Pre-deploy quality control audit against the standard-tier checklist
- `/work-writeup <live-url>` — Creates a `/work` portfolio entry (Project struct + screenshots) from a live client site
