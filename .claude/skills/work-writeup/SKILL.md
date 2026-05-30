---
name: work-writeup
description: "Create a /work portfolio entry (field report) from a live client URL. Scrapes the site, captures desktop + mobile screenshots, interviews the user for the engagement story, drafts the Project struct in the FlintCraft voice, and wires it into the site. Use after completing a client project."
disable-model-invocation: true
allowed-tools: Read Write Edit Bash Glob Grep WebFetch
argument-hint: "<live-url>  e.g. /work-writeup https://newclient.com"
---

# FlintCraft Studio — Work Writeup Skill

Turn a finished client site into a polished `/work/<slug>` field report. The live URL gives
you brand, copy, category, and screenshots. The **engagement story, metrics, and quote come
from the user** — you interview them for those. You produce a draft the user then refines.

A "/work entry" is three things:
1. A `Project` struct in the `Projects` map in `internal/view/portfolio.go`.
2. Two images in `web/static/images/portfolio/`: `<slug>_screenshot.webp` (desktop) and
   `<slug>_screenshot_mobile.webp` (mobile).
3. Wiring: the slug in `ProjectOrder` (front = featured/latest) and a correct `NextSlug`
   cycle so the "Next field report" link works.

Work the phases in order. Do not skip the interview — placeholders are not acceptable output.

---

## Phase 0 — Slug & client

Propose a kebab-case `slug` and a `Client` display name. The slug often does **not** match the
domain — e.g. `nautiluscleaning.com` is published at slug `nautilus-group`. So derive a
candidate from *both* the domain and the client's real business name, then **confirm it with
the user** before using it (a wrong slug bakes into the image filenames and the URL). Read
`internal/view/portfolio.go` first to see existing slugs and rule out a collision.

## Phase 1 — Scrape the live site

WebFetch the homepage, and `/about` + `/services` (or equivalent) if they exist. Extract:
- **Category + location** for the `Location` field, formatted like existing entries:
  `"Coffee roaster · Kennewick, WA"` (use `·` for the middle dot, `—` for em-dash).
- **Tags** — choose from the vocabulary already in use: `Marketing site`, `Brand`,
  `E-commerce`, `SEO`, `Design`. Pick 2–3 that genuinely apply.
- **Brand voice / what they do** — so the Subtitle and Story read true to the business.
- Any visible tech signals (Shopify, WooCommerce, Squarespace, a platform credit in the
  footer, etc.) — but read the caveat below before you interpret them.
Set `LiveURL` to the canonical URL.

> **Caveat — the live site IS the work, not the problem.** What you're scraping is the
> *new* build FlintCraft shipped. Do **not** describe the current stack, copy, or storefront
> as "the problem" — that is exactly backwards. Any platform you detect now (e.g. a custom
> Shopify storefront) is the *solution*, not what was replaced. The scrape also cannot see
> what came before, and a homepage usually has no testimonial. So treat the scrape as the
> source for the present state only (Subtitle, Tags, Location, Deliverables, voice) and
> anchor "the problem", the old stack, the metrics, and the quote **entirely** in the
> Phase 3 interview answers. Never infer the problem from the live site.

## Phase 2 — Capture screenshots

Both commands are validated on this machine. Run from repo root.

Desktop (resize to 800px wide, the dimension existing entries use):
```bash
chromium --headless --disable-gpu --no-sandbox --hide-scrollbars \
  --window-size=1280,832 --screenshot=/tmp/ww_desktop.png "<LIVE_URL>"
magick /tmp/ww_desktop.png -resize 800x -quality 82 \
  web/static/images/portfolio/<slug>_screenshot.webp
```

Mobile (400×866, matching existing mobile shots):
```bash
chromium --headless --disable-gpu --no-sandbox --hide-scrollbars \
  --window-size=400,866 --screenshot=/tmp/ww_mobile.png "<LIVE_URL>"
cwebp -q 82 /tmp/ww_mobile.png -o web/static/images/portfolio/<slug>_screenshot_mobile.webp
```

