<script setup>
import { ref } from "vue";

const workflows = ref([
  {
    id: 1,
    title: "GitHub to Email",
    hero_card: {
      title: "Issue Tracking",
      first_logo: {
        icon: "line-md:github",
        label: "New Issue",
        color: "bg-gray-900",
      },
      trigger: {
        icon: "jam:thunder",
        label: "Trigger",
        color: "bg-tertiary-500",
      },
      second_logo: {
        icon: "material-symbols:mail-outline",
        label: "Notify Team",
        color: "bg-indigo-500",
      },
    },
    trigger: { icon: "line-md:github", name: "New GitHub Issue" },
    action: { icon: "material-symbols:mail-outline", name: "Send Mail" },
    description: "Send an email when a new github issue is created",
  },
  {
    id: 2,
    title: "Email to File Saving",
    hero_card: {
      title: "Save important files",
      first_logo: {
        icon: "material-symbols:mail-outline",
        label: "Email",
        color: "bg-red-500",
      },
      trigger: { icon: "jam:thunder", label: "Save", color: "bg-tertiary-500" },
      second_logo: {
        icon: "material-symbols:attach-file",
        label: "Store File",
        color: "bg-emerald-500",
      },
    },
    trigger: { icon: "material-symbols:mail-outline", name: "New Email" },
    action: { icon: "material-symbols:attach-file", name: "Save files" },
    description: "Save important files received by email in your drive",
  },
  {
    id: 3,
    title: "Spotify Notification",
    hero_card: {
      title: "Song notification",
      first_logo: {
        icon: "line-md:spotify",
        label: "New Song",
        color: "bg-emerald-500",
      },
      trigger: {
        icon: "jam:thunder",
        label: "Verify",
        color: "bg-tertiary-500",
      },
      second_logo: {
        icon: "material-symbols:notifications-outline",
        label: "Notification",
        color: "bg-amber-500",
      },
    },
    trigger: { icon: "line-md:spotify", name: "New Song" },
    action: {
      icon: "material-symbols:notifications-outline",
      name: "Get a Notification",
    },
    description:
      "Get a notification from your favorite artist last song directly on the app",
  },
]);

