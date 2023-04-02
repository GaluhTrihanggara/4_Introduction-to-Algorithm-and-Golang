package main

import (
	"Praktikum/routes"
	"Praktikum/config"
	"Praktikum/middlewares"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := routes.New()
	// implement middleware logger
	middlewares.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
