<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const workflows = ref([
  {
    id: 1,
    title: "GitHub to Email",
    trigger: { icon: 'line-md:github', name: "New GitHub Issue" },
    action: { icon: 'material-symbols:mail-outline', name: "Send Mail" },
    description: "Send an email when a new github issue is created"
  },
  {
    id: 2,
    title: "Email to File Saving",
    trigger: { icon: 'material-symbols:mail-outline', name: "New Email" },
    action: { icon: 'material-symbols:attach-file', name: "Save files" },
    description: "Save important files received by email in your drive"
  },
  {
    id: 3,
    title: "Spotify Notification",
    trigger: { icon: 'line-md:spotify', name: "New Song" },
    action: { icon: 'material-symbols:notifications-outline', name: "Get a Notification" },
    description: "Get a notification from your favorite artist last song directly on the app"
  }
]);

const navigateToRegister = () => {
  useRouter().push('/register');
};
</script>
<template>
  <NavbarComponent />
  <div class="flex flex-col">
    <div class="flex flex-col mt-48 mx-52 m-12" style="height: 50vh;">
      <div class="flex flex-col mx-12">
        <h1 class="text-6xl font-bold mx-2">Automatize</h1>
        <h1 class="text-8xl font-bold text-purple-500 mt-2">Everything.</h1>
        <p class="my-4 text-xl text-gray-700">
          Make your favorite apps connect with each other <br />
          and let us automatize tasks for you !
        </p>
        <ButtonComponent
          text="Start now!"
          bgColor="bg-tertiary-500"
          hoverColor="hover:bg-purple-600"
          textColor="text-fontWhite"
          class="mt-4 inline-block max-w-max rounded-lg"
          @click="navigateToRegister"
        ></ButtonComponent>
      </div>
    </div>

    <div class="max-w-8xl mx-auto px-4 sm:px-6 lg:px-8 items-center mb-16">
      <div class="text-center mb-16">
        <h2 class="text-3xl font-bold text-gray-900 mb-4">Make your day easy</h2>
        <p class="text-gray-600">Discover new ways to create useful workflows</p>
      </div>
      <div class="grid md:grid-cols-3 gap-8">
        <div v-for="workflow in workflows" :key="workflow.id" class="bg-white rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow">
          <h3 class="text-xl font-semibold mb-4">{{ workflow.title }}</h3>
          <div class="flex items-center space-x-4 mb-4">
            <div class="flex items-center">
              <div class="bg-purple-100 p-3 rounded-lg">
                <Icon :name="workflow.trigger.icon" class="h-6 w-6 text-purple-700" />
              </div>
              <span class="ml-2">{{ workflow.trigger.name }}</span>
            </div>
            <Icon name="mdi-arrow-right" class="h-5 w-5 text-gray-400" />
            <div class="flex items-center">
              <div class="bg-purple-100 p-3 rounded-lg">
                <Icon :name="workflow.action.icon" class="h-6 w-6 text-purple-700" />
              </div>
              <span class="ml-2">{{ workflow.action.name }}</span>
            </div>
          </div>
          <p class="text-gray-600">{{ workflow.description }}</p>
          <!-- <button class="mt-4 text-brand-purple hover:text-brand-dark flex items-center">
            Use this workflow
          </button> -->
        </div>
      </div>
    </div>
  </div>
</template>
