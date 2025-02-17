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

var dbase *pgx.Conn

func Init() {
	for _, u := range listaUsuarios {
		u.EncryptPassword()
	}
}

func CreateUser(usuario *models.User) error {
	ctx := context.Background()
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
	_, err := dbase.Exec(ctx, query, usuario.Username, usuario.Password, usuario.Email)
	return err
}

func FindUser(email string) (models.User, error) {
	ctx := context.Background()
	query := "SELECT username, password, email FROM users WHERE email=$1"
	row := dbase.QueryRow(ctx, query, email)
	var usuario models.User
	err := row.Scan(&usuario.Username, &usuario.Password, &usuario.Email)
	if err == pgx.ErrNoRows {
		return models.User{}, errors.New("usuario nao encontrado")
	}
	return usuario, err
}

func ListUsers() []models.User {
	var usuarios []models.User
	ctx := context.Background()
	rows, err := dbase.Query(ctx, "SELECT username, password, email FROM users")
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var usuario models.User
		err := rows.Scan(&usuario.Username, &usuario.Password, &usuario.Email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		usuarios = append(usuarios, usuario)
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
	dbase = conn

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
