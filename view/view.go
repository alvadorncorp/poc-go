package view

import "authentication_api/models"

type ViewUsuario struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewViewUsuario(usuarios []models.User) []ViewUsuario {
	viewUsuarios := make([]ViewUsuario, 0, len(usuarios))
	for _, u := range usuarios {
		viewUsuarios = append(viewUsuarios, ViewUsuario{
			ID:       u.ID,
			Username: *u.Username,
			Email:    *u.Email,
		})
	}

	return viewUsuarios
}
