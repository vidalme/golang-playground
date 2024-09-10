package main

import (
	"github.com/vidalme/api-go-gin/database"
	"github.com/vidalme/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
