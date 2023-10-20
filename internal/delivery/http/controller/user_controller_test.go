package controller

import (
	"heroku/internal/delivery/http/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock User UseCase
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) GetUserInfo(userID int) (*model.UserInfo, error) {
	args := m.Called(userID)
	return args.Get(0).(*model.UserInfo), args.Error(1)
}

func TestUserController_GetUserInfo(t *testing.T) {
	// Create a new instance of the UserController with a mock UserUseCase
	mockUserUseCase := new(MockUserUseCase)
	controller := &UserController{
		UserUseCase: mockUserUseCase,
	}

	// Create a Gin router to simulate the HTTP request
	router := gin.Default()
	router.GET("/v1/users/:user_id", controller.GetUserInfo)

	// Test case 1: Valid user ID
	mockUserUseCase.On("GetUserInfo", 1).Return(&model.UserInfo{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/users/1", nil)
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is 200 (Success)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case 2: Invalid user ID (not a number)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/users/invalid", nil)
	router.ServeHTTP(w, req)

	// Assert that the HTTP status code is 400 (Bad Request)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the Failure method was called for the error case
	mockUserUseCase.AssertExpectations(t)
}
