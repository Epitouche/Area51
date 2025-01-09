<script setup lang="ts">
import { useNotificationStore } from "@/stores/notification";

import type {
  Service,
  WorkflowResponse,
  Workflow,
  AboutResponse
} from "~/src/types";

const notificationStore = useNotificationStore();

function triggerNotification(
  type: "success" | "error" | "warning",
  title: string,
  message: string
) {
  notificationStore.addNotification({
    type,
    title,
    message,
  });
}

const columns = [
  "Name",
  "Action",
  "Reaction",
  "Activity",
  "Creation Date",
];

const filters = ["All Status", "Active", "Inactive", "Selected"];
const sorts = ["Name", "Creation Date", "Action ID", "Reaction ID"];

const services = reactive<Service[]>([]);
const workflowsInList = reactive<Workflow[]>([]);
const lastWorkflowResult = reactive<WorkflowResponse[]>([]);

const actionString = ref("");
const reactionString = ref("");

const selectedFilter = ref("All Status");
const selectedSort = ref("Name");

const isModalActionOpen = ref(false);
const isModalReactionOpen = ref(false);
const token = useCookie("access_token");

const WorkflowName = ref("");

const filteredWorkflows = computed(() => {
  sortWorkflows();
  switch (selectedFilter.value) {
    case "Active":
      return workflowsInList.filter((workflow) => workflow.is_active === true);
    case "Inactive":
      return workflowsInList.filter((workflow) => workflow.is_active === false);
    case "Selected":
      return workflowsInList.filter((workflow) => workflow.checked === true);
    default:
      return workflowsInList;
  }
});

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

const sortWorkflows = () => {
  workflowsInList.sort((a, b) => {
    switch (selectedSort.value) {
      case "Name":
        return a.name.localeCompare(b.name);
      case "Creation Date":
        return a.created_at.localeCompare(b.created_at);
      case "Action ID":
        return a.action_id - b.action_id;
      case "Reaction ID":
        return a.reaction_id - b.reaction_id;
      default:
        return 0;
    }
  });
};

async function fetchServices() {
  try {
    const responseServices = await $fetch<Service[]>(
      "http://localhost:8080/api/user/services",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );


    responseServices.forEach((service: Service) => {
      services.push(service);
    });

    // add actions and reactions to services in about.json fetch
    const responseAbout = await $fetch<AboutResponse>(
      "http://localhost:8080/about.json",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );

    responseAbout.server.services.forEach((service) => {
      const serviceFound = services.find(
        (s) => s.name === service.name
      );

      if (serviceFound) {
        serviceFound.actions = service.actions;
        serviceFound.reactions = service.reactions;
      }
    });

  } catch (error) {
    console.error("Error fetching services:", error);
  }
}

async function fetchWorkflows() {
  try {
    const response = await $fetch<Workflow[]>(
      "http://localhost:8080/api/user/workflows",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );


    workflowsInList.length = 0;
    workflowsInList.push(...response);

    workflowsInList.forEach((workflow) => {
      const dateString = workflow.created_at;
      const date = new Date(dateString);
      const formattedDate = date.toLocaleDateString("en-GB");
      workflow.created_at = formattedDate;
      workflow.checked = false;
    });

    sortWorkflows();
  } catch (error) {
    console.error("Error fetching services:", error);
  }
}

async function addWorkflow() {
  try {
    const actionSelected = services
      .flatMap((service) => service.actions)
      .find((action) => action.name === actionString.value);

    const reactionSelected = services
      .flatMap((service) => service.reactions)
      .find((reaction) => reaction.name === reactionString.value);

    if (actionSelected && reactionSelected) {
      const body: { action_id: number; reaction_id: number; name?: string } = {
        action_id: actionSelected.action_id,
        reaction_id: reactionSelected.reaction_id,
      };

      console.log("WorkflowName", WorkflowName.value);

      if (WorkflowName.value) {
        body.name = WorkflowName.value;
      }

      await $fetch("/api/workflows/addWorkflows", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
        body,
      });

      await fetchWorkflows();

      triggerNotification(
        "success",
        "Workflow added",
        "The workflow has been added successfully"
      );

      actionString.value = "";
      reactionString.value = "";
      WorkflowName.value = "";
    } else {
      triggerNotification(
        "error",
        "Workflow error",
        "Action or reaction not found. Please check your selections."
      );
    }
  } catch (error) {
    console.error("Error adding workflow:", error);
    triggerNotification(
      "error",
      "Error",
      "An error occurred while adding the workflow. Please try again."
    );
  }
}

async function getLastWorkflowResult() {
  try {
    const response = await $fetch<WorkflowResponse[]>(
      "/api/workflows/getLastWorkflow",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token.value}`,
          "Content-Type": "application/json",
        },
      }
    );
    lastWorkflowResult.push(...response);
  } catch (error) {
    console.error("Error getting last workflow:", error);
  }
}

onMounted(() => {
  fetchServices();
  fetchWorkflows();
  getLastWorkflowResult();
});
</script>

<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
  >
    <div class="m-20">
      <h1 class="text-6xl font-bold text-fontBlack dark:text-fontWhite">
        Workflows
      </h1>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
    </div>
    <div class="flex flex-col justify-center m-16 gap-10 items-center">
      <InputComponent
        v-model="WorkflowName"
        type="text"
        label="Workflow Name"
      />
      <div class="flex justify-center gap-5">
        <ButtonComponent
          :text="actionString ? actionString : 'Choose an action'"
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
                v-model="actionString"
                :label="service.name"
                :options="service.actions.map((action) => action.name)"
              />
            </div>
          </div>
        </ModalComponent>
        <ButtonComponent
          :text="reactionString ? reactionString : 'Choose a reaction'"
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
                v-model="reactionString"
                :label="service.name"
                :options="service.reactions.map((reaction) => reaction.name)"
              />
            </div>
          </div>
        </ModalComponent>
      </div>
      <div class="flex justify-center">
        <ButtonComponent
          :class="
            actionString && reactionString
              ? ''
              : 'cursor-not-allowed opacity-50'
          "
          text="Add Workflow"
          :bg-color="
            actionString && reactionString
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
      >
    </div>
    <div class="flex justify-start gap-5 m-20">
      <DropdownComponent
        v-model="selectedFilter"
        :label="selectedFilter"
        :options="filters"
      />
      <DropdownComponent
        v-model="selectedSort"
        :label="selectedSort"
        :options="sorts"
      />
    </div>
    <ListTableComponent
      v-show="columns && filteredWorkflows"
      v-model="workflowsInList"
      :columns="columns"
      :rows="filteredWorkflows"
    />
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      >
    </div>
    <div class="flex justify-center m-20">
      <div
        class="relative flex justify-center bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-2xl w-10/12"
      >
        <button
          class="absolute top-4 right-4 text-fontBlack dark:text-fontWhite hover:text-accent-200 dark:hover:text-accent-500 transition duration-200"
          aria-label="Copier le JSON"
          @click="copyToClipboard(JSON.stringify(lastWorkflowResult, null, 2))"
        >
          <Icon :name="copyIcon" />
        </button>
        <pre
          class="whitespace-pre-wrap break-words text-sm text-primaryWhite-800 dark:text-primaryWhite-200 p-4"
        >
    {{ JSON.stringify(lastWorkflowResult, null, 2) }}
  </pre
        >
      </div>
    </div>
  </div>
</template>
