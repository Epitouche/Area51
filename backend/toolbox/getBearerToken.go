package toolbox

import (
	"area51/schemas"
	"errors"

	"github.com/gin-gonic/gin"
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
