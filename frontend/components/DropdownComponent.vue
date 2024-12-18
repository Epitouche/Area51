<script setup lang="ts">
import { ref } from "vue";
import ButtonComponent from "./ButtonComponent.vue";
import type { Action, Reaction } from "~/src/types";

const props = defineProps<{
  options: (Action | Reaction)[];
  label?: string;
  modelValue: Action | Reaction;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: Action | Reaction): void;
}>();

const isOpen = ref(false);

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const selectOption = (option: Action | Reaction) => {
  emit("update:modelValue", option);
  isOpen.value = false;
};
</script>

<template>
  <div class="relative inline-block text-left">
    <ButtonComponent
      :text="props.label"
      bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
      hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
      text-color="text-fontBlack dark:text-fontWhite"
      @click="toggleDropdown"
    />
    <div
      v-show="isOpen"
      class="absolute left-1/2 transform -translate-x-1/2 mt-2 w-56 bg-white dark:bg-secondaryDark-500 shadow-lg rounded-lg overflow-hidden z-10"
    >
      <div class="flex flex-col divide-y divide-secondaryWhite-700 dark:divide-secondaryDark-700">
        <button
          v-for="(option, index) in props.options"
          :key="index"
          class="text-center px-4 py-2 text-sm font-medium text-fontBlack dark:text-fontWhite hover:bg-accent-100 dark:hover:bg-accent-800 transition duration-300 ease-in-out"
          @click="() => selectOption(option)"
        >
          {{ option.name }}
        </button>
      </div>
    </div>
  </div>
</template>
