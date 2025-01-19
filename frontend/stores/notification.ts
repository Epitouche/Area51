import { defineStore } from 'pinia';

interface Notification {
  id: string;
  title: string;
  message: string;
  type: 'success' | 'error' | 'warning';
}

export const useNotificationStore = defineStore('notification', {
  state: () => ({
    notifications: [] as Notification[],
  }),
  actions: {
    addNotification(notification: Omit<Notification, 'id'>) {
      const id = Date.now().toString();
      this.notifications.push({ id, ...notification });

      setTimeout(() => {
        this.removeNotification(id);
      }, 5000);
    },
    removeNotification(id: string) {
      this.notifications = this.notifications.filter((n) => n.id !== id);
    },
  },
});
