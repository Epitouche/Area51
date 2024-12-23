<script setup lang="ts">
import type { ServerResponse, Service } from "~/src/types";

const services = reactive<Service[]>([]);

// Fetch services from the API
onMounted(async () => {
  const response = await $fetch<ServerResponse>(
    "http://localhost:8080/about.json"
  );
  response.server.services.forEach((service: Service) => {
    services.push(service);
  });

  // add fake services for testing
  services.push({ name: "Service 1", actions: [], reactions: [] });
  services.push({ name: "Service 2", actions: [], reactions: [] });
  services.push({ name: "Service 3", actions: [], reactions: [] });
  services.push({ name: "Service 4", actions: [], reactions: [] });
  services.push({ name: "Service 5", actions: [], reactions: [] });
  services.push({ name: "Service 6", actions: [], reactions: [] });
  services.push({ name: "Service 7", actions: [], reactions: [] });

  services.forEach((service) => {
    service.name = service.name.charAt(0).toUpperCase() + service.name.slice(1);
  });
});
</script>
<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
  >
    <div class="m-20">
      <h1 class="text-6xl font-bold text-fontBlack dark:text-fontWhite">
        Services
      </h1>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
    </div>
    <div class="grid grid-cols-3 gap-4 m-20">
      <div
        v-for="(service, index) in services"
        :key="index"
        class="flex justify-center"
      >
        <div
          class="flex flex-col w-full p-7 bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-lg shadow-lg gap-5"
        >
          <div class="flex items-center justify-between w-full">
            <div
              class="w-12 h-12 bg-primaryWhite-400 dark:bg-secondaryDark-400 rounded-full flex items-center justify-center"
            >
              <p>IMG</p>
            </div>
            <div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" value="" class="sr-only peer" checked>
                <div
                  class="w-11 h-6 bg-primaryWhite-500 dark:bg-secondaryDark-400 rounded-full peer peer-checked:bg-tertiary-500 peer-checked:after:translate-x-5 peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"
                />
              </label>
            </div>
          </div>
          <div class="flex flex-col gap-2">
            <h3
              class="text-2xl font-semibold text-fontBlack dark:text-fontWhite"
            >
              {{ service.name }}
            </h3>
            <p class="text-fontBlack dark:text-fontWhite">
              Description of the service. This text has to be long enough to be able to see the overflow of the text.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
