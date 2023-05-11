package main

import (
	"net/http"

	"github.com/arpushkarev/http-web-app/pkg/config"
	"github.com/arpushkarev/http-web-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(cfg *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/about", handlers.Repo.AboutPage)
	mux.Get("/", handlers.Repo.HomePage)

	return mux
}
