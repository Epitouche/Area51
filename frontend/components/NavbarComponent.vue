<script setup>
import ButtonComponent from "./ButtonComponent.vue";

const token = ref(useCookie("access_token"));
const registered = computed(() => {
  return (
    token.value !== null && token.value !== "" && token.value !== undefined
  );
});

const isModalOpen = ref(false);

async function deleteAccount() {

  try {
    const accessToken = useCookie("access_token");
    const serviceUsedLogin = useCookie("serviceUsedLogin");
  
    
    const response = await $fetch("/api/users/deleteUserData", {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${accessToken.value}`,
        contentType: "application/json",
      },
    });
    
    accessToken.value = null;
    serviceUsedLogin.value = null;

    if (response) {
      navigateTo("/login");
      isModalOpen.value = false;
    } else {
      throw new Error("Failed to delete account");
    }
  } catch (error) {
    console.error(error);
  }
}

</script>
<template>
  <div class="bg-secondaryWhite-500 dark:bg-primaryDark-500" aria-label="Header section">
    <nav class="p-4 border-b border-secondaryDark-100 dark:border-secondaryDark-500" aria-label="Primary navigation">
      <div class="container mx-auto flex justify-between items-center text-fontBlack dark:text-fontWhite" aria-label="Navigation container">
        <div class="flex  gap-4 text-lg font-bold" aria-label="Website logo">
          <NuxtLink v-if="registered" to="/dashboard" aria-label="Dashboard link"><img src="/logo_Area51.png" alt="Logo of the Website 'Area51'" class="h-10 w-auto"></NuxtLink>
          <NuxtLink v-else to="/" aria-label="Homepage link"><img src="/logo_Area51.png" alt="Logo of the Website 'Area51'" class="h-10 w-auto"></NuxtLink>
        </div>
        <div v-if="registered" class="space-x-9 flex items-center" aria-label="User navigation">
          <ThemeSwitch aria-label="Theme switcher: dark/light mode" />
          <NuxtLink to="/dashboard" aria-label="Dashboard page link">Dashboard</NuxtLink>
          <NuxtLink to="/workflows" aria-label="Workflow page link">Workflow</NuxtLink>
          <NuxtLink to="/services" aria-label="Services page link">Services</NuxtLink>
        </div>
        <div v-else class="space-x-9 flex items-center" aria-label="Guest navigation">
          <ThemeSwitch aria-label="Theme switcher: dark/light mode" />
          <NuxtLink to="/login" aria-label="Login page link">Login</NuxtLink>
          <NuxtLink to="/register" aria-label="Registration page link">
            <ButtonComponent
              bg-color="bg-tertiary-500"
              hover-color="hover:bg-purple-600"
              text-color="text-fontWhite"
              text="Sign up"
              aria-label="Sign up button"
            />
          </NuxtLink>
        </div>
        <div v-if="registered" class="flex gap-4">
          <ButtonComponent
          v-if="registered"
          bg-color="bg-tertiary-500"
          hover-color="hover:bg-accent-500"
          text-color="text-fontWhite"
          text="Logout"
            @click="
              () => {
                const accessToken = useCookie('access_token');
                const serviceUsedLogin = useCookie('serviceUsedLogin');
  
                accessToken.value = null;
                serviceUsedLogin.value = null;
  
                navigateTo('/login');
              }
            "
          />
              <ButtonComponent
                  v-if="registered"
                  bg-color="bg-tertiary-500"
                  hover-color="hover:bg-accent-500"
                  text-color="text-fontWhite"
                  text="Delete Account"
                  @click="isModalOpen = true"
                />
        </div>
      </div>
    </nav>
    <ModalComponent
      v-motion-pop
      title="Delete Account"
      :is-open="isModalOpen && registered"
      @close="isModalOpen = false"
      @confirm="deleteAccount"
    >
      <p class="text-fontBlack dark:text-fontWhite text-lg text-center"
      >Are you sure you want to delete your account ? All your data will be lost.</p>
    </ModalComponent>
  </div>
</template>
