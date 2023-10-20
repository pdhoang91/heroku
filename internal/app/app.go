package app

import (
	"heroku/config"
	"heroku/internal/delivery/http"
	controller "heroku/internal/delivery/http/controller"
	repo "heroku/internal/repository"
	uc "heroku/internal/usecase"

	"github.com/gin-gonic/gin"
)

func InitializeHTTPServer(cfg *config.Config) *gin.Engine {
	// Create a new Gin router
	r := gin.New()

	// Enable CORS with custom settings
	http.ConfigureCORS(r)

	// Initialize user-related components
	userController := initUserController()

	// Configure Swagger documentation
	http.ConfigureSwagger(cfg, r)

	// Define API routes
	http.DefineAPIRoutes(r, userController)

	return r
}

func initUserController() controller.UserController {
	// Initialize user repository and use case
	userRepository := repo.NewUserRepository()
	userUseCase := uc.NewUserUseCase(userRepository)

	// Create and return the UserController
	return controller.UserController{UserUseCase: userUseCase}
}
