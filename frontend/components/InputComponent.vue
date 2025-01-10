<script setup lang="ts">

const props = defineProps<{
  id?: string;
  type: string;
  label: string;
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const updateValue = (event: Event) => {
  const input = event.target as HTMLInputElement;
  emit('update:modelValue', input.value);
};
</script>

<template>
  <div class="flex items-center space-x-3 px-4 py-2 bg-transparent rounded-lg">
    <i class="fas fa-user text-fontBlack dark:text-fontWhite" />
    <div class="flex-1">
      <label :for="props.id" class="text-fontBlack dark:text-fontWhite text-sm">{{ props.label }}</label>
      <input
        :id="props.id"
        :type="props.type"
        :value="props.modelValue"
        class="w-full bg-transparent border-b-2 border-primaryDark-500 focus:border-secondaryDark-500 text-fontBlack dark:border-primaryWhite-500 dark:text-fontWhite focus:outline-none focus:ring-0 dark:focus:border-secondaryWhite-500"
        @input="updateValue"
      >
    </div>
  </div>
</template>
