package http

import (
	"heroku/config"
	"heroku/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureSwagger(cfg *config.Config, r *gin.Engine) {
	// Configure Swagger documentation
	docs.SInfo.Title = "API Documentations"
	docs.SInfo.Version = "1.0"
	docs.SInfo.Host = cfg.SwaggerDomain
	docs.SInfo.Schemes = []string{"http"}

	// Enable Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
