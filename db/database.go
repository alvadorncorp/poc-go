package db

import (
	"authentication_api/models"
	"errors"
)

var listaUsuarios = map[string]*models.User{
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

func Init() {
	for _, u := range listaUsuarios {
		u.EncryptPassword()
	}
}

func BuscaUsuario(email string) (models.User, error) {
	if usuario, encontrado := listaUsuarios[email]; encontrado {
		return *usuario, nil
	}
	return models.User{}, errors.New("usuario nao encontrado")
}

func CriaUsuario(usuario models.User) error {
	if _, encontrado := listaUsuarios[usuario.Email]; encontrado {
		return errors.New("email ja cadastrado")
	}
	listaUsuarios[usuario.Email] = &usuario
	return nil
}

func EncontraUsuario(email string) (models.User, error) {
	if usuario, encontrado := listaUsuarios[email]; encontrado {
		return *usuario, nil
	}
	return models.User{}, errors.New("usuario nao encontrado")
}

func ListarUsuarios() []models.User {
	usuarios := make([]models.User, 0, len(listaUsuarios))
	for _, usuario := range listaUsuarios {
		usuarios = append(usuarios, *usuario)
	}

	return usuarios
}
