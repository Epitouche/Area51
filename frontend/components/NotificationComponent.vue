<script setup lang="ts">
import { useNotificationStore } from '@/stores/notification';

const notificationStore = useNotificationStore();

</script>

<template>
  <div class="fixed bottom-5 right-5 flex flex-col gap-4 z-50">
    <div
      v-for="notification in notificationStore.notifications"
      :key="notification.id"
      v-motion-pop
      :class="[
        'flex items-center gap-2 p-4 rounded-md shadow-md transition-transform transform bg-opacity-90 text-fontWhite',
        notification.type === 'success' ? 'bg-success'
        : notification.type === 'error' ? 'bg-error'
        : 'bg-warning',
      ]"
    >
      <Icon
        :name="notification.type === 'success' ? 'material-symbols:check-rounded'
              : notification.type === 'error' ? 'material-symbols:error-rounded'
              : 'material-symbols:warning-rounded'"
        class="w-10 h-10"
      />
      <div>
        <p class="font-semibold">{{ notification.title }}</p>
        <p class="text-sm">{{ notification.message }}</p>
      </div>
    </div>
  </div>
</template>
