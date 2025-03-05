package controllers

import (
	"authentication_api/db"
	"authentication_api/models"
	"authentication_api/view"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken() string {
	var (
		key []byte
		T   *jwt.Token
		S   string
	)
	key = []byte("7zZruvzBoMpYCFFsFYMj")
	T = jwt.New(jwt.SigningMethodHS256)
	S, _ = T.SignedString(key)
	return string(S)
}

func Authentication(c *gin.Context) {
	var usuarioJSON models.User
	if err := c.ShouldBindJSON(&usuarioJSON); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	usuario, err := db.FindUser(*usuarioJSON.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "authentication_failure"})
		return
	}

	if usuario.ComparePassword(*usuarioJSON.Password) {
		c.JSON(200, gin.H{"token": CreateToken()})
	} else {
		c.JSON(401, gin.H{"error": "authentication_failure"})
	}
}

func RegisterUser(c *gin.Context) {
	var usuarioJSON models.User

	if err := c.ShouldBindJSON(&usuarioJSON); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if err := usuarioJSON.EncryptPassword(); err != nil {
		c.JSON(400, gin.H{"error": "Failed to encrypt password"})
		return
	}

	if err := db.CreateUser(&usuarioJSON); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "request failed"})
		return
	}
	c.JSON(201, gin.H{"message": "Usuario criado com sucesso"})

}

func DisplaysUser(c *gin.Context) {
	c.JSON(200, view.NewViewUsuario(db.ListUsers()))
}

func DisplayUserId(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	usuario, err := db.FindUserByID(int(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, view.NewViewUsuarioById(usuario))
}

func ModifyUser(c *gin.Context) {
	var usuarioJSON models.User

	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := c.ShouldBindJSON(&usuarioJSON); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if usuarioJSON.Password != nil {
		if err := usuarioJSON.EncryptPassword(); err != nil {
			c.JSON(400, gin.H{"error": "Failed to encrypt password"})
			return
		}
	}

	if err := db.UpdateUser(int(id), &usuarioJSON); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "request failed"})
		return
	}
	c.JSON(200, gin.H{"message": "Usuario modificado com sucesso"})
}
