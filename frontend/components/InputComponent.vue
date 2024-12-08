<script setup lang="ts">
const props = defineProps<{
  id: string;
  type: string;
  label: string;
}>();
</script>

<template>
  <div class="flex items-center space-x-3 px-4 py-2 bg-transparent rounded-lg">
    <i class="fas fa-user text-white"></i>
    <div class="flex-1">
      <label :for="props.id" class="text-white text-sm">{{
        props.label
      }}</label>
      <input
        :id="props.id"
        :type="props.type"
        class="w-full bg-transparent border-b-2 border-primaryWhite-500 text-fontWhite focus:outline-none focus:ring-0 focus:border-secondaryWhite-500"
      />
    </div>
  </div>
</template>
