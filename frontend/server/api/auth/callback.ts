export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.code || !params.state) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: code, or state",
    });
  }

  try {
    const response = await $fetch(
      `http://server:8080/api/github/callback`,
      {
        method: "POST",
        body: {
          code: params.code,
          state: params.state,
        },
        headers: {
          Authorization: params.authorization ? `${params.authorization}` : "",
        },
      },
    );
    return response;
  } catch (error) {
    console.error(error);
  }
});