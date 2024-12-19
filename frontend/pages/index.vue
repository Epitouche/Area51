<template>
  <NavbarComponent />
  <navigation>
    <div class="flex flex-col mt-32 mr-12 ml-12" style="height: 70vh;">
      <div class="flex flex-col mx-12">
        <h1 class="text-6xl font-bold mx-2">Automatize</h1>
        <h1 class="text-8xl font-bold text-purple-500">Everything.</h1>
        <p class="mt-4 text-xl text-gray-700">
          Make your favorite apps connect with each other <br />
          and let us automatize tasks for you!
        </p>
        <NuxtLink to="/register">
          <ButtonComponent
          text="Start now!"
          bgColor="bg-tertiary-500"
          hoverColor="hover:bg-purple-600"
          textColor="text-fontWhite"
          class="mt-4"
          ></ButtonComponent>
        </NuxtLink>
      </div>
    </div>
  </navigation>
</template>
