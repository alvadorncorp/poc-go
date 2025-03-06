package routes

import (
	"authentication_api/controllers"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func protectedHandler(c *gin.Context) {
	var secretKey = []byte("7zZruvzBoMpYCFFsFYMj")
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token ausente"))
		return
	}
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token inválido"))
		return
	}

	if token.Valid {
		c.Next()
	} else {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token expirado"))
	}
}

func HandleRequests() {
	r := gin.Default()

	r.POST("/auth", controllers.Authentication)

	protected := r.Group("/user")
	protected.Use(protectedHandler)
	{
		protected.GET("/", controllers.DisplaysUser)
		protected.GET("/:id", controllers.DisplayUserId)
		protected.POST("/", controllers.RegisterUser)
		protected.PATCH("/:id", controllers.ModifyUser)
	}

	r.Run()
}
