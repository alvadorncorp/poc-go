package db

import (
	"authentication_api/models"
	"errors"
)

var listaUsuarios = map[string]models.User{
	"pedro@gmail.com": {
		Username: "pedro",
		Password: "senha123",
		Email:    "pedro@gmail.com",
	},

	"igor@gmail.com": {
		Username: "igor",
		Password: "senha456",
		Email:    "igor@gmail.com",
	},
}

func BuscaUsuario(email string) (models.User, error) {
	if usuario, encontrado := listaUsuarios[email]; encontrado {
		return usuario, nil
	}
	return models.User{}, errors.New("usuario nao encontrado")
}
