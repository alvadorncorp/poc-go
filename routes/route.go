package routes

import (
	"authentication_api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.POST("/auth", controllers.Authentication)
	r.GET("/user", controllers.DisplaysUser)
	r.POST("/user", controllers.RegisterUser)
	r.PATCH("/user/:id", controllers.ModifyUser)
	r.Run()
}
