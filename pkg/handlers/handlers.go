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

	render.RenderTemplate(w,"home.page.tmpl",&Models.TemplateData{})
}

//About: is the About page handler
func (m* Repository)Login(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"login.page.tmpl",&Models.TemplateData{
	})
}

func (m* Repository)View(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"view.page.tmpl",&Models.TemplateData{
	})
}
func (m* Repository)Contact(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"contact.page.tmpl",&Models.TemplateData{
	})
}
func (m* Repository)Request(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"request.page.tmpl",&Models.TemplateData{
	})
}