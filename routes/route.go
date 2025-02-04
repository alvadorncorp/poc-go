package routes

import (
	"authentication_api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.POST("/", controllers.Authentication)
	r.Run()
}
