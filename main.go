package main

import (
	"capstoneproject/config"
	"capstoneproject/factory"
	"capstoneproject/routes"

	"capstoneproject/middlewares"
)

func main() {
	dbConn := config.InitDB()

	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":80"))
}
