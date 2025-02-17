package db

import (
	"authentication_api/models"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var listaUsuarios = map[string]*models.User{
	"pedro@gmail.com": {
		Username: "pedro",
		Password: "senha123",
		Email:    "pedro@gmail.com",
	},
}

var db *pgx.Conn

func Init() {
	for _, u := range listaUsuarios {
		u.EncryptPassword()
	}
}

func SearchUser(email string) (models.User, error) {
	if usuario, encontrado := listaUsuarios[email]; encontrado {
		return *usuario, nil
	}
	return models.User{}, errors.New("usuario nao encontrado")
}

func CreateUser(usuario models.User) error {
	if _, encontrado := listaUsuarios[usuario.Email]; encontrado {
		return errors.New("email ja cadastrado")
	}
	listaUsuarios[usuario.Email] = &usuario
	return nil
}

func FindUser(email string) (models.User, error) {
	if usuario, encontrado := listaUsuarios[email]; encontrado {
		return *usuario, nil
	}
	return models.User{}, errors.New("usuario nao encontrado")
}

func ListUsers() []models.User {
	usuarios := make([]models.User, 0, len(listaUsuarios))
	for _, usuario := range listaUsuarios {
		usuarios = append(usuarios, *usuario)
	}

	return usuarios
}

func ConectDB() {
	urlExample := "postgres://root:root@localhost:5432/root"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var xpto bool
	err = conn.QueryRow(context.Background(), "SELECT 1=1").Scan(&xpto)
	if err != nil {
		fmt.Println("QueryRow failed:", err)
		os.Exit(1)
	}

	fmt.Println(xpto)

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL
	);`
	_, err = conn.Exec(context.Background(), createTableQuery)
	if err != nil {
		fmt.Println("Error creating table:", err)
		os.Exit(1)
	}

	fmt.Println("Tabela 'users' criada com sucesso!")

	for _, user := range listaUsuarios {
		insertUserQuery := `
		INSERT INTO users (username, password, email) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (email) DO NOTHING;`

		_, err := conn.Exec(context.Background(), insertUserQuery, user.Username, user.Password, user.Email)
		if err != nil {
			fmt.Println("Failed to insert user:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Usuários inseridos com sucesso!")
}
