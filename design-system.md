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

All color decisions flow from five named tokens. Do not introduce colors outside this set.

| Token | Hex | Role |
|---|---|---|
| `--fc-obsidian` | `#1a1714` | Primary dark background · primary text on light |
| `--fc-flint` | `#2e2924` | Dark surface (cards, nav bar) |
| `--fc-ember` | `#9e5215` | Primary accent — use on light backgrounds |
| `--fc-spark` | `#bf6e2a` | Accent on dark backgrounds (Obsidian/Flint) |
| `--fc-birchwood` | `#faf9f7` | Primary light background |
| `--fc-parchment` | `#f0ece5` | Light surface (cards, section backgrounds) |
| `--fc-trail-dust` | `#6b5f52` | Secondary / muted text on light |
| `--fc-slate-rim` | `#d4cfc9` | Borders, dividers, rules |

### Usage Rules

- **Ember (`#9e5215`)** is the only accent on white/birchwood/parchment backgrounds — buttons, links, ghost underlines, pull quote text, eyebrow rules.
- **Spark (`#bf6e2a`)** replaces Ember wherever it appears on Obsidian or Flint backgrounds. Never use Ember on dark — it loses contrast.
- **Slate Rim (`#d4cfc9`)** is the only border color. Use at `0.5px` weight universally. Exceptions: left-border accent cards use `2.5px` Ember; featured items use `2px` Ember.
- Do not introduce gradients, shadows, or new colors. The palette is intentionally closed.

### CSS Custom Properties

```css
:root {
  --fc-obsidian:   #1a1714;
  --fc-flint:      #2e2924;
  --fc-ember:      #9e5215;
  --fc-spark:      #bf6e2a;
  --fc-birchwood:  #faf9f7;
  --fc-parchment:  #f0ece5;
  --fc-trail-dust: #6b5f52;
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

| Role | Typeface | Size | Weight | Style | Color | Tracking |
|---|---|---|---|---|---|---|
| Display headline | Cormorant Garamond | 48–64px | 400 | Normal or Italic | `--fc-obsidian` | `-0.01em` |
| Section header | Cormorant Garamond | 32–40px | 400 | Normal | `--fc-obsidian` | `0` |
| Card title | Cormorant Garamond | 20–24px | 600 | Normal | `--fc-obsidian` | `0` |
| Pull quote / accent | Cormorant Garamond | 15–18px | 400 | Italic | `--fc-ember` | `0.02em` |
| Eyebrow label | DM Sans | 10–12px | 300 | Normal | `--fc-trail-dust` | `0.18–0.22em` |
| Body copy | DM Sans | 15–16px | 400 | Normal | `#4a4239` | `0` |
| UI label / button | DM Sans | 11–13px | 500 | Normal | Contextual | `0.04–0.08em` |
| Caption / footnote | DM Sans | 11–12px | 400 | Normal | `--fc-trail-dust` | `0` |

### Typography Rules

- **Eyebrows** are always DM Sans, uppercase, tracked (`letter-spacing: 0.18–0.22em`). Never uppercase without tracking.
- **Cormorant italic** is a high-value register — use it for the tagline, hero headline, and pull quotes. Avoid using it for general emphasis in body copy.
- **Ember (`#9e5215`)** is the only color used for Cormorant italic text (pull quotes, accent lines). On dark backgrounds use Spark (`#bf6e2a`).
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
<nav style="background: #1a1714; padding: 0.75rem 1.5rem; display: flex; justify-content: space-between; align-items: center;">
  <div>
    <div style="font-family: 'DM Sans', sans-serif; font-size: 14px; font-weight: 500; letter-spacing: 0.12em; text-transform: uppercase; color: #f0ece5;">
      FlintCraft Studio
    </div>
    <div style="font-family: 'Cormorant Garamond', Georgia, serif; font-size: 11px; font-style: italic; color: #6b5f52; letter-spacing: 0.04em;">
      Still shaped by hand.
    </div>
  </div>
  <div style="display: flex; gap: 1.5rem; align-items: center;">
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #6b5f52; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Work</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #6b5f52; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Process</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #6b5f52; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">About</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 11px; font-weight: 500; color: #bf6e2a; letter-spacing: 0.08em; text-transform: uppercase; text-decoration: none; border: 0.5px solid #bf6e2a; border-radius: 3px; padding: 5px 14px;">
      Start a project
    </a>
  </div>
</nav>
```

**Nav rules:**
- Background is always Obsidian (`#1a1714`). Never light.
- Nav links are Trail Dust (`#6b5f52`) at rest. On hover: Parchment (`#f0ece5`).
- CTA uses Spark (`#bf6e2a`) — border + text. On hover: fill with Spark, text Birchwood.
- Tagline is Cormorant italic. Never DM Sans.

