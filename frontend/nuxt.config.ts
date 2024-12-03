export default defineNuxtConfig({
  css: ['@/assets/css/main.css'],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  modules: ['@nuxt/devtools', '@nuxt/eslint'],
  compatibilityDate: '2024-12-03',
});