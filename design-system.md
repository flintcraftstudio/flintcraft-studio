# FlintCraft Studio — UI Design System
**Version 1.1 · For use by coding agents and developers**

> **Changes from v1.0:** Twelve-token palette (was nine). Body and muted text are now split into two tiers per surface (Trail Dust + Bracken on light; Ash + Smoke on dark). Charcoal added as the highest dark surface. Ember and Spark recalibrated for AA contrast. Buttons and cards now carry **flint geometry** — chipped corners + a struck-edge accent rule.

---

## Brand Identity

**Studio name:** FlintCraft Studio
**Primary tagline (nav/logo):** *Still shaped by hand.*
**Hero headline:** *Carved for purpose.*
**Positioning line:** *As durable as the mountains, and just as unlikely to move to San Francisco.*

**Positioning:** Custom web development for established businesses (years 3–5) that have outgrown their first site. Montana-rooted, boutique-quality, coastal-competitive. Fixed-rate retainers ($150–$350/mo, all in). One named engineer per project, same-day support.

---

## Color Tokens — v1.1

Twelve named tokens, split across three tiers per surface. Do not introduce colors outside this set.

| Tier | Light surface | Dark surface |
|---|---|---|
| Page | Birchwood `#faf9f7` | Obsidian `#1c1e24` |
| Raised (cards, nav) | Parchment `#f0ece5` | Flint `#363b42` |
| Highest (modals, popovers, hover) | *(border on parchment)* | Charcoal `#474d55` |
| Primary text | Obsidian | Birchwood |
| Body / strong muted | Trail Dust `#4a4038` | Ash `#c4bdb2` |
| Subtle muted (captions, eyebrows) | Bracken `#6b5f52` | Smoke `#96897a` |
| Accent | Ember `#a84812` | Spark `#e8873b` |
| Border | Slate Rim `#d4cfc9` | Slate Rim `#d4cfc9` |

### Usage Rules

- **Ember (`#a84812`)** is the only accent on light surfaces — buttons, links, ghost underlines, pull quotes, eyebrow rules.
- **Spark (`#e8873b`)** replaces Ember wherever it appears on Obsidian or Flint. Never use Ember on dark — it loses contrast. On Flint specifically, Spark is only safe at 18px+ or 14px bold; smaller accent text on Flint should use Ash for emphasis instead.
- **Trail Dust** is body / strong muted text on light. **Bracken** is the subtle muted tier (captions, eyebrows, metadata).
- **Ash** is body / strong muted text on dark. **Smoke** is the subtle muted tier on dark.
- **Never:** Trail Dust or Bracken on dark. Ash or Smoke on light. Spark for body-weight text on Flint.
- **Slate Rim** is the only border color, used at `0.5px`. Exceptions: accent cards use `2.5px` left Ember/Spark; featured items use `2px` Ember/Spark.

### CSS Custom Properties

```css
:root {
  /* Dark surfaces */
  --fc-obsidian:   #1c1e24;
  --fc-flint:      #363b42;
  --fc-charcoal:   #474d55;

  /* Light surfaces */
  --fc-birchwood:  #faf9f7;
  --fc-parchment:  #f0ece5;

  /* Light text */
  --fc-trail-dust: #4a4038;  /* body */
  --fc-bracken:    #6b5f52;  /* subtle muted */
  --fc-ember:      #a84812;  /* accent */

  /* Dark text */
  --fc-ash:        #c4bdb2;  /* body */
  --fc-smoke:      #96897a;  /* subtle muted */
  --fc-spark:      #e8873b;  /* accent */

  /* Structure */
  --fc-slate-rim:  #d4cfc9;
}
```

---

## Typography

Two typefaces only. No substitutions.

