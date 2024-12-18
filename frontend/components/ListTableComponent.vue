<script setup lang="ts">
defineProps<{
  columns: string[];
  rows: Record<string, string | number | boolean>[];
}>();
</script>

<template>
  <div class="flex justify-center overflow-x-auto">
    <table
      class="justify-center w-11/12 border-collapse bg-primaryWhite-500 dark:bg-secondaryDark-500"
    >
      <thead>
        <tr class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-full">
          <th
            class="px-6 py-3 text-center text-xs text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            <input type="checkbox">
          </th>
          <th
            v-for="column in columns"
            :key="column"
            class="px-6 py-3 text-center text-xs text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            {{ column }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="row in rows"
          :key="row.id"
          class="odd:bg-secondaryWhite-500 text-center even:bg-bg-primaryWhite-50 dark:odd:bg-primaryDark-500 dark:even:bg-secondaryDark-500"
        >
          <td class="px-6 py-4">
            <input type="checkbox">
          </td>
          <td
            v-for="(value, key) in row"
            :key="key"
            class="px-6 py-4 text-sm"
            :class="
              key === 'is_active'
                ? value
                  ? 'font-bold text-tertiary-500'
                  : 'font-bold text-red-500'
                : 'text-fontBlack dark:text-gray-200'
            "
          >
            {{ key === "is_active" ? (value ? "Active" : "Inactive") : value }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
