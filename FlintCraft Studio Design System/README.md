# FlintCraft Studio — Design System

**Studio name:** FlintCraft Studio
**Location:** Helena, Montana
**Tagline:** *Still shaped by hand.*
**Hero headline:** *Carved for purpose.*

FlintCraft Studio builds custom websites for established small businesses — owners in years three to five who have earned their reputation and outgrown their first site. Everything ships as hand-written code: no WordPress, no page builders, no templates. One named engineer per project, same-day support, fixed-rate retainers ($150–$350/mo, all in). Boutique quality, coastal craft, Montana rate.

Positioning line used in collateral: *"As durable as the mountains, and just as unlikely to move to San Francisco."*

---

## Sources

This system was assembled from the FlintCraft Studio marketing-site repo plus a v1.1 color spec provided directly in the brief.

- **Codebase:** `github.com/flintcraftstudio/flintcraft-studio` — Go + templ + Tailwind CSS (standalone binary). Key files read:
  - `design-system.md` — v1.0 system (nine color tokens; superseded on palette by this system)
  - `CLAUDE.md`, `README.md` — project overview
  - `tailwind/tailwind.config.js` — `fc-*` token names, fluid type plugin, type scale
  - `internal/view/*.templ` — home, services, process, work, contact, nav, footer (canonical component patterns)
  - `logo.svg`, `logo.png`, `logo_wordmark.svg`, `logo_wordmark.png`
- **Brief supplement:** expanded v1.1 color palette — twelve tokens, splitting body vs. muted text tiers on light and dark surfaces, and introducing Charcoal (highest dark surface).

The code patterns shown throughout this system are simplified cosmetic recreations of what ships in the repo — not the production Go + templ source.

---

## Index

Root files:

- `README.md` — this document
- `SKILL.md` — Agent Skills manifest (used when this folder is bundled as a Claude Code skill)
- `colors_and_type.css` — CSS variables for the twelve color tokens, spacing, radii, and semantic typography
- `assets/` — logos (SVG + PNG, silhouette + wordmark)
- `fonts/` — (none; all type loaded from Google Fonts)
- `preview/` — design-system review cards (one per concept, registered to the Design System tab)
- `ui_kits/website/` — marketing-site UI kit (hero, nav, cards, forms, buttons — JSX components + demo)

Design system tab: every card in `preview/` is registered with `register_assets`, grouped by concept (Colors / Type / Spacing / Components / Brand).

---

## Content Fundamentals

**Voice.** Direct, understated, knowledgeable. The studio writes like an experienced tradesperson: confident about craft, uninterested in theatrics. Short declaratives carry most of the weight. Compound sentences are used to add nuance, never to hedge. Hype words (*transform, unleash, revolutionize, journey, seamless*) do not appear.

**Person.** First-person plural for the studio (*we, our*). Second-person for the client (*you, your*). The studio never refers to itself in the third person. Clients are assumed to be adults who run real businesses.

**Casing.**
- Sentence case for sentences and headlines — no title-casing of body prose.
- UPPERCASE with heavy tracking (`0.18–0.22em`) for eyebrow labels and UI chrome (buttons, nav links, badges).
- Never uppercase without tracking.

**Punctuation.** American. Serial comma. Em-dashes spaced or unspaced — consistent within a paragraph. Numerals written as digits (*3–5 years*, *$150–$350*), except at the start of a sentence.

**Emoji.** None. Not in copy, not in UI, not in decorative roles. Unicode arrows (`→`, `—`) are used; nothing else.

**Length.** Short paragraphs (2–4 sentences). Lists over long clauses when items are parallel. Headlines cap around 8 words.

**Italic is a register, not emphasis.** Cormorant italic is reserved for the tagline, hero headline, pull quotes, and one-off emphatic phrases (*"Carved for purpose."*, *"Everything handled."*). Body paragraphs do not italicize for emphasis — they rephrase. Bold (DM Sans 500) is used for parenthetical emphasis at most once per paragraph.

**Calls to action.** Direct action verbs. *Start a project* (not *Get started today*). *See our work* (not *View our portfolio*). *Send it* (not *Submit*). *Read our full terms →* (not *Learn more*). The arrow suffix `→` is standard for link-style CTAs.

**Errors and empty states.** Plain-spoken, human. *"Something's missing here."* *"Got it. You'll hear from Logan. Not a bot. Not a VA. The person who will actually build your project is reading this right now."*

**Representative passages.**

> FlintCraft Studio is a small custom development studio in Helena, Montana. We build websites for businesses that have earned their reputation and are ready for a web presence that reflects it.

> Everything we ship is written from scratch. No WordPress. No plugins. No templates. No page builders. Clean, fast, maintainable code that we host, maintain, and stand behind.

> You've done the hard part. You built something real, earned real customers, and proven the business works. Your website is the one thing you never got right — or the one thing you built when you were still figuring everything out.

**Eyebrow labels** (short, factual): *Services · Our work · Selected work · Featured · More work · Who we are · Who we work with · How we work · When things shift · What it costs · Common questions*.

**Pull-quote treatment** (Cormorant italic, Ember on light / Spark on dark): used once per section, never twice.

