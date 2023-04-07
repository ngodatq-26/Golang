package router

import (
	"github.com/gin-gonic/gin"
	"yogurt-project/router/api"
)

func Setup() *gin.Engine {
	r := gin.New()
	r1 := r.Group("/api/v1/")

	r1.POST("login", api.Login)
	r1.POST("register", api.Register)
	return r
}
