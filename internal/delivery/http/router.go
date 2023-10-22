package http

import (
	"heroku/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

func DefineAPIRoutes(r *gin.Engine, userController controller.UserController) {
	// Define API routes
	// public API
	v1 := r.Group("/v1")
	v1.GET("/users/:user_id", userController.GetUserInfo)

	// private API
	admin := v1.Group("/admin")
	admin.Use(AuthMiddleware)

	admin.GET("/users", userController.GetAllUserInfo)
}
