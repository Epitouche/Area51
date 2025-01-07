import { useNotificationStore } from '@/stores/notification';

export default defineNuxtPlugin(() => {
  return {
    provide: {
      notify: (type: 'success' | 'error' | 'warning', title: string, message: string) => {
        const notificationStore = useNotificationStore();
        notificationStore.addNotification({ type, title, message });
      },
    },
  };
});
