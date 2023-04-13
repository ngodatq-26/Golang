package main

import (
	"fmt"
	"yogurt-project/pkg/setting"
	"yogurt-project/router"
)

func main() {
	fmt.Println("hello")

	r := router.Setup()
	setting.SetEnvironmentInit()

	r.LoadHTMLGlob("resources/templates/html/*")
	r.Static("/css", "resources/templates/css/")
	r.Static("/images", "resources/images/")
	r.Static("/js", "resources/templates/js/")

	r.Run(":3000")
}
