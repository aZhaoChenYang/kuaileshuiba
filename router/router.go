package router

import (
	"github.com/gin-gonic/gin"
	"klsb/controller"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/get", controller.Get)
	r.GET("/api/getjb", controller.Getjb)
	return r
}
