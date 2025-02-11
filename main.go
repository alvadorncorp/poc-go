package main

import (
	"authentication_api/db"
	routes "authentication_api/routes"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	db.Init()

	urlExample := "postgres://root:root@localhost:5432/root"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var xpto bool
	err = conn.QueryRow(context.Background(), "select 1=1").Scan(&xpto)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(xpto)

	routes.HandleRequests()
}
