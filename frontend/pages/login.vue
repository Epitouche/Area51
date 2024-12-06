<script setup>
import { ref } from "vue";
// import { useRouter } from "vue-router";

const username = ref("");
const password = ref("");
const callCount = ref(0);
const didItWork = ref(false);
const body = computed(() => {
  return {
    username: username.value,
    password: password.value,
  };
});
// const router = useRouter();

async function onSubmit() {
  console.log(body.value);
  const { error } = await $fetch("http://localhost:8080/api/auth/login", {
    method: "POST",
    body,
    onResponse() {
      callCount.value++;
    },
  });
  if (!error.value) {
    didItWork.value = true;
  }
}
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-primaryDark-500"
  >
    <h1 class="text-white">{{ callCount }}</h1>
    <div
      class="w-full transform -translate-x-3/4 max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
    >
      <h1>{{ callCount }}</h1>
      <h1>{{ didItWork }}</h1>
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
          <a href="#" class="text-fontWhite underline">Create an account</a>
        </p>
        <p class="text-center text-sm">
          <a href="#" class="text-fontWhite underline">Forgot password?</a>
        </p>
      </div>
    </div>
  </div>
</template>
