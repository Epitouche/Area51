<script setup lang="ts">
const props = defineProps<{
  id?: string;
  type: string;
  label: string;
  modelValue: string;
  forceDark?: boolean;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const updateValue = (event: Event) => {
  const input = event.target as HTMLInputElement;
  emit("update:modelValue", input.value);
};
</script>

<template>
  <div
    :class="[
      'flex items-center space-x-3 px-4 py-2 bg-transparent rounded-lg',
      props.forceDark ? 'dark' : '',
    ]"
  >
    <i
      class="fas fa-user"
      :class="props.forceDark ? 'text-fontWhite' : 'text-fontBlack dark:text-fontWhite'"
    />
    <div class="flex-1">
      <label
        :for="props.id"
        :class="props.forceDark ? 'text-fontWhite' : 'text-fontBlack dark:text-fontWhite'"
      >
        {{ props.label }}
      </label>
      <input
        :id="props.id"
        :type="props.type"
        :value="props.modelValue"
        aria-label="input"
        :class="[
          'w-full bg-transparent border-b-2 border-primaryDark-500 focus:border-secondaryDark-500 focus:outline-none focus:ring-0',
          props.forceDark
          ? 'text-fontWhite dark:border-fontWhite dark:text-fontWhite dark:focus:border-secondaryWhite-500'
          : 'text-fontBlack dark:text-fontWhite border-fontBlack dark:border-fontWhite',
        ]"
        @input="updateValue"
      >
    </div>
  </div>
</template>
