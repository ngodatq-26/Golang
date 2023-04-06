package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yogurt-project/services"
)

type AuthApi struct {
	username string `valid:"Required; MaxSize(50)";json:"username"`
	password string `valid:"Required; MaxSize(50)";json:"password"`
}

func Login(context *gin.Context) {
	var auth AuthApi
	if err := context.Bind(&auth); err != nil {

		authService := services.AuthService{Username: auth.username, Password: auth.password}
		check, _ := authService.Check()

		if check == true {
			data := map[string]interface{}{
				"message": "successfully",
			}
			context.JSON(http.StatusOK, data)
		} else {
			data := map[string]interface{}{
				"message": "Login failure",
			}
			context.JSON(http.StatusBadRequest, data)
		}

	} else {
		data := map[string]interface{}{
			"message": "Error params",
		}
		context.JSON(http.StatusBadRequest, data)
	}
}
