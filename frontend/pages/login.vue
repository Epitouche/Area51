<<<<<<< Updated upstream
<script>
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
=======
<script setup>
import { ref } from "vue";

const username = ref("");
const password = ref("");

async function onSubmit() {
  try {
    const { access_token } = await $fetch("http://localhost:8080/api/auth/login", {
      method: "POST",
      body: {
        username: username.value,
        password: password.value,
      },
    });

    if (access_token) {
      const tokenCookie = useCookie("token");
      tokenCookie.value = access_token;

      navigateTo("/services");
    } else {
      console.error("Access token not received");
    }
  } catch (error) {
    console.error("API response:", error.response?.data || error.message);
  }
}

async function redirectToGitHubOAuth() {
  try {
    const { github_authentication_url } = await $fetch(
      "http://localhost:8080/api/github/auth",
      {
        method: "GET",
      }
    );
    if (github_authentication_url) {
      window.location.href = github_authentication_url;
    } else {
      console.error("GitHub authentication URL not found");
    }
  } catch (error) {
    console.error("Error fetching GitHub OAuth URL:", error);
  }
}
>>>>>>> Stashed changes
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-primaryDark-500"
  >
    <div
      class="w-full transform -translate-x-3/4 max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold text-center">LOG IN</h2>
      <form class="space-y-6">
        <div>
          <Input
            id="email"
            type="email"
            label="Email"
            icon="fas fa-user"
          />
        </div>
        <div>
          <Input
            id="password"
            type="password"
            label="Password"
            icon="fas fa-lock"
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
      <div class="flex justify-around">
        <p class="text-center text-sm">
          <a href="#" class="text-fontWhite underline">Create an account</a>
        </p>
        <p class="text-center text-sm">
          <a href="#" class="text-fontWhite underline">Forgot password?</a>
        </p>
      </div>
    </div>
  </div>
</template>
