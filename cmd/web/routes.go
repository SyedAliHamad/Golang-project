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
	mux.Use(Nosurf)
	mux.Use(SessionLoad)

	mux.Get("/",handlers.Repo.Home)

	mux.Get("/login",handlers.Repo.Login)
	mux.Post("/login",handlers.Repo.PostLogin)

	mux.Get("/logout",handlers.Repo.Logout)

	mux.Get("/signup",handlers.Repo.Signup)
	mux.Post("/signup",handlers.Repo.PostSignup)

	mux.Get("/contact",handlers.Repo.Contact)
	mux.Post("/contact",handlers.Repo.PostContact)

	mux.Route("/upload",func(mux chi.Router){
		mux.Use(Auth)
		mux.Get("/",handlers.Repo.Upload)
		mux.Post("/",handlers.Repo.PostUpload)
	})


	mux.Route("/request",func(mux chi.Router){
		mux.Use(Auth)
		mux.Get("/",handlers.Repo.Request)
		mux.Post("/",handlers.Repo.PostRequest)
	})

	mux.Route("/view",func(mux chi.Router){
		mux.Use(Auth)
		mux.Get("/",handlers.Repo.View)
		mux.Post("/",handlers.Repo.PostView)
	})

	fileServer:= http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}