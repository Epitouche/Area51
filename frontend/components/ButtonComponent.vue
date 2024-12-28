<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  text?: string;
  bgColor?: string;
  textColor?: string;
  hoverColor?: string;
  icon?: string;
  onClick?: () => void;
}>();

const defaultBgColor = "bg-primaryWhite-500 dark:bg-primaryDark-500";
const defaultTextColor = "text-fontBlack dark:text-fontWhite";
const defaultHoverColor =
  "hover:bg-secondaryWhite-500 dark:hover:bg-secondaryDark-500";

const computedBgColor = computed(() => props.bgColor ?? defaultBgColor);
const computedTextColor = computed(() => props.textColor ?? defaultTextColor);
const computedHoverColor = computed(
  () => props.hoverColor ?? defaultHoverColor
);
</script>

<template>
  <button
    :class="[
      'rounded-full px-4 py-2 font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2 flex flex-row gap-2 items-center',
      computedBgColor,
      computedTextColor,
      computedHoverColor,
    ]"
    @click="props.onClick"
  >
    <p>{{ props.text }}</p>
    <Icon v-if="props.icon" :name="props.icon" />
  </button>
</template>
