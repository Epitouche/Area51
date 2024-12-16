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
      class="absolute left-1/2 transform -translate-x-1/2 mt-2 w-56 rounded-md shadow-lg bg-primaryWhite-500 dark:bg-primaryDark-500 z-10"
    >
      <div class="flex flex-col items-center gap-2">
        <ButtonComponent
          v-for="(option, index) in props.options"
          :key="index"
          :text="option.name"
          bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
          @click="() => selectOption(option)"
        />
      </div>
    </div>
  </div>
</template>
