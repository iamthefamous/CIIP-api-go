package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamthefamous/CIIP-api-go/internal/handler"
	"github.com/iamthefamous/CIIP-api-go/internal/middleware"
)

func NewRouter(
	postHandler *handler.PostHandler,
	authHandler *handler.AuthHandler,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	// public routes
	r.POST("/posts", postHandler.Create)
	r.GET("/posts", postHandler.GetAll)

	r.POST("/admin/login", authHandler.Login)

	// protected admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly(jwtSecret))

	admin.GET("/posts", postHandler.GetAll)
	admin.POST("/posts", postHandler.Create)

	return r
}
