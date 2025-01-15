<script setup lang="ts">
// import type { WorkflowResponse } from '~/src/types';
// type ServiceCard = {
//   name: string;
//   description: string;
//   image: string;
//   isConnected: boolean;
// };

const activities = ref([
    {
        name: 'GitHub Issue Notifications',
        time: '2 minutes ago',
        description: 'Successfully notified team about new issue #123'
    },
    {
        name: 'New Spotify song',
        time: '15 minutes ago',
        description: 'Successfully notified of a new song'
    }
])

// const token = useCookie("access_token");

// onMounted(async () => {
//     const response = await $fetch<WorkflowResponse[]>(
//       "/api/workflows/getLastWorkflow",
//       {
//         method: "GET",
//         headers: {
//           Authorization: `Bearer ${token.value}`,
//           "Content-Type": "application/json",
//         },
//       }
//     );
//     console.log(response)
//     const connectedServices = await $fetch<ServiceCard[]>(
//         "http://localhost:8080/api/user/services",
//         {
//             method: "GET",
//             headers: {
//                 Authorization: `Bearer ${token.value}`,
//             },
//         }
//     );
//     console.log(connectedServices)
// })
</script>
<template>
    <div class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6 h-96">
        <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6">Recent Activity</h2>
        <div class="space-y-4">
            <div
                v-for="activity in activities"
                :key="activity.name"
                class="flex items-start space-x-3 p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 dark:hover:border-tertiary-800 transition-colors">
                <div>
                    <p className="text-sm font-medium text-fontBlack dark:text-fontWhite">
                        {{ activity.name }}
                    </p>
                    <p className="text-sm text-gray-500 dark:text-gray-200">{{ activity.description }}</p>
                    <p className="text-xs text-gray-400 dark:text-gray-400 mt-1">{{ activity.time }}</p>
                </div>
            </div>
        </div>
    </div>
</template>
