<script setup>
import RecentActivity from "~/components/dashboard/RecentActivity.vue";
import WorkflowList from "~/components/dashboard/WorkflowList.vue";

// const token = useCookie()
</script>
<!-- First line récap: Active Workflow | Total executions | Last 24h | Last Execution -->
<!-- Second line Tab & Services: Graph taking most space and Services on the right that are connected or not (Button to services page) -->
<!-- Third line Workflow activity: Your workflows | Recent activity (Button to workflows page) -->
<template>
  <div>
    <div
      class="min-h-screen bg-secondaryWhite-500 dark:bg-primaryDark-500 pt-20"
      aria-label="Dashboard page layout"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <h1 
          class="text-2xl font-bold text-fontBlack dark:text-fontWhite mb-8"
          aria-label="Dashboard heading"
        >
          Dashboard
        </h1>

        <div 
          class="grid gap-6 md:grid-cols-2 lg:grid-cols-3 mb-8"
          aria-label="Dashboard statistics section"
        >
          <DashboardStats />
        </div>

        <div 
          class="grid gap-6 lg:grid-cols-2 mb-8"
          aria-label="Dashboard main content layout"
        >
          <div aria-label="Workflow graph section">
            <WorkflowList aria-label="Workflow list" />
          </div>
          <div aria-label="Services and integrations section">
            <DashboardServicesIntegrations />
          </div>
        </div>

        <div 
          class="grid gap-6"
          aria-label="Additional sections with workflows and recent activities"
        >
          <RecentActivity aria-label="Recent activity section" />
        </div>
      </div>
    </div>
  </div>
</template>
