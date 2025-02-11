package controllers

import (
	"authentication_api/db"
	"authentication_api/models"
	"authentication_api/view"

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

	usuario, err := db.BuscaUsuario(usuarioJSON.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "authentication_failure"})
		return
	}

	if usuario.ComparePassword(usuarioJSON.Password) {
		c.JSON(200, gin.H{"token": CreateToken()})
	} else {
		c.JSON(401, gin.H{"error": "authentication_failure"})
	}
}

func RegistaUsuario(c *gin.Context) {
	var usuarioJSON models.User

	if err := c.ShouldBindJSON(&usuarioJSON); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if err := usuarioJSON.EncryptPassword(); err != nil {
		c.JSON(400, gin.H{"error": "Failed to encrypt password"})
		return
	}

	if err := db.CriaUsuario(usuarioJSON); err != nil {
		c.JSON(400, gin.H{"error": "Ivaled request"})
		return
	}
	c.JSON(201, gin.H{"message": "Usuario criado com sucesso"})

}

func ExibeUsuarios(c *gin.Context) {
	c.JSON(200, view.NewViewUsuario(db.ListarUsuarios()))
}
