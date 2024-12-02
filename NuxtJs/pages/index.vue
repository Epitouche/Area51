<template>
  <Navbar />
  <div class="flex flex-col items-center justify-center h-screen bg-gray-50">
    <h1 class="text-4xl font-bold text-center">Automatize <span class="text-purple-500">Everything.</span></h1>
    <p class="mt-4 text-lg text-gray-700 text-center">
      Make your favorite apps connect with each other <br />
      and let us automatize tasks for you!
    </p>
    <NuxtLink to="/login">
      <button
        class="mt-6 bg-purple-500 hover:bg-purple-600 text-white font-bold py-2 px-6 rounded-full shadow-md transition duration-300">
        Start now!
      </button>
    </NuxtLink>
  </div>
</template>
