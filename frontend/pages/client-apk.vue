<template>
  <div aria-label="APK download status">
    <h1 v-if="errorMessage" aria-label="Error message">{{ errorMessage }}</h1>
    <p v-else aria-label="Downloading status">Downloading your APK...</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      errorMessage: null,
    };
  },
  mounted() {
    const apkUrl = "/apk/client.apk";

    fetch(apkUrl, { method: "HEAD" })
      .then((response) => {
        if (response.ok) {
          const link = document.createElement("a");
          link.href = apkUrl;
          link.download = "client.apk";
          link.click();
        } else {
          this.errorMessage = "File not found";
        }
      })
      .catch(() => {
        this.errorMessage = "File not found";
      });
  },
};
</script>
