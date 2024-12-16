<script setup>
import { useRoute, useRouter } from "vue-router";
import { useCookie } from "#app";

const route = useRoute();
const router = useRouter();

onMounted(async () => {
  const code = route.query.code;
  const state = route.query.state;

  console.log("code", code);

  if (code && state) {
    try {
      const { token } = await $fetch(
        "http://localhost:8080/api/github/auth/callback",
        {
          method: "GET",
          params: { code, state },
        }
      );

      if (token) {
        const tokenCookie = useCookie("token");
        tokenCookie.value = token;

        router.push("/dashboard");
      } else {
        console.error("Token not received");
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
