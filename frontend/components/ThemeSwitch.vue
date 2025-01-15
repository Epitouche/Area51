<script setup>
import { ref, onMounted } from 'vue';
const colorMode = useColorMode();
const isMounted = ref(false);

onMounted(() => {
    isMounted.value = true;
});

const toggleTheme = () => {
    colorMode.preference = colorMode.preference === 'dark' ? 'light' : 'dark';
};
</script>

<template>
    <div class="flex flex-row items-center justify-center">
        <Icon
            v-if="isMounted && (colorMode.preference === 'light' || colorMode.preference === 'system')"
            name="mdi:weather-night"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-typography_primary_light dark:text-typography_primary_dark"
            width="28"
            height="28"
            @click="toggleTheme"
        />
        <Icon
            v-if="isMounted && colorMode.preference === 'dark'"
            name="material-symbols:sunny-outline-rounded"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-typography_primary_light dark:text-typography_primary_dark"
            width="28"
            height="28"
            @click="toggleTheme"
        />
    </div>
</template>
