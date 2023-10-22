package controller

import (
	"heroku/internal/delivery/http/model"
	uc "heroku/internal/usecase"

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

	var userRequest model.UserRequest
	if err := ctx.ShouldBindUri(&userRequest); err != nil {
		uc.Failure(ctx, err)
		return
	}

	userInfo, err := uc.UserUseCase.GetUserInfo(*userRequest.UserID)
	if err != nil {
		uc.Failure(ctx, err)
		return
	}

	uc.Success(ctx, userInfo)
}

// @BasePath /v1/admin/users
// PingExample godoc
// @Summary Get all user infor for admin
// @Schemes
// @Description get all user info
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/admin/users [get]
func (uc *UserController) GetAllUserInfo(ctx *gin.Context) {
	userInfo, err := uc.UserUseCase.GetAllUserInfo()
	if err != nil {
		uc.Failure(ctx, err)
		return
	}

	uc.Success(ctx, userInfo)
}
