package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/agung96tm/udemy-modern-go/pkg/config"
	"github.com/agung96tm/udemy-modern-go/pkg/handlers"
	"github.com/agung96tm/udemy-modern-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app Application
var cfg config.AppConfig
var session *scs.SessionManager

type Application struct {
	config config.AppConfig
}

func main() {
	app.config = cfg
	cfg.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.config.InProduction
	cfg.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	cfg.TemplateCache = tc
	cfg.UseCache = false

	handlers.NewHandlers(handlers.NewRepo(&cfg))

	render.NewTemplate(&cfg)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Listening on port %s\n", portNumber)
	log.Fatal(srv.ListenAndServe())
}
