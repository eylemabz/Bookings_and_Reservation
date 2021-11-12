package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eylemabz/go-course/pkg/config"
	"github.com/eylemabz/go-course/pkg/handlers"
	"github.com/eylemabz/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}