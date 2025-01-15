<script setup lang="ts">
import { ref } from "vue";
import { useNotificationStore } from "@/stores/notification";
import type { AboutResponse } from "~/src/types";

type ServiceCard = {
  name: string;
  image: string;
};

const username = ref<string>("");
const password = ref<string>("");

const notificationStore = useNotificationStore();

type NotificationType = "success" | "error" | "warning";

function triggerNotification(
  type: NotificationType,
  title: string,
  message: string
) {
  notificationStore.addNotification({
    type,
    title,
    message,
  });
}

interface LoginResponse {
  access_token: string;
}

const services = ref<ServiceCard[]>([]);

async function fetchOauthServices() {
  try {
    const responseAbout = await $fetch<AboutResponse>(
      "http://localhost:8080/about.json"
    );

    responseAbout.server.services.forEach((service) => {
      services.value.push({
        name: service.name,
        image: service.image || "IMG",
      });
    });

    services.value.forEach((service) => {
      service.name =
        service.name.charAt(0).toUpperCase() + service.name.slice(1);
    });
  } catch (error) {
    console.error("Error fetching services:", error);
  }
}

async function onSubmit() {
  try {
    const { access_token }: LoginResponse = await $fetch(
      "http://localhost:8080/api/auth/login",
      {
        method: "POST",
        body: {
          username: username.value,
          password: password.value,
        },
      }
    );

    if (access_token) {
      const tokenCookie = useCookie("access_token");
      tokenCookie.value = access_token;

      navigateTo("/dashboard");
    } else {
      triggerNotification(
        "error",
        "Login failed",
        "Please check your credentials"
      );
    }
  } catch (error) {
    console.error("Error logging in:", error);
    triggerNotification(
      "error",
      "Login failed",
      "Please check your credentials"
    );
  }
}

interface RedirectResponse {
  service_authentication_url: string;
}

async function redirectToService(index: number) {
  try {
    const selectedService = services.value[index];
    
    const { service_authentication_url }: RedirectResponse = await $fetch(
      `http://localhost:8080/api/${selectedService.name.toLowerCase()}/auth`,
      {
        method: "GET",
      }
    );
    if (service_authentication_url) {
      useCookie("serviceUsedLogin").value = selectedService.name.toLowerCase();
      window.location.href = service_authentication_url;
    } else {
      console.error("${selectedService} authentication URL not found");
    }
  } catch (error) {
    console.error("Error fetching ${selectedService} OAuth URL:", error);
  }
}

onMounted(() => {
  fetchOauthServices();
});
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
          v-for="(service, index) in services"
          :key="index"
          :text="service.name"
          class="w-full"
          bg-color="bg-primaryWhite-500"
          hover-color="hover:bg-secondaryWhite-500"
          text-color="text-fontBlack"
          @click="redirectToService(index)"
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
