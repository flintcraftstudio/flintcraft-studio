---
name: flintcraft-design
description: Use this skill to generate well-branded interfaces and assets for FlintCraft Studio (Helena, MT — custom websites for established small businesses), either for production or throwaway prototypes/mocks/etc. Contains essential design guidelines, colors, type, fonts, assets, and UI kit components for prototyping.
user-invocable: true
---

Read the README.md file within this skill, and explore the other available files.

If creating visual artifacts (slides, mocks, throwaway prototypes, etc.), copy assets out and create static HTML files for the user to view. If working on production code, you can copy assets and read the rules here to become an expert in designing with this brand.

If the user invokes this skill without any other guidance, ask them what they want to build or design, ask some questions, and act as an expert designer who outputs HTML artifacts _or_ production code, depending on the need.

## What's in here

- `README.md` — full content fundamentals, visual foundations, iconography
- `colors_and_type.css` — eight color tokens plus white, spacing, radii, semantic typography; import directly
- `assets/` — logo silhouette + wordmark (SVG + PNG)
- `preview/` — one-concept-per-card design-system review specimens
- `ui_kits/website/` — React/JSX recreation of the marketing site (home, services, work, contact)

## Core rules (do not break)

- **Closed palette:** eight tokens plus white only. Ember Deep is only ever on light; Ember only on dark (with Obsidian text on Ember fills).
- **Two typefaces only:** Cormorant Garamond (display, italic register) + DM Sans (everything else, weights 300/400/500).
- **No gradients, no shadows, no decorative backgrounds.** Atmosphere is typography + negative space + section alternation.
- **No emoji.** Unicode arrows (`→`) and em dashes only.
- **Tight radii:** 3px interactive, 6px cards, 20px pills (status only).
- **Borders are 0.5px Ash** except 2.5px Ember Deep/Ember left-border accent cards and 2px Ember Deep featured items.

## Voice

Direct, understated, knowledgeable. Short declaratives. No hype verbs. CTAs are action verbs ("Start a project", not "Get started today"). Italic is a register (tagline, hero, pull quotes) — never general emphasis.
