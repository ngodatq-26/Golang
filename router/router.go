package router

import (
	"github.com/gin-gonic/gin"
	"yogurt-project/router/api"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r1 := r.Group("/api/v1/")

	r1.POST("login", api.Login)
	return r
}
