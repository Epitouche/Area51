<script setup>
import { useRoute, useRouter } from "vue-router";
import { useCookie } from "#app";

const route = useRoute();
const router = useRouter();

onMounted(async () => {
  const code = route.query.code;
  const state = route.query.state;

  if (code && state) {
    try {
      const response = await $fetch("/api/github", {
        method: "GET",
        params: { code, state },
      });

      console.log("API Response:", response);

      const access_token = response?.access_token;

      if (access_token) {
        const tokenCookie = useCookie("access_token");
        tokenCookie.value = access_token;

        router.push("/dashboard");
      } else {
        console.error("Token not received in API response");
      }
    } catch (error) {
      console.error("Error during GitHub callback:", error);
    }
  } else {
    console.error("Required parameters (code or state) are missing in URL");
  }
});

</script>

<template>
  <div>
    <p>Processing GitHub login...</p>
  </div>
</template>
