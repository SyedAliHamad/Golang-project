package handlers

import (
	"github.com/SyedAliHamad/internproject/pkg/Models"
	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/SyedAliHamad/internproject/pkg/render"

	"net/http"
)

//Repo: The repository used by the handlers
var Repo *Repository


//Repository: creates a new repository
type Repository struct{
	App *config.AppConfig
}


//NewRepo: creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}


//NewHandlers: sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo=r
}


//Home: is the home page handler
func (m* Repository)Home(w http.ResponseWriter, r* http.Request){
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	render.RenderTemplate(w,"home.page.tmpl",&Models.TemplateData{})
}

//About: is the About page handler
func (m* Repository)About(w http.ResponseWriter, r *http.Request){

	stringMap :=make(map[string]string)
	stringMap["test"]="hello,again."

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIP

	render.RenderTemplate(w,"about.page.tmpl",&Models.TemplateData{
		StringMap: stringMap,
	})
}
