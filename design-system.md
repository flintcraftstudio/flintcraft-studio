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

<<<<<<< HEAD
All color decisions flow from twelve named tokens. Do not introduce colors outside this set.
=======
Twelve named tokens, split across three tiers per surface. Do not introduce colors outside this set.
>>>>>>> origin/claude/import-design-docs-QwE1C

| Tier | Light surface | Dark surface |
|---|---|---|
<<<<<<< HEAD
| `--fc-obsidian` | `#1c1e24` | Primary dark background · primary text on light |
| `--fc-flint` | `#363b42` | Raised dark surface (cards, nav bar) |
| `--fc-charcoal` | `#474d55` | Highest dark surface (modals, popovers, hover states) |
| `--fc-ash` | `#c4bdb2` | Strong muted / body text on dark backgrounds |
| `--fc-smoke` | `#96897a` | Subtle muted text on dark — captions, eyebrows, metadata |
| `--fc-spark` | `#e8873b` | Accent on dark (Obsidian only for body text; Flint only at large sizes) |
| `--fc-birchwood` | `#faf9f7` | Primary light background |
| `--fc-parchment` | `#f0ece5` | Raised light surface (cards, section backgrounds) |
| `--fc-trail-dust` | `#4a4038` | Strong muted / body text on light backgrounds |
| `--fc-bracken` | `#6b5f52` | Subtle muted text on light — captions, eyebrows, metadata |
| `--fc-ember` | `#a84812` | Primary accent — use on light backgrounds |
| `--fc-slate-rim` | `#d4cfc9` | Borders, dividers, rules |

### Usage Rules

**Text on light backgrounds (Birchwood, Parchment):**
- Primary text: Obsidian
- Body-weight muted text (long paragraphs, secondary prose): Trail Dust
- Subtle text (captions, eyebrows, timestamps, metadata): Bracken
- Accent (buttons, links, pull quotes, eyebrow rules): Ember
=======
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
>>>>>>> origin/claude/import-design-docs-QwE1C

**Text on dark backgrounds (Obsidian, Flint, Charcoal):**
- Primary text: Birchwood
- Body-weight muted text: Ash
- Subtle text (captions, eyebrows, metadata): Smoke
- Accent on Obsidian: Spark at any size
- Accent on Flint: Spark at 18px+ / 14px bold only (≥3:1 AA Large). For smaller accents on Flint surfaces, use Ash for emphasis — do not use Spark for body-weight accents on Flint.
- Accent on Charcoal: Spark at any size (treat as Obsidian for accent rules)

<<<<<<< HEAD
**Surface hierarchy (dark mode):**
- Page: Obsidian
- Raised (cards, nav): Flint
- Highest (modals, popovers, hovered nav items, dropdown menus): Charcoal

**Surface hierarchy (light mode):**
- Page: Birchwood
- Raised (cards, section backgrounds): Parchment
- Highest: No separate token — use a Slate Rim border on Parchment, or layer via shadow-equivalent borders

**Never:**
- Ember on dark backgrounds (loses contrast — use Spark)
- Spark on light backgrounds (loses contrast — use Ember)
- Trail Dust or Bracken on dark backgrounds
- Ash or Smoke on light backgrounds
- Spark for body-size text on Flint (fails AA)

**Borders:**
- Slate Rim at `0.5px` weight universally
- Left-border accent cards: `2.5px` Ember (light) or `2.5px` Spark (dark)
- Featured items: `2px` Ember (light) or `2px` Spark (dark)

**No gradients, no shadows, no colors outside this set.** The palette is intentionally closed.
=======
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
>>>>>>> origin/claude/import-design-docs-QwE1C

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

<<<<<<< HEAD
| Role | Typeface | Size | Weight | Style | Color | Tracking |
|---|---|---|---|---|---|---|
| Display headline | Cormorant Garamond | 48–64px | 400 | Normal or Italic | `--fc-obsidian` | `-0.01em` |
| Section header | Cormorant Garamond | 32–40px | 400 | Normal | `--fc-obsidian` | `0` |
| Card title | Cormorant Garamond | 20–24px | 600 | Normal | `--fc-obsidian` | `0` |
| Pull quote / accent | Cormorant Garamond | 15–18px | 400 | Italic | `--fc-ember` | `0.02em` |
| Eyebrow label | DM Sans | 10–12px | 300 | Normal | `--fc-bracken` (light) / `--fc-smoke` (dark) | `0.18–0.22em` |
| Body copy | DM Sans | 15–16px | 400 | Normal | `--fc-trail-dust` (`#4a4038`) | `0` |
| UI label / button | DM Sans | 11–13px | 500 | Normal | Contextual | `0.04–0.08em` |
| Caption / footnote | DM Sans | 11–12px | 400 | Normal | `--fc-bracken` (light) / `--fc-smoke` (dark) | `0` |

