package main

import (
	"authentication_api/db"
	routes "authentication_api/routes"
)

func main() {
	db.Init()
	db.ConectDB()
	routes.HandleRequests()
}