Then `identify` both webp files to confirm they're non-trivial in size and correct dimensions.
If a capture is blank/tiny (site uses a cookie wall or slow JS), retry once with a longer
`--virtual-time-budget=8000` flag, then tell the user it may need a manual screenshot.
Clean up: `rm -f /tmp/ww_desktop.png /tmp/ww_mobile.png`.

> **Caveat — do not lift an on-page testimonial into `Quote`.** A review on the client's
> site is a *customer praising the client's business* (e.g. a Google review of the cleaning
> service). That is NOT the project quote. The project `Quote` is the **client endorsing
> FlintCraft's work** — it comes from the Phase 3 interview, never from the scrape. If you do
> spot an on-page headshot that could serve as `Quote.Portrait`, note it but don't
> auto-download — raise it in Phase 3.

## Phase 3 — Interview the user

You now have everything scrapeable. Ask the user — in ONE batched message (a numbered list,
or use AskUserQuestion for the structured bits) — for what only they know:

1. **The problem** — what was the previous site/stack, and what was wrong with it? (Becomes the
   first Story paragraph and the angle of the Title.)
2. **What we did** — how FlintCraft rebuilt it; the approach that distinguishes the work.
3. **Outcomes / metrics** — up to 3 `Results`, each a `{Value, Unit, Label}` (e.g.
   `{"4", "wks", "Launch to first contract"}`). Numbers are best; "0→1" framing works for
   from-scratch builds. Results are optional — omit the field if there are none.
4. **Client quote + attribution** — the client's own words about working with FlintCraft
   (not a customer review of their business — see Phase 2 caveat), plus name/role. Optional.
   Ask whether to use a portrait image.
5. Confirm the **tags**, **location**, and the **Title angle** you're considering.

Keep it tight. Don't ask anything you already scraped.

## Phase 4 — Draft the struct (FlintCraft voice)

Read `CLAUDE.md` and 2–3 existing entries in `portfolio.go` first, then match them exactly.
Voice: **direct, understated, outcome-first.** Em-dashes (`—`), no hype, no exclamation
points. The site already frames every project as a **"Field report."**

- **Title** — a short outcome statement, often two sentences, like the existing ones:
  *"Inherited a plugin store. Replaced it with one worthy of the roast."* /
  *"No web presence. Now they have a pipeline."* Lead with the change, not the feature list.
- **Subtitle** — one line naming the deliverable scope and who it's for.
- **Story** — 2–3 paragraphs: the problem → what we built → the result/character. The
  *problem* paragraph comes from the interview, never from the scrape (see Phase 1 caveat) —
  the live site is what you built, not what was wrong. Use the user's interview answers;
  write in their register, about *this* business specifically.
- **Results** — the metrics from the interview (or omit).
- **Quote** — `&ProjectQuote{Text, Attribution, Portrait}` (Portrait optional).
- **Deliverables** — 5–8 bullets, concrete, em-dash-qualified, matching existing phrasing
  ("Custom-designed marketing site — no templates, no page builders").

Use Go string-escape conventions already in the file: `—` (—), `·` (·), `’` (’),
`“`/`”` (curly quotes), `→` (→), `→1` (→1).

## Phase 5 — Wire it into the site

The new project is the **latest**, so it takes the featured slot (`ProjectOrder[0]`).

1. Add the `Project{...}` literal to the `Projects` map in `internal/view/portfolio.go`.
2. **ProjectOrder**: prepend the new slug to the front of the slice.
3. **NextSlug cycle** (each project links to the next in a loop):
   - Set the new project's `NextSlug` / `NextClient` to the slug that was previously
     `ProjectOrder[0]`.
   - Find the project whose `NextSlug` currently points to that old-first slug and repoint
     its `NextSlug` / `NextClient` to the new slug. This keeps the cycle unbroken.
4. The new entry uses the standard `ProjectPage` template automatically — no handler change
   needed (`WorkProject` in `internal/handler/work.go` only special-cases `rockabilly-roasting`
   for its custom template; leave that alone).

## Phase 6 — Build & verify

```bash
mage GenerateTempl && go build ./... && go vet ./...
```
If it compiles, tell the user to run `mage Dev` and review at `/work` and `/work/<slug>`.
Summarize what you drafted and explicitly flag every field that is your best guess vs. their
own words, so they know what to refine. Do not commit unless asked.
