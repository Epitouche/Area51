<script setup>
import { computed } from 'vue';

const currentYear = computed(() => {
    return new Date().getFullYear();
});
</script>
<template>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div>
                <div class="flex items-center mb-4">
                    <!-- Logo -->
                    <span class="ml-2 text-xl font-bold text-gray-900">AREA</span>
                </div>
                <p class="text-gray-500 text-sm">
                    A project about automating workflows efficiently.
                </p>
            </div>
            <div>
                <h3 class="font-semibold text-gray-900">User</h3>
                <div class="text-gray-500 hover:text-brand-purple transition-colors mt-4"><NuxtLink to="/login">Login</NuxtLink></div>
                <div class="text-gray-500 hover:text-brand-purple transition-colors"><NuxtLink to="/register">Register</NuxtLink></div>
            </div>
            <div>
                <h3 class="font-semibold text-gray-900">Documentation</h3>
                <div class="text-gray-500 hover:text-brand-purple mt-4"><a href="https://github.com/Epitouche/Area51/blob/main/README.md">Github Documentation</a></div>
                <div class="text-gray-500 hover:text-brand-purple"><a href="https://drive.google.com/drive/folders/1Z0oZLYy2zBhhryj8Y1aOzdajEbtKuYpq">Google Drive</a></div>
            </div>
        </div>

        <div class="mt-8 pt-8 border-t border-gray-200">
            <div class="flex flex-col md:flex-row justify-between items-center">
                <p class="text-gray-500 text-sm">
                    © {{ currentYear }} AREA. All rights reserved.
                </p>
                <div class="flex space-x-6 mt-4 md:mt-0">
                    <a href="https://github.com/Epitouche/Area51"><Icon name="line-md:github" /></a>
                </div>
            </div>
        </div>
    </div>
</template>