---

## Visual Foundations

The FlintCraft visual system is **quiet, precise, and intentionally closed.** The identity is carried by typography, negative space, and a single burnt-orange accent against warm neutrals. There are no gradients, no drop shadows, no photographic overlays, no decorative backgrounds, no illustrations beyond the logo.

### Color

Twelve named tokens, split across three tiers on each surface:

| Tier | Light surface | Dark surface |
|---|---|---|
| Page | Birchwood `#faf9f7` | Obsidian `#1c1e24` |
| Raised (cards, nav) | Parchment `#f0ece5` | Flint `#363b42` |
| Highest (modals, popovers, hover) | *(border on parchment)* | Charcoal `#474d55` |
| Primary text | Obsidian | Birchwood |
| Body muted | Trail Dust `#4a4038` | Ash `#c4bdb2` |
| Subtle muted | Bracken `#6b5f52` | Smoke `#96897a` |
| Accent | Ember `#a84812` | Spark `#e8873b` |
| Border | Slate Rim `#d4cfc9` | Slate Rim `#d4cfc9` |

**Accent rules.** Ember and Spark never cross surfaces — Ember on light, Spark on dark. Ember on Flint or Obsidian fails AA; Spark on Birchwood washes out. On Flint specifically, Spark is only safe at 18px+ or 14px bold; smaller accent text on Flint should use Ash for emphasis instead.

**Never.** Trail Dust or Bracken on dark. Ash or Smoke on light. Spark for body-weight text on Flint. Colors outside the set.

### Type

Two families, loaded from Google Fonts, no substitutions.

- **Cormorant Garamond** — display, headlines, pull quotes. 400 and 600; italic used as a register (tagline, hero, pull quotes, one-off emphatic phrases).
- **DM Sans** — all body copy, UI labels, eyebrows, buttons. Three weights only: 300 (body), 400 (default/caption), 500 (labels/buttons). 600 and 700 are not used.

Fluid type scale with `clamp()` between 320px and 1280px viewports. Line-height `1.7` for body, `1.05–1.2` for display. Eyebrow labels always uppercase with `0.18–0.22em` tracking — never uppercase without tracking.

### Spacing

Five-step scale: `4 / 8 / 16 / 32 / 64 px`. Sections are 5rem top/bottom on desktop, 3rem on mobile. Max content width is `1100px`. The canonical grid is 12-column, most content on 8 of 12 centered or a 5/7 split for text+image pairs.

### Radii

Tight. `3px` for buttons, inputs, badges. `6px` for cards. `20px` pill only for status indicators. Nothing larger — the brand is precise, not soft.

### Borders

All borders are `0.5px` Slate Rim. Two exceptions:

- **Accent cards** (testimonials, callouts): `2.5px` left border only. Ember on light, Spark on dark. Radius is `0` — the bar sits flush against the card.
- **Featured items:** `2px` full border. Ember on light, Spark on dark.

### Backgrounds

Atmosphere comes from section alternation, not decoration. Light sections (Birchwood or Parchment) alternate with dark sections (Obsidian or Flint). Never two darks in a row without a transition.

- **No gradients.** The only gradient permitted is a hero scrim (`bg-fc-obsidian/55`) over a full-bleed photograph, and only on the home hero.
- **No grain, noise, or mesh overlays.**
- **No repeating patterns or textures.**
- **Full-bleed photography** is used once per homepage, darkened with an Obsidian scrim. All other imagery sits inside card frames.

### Shadows

**None.** No drop shadows, no box-shadows. The system expresses depth through border weight (`0.5px` rule → `0.5px` card border → `2px` featured border → `2.5px` accent bar) and surface alternation (Obsidian → Flint → Charcoal on dark; Birchwood → Parchment → bordered Parchment on light).

The one exception is the input focus ring: `0 0 0 2px rgba(168, 72, 18, 0.12)`. Functional only.

### Imagery

Photography is warm, natural, and un-retouched in feel — Montana light, real hands, real interiors, no stock polish. Full-bleed treatments are exceedingly rare. When photography appears, it appears inside a `6px` radius card frame alongside text, not as a backdrop.

The logo is a hand-drawn crosshatched arrowhead — the only illustration in the system. Its amber facet is the source of Ember/Spark.

### Motion

Restrained. Entrances fade + translate 12–18px on a 700ms `ease-out-expo` curve (`cubic-bezier(0.16, 1, 0.3, 1)`). Interior micro-interactions use 200–500ms `ease-out-quart` (`cubic-bezier(0.25, 1, 0.5, 1)`).

- No bounce, no spring, no overshoot.
- No scroll-linked parallax beyond the hero's 3s ease-out image zoom-in on mount.
- `prefers-reduced-motion: reduce` turns off all transitions.

### Hover & press states

