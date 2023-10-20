package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureCORS(r *gin.Engine) {
	// Configure CORS
	r.Use(gin.Recovery())
	r.Use(cors.Default()) // Use the default CORS configuration
}
