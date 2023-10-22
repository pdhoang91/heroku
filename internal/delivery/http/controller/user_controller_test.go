package controller

import (
	"heroku/internal/delivery/http/controller/mocks"
	"heroku/internal/delivery/http/model"
	"heroku/internal/usecase"
	iError "heroku/pkg/error"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserController_GetUserInfo(t *testing.T) {
	type fields struct {
		userUseCase usecase.UserUseCase
	}
	type args struct {
		mockCtx func(ctx *gin.Context)
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse string
		expectedStatus   int
	}{
		{
			name: "Happy Case",
			fields: fields{
				userUseCase: func() usecase.UserUseCase {
					mockService := &mocks.UserUseCase{}
					mockService.On("GetUserInfo", 1).Return(&model.UserInfo{
						UserID:  1,
						Name:    "Alice",
						Balance: 145000,
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Params = gin.Params{
						{Key: "user_id", Value: "1"},
					}
					ctx.Request = httptest.NewRequest(
						http.MethodGet,
						"/v1/users/1",
						nil,
					)
				},
			},
			expectedResponse: "{\"status\":\"success\",\"code\":200,\"data\":{\"user_id\":1,\"name\":\"Alice\",\"accounts\":null,\"total_balance\":145000}}",
			expectedStatus:   http.StatusOK,
		},
		{
			name: "Case empty value",
			fields: fields{
				userUseCase: func() usecase.UserUseCase {
					mockService := &mocks.UserUseCase{}
					mockService.On("GetUserInfo", 0).Return(nil, iError.NewErrorHandler(400, "User [0] not found"))
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Params = gin.Params{
						{Key: "user_id", Value: "0"},
					}
					ctx.Request = httptest.NewRequest(
						http.MethodGet,
						"/v1/users/0",
						nil,
					)
				},
			},
			expectedResponse: "{\"status\":\"error\",\"code\":400,\"message\":\"User [0] not found\"}",
			expectedStatus:   http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c := &UserController{
				UserUseCase: tt.fields.userUseCase,
			}
			tt.args.mockCtx(ctx)

			c.GetUserInfo(ctx)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
