package controller

import (
	"errors"
	uc "heroku/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Controller
	UserUseCase uc.UserUseCase
}

// @BasePath /v1/users/{user_id}
// PingExample godoc
// @Summary get user infor to return the name, account list and balances
// @Schemes
// @Description get user info
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/users/{user_id} [get]
func (uc *UserController) GetUserInfo(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		uc.Failure(ctx, errors.New("user_id must is number"))
		return
	}

	userInfo, err := uc.UserUseCase.GetUserInfo(userID)
	if err != nil {
		uc.Failure(ctx, err)
		return
	}

	uc.Success(ctx, userInfo)
}
