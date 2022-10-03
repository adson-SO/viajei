package main

import (
	"api-viajei/src/configuration"
	"api-viajei/src/database"
	"api-viajei/src/routes"
)

func init() {
	configuration.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	server := routes.LoadRouter()

	server.Run()
}
