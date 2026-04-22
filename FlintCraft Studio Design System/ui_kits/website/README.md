# FlintCraft Studio — Website UI Kit

High-fidelity recreation of the FlintCraft Studio marketing site (Helena, MT). Rebuilds the home, services, work, and contact surfaces as static React/JSX components using the v1.1 twelve-token palette.

## Screens covered

- **Home** — Hero with staggered entrance, stats bar, "Who we are" credentialing, ICP statement on Obsidian, WordPress-vs-FlintCraft comparison table, testimonial pull quote, final CTA.
- **Services** — One-relationship hero, six capability cards on Flint, pricing statement, six-month commitment, expandable FAQ.
- **Work** — Field reports hero, featured project (Nautilus Group), secondary project cards.
- **Contact** — Start-a-project form with eyebrow labels and success state.

Navigate between screens with the top nav; tabs are state, not routes.

## Components

Each component is a small JSX file, loaded by `index.html` via Babel. Exported to `window` for cross-file sharing.

- `Nav.jsx` — top nav with brand + tagline
- `Footer.jsx` — footer with contact triad
- `Primitives.jsx` — `Eyebrow`, `RuleAccent`, `PullQuote`, `Button`, `GhostLink`
- `Cards.jsx` — `CardLight`, `CardDark`, `CardAccent`, `CapabilityCard`, `ProjectCard`, `FeaturedProjectCard`
- `Forms.jsx` — `TextInput`, `TextArea`, `ContactForm`
- `Screens.jsx` — `HomeScreen`, `ServicesScreen`, `WorkScreen`, `ContactScreen`

## Caveats

- Production hero photography not included — substituted with Obsidian surface + logo watermark.
- Project screenshots are Obsidian placeholders; swap with `/static/images/portfolio/*.webp` from the repo when shipping.
- Forms are cosmetic (no HTMX/Postmark wire-up).
