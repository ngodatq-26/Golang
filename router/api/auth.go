package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yogurt-project/pkg/util"
	"yogurt-project/services"
)

type AuthApi struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(context *gin.Context) {
	var auth AuthApi

	if err := context.ShouldBind(&auth); err != nil {
		data := map[string]interface{}{
			"message": "Error params",
		}
		context.JSON(http.StatusBadRequest, data)
		return
	}

	authService := services.AuthService{Username: auth.Username, Password: auth.Password}
	check, _ := authService.LoginService()

	if check == true {
		token, _ := util.GenerateToken("test", auth)
		fmt.Println(token)
		data := map[string]interface{}{
			"message": "Login successfully",
			"token":   token,
		}
		context.JSON(http.StatusOK, data)
		return
	}

	data := map[string]interface{}{
		"message": "Login failure",
	}
	context.JSON(http.StatusBadRequest, data)
}

func Register(context *gin.Context) {
	var auth AuthApi

	if err := context.ShouldBind(&auth); err != nil {
		data := map[string]interface{}{
			"message": "Error params",
		}
		context.JSON(http.StatusBadRequest, data)
		return
	}

	authService := services.AuthService{Username: auth.Username, Password: auth.Password}
	check, _ := authService.RegisterService()

	if check == true {
		data := map[string]interface{}{
			"message": "Register successfully",
		}
		context.JSON(http.StatusOK, data)
		return
	}

	data := map[string]interface{}{
		"message": "Register failure",
	}
	context.JSON(http.StatusBadRequest, data)
}
