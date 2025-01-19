<script setup lang="ts">
import { ref } from "vue";
import { useNotificationStore } from "@/stores/notification";
import type { AboutResponse } from "~/src/types";

type ServiceCard = {
  name: string;
  image: string;
};

const username = ref<string>("");
const email = ref<string>("");
const password = ref<string>("");

const services = ref<ServiceCard[]>([]);

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

interface RegisterResponse {
  access_token: string;
}

async function onSubmit() {
  try {
    const { access_token }: RegisterResponse = await $fetch(
      "http://localhost:8080/api/auth/register",
      {
        method: "POST",
        body: {
          username: username.value,
          email: email.value,
          password: password.value,
        },
      }
    );

    if (access_token) {
      const tokenCookie = useCookie("access_token");
      tokenCookie.value = access_token;

      triggerNotification(
        "success",
        "Registration successful",
        "You have successfully registered."
      );
      navigateTo("/login");
    } else {
      triggerNotification(
        "error",
        "Registration failed",
        "Please check your credentials."
      );
    }
  } catch (error) {
    console.error("Error registering user:", error);
    triggerNotification(
      "error",
      "Registration failed",
      "An error occurred during registration."
    );
  }
}

async function fetchOauthServices() {
  try {
    const responseAbout = await $fetch<AboutResponse>(
      "http://localhost:8080/about.json"
    );

    responseAbout.server.services.forEach((service) => {
      if (service.is_oauth) {
        services.value.push({
          name: service.name,
          image: service.image || "IMG",
        });
      }
    });

    services.value.forEach((service) => {
      service.name =
        service.name.charAt(0).toUpperCase() + service.name.slice(1);
    });
  } catch (error) {
    console.error("Error fetching services:", error);
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

onMounted(fetchOauthServices);
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-primaryDark-500"
    aria-label="Registration screen"
  >
    <div
      class="w-full transform max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
      aria-label="Registration form container"
    >
      <h2 
        class="text-2xl font-bold text-center"
        aria-label="Registration heading"
      >
        REGISTER
      </h2>
      <form 
        class="space-y-6"
        aria-label="Registration form"
        @submit.prevent="onSubmit"
      >
        <div>
          <InputComponent
            id="username"
            v-model="username"
            type="text"
            label="Username"
            aria-label="Username input field"
          />
        </div>
        <div>
          <InputComponent
            id="email"
            v-model="email"
            type="email"
            label="Email"
            aria-label="Email input field"
          />
        </div>
        <div>
          <InputComponent
            id="password"
            v-model="password"
            type="password"
            label="Password"
            aria-label="Password input field"
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
            text="Create an account"
            bg-color="bg-primaryWhite-500"
            hover-color="hover:bg-secondaryWhite-500"
            text-color="text-fontBlack"
            aria-label="Submit registration form"
          />
        </div>
      </form>
      <hr class="border-primaryWhite-400" aria-hidden="true">
      <div class="flex justify-around space-x-4" aria-label="Third-party registration buttons">
        <ButtonComponent
          v-for="(service, index) in services"
          :key="index"
          :text="service.name"
          class="w-full"
          bg-color="bg-primaryWhite-500"
          hover-color="hover:bg-secondaryWhite-500"
          text-color="text-fontBlack"
          aria-label="Register with {{ service.name }}"
          @click="redirectToService(index)"
        />
      </div>
      <div class="flex justify-around">
        <p class="text-center text-sm">
          <NuxtLink 
            to="/login" 
            class="text-fontWhite underline"
            aria-label="Navigate to login page"
          >
            Already have an account?
          </NuxtLink>
        </p>
        <p class="text-center text-sm">
          <a href="#" class="text-fontWhite underline" aria-label="Navigate to forgot password page">
            Forgot password?
          </a>
        </p>
      </div>
    </div>
  </div>
</template>
