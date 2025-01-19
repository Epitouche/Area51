export default defineEventHandler(async (event) => {
  try {
    const access_token = event.headers.get("Authorization");

    if (!access_token) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters: token",
      });
    }

    const response = await $fetch("http://server:8080/api/user/account", {
      method: "DELETE",
      headers: {
        Authorization: access_token ? `${access_token}` : "",
        "Content-Type": "application/json",
      },
    });

    return response;
  } catch (error) {
    console.error("Error in API server:", error);
    return { statusCode: 500, message: "Failed to delete user data" };
  }
});