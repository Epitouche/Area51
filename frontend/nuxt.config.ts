export default defineNuxtConfig({
  css: ["@/assets/css/main.css"],

  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      title: 'Area51',
      htmlAttrs: {
        lang: 'en',
      },
    }
  },

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  colorMode: {
    classSuffix: '',
    preference: 'light',
    fallback: 'light',
    storage: 'cookie'
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