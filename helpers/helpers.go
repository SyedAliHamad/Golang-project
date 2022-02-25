package helpers

import (
	"net/http"

	"github.com/SyedAliHamad/internproject/pkg/config"
)



var app *config.AppConfig
func NewHelpers (a *config.AppConfig){
	app =a
}
func IsAuthenticated(r *http.Request)bool{
	exists:=app.Session.Exists(r.Context(),"user_id")
	return exists
}