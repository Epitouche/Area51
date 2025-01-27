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
    <div class="flex flex-row items-center justify-center" role="group" aria-label="Theme switcher">
        <Icon
            v-if="isMounted && (colorMode.preference === 'light' || colorMode.preference === 'system')"
            aria-label="Theme switcher: dark/light mode"
            name="mdi:weather-night"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-fontBlack dark:text-fontWhite"
            @click="toggleTheme"
        />
        <Icon
            v-if="isMounted && colorMode.preference === 'dark'"
            aria-label="Theme switcher: dark/light mode"
            name="material-symbols:sunny-outline-rounded"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-fontBlack dark:text-fontWhite"
            @click="toggleTheme"
        />
    </div>
</template>
