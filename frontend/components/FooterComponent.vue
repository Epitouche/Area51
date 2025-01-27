<script setup>
import { computed } from 'vue';

const currentYear = computed(() => {
    return new Date().getFullYear();
});
</script>
<template>
    <div class="px-4 sm:px-6 lg:px-8 py-12 bg-secondaryWhite-500 dark:bg-primaryDark-500 border-gray-500" aria-label="Footer section">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8" aria-label="Footer content">
            <div aria-label="About Area51">
                <div class="flex items-center mb-4">
                    <img src="/logo_Area51.png" alt="Logo of the Website 'Area51'" class="h-8 w-auto">
                    <span class="ml-2 text-xl font-bold text-fontBlack dark:text-fontWhite">AREA</span>
                </div>
                <p class="text-fontBlack dark:text-fontWhite text-sm" aria-label="Project description">
                    A project about automating workflows efficiently.
                </p>
            </div>
            <div aria-label="User links">
                <h3 class="font-semibold text-fontBlack dark:text-fontWhite" aria-label="User section title">User</h3>
                <div class="text-fontBlack dark:text-fontWhite hover:text-brand-purple transition-colors mt-4">
                    <NuxtLink to="/login" aria-label="Login page link">Login</NuxtLink>
                </div>
                <div class="text-fontBlack dark:text-fontWhite hover:text-brand-purple transition-colors">
                    <NuxtLink to="/register" aria-label="Register page link">Register</NuxtLink>
                </div>
            </div>
            <div aria-label="Documentation links">
                <h3 class="font-semibold text-fontBlack dark:text-fontWhite" aria-label="Documentation section title">Documentation</h3>
                <div class="text-fontBlack dark:text-fontWhite hover:text-brand-purple mt-4">
                    <a href="https://github.com/Epitouche/Area51/blob/main/README.md" aria-label="GitHub documentation link">Github Documentation</a>
                </div>
                <div class="text-fontBlack dark:text-fontWhite hover:text-brand-purple">
                    <a href="https://drive.google.com/drive/folders/1Z0oZLYy2zBhhryj8Y1aOzdajEbtKuYpq" aria-label="Google Drive documentation link">Google Drive</a>
                </div>
            </div>
        </div>

        <div class="mt-8 pt-8 border-t border-gray-500" aria-label="Footer bottom section">
            <div class="flex flex-col md:flex-row justify-between items-center">
                <p class="text-fontBlack dark:text-fontWhite text-sm" aria-label="Copyright notice">
                    © {{ currentYear }} AREA. All rights reserved.
                </p>
                <div class="flex space-x-6 mt-4 md:mt-0 text-fontBlack dark:text-fontWhite">
                    <a href="https://github.com/Epitouche/Area51" aria-label="GitHub repository link">
                        <Icon name="line-md:github" />
                    </a>
                </div>
            </div>
        </div>
    </div>
</template>
