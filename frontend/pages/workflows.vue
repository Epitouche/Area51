<script setup lang="ts">
import { useNotificationStore } from "@/stores/notification";

import type {
  Service,
  WorkflowResponse,
  Workflow,
  AboutResponse,
  OptionWorkflow,
} from "~/src/types";

const notificationStore = useNotificationStore();

// it has to be a table
const actionOption = ref<OptionWorkflow[]>([]);
const reactionOption = ref<OptionWorkflow[]>([]);

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

const columns = ["Name", "Action", "Reaction", "Activity", "Creation Date"];

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

const workflowName = ref("");

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

  const actions =
    services
      .flatMap((service) => service.actions)
      .find((action) => action.name === actionString.value)?.options || "";

  const parsedOptions: Record<string, string> = JSON.parse(actions);

  const transformedOptions: OptionWorkflow[] = Object.entries(
    parsedOptions
  ).map(([name, type]) => ({
    name,
    type,
    input: "",
  }));

  actionOption.value = transformedOptions;
  
};

const openModalReaction = () => {
  isModalReactionOpen.value = true;
};

const closeModalReaction = () => {
  isModalReactionOpen.value = false;
};

const confirmModalReaction = () => {
  closeModalReaction();
  const reactions =
    services
      .flatMap((service) => service.reactions)
      .find((reaction) => reaction.name === reactionString.value)?.options || "";

  const parsedOptions: Record<string, string> = JSON.parse(reactions);

  const transformedOptions: OptionWorkflow[] = Object.entries(
    parsedOptions
  ).map(([name, type]) => ({
    name,
    type,
    input: "",
  }));

  reactionOption.value = transformedOptions;

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
      const serviceFound = services.find((s) => s.name === service.name);

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

async function transformOptions(options: OptionWorkflow[]): Promise<string> {
  const transformed = options.reduce((acc, option) => {
    acc[option.name] = option.input;
    return acc;
  }, {} as Record<string, string>);

  console.log("transformed", transformed);

  return JSON.stringify(transformed);
}

async function addWorkflow() {
  try {
    const actionSelected = services
      .flatMap((service) => service.actions)
      .find((action) => action.name === actionString.value);

    const reactionSelected = services
      .flatMap((service) => service.reactions)
      .find((reaction) => reaction.name === reactionString.value);

    const actionOptions = await transformOptions(actionOption.value);
    const reactionOptions = await transformOptions(reactionOption.value);

    if (actionSelected && reactionSelected) {
      const body: {
        action_id: number;
        reaction_id: number;
        name?: string;
        action_option?: string;
        reaction_option?: string;
      } = {
        action_id: actionSelected.action_id,
        reaction_id: reactionSelected.reaction_id,
        action_option: actionOptions,
        reaction_option: reactionOptions,
        name: workflowName.value,
      };

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
      workflowName.value = "";
      actionOption.value = [];
      reactionOption.value = [];
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

const isConnected = computed(() => {
  return token.value !== undefined;
});

onMounted(() => {
  fetchServices();
  fetchWorkflows();
  getLastWorkflowResult();
});
</script>

<template>
  <div
    class="flex flex-col min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500"
    aria-label="Workflows management screen"
  >
    <div v-if="isConnected" aria-label="Workflows management area">
      <div class="m-5 sm:m-10">
        <h1
          class="text-3xl sm:text-4xl md:text-6xl font-bold text-fontBlack dark:text-fontWhite"
          aria-label="Workflows heading"
        >
          Workflows
        </h1>
      </div>
      <div class="flex justify-center">
        <hr
          class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-full sm:w-11/12"
          aria-hidden="true"
        >
      </div>
      <div
        class="flex flex-col justify-center m-5 sm:m-10 gap-5 sm:gap-10 items-center"
        aria-label="Workflow creation section"
      >
        <InputComponent
          v-model="workflowName"
          type="text"
          label="Workflow Name"
          aria-label="Enter workflow name"
        />
        <div class="flex flex-wrap justify-center gap-3 sm:gap-5">
          <ButtonComponent
            :text="actionString ? actionString : 'Choose an action'"
            bg-color="bg-primaryWhite-500 dark:bg-secondaryDark-500"
            hover-color="hover:bg-accent-100 dark:hover:bg-accent-800"
            text-color="text-fontBlack dark:text-fontWhite"
            :on-click="openModalAction"
            aria-label="Open action selection modal"
          />
          <ModalComponent
            v-motion-pop
            title="Choose an action"
            :is-open="isModalActionOpen"
            @close="closeModalAction"
            @confirm="confirmModalAction"
            aria-label="Action selection modal"
          >
            <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
              <div
                v-for="(service, index) in services"
                :key="index"
                class="flex justify-center"
                aria-label="Service action options"
              >
                <DropdownComponent
                  v-if="service.actions"
                  v-model="actionString"
                  :label="service.name"
                  :options="service.actions.map((action) => action.name)"
                  aria-label="Choose action for {{ service.name }}"
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
            aria-label="Open reaction selection modal"
          />
          <ModalComponent
            v-motion-pop
            title="Choose a reaction"
            :is-open="isModalReactionOpen"
            @close="closeModalReaction"
            @confirm="confirmModalReaction"
            aria-label="Reaction selection modal"
          >
            <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
              <div
                v-for="(service, index) in services"
                :key="index"
                class="flex justify-center"
                aria-label="Service reaction options"
              >
                <DropdownComponent
                  v-if="service.reactions"
                  v-model="reactionString"
                  :label="service.name"
                  :options="service.reactions.map((reaction) => reaction.name)"
                  aria-label="Choose reaction for {{ service.name }}"
                />
              </div>
            </div>
          </ModalComponent>
        </div>
        <div class="flex justify-center gap-3 sm:gap-5">
          <div
            v-if="actionOption.length !== 0"
            class="border-2 rounded-lg border-primaryWhite-500 dark:border-secondaryDark-500 p-2 w-full sm:w-10/12"
            aria-label="Action options"
          >
            <InputComponent
              v-for="(option, index) in actionOption"
              :key="`action-option-${index}`"
              v-model="option.input"
              :type="option.type"
              :label="option.name"
              class="mb-4"
              aria-label="Input for action option {{ option.name }}"
            />
          </div>
          <div
            v-if="reactionOption.length !== 0"
            class="border-2 rounded-lg border-primaryWhite-500 dark:border-secondaryDark-500 p-2 w-full sm:w-10/12"
            aria-label="Reaction options"
          >
            <InputComponent
              v-for="(option, index) in reactionOption"
              :key="`action-option-${index}`"
              v-model="option.input"
              :type="option.type"
              :label="option.name"
              class="mb-4"
              aria-label="Input for reaction option {{ option.name }}"
            />
          </div>
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
            :aria-disabled="!(actionString && reactionString)"
            aria-label="Add workflow button"
          />
        </div>
      </div>
      <div class="flex justify-center">
        <hr
          class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-full sm:w-11/12"
          aria-hidden="true"
        >
      </div>
      <div class="flex flex-wrap justify-start gap-3 sm:gap-5 m-5 sm:m-10">
        <DropdownComponent
          v-model="selectedFilter"
          :label="selectedFilter"
          :options="filters"
          aria-label="Filter workflows dropdown"
        />
        <DropdownComponent
          v-model="selectedSort"
          :label="selectedSort"
          :options="sorts"
          aria-label="Sort workflows dropdown"
        />
      </div>
      <ListTableComponent
        v-show="columns && filteredWorkflows"
        v-model="workflowsInList"
        :columns="columns"
        :rows="filteredWorkflows"
        aria-label="List of workflows"
      />
      <div class="flex justify-center">
        <hr
          class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-full sm:w-11/12"
          aria-hidden="true"
        >
      </div>
      <div class="flex justify-center m-5 sm:m-10">
        <div
          class="relative flex justify-center bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-2xl w-full sm:w-10/12"
          aria-label="Workflow result JSON"
        >
          <button
            class="absolute top-2 right-2 sm:top-4 sm:right-4 text-fontBlack dark:text-fontWhite hover:text-accent-200 dark:hover:text-accent-500 transition duration-200"
            aria-label="Copy workflow result JSON"
            @click="
              copyToClipboard(JSON.stringify(lastWorkflowResult, null, 2))
            "
          >
            <Icon :name="copyIcon" />
          </button>
          <pre
            class="whitespace-pre-wrap break-words text-xs sm:text-sm text-primaryWhite-800 dark:text-primaryWhite-200 p-4"
          >
    {{ JSON.stringify(lastWorkflowResult, null, 2) }}
    </pre
          >
        </div>
      </div>
    </div>
    <div v-else aria-label="Error page">
      <div class="flex flex-col gap-4 justify-center items-center h-full">
        <h1
          class="text-3xl sm:text-4xl md:text-6xl font-bold text-fontBlack dark:text-fontWhite"
          aria-label="Error 404 heading"
        >
          ERROR 404 !
        </h1>
        <h2
          class="text-2xl sm:text-3xl font-bold text-fontBlack dark:text-fontWhite"
          aria-label="Not connected message"
        >
          You are not connected, please log in to access this page.
        </h2>
      </div>
    </div>
  </div>
</template>
