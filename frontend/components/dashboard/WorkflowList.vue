<script setup lang="ts">
import type { Workflow } from '~/src/types';

interface ExtendedWorkflow extends Workflow {
  executions: number;
}

const workflows = reactive<ExtendedWorkflow[]>([]);

const token = useCookie("access_token");

onMounted(async () => {
    const response = await $fetch<Workflow[]>(
      "http://localhost:8080/api/user/workflows",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );
    response.forEach((workflow) => {
        const transformedWorkflow: ExtendedWorkflow = {
            ...workflow,
            executions: 0
        };
        workflows.push(transformedWorkflow);
    });
})
</script>
<template>
    <div 
        class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6"
        aria-label="Workflow Section">
        <h2 
            class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6"
            aria-label="Your Workflow">Your Workflow</h2>
        <div v-if="workflows.values.length === 0" class="space-y-4 max-h-96 overflow-auto">
            <div
                v-for="workflow in workflows"
                :key="workflow.name"
                class="flex items-center justify-between p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 dark:hover:border-tertiary-800 transition-colors"
                aria-label="Workflow item: {{ workflow.name }}">
                <div class="flex items-center space-x-4">
                    <div>
                        <h3 
                            class="font-medium text-fontBlack dark:text-fontWhite"
                            aria-label="Workflow name: {{ workflow.name }}">
                            {{ workflow.name }}
                        </h3>
                        <p 
                            class="text-sm text-gray-500 dark:text-gray-200"
                            aria-label="Action: {{ workflow.action_name }}">
                            Action: {{ workflow.action_name }}
                        </p>
                        <p 
                            class="text-sm text-gray-500 dark:text-gray-200"
                            aria-label="Reaction: {{ workflow.reaction_name }}">
                            Reaction: {{ workflow.reaction_name }}
                        </p>
                    </div>
                </div>
            </div>
        </div>
        <div v-else class="flex flex-col items-center justify-between text-gray-500 dark:text-gray-400" aria-label="No workflows created">
            <p>No Workflow created<br></p>
            <p>Checkout the workflow page 
                <NuxtLink to="/workflows" class="text-tertiary-500" aria-label="Go to workflow page">here</NuxtLink> to create one
            </p>
        </div>
    </div>
</template>
