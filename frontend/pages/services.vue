<script setup lang="ts">
import type {
  Action,
  Reaction,
  ServerResponse,
  Service,
  Workflow,
} from "~/src/types";

const columns = ["Name", "Action", "Reaction", "Status", "Member ID", "Date"];

const rows = [
  {
    id: 1,
    name: "Service",
    action: "On push",
    reaction: "Send message",
    status: "Active",
    memberId: 1,
    date: "12/02/2024",
  },
  {
    id: 2,
    name: "Service 2",
    action: "On pull request",
    reaction: "Play Music",
    status: "Inactive",
    memberId: 2,
    date: "10/25/2024",
  },
  {
    id: 3,
    name: "Service 3",
    action: "On push",
    reaction: "Fill mail",
    status: "Active",
    memberId: 3,
    date: "10/14/2024",
  },
];

const actionSelected = ref(<Action>{});
const reactionSelected = ref(<Reaction>{});
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

const token = useCookie("token");

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
  } catch (error) {
    console.error("Error fetching services:", error);
  }
}

const lastWorkflow = ref<Workflow[]>([]);

async function addWorkflow() {
  try {
    console.log("actionSelected", actionSelected.value.action_id);
    console.log("reactionSelected", reactionSelected.value.reaction_id);
    const response = await $fetch<ServerResponse>(
      "http://localhost:8080/api/workflow",
      {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
        body: {
          action_id: actionSelected.value.action_id,
          reaction_id: reactionSelected.value.reaction_id,
        },
      }
    );
    console.log(response);
  } catch (error) {
    console.error("Error adding workflow:", error);
  }
}

async function getLastWorkflow() {
  try {
    const response = await $fetch<Workflow[]>(
      "http://localhost:8080/api/workflow/reaction",
      {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );
    lastWorkflow.value = response;
    console.log(response);
  } catch (error) {
    console.error("Error getting last workflow:", error);
  }
}

onMounted(() => {
  fetchServices();
  console.log("token", token);
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
      v-show="columns && rows"
      :columns="columns"
      :rows="rows"
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