### Typography Rules

- **Eyebrows** are always DM Sans, uppercase, tracked (`letter-spacing: 0.18–0.22em`). Never uppercase without tracking.
- **Cormorant italic** is a high-value register — use it for the tagline, hero headline, and pull quotes. Avoid using it for general emphasis in body copy.
- **Ember (`#a84812`)** is the only color used for Cormorant italic text on light backgrounds (pull quotes, accent lines). On dark backgrounds use Spark (`#e8873b`).
- **Three DM Sans weights only:** 300 (body), 400 (default/caption), 500 (labels/buttons). Never 600 or 700.
- Line height: `1.7` for body copy, `1.1–1.2` for display/headlines.
=======
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
>>>>>>> origin/claude/import-design-docs-QwE1C

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

<<<<<<< HEAD
### Navigation Bar

Dark bar. Brand name + italic tagline left. Nav links + CTA right.

```html
<nav style="background: #1c1e24; padding: 0.75rem 1.5rem; display: flex; justify-content: space-between; align-items: center;">
  <div>
    <div style="font-family: 'DM Sans', sans-serif; font-size: 14px; font-weight: 500; letter-spacing: 0.12em; text-transform: uppercase; color: #f0ece5;">
      FlintCraft Studio
    </div>
    <div style="font-family: 'Cormorant Garamond', Georgia, serif; font-size: 11px; font-style: italic; color: #96897a; letter-spacing: 0.04em;">
      Still shaped by hand.
    </div>
  </div>
  <div style="display: flex; gap: 1.5rem; align-items: center;">
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #96897a; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Work</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #96897a; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Process</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #96897a; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">About</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 11px; font-weight: 500; color: #e8873b; letter-spacing: 0.08em; text-transform: uppercase; text-decoration: none; border: 0.5px solid #e8873b; border-radius: 3px; padding: 5px 14px;">
      Start a project
    </a>
  </div>
</nav>
```

**Nav rules:**
- Background is always Obsidian (`#1c1e24`). Never light.
- Nav tagline and links are Smoke (`#96897a`) at rest. On hover: Parchment (`#f0ece5`).
- CTA uses Spark (`#e8873b`) — border + text. On hover: fill with Spark, text Birchwood.
- Tagline is Cormorant italic. Never DM Sans.
=======
- **No gradients** beyond the hero plate's dark face (`linear-gradient(150deg, #323742, #1b1d24)`) and the home hero scrim.
- **No drop shadows.** Depth comes from border weight and surface alternation.
- **No grain, noise, mesh overlays, or repeating patterns.**
- **Photography:** warm, natural, un-retouched. Sits inside a 6px (or chipped) frame, not as a backdrop. The home hero is the rare full-bleed exception.
>>>>>>> origin/claude/import-design-docs-QwE1C

---

## Motion

<<<<<<< HEAD
Three variants. Light context and dark context shown separately.

#### On light backgrounds

```css
/* Primary */
.btn-primary {
  background: #a84812;
  color: #faf9f7;
  border-radius: 3px;
  padding: 10px 22px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  border: none;
}

/* Outline */
.btn-outline {
  background: transparent;
  color: #1c1e24;
  border: 1.5px solid #1c1e24;
  border-radius: 3px;
  padding: 10px 22px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

/* Ghost */
.btn-ghost {
  background: transparent;
  color: #a84812;
  border: none;
  border-bottom: 1.5px solid #a84812;
  border-radius: 0;
  padding: 10px 0;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}
```

#### On dark backgrounds (Obsidian / Flint)

```css
/* Primary on dark — Spark fill (Ember is disallowed on dark) */
.btn-primary-dark { background: #e8873b; color: #1c1e24; /* other properties same as .btn-primary */ }

/* Outline — border becomes muted (Charcoal), text becomes light */
.btn-outline-dark {
  background: transparent;
  color: #f0ece5;
  border: 1.5px solid #474d55;
  border-radius: 3px;
  padding: 10px 22px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

/* Ghost — Spark replaces Ember */
.btn-ghost-dark {
  color: #e8873b;
  border-bottom: 1.5px solid #e8873b;
  /* all other properties same as .btn-ghost */
}
```
=======
- Entrances: fade + 12–18px translate, 700ms `cubic-bezier(0.16, 1, 0.3, 1)`.
- Micro-interactions: 200–500ms `cubic-bezier(0.25, 1, 0.5, 1)`.
- **Hero strike animation:** the 2.5px Spark rule draws along the chipped diagonal on load (900ms, 280ms delay, with a soft glow). Headline reveals word-by-word with a translate-up.
- `prefers-reduced-motion: reduce` turns off all transitions.
>>>>>>> origin/claude/import-design-docs-QwE1C

