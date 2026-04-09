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
          obsidian:     "#1a1714",
          flint:        "#2e2924",
          ember:        "#9e5215",
          "ember-dark": "#8a4712",
          spark:        "#bf6e2a",
          birchwood:    "#faf9f7",
          parchment:    "#f0ece5",
          "trail-dust": "#6b5f52",
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
