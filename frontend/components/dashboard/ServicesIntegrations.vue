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
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Gmail',
        icon: 'material-symbols:mail-outline',
        connected: false,
        color: 'bg-red-500'
    }
])
</script>

<template>
    <div class="bg-white rounded-xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6">Integrations</h2>
        <div class="space-y-4 max-h-96 overflow-y-auto">
            <div v-for="systeme in systemes"
                :key="systeme.name"
                class="flex items-center justify-between p-3 rounded-lg border border-gray-100 hover:border-brand-purple/30 transition-colors">
                <div class="flex items-center space-x-3">
                    <IconComponent
                        :bgColor=systeme.color
                        textColor="text-white"
                        :icon=systeme.icon />
                    <span class="font-medium text-fontBlack dar:text-fontWhite">
                        {{ systeme.name }}
                    </span>
                </div>
                <span v-if="systeme.connected" class="text-sm text-green-500 bg-green-50 px-3 py-1 rounded-full">Connected</span>
                <span v-else class="text-sm text-red-500 bg-red-50 px-3 py-1 rounded-full">Disconnected</span>
            </div>
        </div>
    </div>
</template>
