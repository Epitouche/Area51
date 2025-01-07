export default defineNuxtConfig({
  css: ["@/assets/css/main.css"],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: ["@nuxt/devtools", "@nuxt/eslint", "@nuxt/icon", "@nuxtjs/color-mode"],

  colorMode: {
    classSuffix: '',
    preference: 'system',
    fallback: 'light'
  },

  compatibilityDate: "2024-12-03",
});