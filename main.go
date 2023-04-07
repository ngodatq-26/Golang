package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yogurt-project/router"
)

func main() {
	fmt.Println("hello")

	r := router.Setup()

	r.LoadHTMLGlob("resources/templates/html/*")
	r.Static("/css", "resources/templates/css/")
	r.Static("/images", "resources/images/")
	r.Static("/js", "resources/templates/js/")

	r.GET("/login", func(context *gin.Context) {
		context.HTML(200, "auth.tmpl.html", gin.H{})
	})

	r.GET("/home", func(context *gin.Context) {
		context.HTML(200, "index.tmpl.html", gin.H{})
	})

	r.Run(":3000")
}
