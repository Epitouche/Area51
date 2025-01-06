<script setup lang="ts">
import type { ServerResponse, Service } from "~/src/types";

type ServiceCard = {
  name: string;
  description: string;
  image: string;
  isConnected: boolean;
};

const allServices = reactive<ServiceCard[]>([]);

const token = useCookie("access_token");

// Fetch services from the API
onMounted(async () => {

  // fetch the services
  const response = await $fetch<ServerResponse>(
    "http://localhost:8080/about.json"
  );

  // create a new array with the services
  response.server.services.forEach((service: Service) => {
    allServices.push({
      name: service.name,
      description:
        "Description of the service. This text has to be long enough to be able to see the overflow of the text.",
      image: "IMG",
      isConnected: false,
    });
  });

  // fetch the connected services
  const connectedServices = await $fetch<ServiceCard[]>(
    "http://localhost:8080/api/user/services",
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    }
  );

  // if the service is connected, set the isConnected property to true
  allServices.forEach((service) => {
    connectedServices.forEach((connectedService) => {
      if (service.name === connectedService.name) {
        service.isConnected = true;
      }
    });
  });

  // capitalize the first letter of the service name
  allServices.forEach((service) => {
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
        v-for="(service, index) in allServices"
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
              <p>{{ service.image }}</p>
            </div>
            <div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input
                  type="checkbox"
                  :checked="service.isConnected"
                  class="sr-only peer"
                >
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
              {{ service.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
