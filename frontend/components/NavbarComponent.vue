<template>
    <div>
      <nav class="p-4 border-b border-gray-200">
        <div class="container mx-auto flex justify-between items-center">
          <div class="text-lg font-bold">
            <NuxtLink to="/">MyLogo</NuxtLink>
          </div>
          <div class="space-x-9">
            <!-- <NuxtLink to="/">Home</NuxtLink> -->
            <NuxtLink to="/login">Login</NuxtLink>
            <NuxtLink to="/register">
              <ButtonComponent
                bgColor="bg-purple-500"
                textColor="text-white"
                hoverColor="hover:bg-purple-600"
              >
                Sign up
              </ButtonComponent>
                <!-- <button class="bg-[#B454FD] px-4 py-2 text-white rounded-full">Sign up</button> -->
            </NuxtLink>
          </div>
        </div>
      </nav>
    </div>
</template>
