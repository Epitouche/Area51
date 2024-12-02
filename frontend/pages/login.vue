<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-fontBlack"
  >
    <div
      class="w-full max-w-md p-8 space-y-16 bg-gradient-to-b from-tertiary-300 to-tertiary-500 dark:from-primaryDark-600 dark:to-primaryDark-800 text-fontWhite rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold text-center">LOG IN</h2>
      <form class="space-y-6">
        <div>
          <label for="email" class="block text-sm font-medium">Email</label>
          <input
            id="email"
            type="email"
            class="w-full px-4 py-2 mt-1 dark:text-fontBlack bg-primaryWhite-300 border border-primaryWhite-400 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent-500 focus:border-accent-500"
          />
        </div>
        <div>
          <label for="password" class="block text-sm font-medium"
            >Password</label
          >
          <input
            id="password"
            type="password"
            class="w-full px-4 py-2 mt-1 dark:text-fontBlack bg-primaryWhite-300 border border-primaryWhite-400 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent-500 focus:border-accent-500"
          />
        </div>
        <div class="flex items-center">
          <input
            id="remember"
            type="checkbox"
            class="w-4 h-4 text-accent-500 border-primaryWhite-300 rounded focus:ring-accent-500"
          />
          <label for="remember" class="ml-2 text-sm">Remember me</label>
        </div>
        <button
          type="submit"
          class="w-full px-4 py-2 text-lg font-semibold dark:text-fontBlack bg-accent-500 rounded-lg shadow-md hover:bg-accent-600 focus:outline-none focus:ring-2 focus:ring-accent-500"
        >
          Login
        </button>
      </form>
      <div class="flex justify-between space-x-4">
        <button
          class="flex items-center justify-center w-1/2 px-4 py-2 space-x-2 text-fontWhite dark:bg-secondaryDark-600 rounded-lg dark:hover:bg-secondaryDark-500"
        >
          <span>
            <i class="fab fa-github"></i>
          </span>
          <span>GitHub</span>
        </button>
        <button
          class="flex items-center justify-center w-1/2 px-4 py-2 space-x-2 text-fontWhite dark:bg-secondaryDark-500 rounded-lg dark:hover:bg-secondaryDark-400"
        >
          <span>
            <i class="fab fa-google"></i>
          </span>
          <span>Google</span>
        </button>
      </div>
      <p class="text-center text-sm mt-4">
        <a href="#" class="text-fontWhite underline">Forgot password?</a>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginCard",
};
</script>

<style scoped>
/* Vous pouvez ajouter des styles supplémentaires si nécessaire */
</style>
