/** @type {import('tailwindcss').Config} */
const defaultTheme = require("tailwindcss/defaultTheme");

// ─── Fluid Type Utility ───────────────────────────────────────────────────────
//
// Generates a CSS clamp() value that scales linearly between two sizes
// across a defined viewport range.
//
// minSize / maxSize : target font sizes in px at the viewport bounds
// minVw / maxVw     : viewport width bounds in px (default: 320 → 1280)
//
function fluidType(minPx, maxPx, minVw = 320, maxVw = 1280) {
  const minRem = minPx / 16;
  const maxRem = maxPx / 16;
  const slope = (maxRem - minRem) / (maxVw / 16 - minVw / 16);
  const intercept = minRem - slope * (minVw / 16);

  const clampMin = `${minRem.toFixed(4)}rem`;
  const clampVal = `calc(${intercept.toFixed(4)}rem + ${(slope * 100).toFixed(4)}vw)`;
  const clampMax = `${maxRem.toFixed(4)}rem`;

  return `clamp(${clampMin}, ${clampVal}, ${clampMax})`;
}

// ─── Type Scale Definition ────────────────────────────────────────────────────
//
// Two tracks share the same step names:
//   Cormorant Garamond (font-display) — display/headlines, runs optically large
//   DM Sans (font-body)               — body/UI, runs true to size
//
// Naming convention:
//   text-display / text-display-display  — hero headlines
//   text-h1 / text-h1-display            — page titles
//   text-h2 / text-h2-display            — section headers
//   text-h3 / text-h3-display            — subsection headers
//   text-h4 / text-h4-display            — card titles, label headers
//   text-body-lg                         — lead paragraphs (DM Sans only)
//   text-body                            — base body copy (DM Sans only)
//   text-body-sm                         — captions, supporting copy (DM Sans only)
//   text-ui                              — nav, buttons, UI labels (DM Sans only)
//   text-eyebrow                         — section eyebrows (DM Sans only)
//
const typeScale = {
  'display': {
    cormorant: { size: [44, 64], leading: '1.05', tracking: '-0.02em' },
    dmSans:    { size: [48, 72], leading: '1.05', tracking: '-0.02em' },
  },
  'h1': {
    cormorant: { size: [32, 46], leading: '1.1',  tracking: '-0.015em' },
    dmSans:    { size: [36, 52], leading: '1.1',  tracking: '-0.015em' },
  },
  'h2': {
    cormorant: { size: [26, 36], leading: '1.2',  tracking: '-0.01em' },
    dmSans:    { size: [28, 40], leading: '1.2',  tracking: '-0.01em' },
  },
  'h3': {
    cormorant: { size: [20, 28], leading: '1.3',  tracking: '-0.005em' },
    dmSans:    { size: [22, 30], leading: '1.3',  tracking: '-0.005em' },
  },
  'h4': {
    cormorant: { size: [17, 22], leading: '1.4',  tracking: '0em' },
    dmSans:    { size: [18, 24], leading: '1.4',  tracking: '0em' },
  },
  'body-lg': {
    dmSans: { size: [19, 22], leading: '1.65', tracking: '0em' },
  },
  'body': {
    dmSans: { size: [17, 19], leading: '1.7',  tracking: '0em' },
  },
  'body-sm': {
    dmSans: { size: [15, 17], leading: '1.6',  tracking: '0em' },
  },
  'ui': {
    dmSans: { size: [13, 15], leading: '1.4',  tracking: '0.08em' },
  },
  'eyebrow': {
    dmSans: { size: [11, 13], leading: '1.4',  tracking: '0.2em' },
  },
};

// ─── Typography Plugin ────────────────────────────────────────────────────────
//
// Generates two utility classes per heading step:
//   .text-{step}            — DM Sans values (default track)
//   .text-{step}-display    — Cormorant values (display track)
//
// Body/UI steps generate one class each (DM Sans only).
//
function typographyPlugin({ addUtilities }) {
  const utilities = {};

  for (const [step, tracks] of Object.entries(typeScale)) {
    if (tracks.dmSans) {
      const { size, leading, tracking } = tracks.dmSans;
      utilities[`.text-${step}`] = {
        fontSize:      fluidType(...size),
        lineHeight:    leading,
        letterSpacing: tracking,
      };
    }

    if (tracks.cormorant) {
      const { size, leading, tracking } = tracks.cormorant;
      utilities[`.text-${step}-display`] = {
        fontSize:      fluidType(...size),
        lineHeight:    leading,
        letterSpacing: tracking,
      };
    }
  }

  addUtilities(utilities);
}

// ─── Tailwind Config ──────────────────────────────────────────────────────────

module.exports = {
  content: [
    "./internal/view/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        fc: {
          // Dark surfaces
          obsidian:     "#1c1e24",  // page bg
          flint:        "#363b42",  // raised surface (cards, nav)
          charcoal:     "#474d55",  // highest surface (modals, popovers, hover)

          // Dark-mode text
          ash:          "#c4bdb2",  // body / strong muted on dark
          smoke:        "#96897a",  // subtle muted — captions, eyebrows, metadata
          spark:        "#e8873b",  // accent on dark

          // Light surfaces
          birchwood:    "#faf9f7",  // page bg
          parchment:    "#f0ece5",  // raised surface

          // Light-mode text
          "trail-dust": "#4a4038",  // body / strong muted on light
          bracken:      "#6b5f52",  // subtle muted — captions, eyebrows, metadata
          ember:        "#a84812",  // accent on light
          "ember-dark": "#8f3d0f",  // ember hover (darken 8%)

          // Structure
          "slate-rim":  "#d4cfc9",  // borders, dividers
        },
      },
      fontFamily: {
        display: ['"Cormorant Garamond"', ...defaultTheme.fontFamily.serif],
        body:    ['"DM Sans"', ...defaultTheme.fontFamily.sans],
      },
      maxWidth: {
        content: "1100px",
      },
      transitionTimingFunction: {
        "out-expo":  "cubic-bezier(0.16, 1, 0.3, 1)",
        "out-quart": "cubic-bezier(0.25, 1, 0.5, 1)",
      },
    },
  },
  plugins: [typographyPlugin],
}
