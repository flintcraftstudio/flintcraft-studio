# FlintCraft Studio — UI Design System
**Version 1.0 · For use by coding agents and developers**

---

## Brand Identity

**Studio name:** FlintCraft Studio
**Primary tagline (nav/logo):** *Shaped with intent.*
**Hero headline:** *Carved for purpose.*
**Voice lines (supporting copy):**
- *Websites worth keeping.*
- *We handle it.*
- *Sharp work. No shortcuts.*
- *For businesses that have arrived.*

**Positioning:** Custom web development for established businesses (years 3–5) that have outgrown their first site. Montana-rooted, boutique-quality, coastal-competitive.

---

## Color Tokens

All color decisions flow from these eight tokens (Tailwind `fc-*`, defined in `tailwind/tailwind.config.js`), plus plain white for raised light surfaces. Do not introduce colors outside this set.

| Token | Hex | Role |
|---|---|---|
| `--fc-obsidian` | `#17181A` | Primary dark background · primary text on light |
| `--fc-flint` | `#3B3E43` | Dark surface (cards, nav) · body text on light |
| `--fc-slate` | `#666B72` | Muted text on light (captions, eyebrows, metadata, placeholders) |
| `--fc-slate-light` | `#A7ADB5` | Muted text on dark (Obsidian and Flint) |
| `--fc-ash` | `#D5D7DA` | Body text on dark · borders, dividers, rules on light |
| `--fc-chalk` | `#F5F5F3` | Primary light background · primary text on dark |
| `--fc-ember-tint` | `#FCEEDF` | Soft accent background on light (highlights, callouts); Obsidian or Flint text |
| `--fc-ember` | `#EE9440` | Accent on dark backgrounds (Obsidian/Flint) |
| `--fc-ember-deep` | `#A85A14` | Accent on light backgrounds |
| white | `#FFFFFF` | Raised light surface (cards, sections) on a Chalk page |

### Usage Rules

- **Ember Deep (`#A85A14`)** is the only accent on Chalk or white: buttons, links, ghost underlines, pull quote text, eyebrow rules. Buttons are Ember Deep with Chalk text (4.65:1).
- **Ember Tint (`#FCEEDF`)** takes Obsidian or Flint text. Ember Deep on Ember Tint is 4.46:1, so use it there only for large text (24px+) or rules.
- **Ember (`#EE9440`)** replaces Ember Deep wherever it appears on Obsidian or Flint. Never use Ember on light backgrounds (2.2:1 on Chalk). Filled Ember buttons take **Obsidian** text, never Chalk (2.4:1).
- **Slate (`#666B72`)** is muted text on light backgrounds only; on dark use **Slate Light (`#A7ADB5`)**. They always pair this way.
- **Ash (`#D5D7DA`)** is the only border color on light. Use at `0.5px` weight universally. Exceptions: left-border accent cards use `2.5px` Ember Deep; featured items use `2px` Ember Deep.
- Hover on filled buttons: `brightness-90` on Ember Deep, `brightness-110` on Ember. Both keep the text above 4.5:1.
- Do not introduce gradients, shadows, or new colors. The palette is intentionally closed.
- Every text pairing above passes WCAG AA (4.5:1). When adding a combination not listed here, check it.

### CSS Custom Properties

```css
:root {
  --fc-obsidian:    #17181A;
  --fc-flint:       #3B3E43;
  --fc-slate:       #666B72;
  --fc-slate-light: #A7ADB5;
  --fc-ash:         #D5D7DA;
  --fc-chalk:       #F5F5F3;
  --fc-ember-tint:  #FCEEDF;
  --fc-ember:       #EE9440;
  --fc-ember-deep:  #A85A14;
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
| Eyebrow label | DM Sans | 10–12px | 300 | Normal | `--fc-slate` | `0.18–0.22em` |
| Body copy | DM Sans | 15–16px | 400 | Normal | `#4a4239` | `0` |
| UI label / button | DM Sans | 11–13px | 500 | Normal | Contextual | `0.04–0.08em` |
| Caption / footnote | DM Sans | 11–12px | 400 | Normal | `--fc-slate` | `0` |

### Typography Rules

