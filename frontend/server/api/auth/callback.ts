export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.code || !params.state || !params.service) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: code, service or state",
    });
  }

  try {
    const response = await $fetch(
      `http://server:8080/api/${params.service}/callback`,
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