<script setup>
// Three last activities
const workflows = ref([
    {
        name: 'GitHub Issue Notifications',
        lastRun: '2 minutes ago',
        executions: 156
    },
    {
        name: 'New Spotify song',
        lastRun: '15 minutes ago',
        executions: 45
    },
    {
        name: 'Daily Report Generator',
        lastRun: '1 hour ago',
        executions: 89
    }
])
</script>
<template>
    <div class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6">Your Workflow</h2>
        <div class="space-y-4">
            <div 
                v-for="workflow in workflows"
                :key="workflow.name"
                class="flex items-center justify-between p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 transition-colors">
                <div className="flex items-center space-x-4">
                    <div>
                        <h3 className="font-medium text-fontBlack dark:text-fontWhite">{{ workflow.name }}</h3>
                        <p className="text-sm text-gray-500 dark:text-gray-200">
                            Last run: {{ workflow.lastRun }}
                        </p>
                    </div>
                </div>
                <div className="flex items-center space-x-4">
                    <span className="text-sm text-gray-500 dark:text-gray-400">
                        {{ workflow.executions }} runs
                    </span>
                </div>
            </div>
        </div>
    </div>
</template>
