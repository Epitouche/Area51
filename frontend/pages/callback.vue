<script setup lang="ts">
import { useRoute } from "vue-router";
import { useCookie } from "#app";

const access_token = useCookie("access_token");

interface ApiResponse {
  access_token?: string;
}

async function fetchServiceToken(service: string) {
  const route = useRoute();

  const code = route.query.code;
  const state = route.query.state;
  if (code && state) {
    try {
      const response = (await Promise.race([
        $fetch<ApiResponse>("/api/auth/callback", {
          method: "POST",
          body: {
            service: service,
            code: code as string,
            state: state as string,
            authorization: access_token.value
              ? `Bearer ${access_token.value}`
              : "",
          },
        }),
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error("API request timed out")), 5000)
        ),
      ])) as ApiResponse;

      if (access_token) {
        access_token.value = response.access_token;

      } else {
        console.error("Token not received in API response");
      }
    } catch (error) {
      console.error("Error during Service callback:", error);
    }
  } else {
    console.error("Required parameters (code or state) are missing in URL");
  }
}

const loginWithService = async () => {
  const serviceUsedLogin = useCookie("serviceUsedLogin");
  if (serviceUsedLogin.value) {
    await fetchServiceToken(serviceUsedLogin.value as string);
  } else {
    console.error("Service used for login is not defined");
  }
  navigateTo("/dashboard");
};

onMounted(() => {
  if (localStorage.getItem("serviceConnect")) {
    fetchServiceToken(localStorage.getItem("serviceConnect") as string);
    localStorage.removeItem("serviceConnect");
    navigateTo("/services");
  } else {
  loginWithService();
  }
});
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-gradient-to-br from-tertiary-600 to-tertiary-800">
    <div class="text-center">
      <p class="text-fontWhite text-lg md:text-2xl font-semibold animate-pulse">
        Processing Service Login...
      </p>
      <div class="mt-4 flex justify-center">
        <div class="w-8 h-8 border-4 border-white border-t-transparent rounded-full animate-spin" />
      </div>
    </div>
  </div>
</template>