const features = ref([
  {
    icon: "solar:users-group-rounded-linear",
    title: "Our Team",
    description:
      "We are a group of 5 passionate students from Epitech Bordeaux united by our ambition to bring innovative ideas to life",
  },
  {
    icon: "material-symbols:lightbulb-2-outline-rounded",
    title: "The Project",
    description:
      "The goal was to create an AREA website where users like you can create useful workflows",
  },
  {
    icon: "material-symbols:target",
    title: "Our Aim",
    description:
      "Our commitment to excellence and problem-solving drives everything we do. Together, we aim to create something impactful and unforgettable",
  },
]);
</script>
<template>
  <div>
    <div class="flex flex-col bg-secondaryWhite-500 dark:bg-primaryDark-500">
      <!-- Hero Section -->
      <div class="items-center justify-center mx-8" aria-label="Hero Section">
        <div class="mt-32 mb-40 mx-8 sm:mx-16 md:mx-24 lg:mx-48 xl:mx-60">
          <div class="flex flex-col md:flex-row items-center justify-between">
            <!-- Left Content -->
            <div
              class="text-center md:text-left ml-4 sm:ml-8 md:ml-16 lg:ml-24 xl:ml-32"
              aria-label="Introduction and Call to Action"
            >
              <h1
                class="text-4xl sm:text-5xl md:text-6xl font-bold mx-2 text-fontBlack dark:text-fontWhite"
              >
                Automate
              </h1>
              <h1
                class="text-5xl sm:text-6xl md:text-8xl font-bold text-tertiary-500 dark:text-tertiary-300 mt-2"
              >
                Everything.
              </h1>
              <p
                class="my-4 text-lg sm:text-xl md:text-2xl text-gray-700 dark:text-fontWhite"
              >
                Make your favorite apps connect with each other <br>
                and let us automate tasks for you!
              </p>
              <NuxtLink
                to="/register"
                class="text-fontBlack dark:text-fontWhite mt-4 inline-block max-w-max"
                aria-label="Start now button to register"
              >
                <ButtonComponent
                  text="Start now!"
                  bg-color="bg-tertiary-500"
                  hover-color="hover:bg-purple-600"
                  text-color="text-fontWhite"
                />
              </NuxtLink>
            </div>
            <!-- Right Box -->
            <div
              class="flex justify-center items-center w-full md:w-auto mr-8 sm:mr-16 md:mr-24 lg:mr-32 xl:mr-48"
              aria-label="Highlighted Workflows"
            >
              <div class="grid md:grid-rows-3 gap-8">
                <div
                  v-for="workflow in workflows"
                  :key="workflow.id"
                  class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow"
                  aria-label="Workflow card"
                >
                  <div class="mx-6">
                    <h2 class="text-md mb-4 text-fontBlack dark:text-fontWhite" aria-label="Workflow title">
                      {{ workflow.hero_card.title }}
                    </h2>
                    <div class="flex items-center space-x-4 mb-4" aria-label="Workflow details">
                      <div class="flex flex-col items-center">
                        <IconComponent
                          :bg-color="workflow.hero_card.first_logo.color"
                          text-color="text-white"
                          :icon="workflow.hero_card.first_logo.icon"
                          :aria-label="`${workflow.hero_card.first_logo.label} icon`"
                        />
                        <span
                          class="mt-2 text-sm font-medium text-gray-600 dark:text-fontWhite"
                          :aria-label="workflow.hero_card.first_logo.label"
                        >{{ workflow.hero_card.first_logo.label }}</span>
                      </div>
                      <Icon
                        name="mdi-arrow-right"
                        class="h-7 w-7 text-gray-400 mb-5"
                        aria-label="Arrow icon"
                      />
                      <div class="flex flex-col items-center">
                        <IconComponent
                          :bg-color="workflow.hero_card.trigger.color"
                          text-color="text-white"
                          :icon="workflow.hero_card.trigger.icon"
                          :aria-label="`${workflow.hero_card.trigger.label} icon`"
                        />
                        <span
                          class="mt-2 text-sm font-medium text-gray-600 dark:text-fontWhite"
                          :aria-label="workflow.hero_card.trigger.label"
                        >{{ workflow.hero_card.trigger.label }}</span>
                      </div>
                      <Icon
                        name="mdi-arrow-right"
                        class="h-7 w-7 text-gray-400 mb-5"
                        aria-label="Arrow icon"
                      />
                      <div class="flex flex-col items-center">
                        <IconComponent
                          :bg-color="workflow.hero_card.second_logo.color"
                          text-color="text-white"
                          :icon="workflow.hero_card.second_logo.icon"
                          :aria-label="`${workflow.hero_card.second_logo.label} icon`"
                        />
                        <span
                          class="mt-2 text-sm font-medium text-gray-600 dark:text-fontWhite"
                          :aria-label="workflow.hero_card.second_logo.label"
                        >{{ workflow.hero_card.second_logo.label }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Workflow example -->
      <div
        class="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 items-center mb-24"
        aria-label="Workflow examples"
      >
        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-gray-900 mb-4 dark:text-fontWhite" aria-label="Section title">
            Make your day easy
          </h2>
          <p class="text-fontBlack dark:text-fontWhite" aria-label="Section description">
            Discover new ways to create useful workflows
          </p>
        </div>
        <div class="grid md:grid-cols-3 gap-8">
          <div
            v-for="workflow in workflows"
            :key="workflow.id"
            class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow"
            aria-label="Workflow example card"
          >
            <h2
              class="text-xl font-semibold mb-4 text-fontBlack dark:text-fontWhite"
              aria-label="Workflow example title"
            >
              {{ workflow.title }}
            </h2>
            <div class="flex items-center space-x-4 mb-4" aria-label="Workflow example details">
              <div class="flex items-center">
                <IconComponent
                  bg-color="bg-purple-100"
                  text-color="text-purple-700"
                  :icon="workflow.trigger.icon"
                  :aria-label="`${ workflow.trigger.name } icon`"
                />
                <span class="ml-4 text-fontBlack dark:text-fontWhite" :aria-label="workflow.trigger.name">
                  {{ workflow.trigger.name }}
                </span>
              </div>
              <Icon
                name="mdi-arrow-right"
                class="h-5 w-5 text-gray-400"
                aria-label="Arrow icon"
              />
              <div class="flex items-center">
                <IconComponent
                  bg-color="bg-purple-100"
                  text-color="text-purple-700"
                  :icon="workflow.action.icon"
                  :aria-label="`${ workflow.action.name } icon`"
                />
                <span class="ml-4 text-fontBlack dark:text-fontWhite" :aria-label="workflow.action.name">
                  {{ workflow.action.name }}
                </span>
              </div>
            </div>
            <p class="text-gray-600 dark:text-fontWhite" aria-label="Workflow example description">
              {{ workflow.description }}
            </p>
          </div>
        </div>
      </div>

      <!-- About us -->
      <div class="bg-tertiary-100 dark:bg-tertiary-700 my-8" aria-label="About us section">
        <div class="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 mt-24 mb-20">
          <div class="text-center mb-16">
            <h2
              class="text-4xl font-bold text-gray-900 dark:text-fontWhite mb-4"
              aria-label="About us title"
            >
              About Us
            </h2>
            <p
              class="text-xl text-gray-600 max-w-2xl mx-auto dark:text-fontWhite"
              aria-label="About us description"
            >
              We're on a mission to make workflow automation accessible to
              everyone
            </p>
          </div>
          <div class="grid md:grid-cols-3 gap-8">
            <div
              v-for="feature in features"
              :key="feature.title"
              class="bg-primaryWhite-500 dark:bg-secondaryDark-500 rounded-xl p-6 shadow-md hover:shadow-lg transition-shadow"
              aria-label="Feature card"
            >
              <IconComponent
                bg-color="bg-purple-100"
                text-color="text-tertiary-700"
                :icon="feature.icon"
                class="w-14 h-14 flex justify-center mb-4"
                :aria-label="`${ feature.title } icon`"
              />
              <h2
                class="text-xl font-semibold mb-2 text-fontBlack dark:text-fontWhite"
                aria-label="Feature title"
              >
                {{ feature.title }}
              </h2>
              <p class="text-gray-600 dark:text-fontWhite" aria-label="Feature description">
                {{ feature.description }}
              </p>
            </div>
          </div>
          <p
            class="text-center text-xl text-gray-600 dark:text-fontWhite max-w-3xl mx-auto mt-16"
            aria-label="Documentation link"
          >
            You can see the documentation
            <a
              class="underline-offset-2 underline"
              href="https://github.com/Epitouche/Area51/blob/main/README.md"
              aria-label="Link to documentation"
            >here</a
            >
            to learn more about us !
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
