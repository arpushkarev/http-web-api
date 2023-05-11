package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/arpushkarev/http-web-app/pkg/config"
	"github.com/arpushkarev/http-web-app/pkg/handlers"
	"github.com/arpushkarev/http-web-app/pkg/render"
)

const port = ":8080"

var cfg config.AppConfig
var session *scs.SessionManager

func main() {

	//change this into true before the production
	cfg.ProductionMode = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	cfg.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Failed creating template's cache:%s", err.Error())
	}

	cfg.TemplateCache = tc
	cfg.UseCache = false

	repo := handlers.NewRepo(&cfg)
	handlers.HandlersRepo(repo)

	render.NewTemplates(&cfg)

	fmt.Printf("HTTP server is running on port %s", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&cfg),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
