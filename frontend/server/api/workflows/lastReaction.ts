export default defineEventHandler(async (event) => {
  try {
    const access_token = event.headers.get("Authorization");
    const params = getQuery(event);
    if (!access_token || !params.workflow_id) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters: token or workflow_id",
      });
    }

    const response = await $fetch("http://server:8080/api/workflow/reaction/latest", {
      method: "GET",
      headers: {
        Authorization: access_token ? `${access_token}` : "",
        "Content-Type": "application/json",
      },
      params: {
        workflow_id: params.workflow_id,
      },
    });

    return response;
  } catch (error) {
    console.error("Error in API server:", error);
    return { statusCode: 500, message: "Failed to get reaction" };
  }
});
