export default defineEventHandler(async (event) => {
    const access_token = event.headers.get("Authorization");
    const contentTypes = event.headers.get("Content-Type");

    if (!access_token || !contentTypes) {
        throw createError({
            statusCode: 400,
            message: "Missing parameters: token or content type",
        });
    }
    try {
        const response = await $fetch("http://server:8080/api/workflow/reactions", {
            method: "GET",
            headers: {
                "Authorization": access_token ? `${access_token}` : "",
                "Content-Type": contentTypes ? `${contentTypes}` : "",
            },
        });

        return response;
    } catch (error) {
        console.error("Error in API server:", error);
        return { statusCode: 500, message: "Failed to get last workflow" };
    }
});