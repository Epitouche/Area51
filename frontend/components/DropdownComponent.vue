<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import ButtonComponent from "./ButtonComponent.vue";

const props = defineProps<{
  options: string[];
  label?: string;
  labelKey?: string;
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const isOpen = ref(false);
const dropdownRef = ref<HTMLElement | null>(null);

const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

const selectOption = (option: string) => {
  emit("update:modelValue", option);
  isOpen.value = false;
};

const handleClickOutside = (event: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener("mousedown", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("mousedown", handleClickOutside);
});
</script>
<template>
  <div ref="dropdownRef" class="relative inline-block text-left">
    <ButtonComponent
      :text="props.label || 'Select an option'"
      bg-color="bg-primaryWhite-300 dark:bg-secondaryDark-500"
      hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
      text-color="text-fontBlack dark:text-fontWhite"
      icon="material-symbols:keyboard-arrow-down-rounded"
      @click="toggleDropdown"
    />

    <div
      v-if="isOpen"
      class="absolute left-1/2 transform -translate-x-1/2 mt-2 w-56 bg-secondaryWhite-500 dark:bg-secondaryDark-500 shadow-lg rounded-lg overflow-hidden z-10"
    >
      <div
        class="flex flex-col divide-y divide-secondaryWhite-700 dark:divide-secondaryDark-700"
      >
        <button
          v-for="(option, index) in props.options"
          :key="index"
          class="text-center px-4 py-2 text-sm font-medium text-fontBlack dark:text-fontWhite hover:bg-accent-100 dark:hover:bg-accent-800 transition duration-300 ease-in-out"
          @click="() => selectOption(option)"
        >
          {{ props.labelKey ? option[props.labelKey] : option }}
        </button>
      </div>
    </div>
  </div>
</template>