---

## Editorial Work Page

The Work screen is structured as a **printed journal** of field reports, not a portfolio grid:

<<<<<<< HEAD
#### Light card (default)

```css
.card-light {
  background: #faf9f7;
  border: 0.5px solid #d4cfc9;
  border-radius: 6px;
  padding: 1.25rem;
}
```

#### Dark card

```css
.card-dark {
  background: #1c1e24;
  border: 0.5px solid #363b42;
  border-radius: 6px;
  padding: 1.25rem;
}
```

#### Ruled / accent card (testimonials, callouts)

```css
.card-ruled {
  background: #faf9f7;
  border-left: 2.5px solid #a84812;
  border-radius: 0;
  padding: 1rem 1.25rem;
}
```

**Card anatomy:**
- **Eyebrow:** DM Sans, 10px, uppercase, tracked, Bracken (`#6b5f52`) on light card / Smoke (`#96897a`) on dark card
- **Title:** Cormorant Garamond, 20–24px, weight 600 (light) or 400 (dark), color Obsidian (light) or Birchwood (dark)
- **Body:** DM Sans, 13px, weight 400, line-height 1.6, Trail Dust (`#4a4038`) on light / Ash (`#c4bdb2`) on dark
- **CTA link:** DM Sans, 11px, uppercase, tracked, Ember (light) or Spark (dark), no underline, arrow suffix ` →`
=======
1. **Masthead** — Vol. 7 / Spring 2026 / Helena, MT marginalia + giant italic *"Field Reports."*
2. **Ledger** — table of contents with project numerals and page numbers.
3. **Field reports** — alternating Birchwood ↔ Obsidian surfaces. Each report has a 240px marginalia column (№ numeral, Client, Location, Scope, Shipped) beside an editorial column with image plate + run-in `The brief —` / `What we did —` / `The outcome —` paragraphs. Body copy on dark is Birchwood (not Ash) for primary readability.
4. **Pull quote** — 2.5px Ember rule, Cormorant italic.
5. **Results strip** — display-style numerals with Ember/Spark unit suffix.
6. **Colophon** — italic closing aphorism.
>>>>>>> origin/claude/import-design-docs-QwE1C

---

## Hero Concepts

<<<<<<< HEAD
```css
.form-label {
  font-family: 'DM Sans', sans-serif;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: #6b5f52; /* Bracken */
  display: block;
  margin-bottom: 6px;
}

.form-input {
  width: 100%;
  padding: 10px 14px;
  background: #faf9f7;
  border: 0.5px solid #d4cfc9;
  border-radius: 3px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  color: #1c1e24;
  outline: none;
}

.form-input::placeholder {
  color: #b4ada6; /* Bracken @ ~50% luminance — derived, not a named token */
}

.form-input:focus {
  border-color: #a84812;
  box-shadow: 0 0 0 2px rgba(168, 72, 18, 0.12);
}

textarea.form-input {
  resize: vertical;
  min-height: 120px;
  line-height: 1.6;
}
```

---

### Badges

```css
/* On dark surface — Spark-tinted (not a named token, derived) */
.badge-ember {
  background: #33180a;
  color: #e8a84a;
  padding: 3px 10px;
  border-radius: 3px;
  font-family: 'DM Sans', sans-serif;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.05em;
}

/* On dark surface — neutral (Flint bg + Slate Rim text) */
.badge-flint {
  background: #363b42;
  color: #d4cfc9;
  /* same padding/radius/font as above */
}

/* On light surface — neutral (Parchment-shaded bg, derived) */
.badge-birch {
  background: #e8e4db;
  color: #4a4038; /* Trail Dust */
  /* same padding/radius/font as above */
}

/* On light surface — outlined (Ember) */
.badge-outline {
  background: transparent;
  color: #a84812;
  border: 0.5px solid #a84812;
  /* same padding/radius/font as above */
}
```

---

### Section Eyebrow Pattern

Every major section opens with a short rule + uppercase label before the headline.

```html
<div style="margin-bottom: 1rem;">
  <div style="width: 28px; height: 1px; background: #a84812; margin-bottom: 8px;"></div>
  <div style="font-family: 'DM Sans', sans-serif; font-size: 10px; font-weight: 500; letter-spacing: 0.18em; text-transform: uppercase; color: #6b5f52;">
    Our work
  </div>
</div>
```