- **Eyebrows** are always DM Sans, uppercase, tracked (`letter-spacing: 0.18–0.22em`). Never uppercase without tracking.
- **Cormorant italic** is a high-value register — use it for the tagline, hero headline, and pull quotes. Avoid using it for general emphasis in body copy.
- **Ember Deep (`#A85A14`)** is the only color used for Cormorant italic text (pull quotes, accent lines). On dark backgrounds use Ember (`#EE9440`).
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
<nav style="background: #17181A; padding: 0.75rem 1.5rem; display: flex; justify-content: space-between; align-items: center;">
  <div>
    <div style="font-family: 'DM Sans', sans-serif; font-size: 14px; font-weight: 500; letter-spacing: 0.12em; text-transform: uppercase; color: #F5F5F3;">
      FlintCraft Studio
    </div>
    <div style="font-family: 'Cormorant Garamond', Georgia, serif; font-size: 11px; font-style: italic; color: #A7ADB5; letter-spacing: 0.04em;">
      Shaped with intent.
    </div>
  </div>
  <div style="display: flex; gap: 1.5rem; align-items: center;">
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #666B72; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Work</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #666B72; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">Process</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 12px; color: #666B72; letter-spacing: 0.06em; text-transform: uppercase; text-decoration: none;">About</a>
    <a style="font-family: 'DM Sans', sans-serif; font-size: 11px; font-weight: 500; color: #EE9440; letter-spacing: 0.08em; text-transform: uppercase; text-decoration: none; border: 0.5px solid #EE9440; border-radius: 3px; padding: 5px 14px;">
      Start a project
    </a>
  </div>
</nav>
```

**Nav rules:**
- Background is always Obsidian (`#17181A`). Never light.
- Nav links are Ash (`#D5D7DA`) at rest. On hover: Chalk (`#F5F5F3`).
- CTA uses Ember (`#EE9440`) — border + text. On hover: fill with Ember, text Chalk.
- Tagline is Cormorant italic. Never DM Sans.

---

### Buttons

Three variants. Light context and dark context shown separately.

#### On light backgrounds

```css
/* Primary */
.btn-primary {
  background: #A85A14;
  color: #F5F5F3;
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
  color: #17181A;
  border: 1.5px solid #17181A;
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
  color: #A85A14;
  border: none;
  border-bottom: 1.5px solid #A85A14;
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
.btn-primary { background: #A85A14; color: #F5F5F3; /* unchanged */ }

/* Outline — border becomes muted, text becomes light */
.btn-outline-dark {
  background: transparent;
  color: #F5F5F3;
  border: 1.5px solid #5a4f43;
  border-radius: 3px;
  padding: 10px 22px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

/* Ghost — Ember replaces Ember Deep */
.btn-ghost-dark {
  color: #EE9440;
  border-bottom: 1.5px solid #EE9440;
  /* all other properties same as .btn-ghost */
}
```

---

### Cards

Three card types.

#### Light card (default)

```css
.card-light {
  background: #F5F5F3;
  border: 0.5px solid #D5D7DA;
  border-radius: 6px;
  padding: 1.25rem;
}
```

#### Dark card

```css
.card-dark {
  background: #17181A;
  border: 0.5px solid #3B3E43;
  border-radius: 6px;
  padding: 1.25rem;
}
```

#### Ruled / accent card (testimonials, callouts)

```css
.card-ruled {
  background: #F5F5F3;
  border-left: 2.5px solid #A85A14;
  border-radius: 0;
  padding: 1rem 1.25rem;
}
```

**Card anatomy:**
- **Eyebrow:** DM Sans, 10px, uppercase, tracked, `#8a8278` (light card) or `#5a4f43` (dark card)
- **Title:** Cormorant Garamond, 20–24px, weight 600 (light) or 400 (dark), color Obsidian (light) or Chalk (dark)
- **Body:** DM Sans, 13px, weight 400, line-height 1.6, Slate (`#666B72`) light / `#A7ADB5` dark
- **CTA link:** DM Sans, 11px, uppercase, tracked, Ember Deep (light) or Ember (dark), no underline, arrow suffix ` →`

---

### Form Elements

