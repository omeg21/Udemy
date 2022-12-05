package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/omeg21/udemy-throwaway/pkg/config"
	"github.com/omeg21/udemy-throwaway/pkg/controller"
	"github.com/omeg21/udemy-throwaway/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := controller.NewRepo(&app)
	controller.NewController(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", controller.Repo.Home)
	http.HandleFunc("/about", controller.Repo.About)
	// http.HandleFunc("/new", controller.New)
	_ = http.ListenAndServe(portNumber, nil)
	fmt.Println("what")
}
