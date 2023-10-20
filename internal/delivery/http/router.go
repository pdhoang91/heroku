package http

import (
	"heroku/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

func DefineAPIRoutes(r *gin.Engine, userController controller.UserController) {
	// Define API routes
	v1 := r.Group("/v1")
	v1.GET("/users/:user_id", userController.GetUserInfo)
}