---

### Buttons

Three variants. Light context and dark context shown separately.

#### On light backgrounds

```css
/* Primary */
.btn-primary {
  background: #9e5215;
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
  color: #1a1714;
  border: 1.5px solid #1a1714;
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
  color: #9e5215;
  border: none;
  border-bottom: 1.5px solid #9e5215;
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
/* Primary — same as light */
.btn-primary { background: #9e5215; color: #faf9f7; /* unchanged */ }

/* Outline — border becomes muted, text becomes light */
.btn-outline-dark {
  background: transparent;
  color: #f0ece5;
  border: 1.5px solid #5a4f43;
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
  color: #bf6e2a;
  border-bottom: 1.5px solid #bf6e2a;
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
  background: #1a1714;
  border: 0.5px solid #2e2924;
  border-radius: 6px;
  padding: 1.25rem;
}
```

#### Ruled / accent card (testimonials, callouts)

```css
.card-ruled {
  background: #faf9f7;
  border-left: 2.5px solid #9e5215;
  border-radius: 0;
  padding: 1rem 1.25rem;
}
```

**Card anatomy:**
- **Eyebrow:** DM Sans, 10px, uppercase, tracked, `#8a8278` (light card) or `#5a4f43` (dark card)
- **Title:** Cormorant Garamond, 20–24px, weight 600 (light) or 400 (dark), color Obsidian (light) or Parchment (dark)
- **Body:** DM Sans, 13px, weight 400, line-height 1.6, Trail Dust (`#6b5f52`) light / `#8a7d6e` dark
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
  color: #6b5f52;
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
  color: #1a1714;
  outline: none;
}

.form-input::placeholder {
  color: #b4ada6;
}

.form-input:focus {
  border-color: #9e5215;
  box-shadow: 0 0 0 2px rgba(158, 82, 21, 0.12);
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
/* On dark surface */
.badge-ember {
  background: #3d2208;
  color: #e8a84a;
  padding: 3px 10px;
  border-radius: 3px;
  font-family: 'DM Sans', sans-serif;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.05em;
}

/* On dark surface — neutral */
.badge-flint {
  background: #2e2924;
  color: #d4cfc9;
  /* same padding/radius/font as above */
}

/* On light surface — neutral */
.badge-birch {
  background: #e8e4db;
  color: #4a4239;
  /* same padding/radius/font as above */
}

/* On light surface — outlined */
.badge-outline {
  background: transparent;
  color: #9e5215;
  border: 0.5px solid #9e5215;
  /* same padding/radius/font as above */
}
```

---

### Section Eyebrow Pattern

Every major section opens with a short rule + uppercase label before the headline.

```html
<div style="margin-bottom: 1rem;">
  <div style="width: 28px; height: 1px; background: #9e5215; margin-bottom: 8px;"></div>
  <div style="font-family: 'DM Sans', sans-serif; font-size: 10px; font-weight: 500; letter-spacing: 0.18em; text-transform: uppercase; color: #6b5f52;">
    Our work
  </div>
</div>
```

On dark backgrounds, the rule stays Ember (`#9e5215`) but the label shifts to `#8a7d6e`.

---

### Dividers

```css
/* Standard rule */
.divider {
  border: none;
  border-top: 0.5px solid #d4cfc9;
  margin: 2.5rem 0;
}

/* Accent rule — use sparingly, once per section max */
.divider-accent {
  border: none;
  border-top: 1px solid #9e5215;
  width: 48px;
  margin: 2rem 0;
}
```

---

## Layout Principles

- **Section alternation:** Light sections (`#faf9f7`) alternate with dark sections (`#1a1714`). Never two dark sections in a row without a transition.
- **Max content width:** `1100px`, centered.
- **Section padding:** `5rem` top/bottom on desktop, `3rem` on mobile.
- **Grid:** 12-column. Most content uses 8 of 12 centered, or a 5/7 split for text/image pairs.
- **No decorative backgrounds** — no gradients, no grain overlays, no mesh textures. Atmosphere comes from typography and negative space.

---

## Interaction States

| Element | Default | Hover | Active/Focus |
|---|---|---|---|
| Nav links | `#6b5f52` | `#f0ece5` | `#f0ece5` |
| Nav CTA | Spark border + text | Spark fill, Birchwood text | Scale `0.98` |
| Primary button | Ember fill | Darken 8% | Scale `0.97` |
| Outline button | Dark border | Dark fill, Birchwood text | Scale `0.97` |
| Ghost button | Ember underline | Ember text brightens | Scale `0.98` |
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
- No additional accent colors beyond the eight tokens above

---

*FlintCraft Studio · Helena, Montana*
*Design system v1.0 — produced by Firefly Software*
