# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

FlintCraft Studio website — a rebrand of the existing Firefly Software site (`../fs-website`). Built with Go stdlib `net/http`, templ (type-safe HTML templating), Tailwind CSS (standalone CLI), HTMX, and Alpine.js.

Same content as the original (pricing, portfolio, team, process, contact). Reference `../fs-website` for copy and page structure. Full design system is in `./design-system.md`.

## Commands

```bash
# Development
mage Dev               # Generate templ, build CSS, start server on :8080

# One-time setup
mage InstallTailwind   # Download platform-specific Tailwind CLI binary

# Production build
mage Build             # BuildCSS + GenerateTempl + BuildGo -> ./bin/server

# Individual build steps
mage BuildCSS          # Compile + minify Tailwind CSS
mage GenerateTempl     # Run templ code generation
mage BuildGo           # GenerateTempl + go build -> ./bin/server

# Linting
go vet ./...
golangci-lint run

# Docker
docker compose up
```

## Architecture

**Routing**: Go 1.22+ pattern-based routing via `http.NewServeMux()` in `cmd/server/main.go`. Routes use method+pattern syntax (e.g., `"GET /{$}"`, `"POST /contact"`).

**Handler pattern**: Handlers in `internal/handler/` return `http.HandlerFunc` closures, keeping dependencies (mailer, config) in closure scope.

**Templating**: `.templ` files in `internal/view/` compile to Go code via `templ generate`. Templates are type-safe and checked at compile time. The base layout is `layout.templ`; pages compose into it. Shared constants (SiteName, tracking IDs) live in `view/shared.go`.

**Styling**: Tailwind CSS uses a standalone CLI binary (not npm). Config is at `tailwind/tailwind.config.js` with content paths scoped to `./internal/view/**/*.templ`. FlintCraft color palette uses `fc-*` prefix. Output goes to `web/static/css/site.css` (gitignored).

## Design System

Full spec in `./design-system.md`. Key points:

**Color tokens** (12 total, closed palette — no gradients, shadows, or additional colors):
- `fc-obsidian` (#1c1e24) — primary dark bg / primary text on light
- `fc-flint` (#363b42) — raised dark surface (cards, nav)
- `fc-charcoal` (#474d55) — highest dark surface (modals, popovers, hover)
- `fc-ash` (#c4bdb2) — strong muted / body-weight text on dark
- `fc-smoke` (#96897a) — subtle muted on dark (captions, eyebrows, nav links)
- `fc-spark` (#e8873b) — accent on dark (never ember on dark)
- `fc-birchwood` (#faf9f7) — primary light bg
- `fc-parchment` (#f0ece5) — raised light surface
- `fc-trail-dust` (#4a4038) — strong muted / body-weight text on light
- `fc-bracken` (#6b5f52) — subtle muted on light (captions, eyebrows, form labels)
- `fc-ember` (#a84812) — accent on light (never spark on light)
- `fc-slate-rim` (#d4cfc9) — borders/dividers (0.5px weight)

Strong-muted vs subtle-muted: use trail-dust/ash for body-weight prose; use bracken/smoke for captions, eyebrows, metadata. Never trail-dust or bracken on dark; never ash or smoke on light.

One derived utility tint (not a palette token): `fc-ember-dark` (#8a3a0e) — primary-button hover only.

**Typography**: Cormorant Garamond (display/headlines) + DM Sans (body/UI). Three DM Sans weights only: 300, 400, 500. Cormorant italic reserved for tagline, hero, pull quotes.

**Brand voice**: "Still shaped by hand." (tagline), "Carved for purpose." (hero). Direct, understated. CTA verbs: "Start a project", "See our work".

**Layout**: Light/dark section alternation. Max width 1100px. Tight radii (3px interactive, 6px containers). No decorative backgrounds.

**Logos**: `logo.svg`/`logo.png` (clean silhouette, nav/small), `logo_wordmark.svg`/`logo_wordmark.png` (full arrowhead, hero/large). Favicons in `./favicon/`.

**Graceful degradation**: All integrations (Postmark email, Cloudflare Turnstile, Google Analytics, Facebook Pixel) are optional. If env vars are missing, features disable with log warnings rather than failing.

**Middleware**: `internal/middleware/logging.go` wraps the mux with structured JSON logging via `slog`, generates request IDs.

**Static files**: Served from `web/static/` at `/static/`. Includes vendored HTMX and Alpine.js.

## Key Conventions

- No external Go frameworks — stdlib `net/http` only
- No npm/Node.js dependency — Tailwind runs as a standalone binary via Mage
- Mage (Go-based task runner) instead of Make
- Environment config loaded in `internal/config/config.go`; all vars optional with defaults
- `.env` file parsed by hand in `main.go` (no third-party dotenv library)
- Module path: `github.com/firefly-software-mt/standard-template`

## Custom Skills

- `/two-variation-site` — Builds a split-panel site presenting two brand directions for client comparison
- `/qc` — Pre-deploy quality control audit against the standard-tier checklist
