<script setup lang="ts">
defineProps({
  isOpen: { type: Boolean, required: true },
  title: { type: String, required: true },
  showFooter: { type: Boolean, default: true },
});

const emit = defineEmits(["close", "confirm"]);

const closeModal = () => {
  emit("close");
};

const confirmAction = () => {
  emit("confirm");
};
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50"
  >
    <div
      class="bg-primaryWhite-500 dark:bg-primaryDark-500 rounded-xl shadow-lg w-full max-w-md"
    >
      <header class="flex justify-between items-center px-4 py-2">
        <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite">
          {{ title }}
        </h2>
        <button
          class="text-fontBlack dark:text-fontWhite text-2xl"
          @click="closeModal"
        >
          &times;
        </button>
      </header>
      <main class="px-4 py-6">
        <slot />
      </main>
      <footer v-if="showFooter" class="flex justify-end gap-2 px-4 py-2">
        <ButtonComponent
          text="Cancel"
          bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
          @click="closeModal" />
        <ButtonComponent
          text="Confirm"
          bg-color="bg-tertiary-500"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontWhite"
          @click="confirmAction"
        />
      </footer>
    </div>
  </div>
</template>
