<template>
    <div class="bg-primaryWhite-500 dark:bg-primaryDark-500">
      <nav class="p-4 border-b border-gray-200 dark:border-secondaryDark-500">
        <div class="container mx-auto flex justify-between items-center">
          <div class="text-lg font-bold">
            <NuxtLink to="/">MyLogo</NuxtLink>
          </div>
          <div class="space-x-9">
            <!-- <NuxtLink to="/">Home</NuxtLink> -->
            <NuxtLink to="/login" class="text-fontBlack dark:text-fontWhite">Login</NuxtLink>
            <NuxtLink to="/register" class="text-fontBlack dark:text-fontWhite">
              <ButtonComponent
                bgColor="bg-tertiary-500"
                hoverColor="hover:bg-purple-600"
              >
                Sign up
              </ButtonComponent>
            </NuxtLink>
          </div>
        </div>
      </nav>
    </div>
</template>
