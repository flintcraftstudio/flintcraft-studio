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
