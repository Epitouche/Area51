package toolbox

import (
	"errors"

	"github.com/gin-gonic/gin"

	"area51/schemas"
)

type Bearer interface {
	GetBearerToken() (string, error)
}

func GetBearerToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", schemas.ErrNoAuthorizationHeaderFound
	}
	if len("Bearer ") >= len(authHeader) {
		return "", errors.New("invalid authorization")
	}
	return authHeader[len("Bearer "):], nil
}
