<template>
  <div>
    <NavbarComponent />
    <div class="flex flex-col items-center justify-center h-screen bg-gray-50">
      <!-- <div class="absolute inset-0">
        <svg
          class="w-full h-full"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 1440 320"
        >
          <path
            fill="none"
            stroke="#A78BFA"
            stroke-width="1"
            d="M0,0 C360,320 1080,320 1440,300"
          ></path>
          <path
            fill="none"
            stroke="#D8B4FE"
            stroke-width="1"
            d="M0,50 C400,250 1040,350 1440,280"
          ></path>
          <path
            fill="none"
            stroke="#7C3AED"
            stroke-width="1"
            d="M0,150 C500,50 940,400 1440,200"
          ></path>
          <path
            fill="none"
            stroke="#A78BFA"
            stroke-width="1"
            d="M0,100 C480,200 960,200 1440,150"
          ></path>
        </svg>
      </div> -->
      <h1 class="text-5xl font-bold text-center">Automatize <span class="text-purple-500">Everything.</span></h1>
      <p class="mt-4 text-lg text-gray-700 text-center">
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
</template>
