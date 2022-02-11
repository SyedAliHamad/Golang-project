package main

import (
	"net/http"

	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/SyedAliHamad/internproject/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig)http.Handler{

	mux:=chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoServe)
	mux.Use(SessionLoad)

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/login",handlers.Repo.Login)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Get("/view",handlers.Repo.View)
	mux.Post("/view",handlers.Repo.PostView)
	mux.Get("/request",handlers.Repo.Request)

	fileServer:= http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}