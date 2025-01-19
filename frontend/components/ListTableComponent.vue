<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";
import type { Workflow } from "~/src/types";
import ModalComponent from "./ModalComponent.vue";

const props = defineProps<{
  columns: string[];
  rows: Workflow[];
  modelValue: Workflow[];
}>();

const modalOpen = ref(false);

const copyIcon = ref("material-symbols:content-copy-outline-rounded");

const workflowReaction = reactive<unknown[]>([]);

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    copyIcon.value = "material-symbols:check-rounded";
  } catch (err) {
    console.error("Erreur lors de la copie :", err);
  }
};

const filteredWorkflows = computed(() =>
  props.rows.map(
    ({ checked, action_id, reaction_id, workflow_id, ...rest }) => rest
  )
);

const emit = defineEmits<{
  (e: "update:modelValue", value: Workflow[]): void;
}>();

const headCheckbox = ref(false);

const checkAll = () => {
  headCheckbox.value = !headCheckbox.value;
  props.rows.forEach((row) => (row.checked = headCheckbox.value));
  emitCheckboxes();
};

const emitCheckboxes = () => {
  emit("update:modelValue", props.rows);
};

const toggleCheckbox = (workflow: Workflow) => {
  workflow.checked = !workflow.checked;
  emitCheckboxes();
};

const activeDropdownIndex = ref<number | null>(null);

const toggleDropdown = (index: number) => {
  activeDropdownIndex.value =
    activeDropdownIndex.value === index ? null : index;
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

const token = useCookie("access_token");

async function launchAction(option: string, workflow: Workflow) {
  switch (option) {
    case "Last Reaction":
      await fetchWorkflowReaction(workflow);
      modalOpen.value = true;
      activeDropdownIndex.value = null;
      break;
    case "Switch Activity":
      await switchState(workflow);
      activeDropdownIndex.value = null;
      break;
    case "Delete":
      await deleteWorkflow(workflow);
      activeDropdownIndex.value = null;
      break;
    default:
      break;
  }
}

async function deleteWorkflow(workflow: Workflow) {
  try {
    await $fetch("/api/workflows/deleteWorkflow", {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token.value}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        workflow_id: workflow.workflow_id,
        name: workflow.name,
        action_id: workflow.action_id,
        reaction_id: workflow.reaction_id,
      }),
    });
    window.location.reload();
  } catch (error) {
    console.error(error);
  }
}

