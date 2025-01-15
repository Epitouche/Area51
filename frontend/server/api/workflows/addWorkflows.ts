export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    const access_token = event.headers.get("Authorization");
    if (!access_token || !params.action_id || !params.reaction_id) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters: token, action_id or reaction_id",
      });
    }
    const response = await $fetch("http://server:8080/api/workflow", {
      method: "POST",
      headers: {
        Authorization: access_token ? `${access_token}` : "",
        "Content-Type": "application/json",
      },
      body: {
        action_id: params.action_id,
        reaction_id: params.reaction_id,
        name: params.name,
        action_options: params.action_options ? params.action_options : "",
        reaction_options: params.reaction_options ? params.reaction_options : "",
      },
    });

    return response;
  } catch (error) {
    console.error("Error in API server:", error);
    return { statusCode: 500, message: "Failed to add workflow" };
  }
});