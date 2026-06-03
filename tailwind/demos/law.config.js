/** @type {import('tailwindcss').Config} */
//
// Per-demo Tailwind config for the Granite Peak Trial Lawyers demo (the second
// vertical, after the chiropractor). Mirrors that build: each demo gets its own
// config so it can carry its own palette without forking the shared build. The
// bespoke design system lives in ./law.css as raw CSS; this config exposes the
// same tokens as Tailwind utilities for any incidental utility use in .templ.
//
// Built via `mage BuildLawCSS` -> demos/law/static/css/site.css
module.exports = {
  content: ["./demos/law/**/*.templ"],
  theme: {
    extend: {
      colors: {
        pine: "#373f3c",
        evergreen: "#2c3431",
        ink: "#232826",
        brass: {
          100: "#f1e7d3", 200: "#e3d0a9", 400: "#c9a86a",
          500: "#a57633", 600: "#8f6429", 700: "#75501f",
        },
        bone: "#f6f3ec",
        paper: "#fbfaf5",
        stone: { 100: "#ece7db", 200: "#ded7c6", 300: "#c8bfa9" },
        fg: { 1: "#232826", 2: "#54615c", 3: "#646f68" },
        success: "#2f6b4f",
        error: "#a13d2d",
      },
      fontFamily: {
        display: ['"Spectral"', "Georgia", "serif"],
        sans: ['"Libre Franklin"', "system-ui", "sans-serif"],
      },
      maxWidth: {
        content: "1200px",
      },
    },
  },
  plugins: [],
};
