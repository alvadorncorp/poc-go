package routes

import (
	"authentication_api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.POST("/auth", controllers.Authentication)
	r.GET("/usuarios", controllers.ExibeUsuarios)
	r.POST("/usuarios", controllers.RegistaUsuario)
	r.Run()
}
