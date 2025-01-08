export default defineEventHandler(async (event) => {
  const access_token = event.headers.get("Authorization");
  const contentTypes = event.headers.get("Content-Type");
  const params = await readBody(event);
  if (!access_token || !contentTypes || !params.workflow_id || params.workflow_state === undefined) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token or workflow_id or workflow_state",
    });
  }
  try {
    const response = await $fetch(
      "http://server:8080/api/workflow/activation",
      {
        method: "PUT",
        headers: {
          Authorization: access_token ? `${access_token}` : "",
          "Content-Type": contentTypes ? `${contentTypes}` : "",
        },
        body: {
          workflow_id: params.workflow_id,
          workflow_state: params.workflow_state,
        },
      }
    );
    return response;
  } catch (error) {
    console.error("Error in API server:", error);
    return { statusCode: 500, message: "Failed to switch workflow state" };
  }
});