| Typeface | Source | Role |
|---|---|---|
| Cormorant Garamond | Google Fonts | Display, headlines, pull quotes |
| DM Sans | Google Fonts | All body copy, UI labels, eyebrows, buttons |

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Cormorant+Garamond:ital,wght@0,400;0,600;1,400;1,600&family=DM+Sans:opsz,wght@9..40,300;9..40,400;9..40,500&display=swap" rel="stylesheet">
```

### Type Scale

Fluid type via `clamp()` between 320px and 1280px viewports.

| Role | Typeface | Size (min → max) | Weight | Style | Color |
|---|---|---|---|---|---|
| Display headline | Cormorant Garamond | 48–72px | 400 | Normal or Italic | Obsidian / Birchwood |
| Hero headline (Strike) | Cormorant Garamond | 64–128px | 400 | Italic on accent word | Birchwood |
| Section header (h2) | Cormorant Garamond | 28–40px | 400 | Normal | Obsidian / Birchwood |
| Subsection (h3) | Cormorant Garamond | 22–30px | 400 | Normal | Obsidian / Birchwood |
| Card title (h4) | Cormorant Garamond | 18–24px | 600 | Normal | Obsidian / Birchwood |
| Pull quote | Cormorant Garamond | 18–22px | 400 | Italic | Ember / Spark |
| Body lg | DM Sans | 19–22px | 300 | Normal | Trail Dust / Ash |
| Body | DM Sans | 17–19px | 300 | Normal | Trail Dust / Ash |
| Body sm | DM Sans | 15–17px | 400 | Normal | Trail Dust / Ash |
| UI label / button | DM Sans | 13–15px | 500 | Uppercase | Contextual |
| Eyebrow | DM Sans | 11–13px | 500 | Uppercase, `0.18–0.22em` | Bracken / Smoke |

### Typography Rules

- **Eyebrows** are always DM Sans, uppercase, tracked (`0.18–0.22em`). Never uppercase without tracking.
- **Cormorant italic is a register, not emphasis.** Reserved for tagline, hero headline, pull quotes, one-off emphatic phrases. Body paragraphs do not italicize for emphasis — they rephrase.
- **Three DM Sans weights only:** 300, 400, 500. Never 600 or 700.
- Line height: `1.7` for body, `1.05–1.2` for display.

---

## Flint Geometry

The signature differentiator. Flint and obsidian are knapped — struck with a hammerstone, producing sharp, irregular, **angular edges with conchoidal fractures**. The system expresses this through asymmetric clipped corners and a 2.5px accent rule running the chipped diagonal.

### Chip primitives (CSS classes, defined in `tailwind/input.css`)

| Class | Use | Chip size |
|---|---|---|
| `.fc-chip-btn` | Buttons, badges | 8px top-right |
| `.fc-chip-tr` | Standard cards | 14px top-right |
| `.fc-chip-tr-bl` | Featured cards (twin chip) | 18px top-right + bottom-left |
| `.fc-chip-tr-lg` | Editorial image plates | 28px top-right |
| `.fc-chip-plate-sm` | Smaller editorial plates | 28px top-right |
| `.fc-chip-plate` | Hero plate (arrowhead silhouette) | 96px tr + 64px bl |

### Strike rule (the "struck edge")

A 2.5px accent line drawn along the chipped diagonal — the move that reads as knapped stone, not a rounded-corner variant. Position-absolute inside a `.fc-chip-*` parent.

| Class | Width | Position |
|---|---|---|
| `.fc-strike .fc-strike--card` | 24px | Cards (14px chip) |
| `.fc-strike .fc-strike--feature` | 44px | Featured cards (18px chip) |
| `.fc-strike .fc-strike--lg` | 60px | Editorial plates (28px chip) |
| `.fc-strike .fc-strike--hero` | 136px | Hero plate (96px chip) |
| `.fc-strike--dark` modifier | — | Switch fill from Ember to Spark |

### Hover behavior

Buttons: `filter: brightness(1.08)` + `translateY(-1px)`. Active: `scale(0.97)`. The chip shape is static — no animation on hover (the deepening-chip idea was tested and felt over-the-top).

---

## Components

### Navigation

Dark bar (Obsidian). Brand silhouette + wordmark + italic tagline left. Nav links + Spark "Start a project" CTA right. The CTA carries `.fc-chip-btn` for the chipped corner.

- Nav background: always Obsidian
- Nav links at rest: Smoke (`#96897a`); active/hover: Birchwood
- CTA: Spark border + text; on hover, Spark fill + Birchwood text

### Buttons

All chipped buttons carry `.fc-chip-btn` for the 8px top-right notch.

