export default {
  content: [
    "./components/**/*.{vue,js}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./app.vue",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        fontWhite: "#f5f5f5",
        fontBlack: "#0a0a0a",
        primaryWhite: {
          100: "#fafbfb",
          200: "#f6f6f6",
          300: "#f1f2f2",
          400: "#ededed",
          500: "#e8e9e9",
          600: "#bababa",
          700: "#8b8c8c",
          800: "#5d5d5d",
          900: "#2e2f2f",
        },
        secondaryWhite: {
          100: "#fdfefe",
          200: "#fcfdfd",
          300: "#fafcfd",
          400: "#f9fbfc",
          500: "#f7fafb",
          600: "#c6c8c9",
          700: "#949697",
          800: "#636464",
          900: "#313232",
        },
        primaryDark: {
          100: "#d1d1d1",
          200: "#a3a3a3",
          300: "#767676",
          400: "#484848",
          500: "#1a1a1a",
          600: "#151515",
          700: "#101010",
          800: "#0a0a0a",
          900: "#050505",
        },
        secondaryDark: {
          100: "#d3d4d6",
          200: "#a7a9ad",
          300: "#7a7e83",
          400: "#4e535a",
          500: "#222831",
          600: "#1b2027",
          700: "#14181d",
          800: "#0e1014",
          900: "#07080a",
        },
        tertiary: {
          100: "#f0ddff",
          200: "#e1bbfe",
          300: "#d298fe",
          400: "#c376fd",
          500: "#b454fd",
          600: "#9043ca",
          700: "#6c3298",
          800: "#482265",
          900: "#241133",
        },
        accent: {
          100: "#f6eaff",
          200: "#edd5ff",
          300: "#e3c0fe",
          400: "#daabfe",
          500: "#d196fe",
          600: "#a778cb",
          700: "#7d5a98",
          800: "#543c66",
          900: "#2a1e33",
        },
      },
    },
  },
  plugins: [],
};
