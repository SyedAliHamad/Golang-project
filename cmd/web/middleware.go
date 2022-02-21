package main

import (
	//"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//Nosurf: adds CSRF protection to all POST requests
func Nosurf(next http.Handler) http.Handler{
	csrfHandler :=nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:"/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
return csrfHandler
}

//Sessionoads: Loads and saves the session to every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}