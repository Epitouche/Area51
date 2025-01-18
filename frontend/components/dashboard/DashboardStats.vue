<script setup lang="ts">

import type { Workflow } from "~/src/types";

const stats = ref([
    {
        name: 'Active Workflows',
        value: 0,
        icon: 'mynaui:activity-solid',
        color: 'bg-amber-500'
    },
    {
        name: 'Total Executions',
        value: 12,
        icon: 'jam:thunder',
        color: 'bg-tertiary-500'
    },
    {
        name: 'Last 24h Executions',
        value: 4,
        icon: 'ic:round-loop',
        color: 'bg-green-500'
    },
    {
        name: 'Last Executions',
        value: '2m ago',
        icon: 'solar:history-bold',
        color: 'bg-blue-500'
    }
])

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
    response.forEach(() => {
        stats.value[0].value = Number(stats.value[0].value) + 1;
    })
    // const reaction = await $fetch<Reaction[]>(
    //   "http://localhost:8080/api/workflow/reaction",
    //   {
    //     method: "GET",
    //     headers: {
    //       Authorization: `Bearer ${token.value}`,
    //       "Content-Type": "application/json",
    //     },
    //   }
    // );
})
</script>
<template>
    <div 
        v-for="stat in stats" 
        :key="stat.name" 
        class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6 hover:shadow-md transition-shadow" 
        :aria-label="`Statistic for ${stat.name}`">
        <div class="flex items-center" aria-label="Statistic details">
            <IconComponent
                :bg-color="stat.color"
                text-color="text-white"
                :icon="stat.icon"
                :aria-label="`${stat.name} icon`" />
            <div class="ml-4" aria-label="Statistic text">
                <p class="text-sm font-medium text-gray-600 dark:text-gray-400" :aria-label="`Name: ${stat.name}`">{{ stat.name }}</p>
                <p class="text-2xl font-semibold text-fontBlack dark:text-fontWhite" :aria-label="`Value: ${stat.value}`">{{ stat.value }}</p>
            </div>
        </div>
    </div>
</template>
