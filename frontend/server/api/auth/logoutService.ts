export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  const access_token = event.headers.get("Authorization");

  const jsonParams = JSON.parse(params);

  if (!access_token || !jsonParams.service_name) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token or service",
    });
  }

  try {
    const response = await $fetch(`http://server:8080/api/user/service/logout`, {
      method: "PUT",
      headers: {
        Authorization: access_token ? `${access_token}` : "",
        "Content-Type": "application/json",
      },
      body: {
        service_name: jsonParams.service_name,
      },
    });

    return response;
  } catch (error) {
    console.error(error);
  }
});
