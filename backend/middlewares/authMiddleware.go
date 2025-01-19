package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := toolbox.GetBearerToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, schemas.BasicResponse{
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}
		token, _ := services.NewJWTService().ValidateJWTToken(tokenString)

		if token.Valid {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, schemas.BasicResponse{
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}
