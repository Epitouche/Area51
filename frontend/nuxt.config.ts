export default defineNuxtConfig({
  css: ["@/assets/css/main.css"],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  colorMode: {
    classSuffix: '',
    preference: 'system',
    fallback: 'light'
  },
  modules: [
    "@nuxt/devtools",
    "@nuxt/eslint",
    "@nuxt/icon",
    "@pinia/nuxt",
    "@vueuse/motion/nuxt",
    "@nuxtjs/color-mode",
  ],

  compatibilityDate: "2024-12-03",
});