<script setup lang="ts">
import type { Service } from "~/src/types";

const services = ref<Service[]>([]);
</script>
<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
  >
    <div class="m-20">
      <h1 class="text-6xl font-bold text-fontBlack dark:text-fontWhite">
        Disconnected Services
      </h1>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
    </div>
    <div class="m-20">
      <h1 class="text-5xl font-bold text-fontBlack dark:text-fontWhite">
        Connected Services
      </h1>
    </div>
    <div class="grid grid-cols-3 gap-4 m-20">
      <div
        v-for="(service, index) in services"
        :key="index"
        class="flex flex-col items-center justify-center p-4 bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-lg shadow-lg"
      >
        <h2 class="text-2xl font-bold text-fontBlack dark:text-fontWhite">
          {{ service.name }}
        </h2>
        <div class="flex justify-center mt-4">
          <button
            class="px-4 py-2 bg-primaryBlue-500 dark:bg-secondaryBlue-500 text-fontWhite dark:text-fontWhite rounded-lg shadow-lg ml-4"
          >
            Connect
          </button>
        </div>
      </div>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
    </div>
    <div class="m-20">
      <h1 class="text-5xl font-bold text-fontBlack dark:text-fontWhite">
        Connected Services
      </h1>
    </div>
    <div class="grid grid-cols-3 gap-4 m-20">
      <div
        v-for="(service, index) in services"
        :key="index"
        class="flex flex-col items-center justify-center p-4 bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-lg shadow-lg"
      >
        <h2 class="text-2xl font-bold text-fontBlack dark:text-fontWhite">
          {{ service.name }}
        </h2>
        <div class="flex justify-center mt-4">
          <button
            class="px-4 py-2 bg-primaryBlue-500 dark:bg-secondaryBlue-500 text-fontWhite dark:text-fontWhite rounded-lg shadow-lg ml-4"
          >
            Disconnect
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