```css
.form-label {
  font-family: 'DM Sans', sans-serif;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: #666B72;
  display: block;
  margin-bottom: 6px;
}

.form-input {
  width: 100%;
  padding: 10px 14px;
  background: #F5F5F3;
  border: 0.5px solid #D5D7DA;
  border-radius: 3px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  color: #17181A;
  outline: none;
}

.form-input::placeholder {
  color: #b4ada6;
}

.form-input:focus {
  border-color: #A85A14;
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
  background: #3B3E43;
  color: #D5D7DA;
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
  color: #A85A14;
  border: 0.5px solid #A85A14;
  /* same padding/radius/font as above */
}
```

---

### Section Eyebrow Pattern

Every major section opens with a short rule + uppercase label before the headline.

```html
<div style="margin-bottom: 1rem;">
  <div style="width: 28px; height: 1px; background: #A85A14; margin-bottom: 8px;"></div>
  <div style="font-family: 'DM Sans', sans-serif; font-size: 10px; font-weight: 500; letter-spacing: 0.18em; text-transform: uppercase; color: #666B72;">
    Our work
  </div>
</div>
```

On dark backgrounds, the rule shifts to Ember (`#EE9440`) and the label to Slate Light (`#A7ADB5`).

---

### Dividers

```css
/* Standard rule */
.divider {
  border: none;
  border-top: 0.5px solid #D5D7DA;
  margin: 2.5rem 0;
}

/* Accent rule — use sparingly, once per section max */
.divider-accent {
  border: none;
  border-top: 1px solid #A85A14;
  width: 48px;
  margin: 2rem 0;
}
```

---

## Layout Principles

- **Section alternation:** Light sections (`#F5F5F3`) alternate with dark sections (`#17181A`). Never two dark sections in a row without a transition.
- **Max content width:** `1100px`, centered.
- **Section padding:** `5rem` top/bottom on desktop, `3rem` on mobile.
- **Grid:** 12-column. Most content uses 8 of 12 centered, or a 5/7 split for text/image pairs.
- **No decorative backgrounds** — no gradients, no grain overlays, no mesh textures. Atmosphere comes from typography and negative space.

---

## Interaction States

| Element | Default | Hover | Active/Focus |
|---|---|---|---|
| Nav links | `#D5D7DA` | `#F5F5F3` | `#F5F5F3` |
| Nav CTA | Ember border + text | Ember fill, Obsidian text | Scale `0.98` |
| Primary button | Ember Deep fill | `brightness-90` | Scale `0.97` |
| Outline button | Dark border | Dark fill, Chalk text | Scale `0.97` |
| Ghost button | Ember Deep underline | Ember Deep text brightens | Scale `0.98` |
| Text input | Ash border | Ash border | Ember Deep border + soft glow ring |
| Card (linked) | No shadow | Lift: `translateY(-2px)` | `translateY(0)` |

---

## Logo Usage

**Two logo assets provided:**
- `Asset_3.png` — No tang, clean silhouette. **Use for nav/small contexts.**
- `Asset_7.png` — Full arrowhead with tang. **Use for hero and large contexts.**

Both feature the illustrated hand-drawn style with crosshatching. Do not recolor or modify the logo files. The amber/orange facet in the logo is the source of the Ember Deep/Ember palette — do not replace it with a flat color version.

**Logo on dark (Obsidian) backgrounds:** Use as-is — the black outline reads on dark.
**Logo on light backgrounds:** Use as-is — the dark flint reads on Chalk and white.

---

## Voice Reference (for copy in components)

- **CTA buttons:** Direct action verbs. "Start a project" not "Get started today." "See our work" not "View portfolio."
- **Eyebrow labels:** Factual, understated. "Our work" not "What we've built." "Process" not "How it works."
- **Placeholder text:** Specific and inviting. "yourbusiness.com" not "Enter URL." "What's holding your current site back?" not "Your message."
- **Error states:** Plain and helpful. "Something's missing here." not "This field is required."

---

## What This System Is Not

- No gradients or shadows (except functional focus rings)
- No warm tones beyond the Ember family (Ember Tint, Ember, Ember Deep); neutrals stay cool grey, never beige or brown
- No full-bleed photography treatments or overlaid text on images
- No rounded corners beyond `r-6`
- No system fonts — always load Cormorant Garamond + DM Sans
- No additional accent colors beyond the eight tokens above

---

*FlintCraft Studio · Helena, Montana*
*Design system v1.0 — produced by Firefly Software*
