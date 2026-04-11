/** @type {import('tailwindcss').Config} */
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: [
    "./internal/view/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        fc: {
          obsidian:     "#232226",
          flint:        "#2e3338",
          ember:        "#cc5e1a",
          "ember-dark": "#a84b12",
          spark:        "#db7f36",
          birchwood:    "#faf9f7",
          parchment:    "#f0ece5",
          "trail-dust": "#6b5f52",
          ash:          "#8a7d6e",
          "slate-rim":  "#d4cfc9",
        },
      },
      fontFamily: {
        display: ['"Cormorant Garamond"', ...defaultTheme.fontFamily.serif],
        body:    ['"DM Sans"', ...defaultTheme.fontFamily.sans],
      },
      maxWidth: {
        content: "1100px",
      },
    },
  },
  plugins: [],
}
