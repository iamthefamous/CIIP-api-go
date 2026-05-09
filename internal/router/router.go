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
	r.GET("/posts", postHandler.GetAllPublic)
	r.GET("/posts/:id", postHandler.GetByID)
	r.GET("/universities", universityHandler.GetPublished)
	r.GET("/universities/:id", universityHandler.GetByID)
	r.GET("/faculties", facultyHandler.GetPublished)
	r.GET("/faculties/:id", facultyHandler.GetByID)
	r.GET("/programs", programHandler.GetPublished)
	r.GET("/programs/:id", programHandler.GetByID)

	r.POST("/admin/login", authHandler.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly(jwtSecret, supabaseJWTSecret, getProfileRoleByID))

	admin.GET("/posts", postHandler.GetAllAdmin)
	admin.GET("/posts/:id", postHandler.GetByID)
	admin.POST("/posts", postHandler.Create)
	admin.PUT("/posts/:id", postHandler.Update)
	admin.DELETE("/posts/:id", postHandler.Delete)

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
