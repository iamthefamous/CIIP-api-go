package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamthefamous/CIIP-api-go/internal/handler"
)

type Router struct {
	postHandler *handler.PostHandler
}

func NewRouter(ph *handler.PostHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/posts", ph.Create)
	r.GET("/posts", ph.GetAll)

	return r
}
