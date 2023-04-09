package main

import (
	"github.com/felipesantoos/tcms/src/api/router"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres"
	"github.com/gin-contrib/cors"
)

func main() {
	configurateDatabase()
	configurateServer()
}

func configurateDatabase() {
	err := postgres.CreateTables()
	if err != nil {
		panic(err)
	}
	err = postgres.TruncateTables()
	if err != nil {
		panic(err)
	}
	err = postgres.LoadFakeData()
	if err != nil {
		panic(err)
	}
	postgres.ResetTableObjects()
}

func configurateServer() {
	server := router.New()
	server.Use(cors.Default())
	server.LoadRoutes()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
