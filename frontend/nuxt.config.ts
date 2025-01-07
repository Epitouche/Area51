export default defineNuxtConfig({
  css: ["@/assets/css/main.css"],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: [
    "@nuxt/devtools",
    "@nuxt/eslint",
    "@nuxt/icon",
    "@pinia/nuxt",
    "@vueuse/motion/nuxt",
  ],

  compatibilityDate: "2024-12-03",
});