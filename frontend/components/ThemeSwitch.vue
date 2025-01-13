<template>
    <div class="flex flex-row items-center justify-center">
        <Icon
            v-show="colorMode.preference === 'light'"
            name="mdi:weather-night"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-fontBlack dark:text-fontWhite"
            @click="toggleTheme"
        />
        <Icon
            v-show="colorMode.preference === 'dark'"
            name="material-symbols:sunny-outline-rounded"
            class="w-7 h-7 cursor-pointer transition-transform duration-100 hover:scale-120 text-fontBlack dark:text-fontWhite"
            @click="toggleTheme"
        />
    </div>
</template>

<script setup>
const colorMode = useColorMode();
const toggleTheme = () => {
    colorMode.preference = colorMode.preference === 'dark' ? 'light' : 'dark';
};
</script>
