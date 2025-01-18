<script setup lang="ts">
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
</script>
<template>
    <div 
        class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6 h-96" 
        aria-label="Recent Activity Section">
        <h2 
            class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6" 
            aria-label="Section Title: Recent Activity">
            Recent Activity
        </h2>
        <div class="space-y-4">
            <div
                v-for="activity in activities"
                :key="activity.name"
                class="flex items-start space-x-3 p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 dark:hover:border-tertiary-800 transition-colors"
                :aria-label="`Activity: ${activity.name}`">
                <div aria-label="Activity Details">
                    <p 
                        className="text-sm font-medium text-fontBlack dark:text-fontWhite" 
                        :aria-label="`Name: ${activity.name}`">
                        {{ activity.name }}
                    </p>
                    <p 
                        className="text-sm text-gray-500 dark:text-gray-200" 
                        :aria-label="`Description: ${activity.description}`">
                        {{ activity.description }}
                    </p>
                    <p 
                        className="text-xs text-gray-400 dark:text-gray-400 mt-1" 
                        :aria-label="`Time: ${activity.time}`">
                        {{ activity.time }}
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>
