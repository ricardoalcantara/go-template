package foo

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/go-template/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/foo", list)
	protected.POST("/foo", post)
	protected.GET("/foo/:fooId", get)
	protected.PUT("/foo/:fooId", put)
	protected.DELETE("/foo/:fooId", delete)
}
