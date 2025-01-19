<script setup lang="ts">
import type { Reaction } from '~/src/types';

const activities = ref<Reaction[]>([]);

const token = useCookie("access_token");

onMounted(async () => {
    const reactions = await $fetch<Reaction[]>(
      "http://localhost:8080/api/workflow/reactions",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );
    if (reactions != null) {
        reactions.forEach((reaction) => {
            activities.value.push(reaction)
        })
    }
})
</script>
<template>
    <div
        class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6"
        aria-label="Recent Activity Section">
        <h2
            class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6"
            aria-label="Section Title: Recent Activity">
            Recent Activity
        </h2>
        <div v-if="activities.values.length === 0" class="space-y-4 max-h-96 overflow-auto">
            <div
                v-for="activity in activities"
                :key="activity.name"
                class="flex items-start space-x-3 p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 dark:hover:border-tertiary-800 transition-colors"
                :aria-label="`Activity: ${activity.name}`">
                <div aria-label="Activity Details">
                    <pre
              class="whitespace-pre-wrap break-words text-xs sm:text-sm text-primaryWhite-800 dark:text-primaryWhite-200"
              >{{ JSON.stringify(activities, null, 2) }}
      </pre
            >
                    <!-- <p
                        className="text-sm font-medium text-fontBlack dark:text-fontWhite"
                        :aria-label="`Name: ${activity.name}`">
                        {{ activity.name }}
                    </p>
                    <p
                        className="text-sm text-gray-500 dark:text-gray-200"
                        :aria-label="`Description: ${activity.description}`">
                        {{ activity.description }}
                    </p> -->
                </div>
            </div>
        </div>
        <div v-else class="flex flex-col items-center justify-between text-gray-500 dark:text-gray-400" aria-label="No workflows created">
            <p>No Reactions of workflow<br></p>
        </div>
    </div>
</template>
