<script setup lang="ts">
import type { AboutResponse } from "~/src/types";

type ServiceCard = {
  name: string;
  description: string;
  image: string;
  isConnected: boolean;
};

const allServices = reactive<ServiceCard[]>([]);

const token = useCookie("access_token");

const isConnected = computed(() => {
  return token.value !== undefined;
});

onMounted(async () => {
  const response = await $fetch<AboutResponse>(
    "http://localhost:8080/about.json"
  );

  response.server.services.forEach((service) => {
    allServices.push({
      name: service.name,
      description:
        service.description || "No description available for this service.",
      image: service.image || "IMG",
      isConnected: false,
    });
  });

  const connectedServices = await $fetch<ServiceCard[]>(
    "http://localhost:8080/api/user/services",
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }
  );

  allServices.forEach((service) => {
    connectedServices.forEach((connectedService) => {
      if (service.name === connectedService.name) {
        service.isConnected = true;
      }
    });
  });

  allServices.forEach((service) => {
    service.name = service.name.charAt(0).toUpperCase() + service.name.slice(1);
  });
});
</script>
<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
  >
    <NuxtLayout />
    <div v-if="isConnected">
      <div class="m-5 sm:m-10">
        <h1
          class="text-3xl sm:text-4xl md:text-6xl font-bold text-fontBlack dark:text-fontWhite"
        >
          Services
        </h1>
      </div>
      <div class="flex justify-center">
        <hr
          class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-full sm:w-11/12"
        >
      </div>
      <div
        class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 m-5 sm:m-10"
      >
        <div
          v-for="(service, index) in allServices"
          :key="index"
          class="flex justify-center"
        >
          <div
            class="flex flex-col w-full p-5 sm:p-7 bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-lg shadow-lg gap-4 sm:gap-5"
          >
            <div class="flex items-center justify-between w-full">
              <div
                class="w-10 h-10 sm:w-12 sm:h-12 bg-primaryWhite-400 dark:bg-secondaryDark-400 rounded-full flex items-center justify-center"
              >
                <img
                  v-if="service.image !== 'IMG'"
                  :src="service.image"
                  alt="service image"
                >
                <p
                  v-else
                  class="text-lg sm:text-xl text-fontBlack dark:text-fontWhite"
                >
                  {{ service.name.charAt(0) }}
                </p>
              </div>
              <div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input
                    type="checkbox"
                    :checked="service.isConnected"
                    class="sr-only peer"
                  >
                  <div
                    class="w-10 h-5 sm:w-11 sm:h-6 bg-secondaryWhite-500 dark:bg-secondaryDark-400 rounded-full peer peer-checked:bg-tertiary-500 peer-checked:after:translate-x-4 sm:peer-checked:after:translate-x-5 peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 sm:after:h-5 after:w-4 sm:after:w-5 after:transition-all"
                  />
                </label>
              </div>
            </div>
            <div class="flex flex-col gap-1 sm:gap-2">
              <h3
                class="text-lg sm:text-xl md:text-2xl font-semibold text-fontBlack dark:text-fontWhite"
              >
                {{ service.name }}
              </h3>
              <p
                class="text-sm sm:text-base text-fontBlack dark:text-fontWhite"
              >
                {{ service.description }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <div class="flex flex-col gap-4 justify-center items-center h-full">
        <h1
          class="text-3xl sm:text-4xl md:text-6xl font-bold text-fontBlack dark:text-fontWhite"
        >
          ERROR 404 !
        </h1>
        <h2 class="text-2xl sm:text-3xl font-bold text-fontBlack dark:text-fontWhite">
          You are not connected, please log in to access this page.
        </h2>
      </div>
    </div>
  </div>
</template>
