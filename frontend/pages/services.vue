<script setup lang="ts">
import type { ServerResponse, Service } from "~/src/types";

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

const actionSelected = ref("");
const reactionSelected = ref("");
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

onMounted(fetchServices);
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
      >
    </div>
    <div class="flex flex-col justify-center m-16 gap-10">
      <div class="flex justify-center gap-5">
        <ButtonComponent
          :text="actionSelected ? actionSelected : 'Choose an action'"
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
                :options="service.actions.map((action) => action.name)"
              />
            </div>
          </div>
        </ModalComponent>
        <ButtonComponent
          :text="reactionSelected ? reactionSelected : 'Choose a reaction'"
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
                :options="service.reactions.map((reaction) => reaction.name)"
              />
            </div>
          </div>
        </ModalComponent>
      </div>
      <div class="flex justify-center">
        <ButtonComponent
          :class="actionSelected && reactionSelected ? '' : 'cursor-not-allowed opacity-50'"
          text="Add Workflow"
          :bg-color="actionSelected && reactionSelected ? 'bg-tertiary-500' : 'bg-primaryWhite-500 dark:bg-secondaryDark-500'"
          hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
          text-color="text-fontBlack dark:text-fontWhite"
        />
      </div>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
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
  </div>
</template>
