package main

import (
	//"fmt"
	"net/http"

	"github.com/SyedAliHamad/internproject/helpers"
	"github.com/justinas/nosurf"
	"tawesoft.co.uk/go/dialog"
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

func Auth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		if !helpers.IsAuthenticated(r){
			session.Put(r.Context(),"error","Log in first")
			http.Redirect(w,r,"/login",http.StatusSeeOther)
			dialog.Alert("Log in first")
			return
		}
	next.ServeHTTP(w,r)
	})
}