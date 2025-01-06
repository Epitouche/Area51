<script setup>
const registered = false
</script>
<template>
  <div class="bg-primaryWhite-500 dark:bg-primaryDark-500">
    <nav v-if="registered" class="p-4 border-b border-gray-200 dark:border-secondaryDark-500">
      <div class="container mx-auto flex justify-between items-center text-fontBlack dark:text-fontWhite">
        <div class="text-lg font-bold">
          <NuxtLink to="/">MyLogo</NuxtLink>
        </div>
        <div class="space-x-9">
          <NuxtLink to="/">Dashboard</NuxtLink>
          <NuxtLink to="/services">Services</NuxtLink>
        </div>
      </div>
    </nav>
    <nav v-else class="p-4 border-b border-gray-200 dark:border-secondaryDark-500">
      <div class="container mx-auto flex justify-between items-center text-fontBlack dark:text-fontWhite">
        <div class="text-lg font-bold">
          <NuxtLink to="/">MyLogo</NuxtLink>
        </div>
        <div class="space-x-9 flex items-center">
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
