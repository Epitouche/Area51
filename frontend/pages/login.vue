<script>
import Button from "@/components/Button.vue";
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-fontBlack"
  >
    <div
      class="w-full max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
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
      </form>
      <div class="flex items-center">
        <input
          id="remember"
          type="checkbox"
          class="w-4 h-4 text-accent-500 border-primaryWhite-300 rounded focus:ring-accent-500"
        />
        <label for="remember" class="ml-2 text-sm">Remember me</label>
      </div>
      <div class="flex justify-center">
        <Button
          text="Log In"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
      </div>
      <hr class="border-primaryWhite-400" />
      <div class="flex justify-around space-x-4">
        <Button
          text="Github"
          class="w-full"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
        <Button
          text="Google"
          class="w-full"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
      </div>
      <p class="text-center text-sm">
        <a href="#" class="text-fontWhite underline">Forgot password?</a>
      </p>
    </div>
  </div>
</template>
