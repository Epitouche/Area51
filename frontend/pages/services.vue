<script setup lang="ts">
import type {
  Action,
  Reaction,
  ServerResponse,
  Service,
  WorkflowResponse,
  Workflow,
} from "~/src/types";

const columns = [
  "Name",
  "Action ID",
  "Reaction ID",
  "Activity",
  "Creation Date",
];

const services = reactive<Service[]>([]);
const workflowsInList = reactive<Workflow[]>([]);
const lastWorkflow = reactive<WorkflowResponse[]>([]);
const actionSelected = ref(<Action>{});
const reactionSelected = ref(<Reaction>{});
const isModalActionOpen = ref(false);
const isModalReactionOpen = ref(false);
const token = useCookie("access_token");
const copyIcon = ref("material-symbols:content-copy-outline-rounded");

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    copyIcon.value = "material-symbols:check-rounded";
  } catch (err) {
    console.error("Erreur lors de la copie :", err);
  }
};

const openModalAction = () => {
  isModalActionOpen.value = true;
};

const closeModalAction = () => {
  isModalActionOpen.value = false;
};

const confirmModalAction = () => {
  closeModalAction();
};

const openModalReaction = () => {
  isModalReactionOpen.value = true;
};

const closeModalReaction = () => {
  isModalReactionOpen.value = false;
};

const confirmModalReaction = () => {
  closeModalReaction();
};

async function fetchServices() {
  try {
    const response = await $fetch<ServerResponse>(
      "http://localhost:8080/about.json",
      {
        method: "GET",
      }
    );

    response.server.services.forEach((service: Service) => {
      services.push(service);
    });

    workflowsInList.length = 0;
    workflowsInList.push(...response.server.workflows);

    workflowsInList.forEach((workflow) => {
      const dateString = workflow.created_at;
      const date = new Date(dateString);
      const formattedDate = date.toLocaleDateString("en-GB");
      workflow.created_at = formattedDate;
    });
  } catch (error) {
    console.error("Error fetching services:", error);
  }
}
async function addWorkflow() {
  try {
    await $fetch<ServerResponse>("/api/workflows", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token.value}`,
        "Content-Type": "application/json",
      },
      body: {
        action_id: actionSelected.value.action_id,
        reaction_id: reactionSelected.value.reaction_id,
      },
    });
    fetchServices();
  } catch (error) {
    console.error("Error adding workflow:", error);
  }
}

async function getLastWorkflow() {
  try {
    const response = await $fetch<WorkflowResponse[]>(
      "http://localhost:8080/api/workflow/reaction",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );
    lastWorkflow.push(...response);
  } catch (error) {
    console.error("Error getting last workflow:", error);
  }
}

onMounted(() => {
  fetchServices();
  getLastWorkflow();
});
</script>

<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
  >
    <div class="m-20">
      <h1 class="text-6xl font-bold text-fontBlack dark:text-fontWhite">
        Services
      </h1>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      />
    </div>
    <div class="flex flex-col justify-center m-16 gap-10">
      <div class="flex justify-center gap-5">
        <ButtonComponent
          :text="actionSelected.name ? actionSelected.name : 'Choose an action'"
          bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
          :on-click="openModalAction"
        />
        <ModalComponent
          title="Choose an action"
          :is-open="isModalActionOpen"
          @close="closeModalAction"
          @confirm="confirmModalAction"
        >
          <div class="grid grid-cols-3 gap-4">
            <div
              v-for="(service, index) in services"
              :key="index"
              class="flex justify-center"
            >
              <DropdownComponent
                v-if="service.actions"
                v-model="actionSelected"
                :label="service.name"
                :options="service.actions.map((action) => action)"
              />
            </div>
          </div>
        </ModalComponent>
        <ButtonComponent
          :text="
            reactionSelected.name ? reactionSelected.name : 'Choose a reaction'
          "
          bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
          :on-click="openModalReaction"
        />
        <ModalComponent
          title="Choose a reaction"
          :is-open="isModalReactionOpen"
          @close="closeModalReaction"
          @confirm="confirmModalReaction"
        >
          <div class="grid grid-cols-3 gap-4">
            <div
              v-for="(service, index) in services"
              :key="index"
              class="flex justify-center"
            >
              <DropdownComponent
                v-if="service.reactions"
                v-model="reactionSelected"
                :label="service.name"
                :options="service.reactions.map((reaction) => reaction)"
              />
            </div>
          </div>
        </ModalComponent>
      </div>
      <div class="flex justify-center">
        <ButtonComponent
          :class="
            actionSelected && reactionSelected
              ? ''
              : 'cursor-not-allowed opacity-50'
          "
          text="Add Workflow"
          :bg-color="
            actionSelected && reactionSelected
              ? 'bg-tertiary-500'
              : 'bg-primaryWhite-500 dark:bg-secondaryDark-500'
          "
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
          :on-click="addWorkflow"
        />
      </div>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      />
    </div>
    <div class="flex justify-start gap-5 m-20">
      <ButtonComponent
        text="Filter"
        bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
        text-color="text-fontBlack dark:text-fontWhite"
      />
      <ButtonComponent
        text="All Status"
        bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
        text-color="text-fontBlack dark:text-fontWhite"
      />
    </div>
    <ListTableComponent
      v-show="columns && workflowsInList"
      :columns="columns"
      :rows="workflowsInList"
    />
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      />
    </div>
    <div class="flex justify-center m-20">
      <div
        class="relative flex justify-center bg-primaryWhite-100 dark:bg-secondaryDark-500 rounded-2xl w-10/12"
      >
        <!-- Bouton pour copier le JSON, turn accent after copy -->
        <button
          @click="copyToClipboard(JSON.stringify(lastWorkflow, null, 2))"
          class="absolute top-4 right-4 text-fontBlack dark:text-fontWhite hover:text-accent-200 dark:hover:text-accent-500 transition duration-200"
          aria-label="Copier le JSON"
        >
          <Icon :name="copyIcon" />
        </button>
        <pre
          class="whitespace-pre-wrap break-words text-sm text-primaryWhite-800 dark:text-primaryWhite-200 p-4"
        >
    {{ JSON.stringify(lastWorkflow, null, 2) }}
  </pre
        >
      </div>
    </div>
  </div>
</template>
