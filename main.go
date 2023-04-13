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

	r.Run(":3000")
}
