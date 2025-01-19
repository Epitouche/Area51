<script setup lang="ts">
type ServiceCard = {
  name: string;
  description: string;
  image: string;
  isConnected: boolean;
};

const systemes = ref([
    {
        name: 'Github',
        icon: 'line-md:github',
        connected: false,
        color: 'bg-gray-900'
    },
    {
        name: 'Google',
        icon: 'mdi:google',
        connected: false,
        color: 'bg-red-500'
    },
    {
        name: 'Spotify',
        icon: 'line-md:spotify',
        connected: false,
        color: 'bg-emerald-500'
    },
    {
        name: 'Microsoft',
        icon: 'mdi:microsoft',
        connected: false,
        color: 'bg-blue-500'
    }
]);

const token = useCookie("access_token");

onMounted(async () => {
    const connectedServices = await $fetch<ServiceCard[]>(
        "http://localhost:8080/api/user/services",
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token.value}`,
            },
        }
    );
    connectedServices.forEach((service) => {
        service.name = service.name.charAt(0).toUpperCase() + service.name.slice(1);
    });
    systemes.value.forEach((service) => {
        connectedServices.forEach((connectedService) => {
            if (service.name === connectedService.name) {
                service.connected = true
            }
        })
    })
})

const shouldOverflow = computed(() => systemes.value.length > 4);
</script>

<template>
    <div 
        class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-sm p-6" 
        aria-label="Integrations Section">
        <h2 
            class="text-lg font-semibold text-fontBlack dark:text-fontWhite mb-6" 
            aria-label="Section Title: Integrations">
            Integrations
        </h2>
        <div :class="['space-y-4', shouldOverflow ? 'overflow-y-auto max-h-96' : 'h-96']">
            <div 
                v-for="systeme in systemes"
                :key="systeme.name"
                class="flex items-center justify-between p-3 rounded-lg border border-gray-100 dark:border-gray-900 hover:border-purple-200 dark:hover:border-tertiary-800 transition-colors"
                :aria-label="`Integration: ${systeme.name}`">
                <div 
                    class="flex items-center space-x-3"
                    :aria-label="`Integration Details for: ${systeme.name}`">
                    <IconComponent
                        :bg-color=systeme.color
                        text-color="text-white"
                        :icon=systeme.icon 
                        :aria-label="`Icon for: ${systeme.name}`" />
                    <span 
                        class="font-medium text-fontBlack dark:text-fontWhite"
                        :aria-label="`Integration Name: ${systeme.name}`">
                        {{ systeme.name }}
                    </span>
                </div>
                <span 
                    v-if="systeme.connected" 
                    class="text-sm text-green-500 bg-green-50 px-3 py-1 rounded-full" 
                    aria-label="Status: Connected">
                    Connected
                </span>
                <span 
                    v-else 
                    class="text-sm text-red-500 bg-red-50 px-3 py-1 rounded-full" 
                    aria-label="Status: Disconnected">
                    Disconnected
                </span>
            </div>
        </div>
    </div>
</template>
