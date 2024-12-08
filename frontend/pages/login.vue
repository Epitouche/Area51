<script setup>
import { ref } from "vue";
// import { useRouter } from "vue-router";

const username = ref("");
const password = ref("");
const router = useRouter();

async function onSubmit() {
  try {
    const { token } = await $fetch("http://localhost:8080/api/auth/login", {
      method: "POST",
      body: {
        username: username.value,
        password: password.value,
      },
    });
    if (token) {
      localStorage.setItem("token", token);
      router.push("/services");
    }
  } catch (error) {
    console.error(error);
  }
}

async function redirectToGitHubOAuth() {
  try {
    const { github_authentication_url } = await $fetch("http://localhost:8080/api/github/auth", {
      method: "GET",
    });
    if (github_authentication_url) {
      window.location.href = github_authentication_url;
    } else {
      console.error("GitHub authentication URL not found");
    }
  } catch (error) {
    console.error("Error fetching GitHub OAuth URL:", error);
  }
}

</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-primaryDark-500"
  >
    <div
      class="w-full transform -translate-x-3/4 max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold text-center">LOG IN</h2>
      <form class="space-y-6" @submit.prevent="onSubmit">
        <div>
          <InputComponent
            id="username"
            v-model="username"
            type="text"
            label="Username"
            icon="fas fa-user"
          />
        </div>
        <div>
          <InputComponent
            id="password"
            v-model="password"
            type="password"
            label="Password"
            icon="fas fa-lock"
          />
        </div>
        <!-- <div class="flex items-center gap-2">
          <InputComponent
            id="remember"
            v-model="rememberMe"
            type="checkbox"
            class="w-4 h-4 text-accent-500 border-primaryWhite-300 rounded focus:ring-accent-500"
            label=""
          />
          <label for="remember" class="ml-2 text-sm">Remember me</label>
        </div> -->
        <div class="flex justify-center">
          <ButtonComponent
            text="Log In"
            bg-color="bg-primaryWhite-500"
            hover-color="hover:bg-secondaryWhite-500"
            text-color="text-fontBlack"
          />
        </div>
      </form>
      <hr class="border-primaryWhite-400">
      <div class="flex justify-around space-x-4">
        <ButtonComponent
          text="Github"
          class="w-full"
          bg-color="bg-primaryWhite-500"
          hover-color="hover:bg-secondaryWhite-500"
          text-color="text-fontBlack"
          :on-click="redirectToGitHubOAuth"
        />
        <ButtonComponent
          text="Google"
          class="w-full"
          bg-color="bg-primaryWhite-500"
          hover-color="hover:bg-secondaryWhite-500"
          text-color="text-fontBlack"
        />
      </div>
      <div class="flex justify-around">
        <p class="text-center text-sm">
          <NuxtLink to="/register" class="text-fontWhite underline">
            Create an account
          </NuxtLink>
        </p>
        <p class="text-center text-sm">
          <a href="#" class="text-fontWhite underline">Forgot password?</a>
        </p>
      </div>
    </div>
  </div>
</template>
