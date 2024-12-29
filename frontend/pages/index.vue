<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import FooterComponent from '~/components/FooterComponent.vue';

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

const features = ref([
  {
    icon: 'solar:users-group-rounded-linear',
    title: 'Our Team',
    description: 'We are a group of 5 passionate students from Epitech Bordeaux united by our ambition to bring innovative ideas to life'
  },
  {
    icon: 'material-symbols:lightbulb-2-outline-rounded',
    title: 'The Project',
    description: 'The goal was to create an AREA website where users like you can create useful workflows'
  },
  {
    icon: 'material-symbols:target',
    title: 'Our Aim',
    description: 'Our commitment to excellence and problem-solving drives everything we do. Together, we aim to create something impactful and unforgettable'
  }
]);

const navigateToRegister = () => {
  useRouter().push('/register');
};
</script>
<template>
<!-- Navbar -->
  <NavbarComponent />

  <!-- Hero Section -->
  <div class="flex flex-col bg-gray-50">
    <div class="items-center justify-center h-2/3 bg-gray-50 mx-8">
      <div class="mt-32 mb-48 mx-16 sm:mx-24 md:mx-32 lg:mx-48 xl:mx-60">
        <div class="flex flex-col md:flex-row items-center justify-between">
          <!-- Left Content -->
          <div class="text-center md:text-left">
            <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold mx-2">Automatize</h1>
            <h1 class="text-5xl sm:text-6xl md:text-8xl font-bold text-purple-500 mt-2">Everything.</h1>
            <p class="my-4 text-lg sm:text-xl md:text-2xl text-gray-700">
              Make your favorite apps connect with each other <br />
              and let us automatize tasks for you!
            </p>
            <ButtonComponent
              text="Start now!"
              bgColor="bg-tertiary-500"
              hoverColor="hover:bg-purple-600"
              textColor="text-fontWhite"
              class="mt-4 inline-block max-w-max"
              @click="navigateToRegister"
            ></ButtonComponent>
          </div>
          <!-- Right Box -->
          <div class="flex justify-center items-center w-full md:w-auto mr-8 sm:mr-16 md:mr-24 lg:mr-32 xl:mr-48">
            <div class="grid md:grid-rows-3 gap-8">
              <div v-for="workflow in workflows" :key="workflow.id" class="bg-white rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow">
                <h3 class="text-sm mb-4">{{ workflow.title }}</h3>
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
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

<!-- Workflow example -->
    <div class="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 items-center mb-24">
      <div class="text-center mb-16">
        <h2 class="text-4xl font-bold text-gray-900 mb-4">Make your day easy</h2>
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
        </div>
      </div>
    </div>

<!-- About us -->
    <div class="bg-purple-50 my-8">
      <div class="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 mt-24 mb-20">
        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-gray-900 mb-4">About Us</h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto">
            We're on a mission to make workflow automation accessible to everyone
          </p>
        </div>
        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="feature in features" class="bg-white rounded-xl p-6 shadow-md hover:shadow-lg transition-shadow">
            <div class="bg-purple-100 w-12 h-12 rounded-lg flex items-center justify-center mb-4">
              <Icon :name="feature.icon" class="h-6 w-6 text-purple-700" />
            </div>
            <h3 class="text-xl font-semibold mb-2">{{ feature.title }}</h3>
            <p class="text-gray-600">{{ feature.description }}</p>
          </div>
        </div>
        <p class="text-center text-xl text-gray-600 max-w-3xl mx-auto mt-16">
          You can see the documentation <a class="text-purple-700" href="https://github.com/Epitouche/Area51/blob/main/README.md">here</a> to learn more about us !
        </p>
      </div>
    </div>

<!-- Footer -->
    <FooterComponent />
  </div>
</template>
