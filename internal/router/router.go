package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamthefamous/CIIP-api-go/internal/handler"
	"github.com/iamthefamous/CIIP-api-go/internal/middleware"
)

func NewRouter(
	postHandler *handler.PostHandler,
	authHandler *handler.AuthHandler,
	universityHandler *handler.UniversityHandler,
	facultyHandler *handler.FacultyHandler,
	programHandler *handler.ProgramHandler,
	jwtSecret string,
	supabaseJWTSecret string,
	getProfileRoleByID func(userID string) (string, error),
) *gin.Engine {
	r := gin.Default()

	// public routes
	r.POST("/posts", postHandler.Create)
	r.GET("/posts", postHandler.GetAll)

	r.POST("/admin/login", authHandler.Login)

	r.GET("/universities", universityHandler.GetAll)
	r.GET("/universities/:id", universityHandler.GetByID)
	r.GET("/faculties", facultyHandler.GetAll)
	r.GET("/faculties/:id", facultyHandler.GetByID)
	r.GET("/programs", programHandler.GetAll)
	r.GET("/programs/:id", programHandler.GetByID)

	// protected admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly(jwtSecret, supabaseJWTSecret, getProfileRoleByID))

	admin.GET("/posts", postHandler.GetAll)
	admin.POST("/posts", postHandler.Create)

	admin.GET("/universities", universityHandler.GetAll)
	admin.GET("/universities/:id", universityHandler.GetByID)
	admin.POST("/universities", universityHandler.Create)
	admin.PUT("/universities/:id", universityHandler.Update)
	admin.DELETE("/universities/:id", universityHandler.Delete)

	admin.GET("/faculties", facultyHandler.GetAll)
	admin.GET("/faculties/:id", facultyHandler.GetByID)
	admin.POST("/faculties", facultyHandler.Create)
	admin.PUT("/faculties/:id", facultyHandler.Update)
	admin.DELETE("/faculties/:id", facultyHandler.Delete)

	admin.GET("/programs", programHandler.GetAll)
	admin.GET("/programs/:id", programHandler.GetByID)
	admin.POST("/programs", programHandler.Create)
	admin.PUT("/programs/:id", programHandler.Update)
	admin.DELETE("/programs/:id", programHandler.Delete)

	return r
}
