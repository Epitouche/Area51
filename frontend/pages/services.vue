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

const actionSelected = ref(<Action>{});
const reactionSelected = ref(<Reaction>{});
const workflowsInList = ref<Workflow[]>([]);
const isModalActionOpen = ref(false);
const isModalReactionOpen = ref(false);

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

const services = ref<Service[]>([]);

const token = useCookie("access_token");

async function fetchServices() {
  try {
    const response = await $fetch<ServerResponse>(
      "http://localhost:8080/about.json",
      {
        method: "GET",
      }
    );
    response.server.services.forEach((service: Service) => {
      services.value.push(service);
    });
    console.log("services", services.value);
    workflowsInList.value = response.server.workflows;
    workflowsInList.value.forEach((workflow) => {
      const dateString = workflow.created_at;
      const date = new Date(dateString);

      const formattedDate = date.toLocaleDateString("en-GB");
      workflow.created_at = formattedDate;
    });
  } catch (error) {
    console.error("Error fetching services:", error);
  }
}

const lastWorkflow = ref<WorkflowResponse[]>([]);

async function addWorkflow() {
  try {
    const response = await $fetch<ServerResponse>(
      "http://localhost:8080/api/workflow",
      {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
        body: {
          action_id: actionSelected.value.action_id,
          reaction_id: reactionSelected.value.reaction_id,
        },
      }
    );
    // Update the list of workflows
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
    lastWorkflow.value = response;
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
    <div class="flex flex-col justify-center m-20 p-5 rounded-xl">
      <p
        v-for="workflow in lastWorkflow"
        class="text-2xl font-bold text-fontBlack dark:text-fontWhite"
      >
        BODY: {{ workflow.body }}
      </p>
    </div>
  </div>
</template>
