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
      <footer
        v-if="showFooter"
        class="flex justify-end gap-2 px-4 py-2"
      >
        <button
          class="px-4 py-2 text-sm bg-primaryWhite-500 rounded hover:bg-secondaryWhite-500 dark:text-fontWhite dark:bg-primaryDark-500 dark:hover:bg-accent-800"
          @click="closeModal"
        >
          Cancel
        </button>
        <button
          class="px-4 py-2 text-sm bg-tertiary-500 text-fontWhite rounded hover:bg-accent-500"
          @click="confirmAction"
        >
          Confirm
        </button>
      </footer>
    </div>
  </div>
</template>