- **Primary on light:** Ember fill, Birchwood text. Hover: darken 8%, lift 1px.
- **Primary on dark:** Spark fill, Obsidian text. Same hover.
- **Outline on light:** transparent, Obsidian text, 1.5px Obsidian border (use `box-shadow: inset 0 0 0 1.5px` so the chip clip-path doesn't truncate the border).
- **Outline on dark:** transparent, Birchwood text, 1.5px Charcoal border.
- **Ghost (link):** No chip. Ember/Spark text + 1.5px underline + `→` suffix.

### Cards

Three card types. All non-ghost cards carry a chip + a strike rule.

- **Light card:** Birchwood or Parchment background, `.fc-chip-tr`, 0.5px Slate Rim border, `.fc-strike .fc-strike--card`.
- **Dark card:** Flint background, `.fc-chip-tr`, 0.5px charcoal border, `.fc-strike .fc-strike--card.fc-strike--dark`.
- **Featured card:** twin-chip `.fc-chip-tr-bl`, `.fc-strike--feature`. Reserved for hero / lead position.
- **Accent card** (testimonial / callout): Parchment + 2.5px left Ember border, radius `0`. **No chip on accent cards** — the bar carries the "struck" message instead.

### Eyebrow pattern

Every major section opens with: 28×1px Ember rule + 8px gap + uppercase tracked label + headline. Label color: Bracken on light, Smoke on dark.

---

## Spacing & Radii

```css
:root {
  --fc-space-xs:  4px;
  --fc-space-sm:  8px;
  --fc-space-md:  16px;
  --fc-space-lg:  32px;
  --fc-space-xl:  64px;

  --fc-radius-sm:   3px;   /* (legacy) buttons before chip primitives — most CTAs now use .fc-chip-btn */
  --fc-radius-md:   6px;   /* cards (when not chipped, e.g. accent cards) */
  --fc-radius-pill: 20px;  /* status indicators only */
}
```

Tight radii. The brand is precise, not soft.

---

## Backgrounds, Shadows, Imagery

- **No gradients** beyond the hero plate's dark face (`linear-gradient(150deg, #323742, #1b1d24)`) and the home hero scrim.
- **No drop shadows.** Depth comes from border weight and surface alternation.
- **No grain, noise, mesh overlays, or repeating patterns.**
- **Photography:** warm, natural, un-retouched. Sits inside a 6px (or chipped) frame, not as a backdrop. The home hero is the rare full-bleed exception.

---

## Motion

- Entrances: fade + 12–18px translate, 700ms `cubic-bezier(0.16, 1, 0.3, 1)`.
- Micro-interactions: 200–500ms `cubic-bezier(0.25, 1, 0.5, 1)`.
- **Hero strike animation:** the 2.5px Spark rule draws along the chipped diagonal on load (900ms, 280ms delay, with a soft glow). Headline reveals word-by-word with a translate-up.
- `prefers-reduced-motion: reduce` turns off all transitions.

---

## Editorial Work Page

The Work screen is structured as a **printed journal** of field reports, not a portfolio grid:

1. **Masthead** — Vol. 7 / Spring 2026 / Helena, MT marginalia + giant italic *"Field Reports."*
2. **Ledger** — table of contents with project numerals and page numbers.
3. **Field reports** — alternating Birchwood ↔ Obsidian surfaces. Each report has a 240px marginalia column (№ numeral, Client, Location, Scope, Shipped) beside an editorial column with image plate + run-in `The brief —` / `What we did —` / `The outcome —` paragraphs. Body copy on dark is Birchwood (not Ash) for primary readability.
4. **Pull quote** — 2.5px Ember rule, Cormorant italic.
5. **Results strip** — display-style numerals with Ember/Spark unit suffix.
6. **Colophon** — italic closing aphorism.

---

## Hero Concepts

The home hero is **"The Strike"**: full-bleed Obsidian, twin-chipped plate (96px tr / 64px bl), the 2.5px Spark strike rule drawing along the top-right diagonal on load (with a soft ember glow). Headline reveals word-by-word. Right column carries a museum-style "Plate 01 — the mark" specimen card and a four-row stats ledger.

Two additional concepts live as standalone explorations and can be promoted to other pages if useful:

- **Tectonic** — typographic. One sentence broken across three staggered lines, no imagery.
- **Specimen Case** — the arrowhead PNG presented as a museum plate with hairline callout rules pointing to brand values.

---

## What This System Is Not

- No gradients or shadows (except functional focus rings + the hero plate face)
- No warm amber/golden tones — Ember and Spark are burnt, not bright
- No additional accent colors beyond the twelve tokens above
- No icon font / icon library by default — `→` Unicode arrow + the two inline SVGs in `nav.templ` (hamburger, close) are the entire icon set. Lucide is the substitution if a set is needed.
- No emoji, anywhere

---

*FlintCraft Studio · Helena, Montana · Design system v1.1*
