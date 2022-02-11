package handlers

import (
	"fmt"

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

	render.RenderTemplate(w,r,"home.page.tmpl",&Models.TemplateData{})
}

//About: is the About page handler
func (m* Repository)Login(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"login.page.tmpl",&Models.TemplateData{
	})
}

func (m* Repository)View(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"view.page.tmpl",&Models.TemplateData{
	})
}


func (m* Repository)PostView(w http.ResponseWriter, r *http.Request){
	
	loginname :=r.Form.Get("login_name")
	loginpasswod:=r.Form.Get("login_password")

	signupname:= r.Form.Get("signup_name")
	signupemail:=r.Form.Get("signup_email")
	signuppassword:=r.Form.Get("signup_password")

	w.Write([]byte(fmt.Sprintf("login name: %s      password is: %s       signup name is: %s        signup password : %s         signup password : %s",loginname,loginpasswod,signupname,signupemail,signuppassword)))
}

func (m* Repository)Contact(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"contact.page.tmpl",&Models.TemplateData{
	})
}
func (m* Repository)Request(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"request.page.tmpl",&Models.TemplateData{
	})
}