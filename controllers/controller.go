package controllers

import (
	"authentication_api/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var authorizedUser = "pedro"
var authorizedPassword = "senha123"

var (
	key []byte
	T   *jwt.Token
	S   string
)

func CreatToken() string {
	key = []byte("7zZruvzBoMpYCFFsFYMj")
	T = jwt.New(jwt.SigningMethodHS256)
	S, _ = T.SignedString(key)
	return string(S)
}

func Authentication(c *gin.Context) {
	var username models.User
	if err := c.ShouldBindJSON(&username); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println(username)

	if username.Username == authorizedUser && username.Password == authorizedPassword {
		c.JSON(202, gin.H{"message": CreatToken()})
	} else {
		c.JSON(401, gin.H{"error": "authentication_failure"})
	}
}