On dark backgrounds, the rule becomes Spark (`#e8873b`) — Ember is reserved for light — and the label shifts to Smoke (`#96897a`).

---

### Dividers

```css
/* Standard rule */
.divider {
  border: none;
  border-top: 0.5px solid #d4cfc9;
  margin: 2.5rem 0;
}

/* Accent rule — use sparingly, once per section max. Use Spark on dark. */
.divider-accent {
  border: none;
  border-top: 1px solid #a84812;
  width: 48px;
  margin: 2rem 0;
}
```

---

## Layout Principles

- **Section alternation:** Light sections (`#faf9f7`) alternate with dark sections (`#1c1e24`). Never two dark sections in a row without a transition.
- **Max content width:** `1100px`, centered.
- **Section padding:** `5rem` top/bottom on desktop, `3rem` on mobile.
- **Grid:** 12-column. Most content uses 8 of 12 centered, or a 5/7 split for text/image pairs.
- **No decorative backgrounds** — no gradients, no grain overlays, no mesh textures. Atmosphere comes from typography and negative space.

---

## Interaction States

| Element | Default | Hover | Active/Focus |
|---|---|---|---|
| Nav links | Smoke `#96897a` | Parchment `#f0ece5` | Parchment `#f0ece5` |
| Nav CTA | Spark border + text | Spark fill, Obsidian text | Scale `0.98` |
| Primary button | Ember fill (light) / Spark fill (dark) | Darken 8% | Scale `0.97` |
| Outline button | Obsidian border (light) / Charcoal border (dark) | Dark fill, Birchwood text | Scale `0.97` |
| Ghost button | Ember underline (light) / Spark underline (dark) | Accent text brightens | Scale `0.98` |
| Text input | Slate Rim border | Slate Rim border | Ember border + soft glow ring |
| Card (linked) | No shadow | Lift: `translateY(-2px)` | `translateY(0)` |

---

## Logo Usage

**Two logo assets provided:**
- `Asset_3.png` — No tang, clean silhouette. **Use for nav/small contexts.**
- `Asset_7.png` — Full arrowhead with tang. **Use for hero and large contexts.**

Both feature the illustrated hand-drawn style with crosshatching. Do not recolor or modify the logo files. The amber/orange facet in the logo is the source of the Ember/Spark palette — do not replace it with a flat color version.

**Logo on dark (Obsidian) backgrounds:** Use as-is — the black outline reads on dark.
**Logo on light backgrounds:** Use as-is — the dark flint reads on Birchwood/Parchment.

---

## Voice Reference (for copy in components)

- **CTA buttons:** Direct action verbs. "Start a project" not "Get started today." "See our work" not "View portfolio."
- **Eyebrow labels:** Factual, understated. "Our work" not "What we've built." "Process" not "How it works."
- **Placeholder text:** Specific and inviting. "yourbusiness.com" not "Enter URL." "What's holding your current site back?" not "Your message."
- **Error states:** Plain and helpful. "Something's missing here." not "This field is required."
=======
The home hero is **"The Strike"**: full-bleed Obsidian, twin-chipped plate (96px tr / 64px bl), the 2.5px Spark strike rule drawing along the top-right diagonal on load (with a soft ember glow). Headline reveals word-by-word. Right column carries a museum-style "Plate 01 — the mark" specimen card and a four-row stats ledger.

Two additional concepts live as standalone explorations and can be promoted to other pages if useful:

- **Tectonic** — typographic. One sentence broken across three staggered lines, no imagery.
- **Specimen Case** — the arrowhead PNG presented as a museum plate with hairline callout rules pointing to brand values.
>>>>>>> origin/claude/import-design-docs-QwE1C

---

## What This System Is Not

- No gradients or shadows (except functional focus rings + the hero plate face)
- No warm amber/golden tones — Ember and Spark are burnt, not bright
<<<<<<< HEAD
- No full-bleed photography treatments or overlaid text on images
- No rounded corners beyond `r-6`
- No system fonts — always load Cormorant Garamond + DM Sans
- No additional accent colors beyond the twelve tokens above
=======
- No additional accent colors beyond the twelve tokens above
- No icon font / icon library by default — `→` Unicode arrow + the two inline SVGs in `nav.templ` (hamburger, close) are the entire icon set. Lucide is the substitution if a set is needed.
- No emoji, anywhere
>>>>>>> origin/claude/import-design-docs-QwE1C

---

*FlintCraft Studio · Helena, Montana · Design system v1.1*
