export default defineEventHandler(async (event) => {
  const { code, state } = getQuery(event);

  if (!code || !state) {
    throw createError({
      statusCode: 400,
      statusMessage: "Missing required parameters: code or state",
    });
  }

  try {
    const response = await $fetch("http://localhost:8080/api/github/callback", {
      params: {
        code: code,
        state: state,
      },
      headers: {
        "Content-Type": "application/json",
      },
      method: "GET",
    });

    if (!response) {
      throw createError({
        statusCode: 500,
        statusMessage: "Access token not found in response",
      });
    }

    console.log("Response from callback:", response);
    return response;
  } catch (error) {
    console.error("Error during GitHub callback:", error);

    throw createError({
      statusCode: 502,
      statusMessage: "Failed to fetch from GitHub callback endpoint",
    });
  }
});
