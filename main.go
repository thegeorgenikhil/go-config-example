package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thegeorgenikhil/go-config-example/config"
	"github.com/thegeorgenikhil/go-config-example/database"
	"github.com/thegeorgenikhil/go-config-example/handlers"
)

func main() {
	var app config.AppConfig
	app.DatabaseURI = "mongodb://127.0.0.1:27017/myapp"
	app.DatabaseName = "userDB"

	dbClient := database.Connect(app.DatabaseURI)

	app.Client = dbClient

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/users", handlers.Repo.UserHandler)

	fmt.Println("Server started at port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
