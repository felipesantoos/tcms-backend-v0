package main

import (
	"github.com/felipesantoos/tcms/src/api/router"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres"
)

func main() {
	configurateDatabase()
	configurateServer()
}

func configurateDatabase() {
	err := postgres.NewConnector().CreateTables()
	if err != nil {
		panic(err)
	}
}

func configurateServer() {
	server := router.New()
	server.LoadRoutes()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
