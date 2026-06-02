/** @type {import('tailwindcss').Config} */
//
// Per-demo Tailwind config for the Alpine Spine chiropractor demo.
// The marketing site has its own config (../tailwind.config.js); each demo
// gets its own so a vertical can carry its own palette without forking the
// shared build. The bespoke design system lives in ./chiropractor.css as raw
// CSS; this config exposes the same tokens as Tailwind utilities for any
// incidental utility use in the demo's .templ files.
//
// Built via `mage BuildChiroCSS` -> demos/chiropractor/static/css/site.css
module.exports = {
  content: ["./demos/chiropractor/**/*.templ"],
  theme: {
    extend: {
      colors: {
        paper: "#FAF6EC",
        cream: "#F1EBDB",
        linen: "#ECE4D2",
        sand: "#E2D8C1",
        stone: "#D3C7AC",
        ink: { DEFAULT: "#1E2E2C", 2: "#44524D", 3: "#5E665C" },
        teal: {
          900: "#102B2E", 800: "#15363A", 700: "#1C4347", 600: "#235257",
          500: "#2F666B", 300: "#8FB0AF", 100: "#DCE8E5", 50: "#ECF2EF",
        },
        pine: { 700: "#324A2A", 600: "#3E5236" },
        sage: { 500: "#6E8A5C", 300: "#A9BE98", 100: "#E6ECDC" },
      },
      fontFamily: {
        display: ['"Newsreader"', "Georgia", "serif"],
        sans: ['"DM Sans"', "system-ui", "sans-serif"],
      },
      maxWidth: {
        content: "1200px",
      },
    },
  },
  plugins: [],
};
