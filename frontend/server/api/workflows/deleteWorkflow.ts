export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  const access_token = event.headers.get("Authorization");
  console.log("params", params);
  if (!params.action_id || !access_token || !params.reaction_id || !params.name || !params.workflow_id) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: action_id, reaction_id, name, or workflow_id",
    });
  }

  try {
    const response = await $fetch(
      `http://server:8080/api/workflow`,
      {
        method: "DELETE",
        headers: {
          Authorization: access_token ? `${access_token}` : "",
          "Content-Type": "application/json",
        },
        body: {
          action_id: params.action_id,
          reaction_id: params.reaction_id,
          name: params.name,
          workflow_id: params.workflow_id
        },
      }
    );
    return response;
  } catch (error) {
    console.error(error);
  }
});
