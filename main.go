package main

import (
	"github.com/DudhaneShrey86/cake_app_back/connection"
	"github.com/DudhaneShrey86/cake_app_back/routes"
)

func main() {
	closeConnection := connection.ConnectToDB()
	defer closeConnection()
	routes.RegisterRoutes()
}
