package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/simplifywoopii88/modern_web_applications_with_go/pkg/config"
	"github.com/simplifywoopii88/modern_web_applications_with_go/pkg/handlers"
	"github.com/simplifywoopii88/modern_web_applications_with_go/pkg/render"
)

const portNumber = ":8080"



func main(){
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create temlate cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	 
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("listening server http://localhost%s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}