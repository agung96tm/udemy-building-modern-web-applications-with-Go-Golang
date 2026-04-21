package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agung96tm/udemy-modern-go/pkg/config"
	"github.com/agung96tm/udemy-modern-go/pkg/handlers"
	"github.com/agung96tm/udemy-modern-go/pkg/render"
)

const portNumber = ":8080"

type Application struct {
	config config.AppConfig
}

func main() {
	var app Application
	var cfg config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	cfg.TemplateCache = tc
	cfg.UseCache = false
	app.config = cfg

	handlers.NewHandlers(handlers.NewRepo(&cfg))

	render.NewTemplate(&cfg)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Listening on port %s\n", portNumber)
	log.Fatal(srv.ListenAndServe())
}
