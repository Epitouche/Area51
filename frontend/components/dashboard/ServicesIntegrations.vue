<script setup>
const systemes = ref([
    {
        name: 'Github',
        icon: 'line-md:github',
        connected: true,
        color: 'bg-gray-900'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    }
]);

const shouldOverflow = computed(() => systemes.value.length > 4);
</script>

<template>
    <div class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6">Integrations</h2>
        <div :class="['space-y-4', shouldOverflow ? 'overflow-y-auto max-h-96' : 'h-96']">
            <div v-for="systeme in systemes"
                :key="systeme.name"
                class="flex items-center justify-between p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 transition-colors">
                <div class="flex items-center space-x-3">
                    <IconComponent
                        :bgColor=systeme.color
                        textColor="text-white"
                        :icon=systeme.icon />
                    <span class="font-medium text-fontBlack dark:text-fontWhite">
                        {{ systeme.name }}
                    </span>
                </div>
                <span v-if="systeme.connected" class="text-sm text-green-500 bg-green-50 px-3 py-1 rounded-full">Connected</span>
                <span v-else class="text-sm text-red-500 bg-red-50 px-3 py-1 rounded-full">Disconnected</span>
            </div>
        </div>
    </div>
</template>