async function fetchWorkflowReaction(workflow: Workflow) {
  try {
    const response = await $fetch("/api/workflows/lastReaction", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token.value}`,
        "Content-Type": "application/json",
      },
      params: {
        workflow_id: workflow.workflow_id,
      },
    });

    if (response !== undefined) {
      workflowReaction.push(response);
    }
  } catch (error) {
    console.error(error);
  }
}

async function switchState(workflow: Workflow) {
  try {
    workflow.is_active = !workflow.is_active;
    await $fetch("/api/workflows/switchState", {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token.value}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        workflow_id: workflow.workflow_id,
        workflow_state: workflow.is_active,
      }),
    });
  } catch (error) {
    workflow.is_active = !workflow.is_active;
    console.error(error);
  }
}

onMounted(() => {
  window.addEventListener("click", handleClickOutside);
});

onBeforeUnmount(() => {
  window.removeEventListener("click", handleClickOutside);
});
</script>

<template>
  <div class="flex justify-center overflow-x-auto p-16">
    <table
      class="w-full sm:w-11/12 border-collapse bg-primaryWhite-500 dark:bg-secondaryDark-500 text-sm sm:text-base"
    >
      <thead>
        <tr class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-t-lg">
          <th
            class="px-3 py-2 sm:px-6 sm:py-3 text-center text-xs sm:text-sm text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            <input type="checkbox" :checked="headCheckbox" @click="checkAll">
          </th>
          <th
            v-for="column in columns"
            :key="column"
            class="px-3 py-2 sm:px-6 sm:py-3 text-center text-xs sm:text-sm text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            {{ column }}
          </th>
          <th
            class="px-3 py-2 sm:px-6 sm:py-3 text-center text-xs sm:text-sm text-fontBlack dark:text-gray-300 uppercase tracking-wider"
          >
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(row, i) in rows"
          :key="i"
          class="odd:bg-secondaryWhite-500 even:bg-bg-primaryWhite-50 dark:odd:bg-primaryDark-500 dark:even:bg-secondaryDark-500 text-center"
        >
          <td class="px-3 py-2 sm:px-6 sm:py-4">
            <input
              type="checkbox"
              :checked="row.checked"
              @change="toggleCheckbox(row)"
            >
          </td>
          <td
            v-for="(value, key) in filteredWorkflows[i]"
            :key="key"
            class="px-3 py-2 sm:px-6 sm:py-4 text-xs sm:text-sm"
          >
            <p
              class="text-fontBlack dark:text-fontWhite p-1 rounded-full"
              :class="
                key === 'is_active'
                  ? value
                    ? 'font-bold text-fontWhite bg-tertiary-500'
                    : 'font-bold bg-error text-fontWhite'
                  : 'text-fontBlack dark:text-gray-200'
              "
            >
              {{
                key === "is_active" ? (value ? "Active" : "Inactive") : value
              }}
            </p>
          </td>
          <td class="relative dropdown">
            <Icon
              name="material-symbols:more-vert"
              class="cursor-pointer h-5 w-5 sm:h-6 sm:w-6 text-fontBlack dark:text-fontWhite"
              @click.stop="toggleDropdown(i)"
            />
            <div
              v-show="activeDropdownIndex === i"
              :class="{
                '-translate-y-40': i >= Math.floor(rows.length / 2),
                '': i < Math.floor(rows.length / 2),
              }"
              class="absolute left-1/2 transform -translate-x-1/2 mt-2 w-28 sm:w-32 bg-white dark:bg-secondaryDark-500 shadow-lg rounded-lg overflow-hidden z-10"
            >
              <div
                class="flex flex-col divide-y divide-secondaryWhite-700 dark:divide-secondaryDark-700"
              >
                <button
                  v-for="(option, index) in [
                    'Last Reaction',
                    'Switch Activity',
                    'Delete',
                  ]"
                  :key="index"
                  class="text-center px-2 sm:px-4 py-2 text-xs sm:text-sm font-medium text-fontBlack dark:text-fontWhite hover:bg-accent-100 dark:hover:bg-accent-800 transition duration-300 ease-in-out"
                  :class="
                    option.includes('Delete')
                      ? 'hover:bg-error text-error hover:text-fontWhite dark:hover:bg-error'
                      : ''
                  "
                  @click="launchAction(option, row)"
                >
                  <p>
                    {{ option }}
                  </p>
                </button>
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
    <ModalComponent
      v-motion-pop
      title="Last Reaction"
      :is-open="modalOpen"
      @close="modalOpen = false"
      @confirm="modalOpen = false"
    >
      <div class="flex justify-center m-5 sm:m-10">
        <div
          class="relative flex justify-center bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-2xl w-full sm:w-10/12"
          aria-label="Workflow result JSON"
        >
          <button
            class="absolute top-2 right-2 sm:top-4 sm:right-4 text-fontBlack dark:text-fontWhite hover:text-accent-200 dark:hover:text-accent-500 transition duration-200"
            aria-label="Copy workflow result JSON"
            @click="copyToClipboard(JSON.stringify(workflowReaction, null, 2))"
          >
            <Icon :name="copyIcon" />
          </button>
          <div
            class="max-h-96 overflow-auto w-full p-4"
            aria-label="Scrollable content"
          >
            <pre
              class="whitespace-pre-wrap break-words text-xs sm:text-sm text-primaryWhite-800 dark:text-primaryWhite-200"
              >{{ JSON.stringify(workflowReaction, null, 2) }}
        </pre
            >
          </div>
        </div>
      </div>
    </ModalComponent>
  </div>
</template>
