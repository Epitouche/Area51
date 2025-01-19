package toolbox

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/schemas"
)

type ErrorManagement interface {
	HandleError(err error, statusOKValue interface{})
}

func HandleError(ctx *gin.Context, err error, statusOKValue interface{}) {
	switch err {
	case schemas.ErrorBadParameter:
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Message: err.Error(),
		})
		return
	case schemas.ErrorNoWorkflowFound:
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Message: err.Error(),
		})
		return
	case schemas.ErrorAlreadyExistingRessource:
		ctx.JSON(http.StatusConflict, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrReactionNotFound:
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrActionNotFound:
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrUserNotFound:
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrorNoServiceFound:
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrorInvalidToken:
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrWhileLinking:
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrorHashingPassword:
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Message: err.Error(),
		})
	case schemas.ErrNoAuthorizationHeaderFound:
		return
	default:
		ctx.JSON(http.StatusOK, statusOKValue)
	}
}
