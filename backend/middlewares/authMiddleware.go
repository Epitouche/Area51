package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) < len("Bearer ") {
			ctx.JSON(http.StatusUnauthorized, schemas.BasicResponse{
				Message: "Unauthorized because no token provided",
			})
			return
		}
		tokenString := authHeader[len("Bearer "):]
		token, _ := services.NewJWTService().ValidateJWTToken(tokenString)

		if token.Valid {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, schemas.BasicResponse{
				Message: "Unauthorized",
			})
		}
	}
}
