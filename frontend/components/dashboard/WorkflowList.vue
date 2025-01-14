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
                            Action: {{ workflow.action_name }}
                        </p>
                        <p className="text-sm text-gray-500 dark:text-gray-200">
                            Reaction: {{ workflow.reaction_name }}
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
