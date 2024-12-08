<script>
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-primaryWhite-500 dark:bg-primaryDark-500"
  >
    <div
      class="w-full transform -translate-x-3/4 max-w-md p-8 space-y-10 bg-gradient-to-b from-tertiary-500 to-tertiary-600 dark:from-tertiary-600 dark:to-tertiary-500 text-fontWhite rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold text-center">REGISTER</h2>
      <form class="space-y-6">
        <div>
          <InputComponent id="username" type="text" label="Username" />
        </div>
        <div>
          <InputComponent id="email" type="email" label="Email" />
        </div>
        <div>
          <InputComponent id="password" type="password" label="Password" />
        </div>
      </form>
      <div class="flex items-center">
        <InputComponent
          id="remember"
          type="checkbox"
          class="w-4 h-4 text-accent-500 border-primaryWhite-300 rounded focus:ring-accent-500"
        />
        <label for="remember" class="ml-2 text-sm">Remember me</label>
      </div>
      <div class="flex justify-center">
        <ButtonComponent
          text="Create Account"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
      </div>
      <hr class="border-primaryWhite-400" />
      <div class="flex justify-around space-x-4">
        <ButtonComponent
          text="Github"
          class="w-full"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
        <ButtonComponent
          text="Google"
          class="w-full"
          bgColor="bg-primaryWhite-500"
          hoverColor="hover:bg-secondaryWhite-500"
          textColor="text-fontBlack"
        />
      </div>
    </div>
  </div>
</template>
