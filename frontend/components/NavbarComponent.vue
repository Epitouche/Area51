<script setup>
const token = ref(useCookie("access_token"));
const registered = computed(() => {
  return token.value !== null && token.value !== '' && token.value !== undefined;
})
</script>
<template>
  <div class="bg-secondaryWhite-500 dark:bg-primaryDark-500">
    <nav class="p-4 border-b border-primaryWhite-500 dark:border-secondaryDark-500">
      <div class="container mx-auto flex justify-between items-center text-fontBlack dark:text-fontWhite">
        <div class="text-lg font-bold">
          <NuxtLink to="/">MyLogo</NuxtLink>
        </div>
        <div v-if="registered" class="space-x-9 flex items-center">
          <ThemeSwitch/>
          <NuxtLink to="/">Dashboard</NuxtLink>
          <NuxtLink to="/workflows">Workflow</NuxtLink>
          <NuxtLink to="/services">Services</NuxtLink>
        </div>
        <div v-else class="space-x-9 flex items-center">
          <ThemeSwitch/>
          <NuxtLink to="/login">Login</NuxtLink>
          <NuxtLink to="/register">
            <ButtonComponent
              bg-color="bg-tertiary-500"
              hover-color="hover:bg-purple-600"
              text="Sign up"
            />
          </NuxtLink>
        </div>
      </div>
    </nav>
  </div>
</template>
