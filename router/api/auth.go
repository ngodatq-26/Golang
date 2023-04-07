package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yogurt-project/services"
)

type AuthApi struct {
	Username string `valid:"Required; MaxSize(50)";json:"username"`
	Password string `valid:"Required; MaxSize(50)";json:"password"`
}

func Login(context *gin.Context) {
	var auth AuthApi
	if err := context.ShouldBindJSON(&auth); err != nil {
		data := map[string]interface{}{
			"message": "Error params",
		}
		context.JSON(http.StatusBadRequest, data)

	} else {
		fmt.Println(auth)
		authService := services.AuthService{Username: auth.Username, Password: auth.Password}
		check, _ := authService.LoginService()

		if check == true {
			data := map[string]interface{}{
				"message": "Login successfully",
			}
			context.JSON(http.StatusOK, data)
		} else {
			data := map[string]interface{}{
				"message": "Login failure",
			}
			context.JSON(http.StatusBadRequest, data)
		}
	}
}

func Register(context *gin.Context) {
	var auth AuthApi
	if err := context.ShouldBindJSON(&auth); err != nil {
		data := map[string]interface{}{
			"message": "Error params",
		}
		context.JSON(http.StatusBadRequest, data)

	} else {
		fmt.Println(auth)
		authService := services.AuthService{Username: auth.Username, Password: auth.Password}
		check, _ := authService.RegisterService()

		if check == true {
			data := map[string]interface{}{
				"message": "Register successfully",
			}
			context.JSON(http.StatusOK, data)
		} else {
			data := map[string]interface{}{
				"message": "Register failure",
			}
			context.JSON(http.StatusBadRequest, data)
		}
	}

}
