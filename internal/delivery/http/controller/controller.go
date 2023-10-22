package controller

import (
	"heroku/internal/delivery/http/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (h *Controller) Success(ctx *gin.Context, data interface{}) {
	response := model.SuccessResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   data,
	}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}

func (h *Controller) Failure(ctx *gin.Context, err error) {
	response := model.ErrorResponse{
		Status:  "error",
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
	ctx.JSON(http.StatusBadRequest, response)
	ctx.Abort()
}
