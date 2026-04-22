# FlintCraft Studio — UI Design System
**Version 1.0 · For use by coding agents and developers**

---

## Brand Identity

**Studio name:** FlintCraft Studio
**Primary tagline (nav/logo):** *Still shaped by hand.*
**Hero headline:** *Carved for purpose.*
**Voice lines (supporting copy):**
- *Websites worth keeping.*
- *We handle it.*
- *Sharp work. No shortcuts.*
- *For businesses that have arrived.*

**Positioning:** Custom web development for established businesses (years 3–5) that have outgrown their first site. Montana-rooted, boutique-quality, coastal-competitive.

---

## Color Tokens

All color decisions flow from twelve named tokens. Do not introduce colors outside this set.

| Token | Hex | Role |
|---|---|---|
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

**Text on dark backgrounds (Obsidian, Flint, Charcoal):**
- Primary text: Birchwood
- Body-weight muted text: Ash
- Subtle text (captions, eyebrows, metadata): Smoke
- Accent on Obsidian: Spark at any size
- Accent on Flint: Spark at 18px+ / 14px bold only (≥3:1 AA Large). For smaller accents on Flint surfaces, use Ash for emphasis — do not use Spark for body-weight accents on Flint.
- Accent on Charcoal: Spark at any size (treat as Obsidian for accent rules)

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

---

## Spacing & Radius Tokens

```css
:root {
  --fc-space-xs:  4px;
  --fc-space-sm:  8px;
  --fc-space-md:  16px;
  --fc-space-lg:  32px;
  --fc-space-xl:  64px;

  --fc-radius-none: 0px;
  --fc-radius-sm:   3px;   /* buttons, inputs, badges */
  --fc-radius-md:   6px;   /* cards */
  --fc-radius-pill: 20px;  /* use sparingly — status badges only */
}
```

Prefer tight radii. The brand is precise and crafted, not soft and friendly. `r-3` for interactive elements, `r-6` for containers. Pill only for status indicators.

---

## Components

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

---

### Buttons

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

---

### Cards

Three card types.

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

---

### Form Elements

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

---

## What This System Is Not

- No gradients or shadows (except functional focus rings)
- No warm amber/golden tones — Ember and Spark are burnt, not bright
- No full-bleed photography treatments or overlaid text on images
- No rounded corners beyond `r-6`
- No system fonts — always load Cormorant Garamond + DM Sans
- No additional accent colors beyond the twelve tokens above

---

*FlintCraft Studio · Helena, Montana*
*Design system v1.0 — produced by Firefly Software*
