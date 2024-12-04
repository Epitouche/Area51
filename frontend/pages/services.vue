<script setup lang="ts">
import ListTable from "~/components/ListTable.vue";

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
    <div class="flex justify-center space-x-10 m-20">
      <Button
        text="Choose an action"
        bgColor="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hoverColor="hover:bg-accent-900 dark:hover:bg-accent-800"
        textColor="text-fontBlack dark:text-fontWhite"
      >
      </Button>
      <Button
        text="Choose a reaction"
        bgColor="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hoverColor="hover:bg-accent-900 dark:hover:bg-accent-800"
        textColor="text-fontBlack dark:text-fontWhite"
      ></Button>
    </div>
    <div class="flex justify-center">
      <hr
        class="border-primaryWhite-500 dark:border-secondaryDark-500 border-2 w-11/12"
      />
    </div>
    <div class="flex justify-start gap-5 m-20">
      <Button
        text="Filter"
        bgColor="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hoverColor="hover:bg-accent-900 dark:hover:bg-accent-800"
        textColor="text-fontBlack dark:text-fontWhite"
      ></Button>
      <Button
        text="All Status"
        bgColor="bg-primaryWhite-500 dark:bg-secondaryDark-500"
        hoverColor="hover:bg-accent-900 dark:hover:bg-accent-800"
        textColor="text-fontBlack dark:text-fontWhite"
      ></Button>
    </div>
    <ListTable v-show="columns && rows" :columns="columns" :rows="rows" />
  </div>
</template>
