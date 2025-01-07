<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";
import DropdownComponent from "./DropdownComponent.vue";

const props = defineProps<{
  columns: string[];
  rows: Record<string, string | number | boolean>[];
  modelValue: boolean[];
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: boolean[]): void;
}>();

const checkboxList = ref(props.modelValue);

const headCheckbox = ref(false);

const checkAll = () => {
  headCheckbox.value = !headCheckbox.value;
  checkboxList.value = checkboxList.value.map(() => headCheckbox.value);
  emitCheckboxes();
};

const emitCheckboxes = () => {
  emit("update:modelValue", checkboxList.value);
};

const toggleCheckbox = (index: number) => {
  checkboxList.value[index] = !checkboxList.value[index];
  emitCheckboxes();
};

// Gestion de l'état du dropdown
const activeDropdownIndex = ref<number | null>(null);

const toggleDropdown = (index: number) => {
  activeDropdownIndex.value = activeDropdownIndex.value === index ? null : index;
};

const handleClickOutside = (event: MouseEvent) => {
  const dropdowns = document.querySelectorAll(".dropdown");
  let clickedInside = false;

  dropdowns.forEach((dropdown) => {
    if (dropdown.contains(event.target as Node)) {
      clickedInside = true;
    }
  });

  if (!clickedInside) {
    activeDropdownIndex.value = null;
  }
};

// Ajouter et retirer l'événement global
onMounted(() => {
  window.addEventListener("click", handleClickOutside);
});

onBeforeUnmount(() => {
  window.removeEventListener("click", handleClickOutside);
});
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
            <input type="checkbox" :checked="headCheckbox" @click="checkAll" >
          </th>
          <th
            v-for="column in columns"
            :key="column"
            class="px-6 py-3 text-center text-xs text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            {{ column }}
          </th>
          <th
            class="px-6 py-3 text-center text-xs text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(row, i) in rows"
          :key="i"
          class="odd:bg-secondaryWhite-500 text-center even:bg-bg-primaryWhite-50 dark:odd:bg-primaryDark-500 dark:even:bg-secondaryDark-500"
        >
          <td class="px-6 py-4">
            <input
              type="checkbox"
              :checked="checkboxList[i]"
              @change="toggleCheckbox(i)"
            >
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
          <td class="relative dropdown">
            <Icon
              name="material-symbols:more-vert"
              class="cursor-pointer h-6 w-6 text-fontBlack dark:text-fontWhite"
              @click.stop="toggleDropdown(i)"
            />
            <div
              v-show="activeDropdownIndex === i"
              class="absolute left-1/2 transform -translate-x-1/2 mt-2 w-56 bg-white dark:bg-secondaryDark-500 shadow-lg rounded-lg overflow-hidden z-10"
            >
              <div
                class="flex flex-col divide-y divide-secondaryWhite-700 dark:divide-secondaryDark-700"
              >
                <button
                  v-for="(option, index) in ['Edit', 'Delete', 'View Details']"
                  :key="index"
                  class="text-center px-4 py-2 text-sm font-medium text-fontBlack dark:text-fontWhite hover:bg-accent-100 dark:hover:bg-accent-800 transition duration-300 ease-in-out"
                  @click="() => console.log(`Action ${option} for row ${i}`)"
                >
                  {{ option }}
                </button>
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
