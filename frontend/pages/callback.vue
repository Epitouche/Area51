<script setup lang="ts">
import { useRoute } from "vue-router";
import { useCookie } from "#app";

const access_token = useCookie("access_token");

interface ApiResponse {
  access_token?: string;
}

async function fetchGitHubToken() {
  const route = useRoute();

  const code = route.query.code;
  const state = route.query.state;
  if (code && state) {
    try {
      const response = await Promise.race([
        $fetch<ApiResponse>("/api/auth/callback", {
          method: "POST",
          body: {
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
      ]) as ApiResponse;

      if (access_token) {
        access_token.value = response.access_token;

        navigateTo("/workflows");
      } else {
        console.error("Token not received in API response");
      }
    } catch (error) {
      console.error("Error during GitHub callback:", error);
    }
  } else {
    console.error("Required parameters (code or state) are missing in URL");
  }
}

onMounted(fetchGitHubToken);
</script>

<template>
  <div>
    <p>Processing GitHub login...</p>
  </div>
</template>