| Surface | Hover | Active |
|---|---|---|
| Nav link (Bracken) | Parchment — colour only | — |
| Primary button (Ember) | Darken 8% (Ember-dark `#8f3d0f`) + `translateY(-1px)` | `scale(0.97)` |
| Outline button | Fill with border colour, text flips to Birchwood | `scale(0.97)` |
| Ghost button (Ember underline) | `filter: brightness(1.1)` | `scale(0.98)` |
| Card (linked) | No shadow; `translateY(-2px)` | `translateY(0)` |
| Image in card | `scale(1.02)` over 500ms | — |
| Input | Slate Rim border → Ember border + `0 0 0 2px rgba(168, 72, 18, 0.12)` focus ring | — |

Brightness and translate are preferred over colour shifts for non-accent elements. Scale-down on press is universal for buttons (`0.97`) and ghost links (`0.98`).

### Transparency & blur

Effectively none. The only semi-transparent elements:

- Hero image scrims: `bg-fc-obsidian/25` (global tint) and `bg-fc-obsidian/55` (text-side panel).
- Focus ring: `rgba(168, 72, 18, 0.12)`.
- Tag chips on dark cards: `border-fc-ash/30`, `bg-fc-ash/20`.

No `backdrop-filter`, no frosted glass, no blurred surfaces.

### Cards

- **Light card** — `bg-Birchwood` or `bg-Parchment`, `0.5px` Slate Rim border, `6px` radius, `20–24px` internal padding, no shadow.
- **Dark card** — `bg-Flint`, `0.5px` border matched to Flint (effectively invisible), `6px` radius, same padding.
- **Accent card** (testimonial / callout) — `bg-Parchment`, `2.5px` left border Ember (or Spark on dark), radius `0`.

Card anatomy is consistent: eyebrow (uppercase, tracked) → Cormorant title → DM Sans body → `→`-suffix CTA link. Card hover is a `translateY(-2px)` lift and optional border-colour shift to a 30%-opacity accent; never a shadow.

### Layout rules

- Fixed top nav on Obsidian, scrolls with page.
- Section padding `5rem` desktop / `3rem` mobile. Top + bottom, matched.
- Divider rule: `0.5px` Slate Rim, `2.5rem` vertical margin.
- Accent divider (sparingly): `1px` Ember, `48px` wide, aligned left.
- Section eyebrow pattern: short `28×1px` Ember rule + `8px` gap + uppercase tracked label + headline.

---

## Iconography

FlintCraft's approach to iconography is **minimal by design.** The shipping site contains essentially no decorative icons — the only recurring glyph is the `→` Unicode right arrow used on ghost/link-style CTAs. There is no installed icon font, no Heroicons/Lucide/Feather import, no proprietary SVG set.

Where icons do appear, they are inline SVGs written directly in markup at `1.5px` stroke weight, currentColor-driven, sized to 5×5 Tailwind units (20px). Two such icons exist in the codebase:

- **Hamburger menu** — three horizontal `path` strokes (mobile nav toggle).
- **Close** — two diagonal `path` strokes (mobile nav close).

Both are defined inline in `internal/view/nav.templ`. Stroke weight `1.5`, `fill="none"`, `stroke="currentColor"`.

**For future icon needs**, this system substitutes **[Lucide](https://lucide.dev)** — the closest open-source match to the existing two SVGs (same 1.5px stroke, same outline style, same 24×24 viewBox). Lucide ships as both a web component and individual SVGs. **This is a substitution, not an established choice**, and should be flagged to the user before shipping production icons.

- CDN: `<script src="https://unpkg.com/lucide@latest"></script>` and `<i data-lucide="arrow-right"></i>` + `lucide.createIcons()`.

**Emoji are not used** anywhere in the system — not in copy, not in UI, not in error states. **Unicode glyphs** are used sparingly: `→` (right arrow, CTAs), `—` (em dash, punctuation), `·` (interpunct, metadata separators), `&mdash;`, `&rdquo;`, `&ldquo;`.

**Brand assets** in `assets/`:
- `logo.svg` / `logo.png` — clean silhouette, no tang. Use at small sizes (nav, favicons, body references).
- `logo_wordmark.svg` / `logo_wordmark.png` — full arrowhead with tang. Use at hero scale.

Both feature hand-drawn crosshatching and an amber facet. Do not recolour the logo or produce a flat-colour variant — the illustration carries the "shaped by hand" message; flattening it loses the brand.

**Placeholder imagery.** The codebase references `/static/images/hero.webp` (Montana mountains at sunset with circuit traces carved into stone) and per-client screenshots. These production images are **not included** in this design system package; HTML examples in `ui_kits/website/` use a solid Obsidian block with the logo watermark as a placeholder where a hero photograph would live. Request production imagery from the studio when building anything shipping.

---

## Flags & substitutions

- **Icon set** — Lucide is a substitution for a not-yet-established set. Confirm before shipping production icons.
- **Hero photography** — real assets not included. Current placeholders are logo watermarks on Obsidian.
- **Color palette version** — this system uses the **v1.1 twelve-token palette** from the brief. The committed `design-system.md` in the repo is v1.0 (nine tokens, different hex values for Ember/Obsidian/Flint/Ash/Trail Dust). The v1.1 version supersedes it here. If the repo's `design-system.md` is updated, revisit this file.

---

*FlintCraft Studio · Helena, Montana · Design system v1.1*
