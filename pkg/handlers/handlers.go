package handlers

import (
	"time"

	"github.com/SyedAliHamad/internproject/helpers"
	"github.com/SyedAliHamad/internproject/internal/driver"
	"github.com/SyedAliHamad/internproject/internal/forms"
	"github.com/SyedAliHamad/internproject/internal/repository"
	"github.com/SyedAliHamad/internproject/internal/repository/dbrepo"
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
	DB repository.DatabaseRepo
}


//NewRepo: creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository{
	return &Repository{
		App:a,
		DB: dbrepo.NewPostgresRepo(db.SQL,a),
	}
}


//NewHandlers: sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo=r
}


//Home: is the home page handler
func (m* Repository)Home(w http.ResponseWriter, r* http.Request){

	render.Template(w,r,"home.page.tmpl",&Models.TemplateData{})
}

//Login: is the About page handler
func (m* Repository)Login(w http.ResponseWriter, r *http.Request){

	var emptylogin Models.Loginform
	data:= make(map[string]interface{})
	data["loginform"] = emptylogin

	render.Template(w,r,"login.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m* Repository)Signup(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"signup.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m* Repository)PostSignup(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err!=nil{
		helpers.ServerError(w,err)
		return
	}

	signup:=Models.Student_info{
		Username :r.Form.Get("signup_name"),
		Email :r.Form.Get("signup_email"),
		University :r.Form.Get("signup_university"),
		Password :r.Form.Get("signup_password"),
		Created : time.Now(),
		Status :false,
		Hash :"not set rn",

	}
	form :=forms.New(r.PostForm)
	form.Required("signup_name","signup_email","signup_university","signup_password")
	form.Minlength("signup_name",3,r)
	form.IsEmail("signup_email")

	if !form.Valid(){
		data:=make(map[string]interface{})
		data["signupform"]=signup
		render.Template(w,r,"signup.page.tmpl",&Models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	} 
	
	err =m.DB.InsertStudentinfo(signup)
	if err !=nil{
		helpers.ServerError(w,err)
	} 
}
//PostLogin: Handles the postin of the form
func (m* Repository)PostLogin(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err!=nil{
		helpers.ServerError(w,err)
		return
	}

	login := Models.Loginform{

		LoginEmail: r.Form.Get("login_email"),
		LoginPassword: r.Form.Get("login_password"),
	}


	form :=forms.New(r.PostForm)
	//form.Has("login_email",r)
	form.Required("login_email","login_password")
	form.Minlength("login_password",8,r)
	form.IsEmail("login_email")

	if !form.Valid(){
		data:=make(map[string]interface{})
		data["loginform"]=login

		render.Template(w,r,"login.page.tmpl",&Models.TemplateData{
			Form: form,
			Data: data,
		})
	
	return
	}
}

func (m* Repository)View(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"view.page.tmpl",&Models.TemplateData{
	})
}


func (m* Repository)PostView(w http.ResponseWriter, r *http.Request){
	/*
	loginname :=r.Form.Get("login_name")
	loginpasswod:=r.Form.Get("login_password")

	signupname:= r.Form.Get("signup_name")
	signupemail:=r.Form.Get("signup_email")
	signuppassword:=r.Form.Get("signup_password")

	w.Write([]byte(fmt.Sprintf("login name: %s      password is: %s       signup name is: %s        signup password : %s         signup password : %s",loginname,loginpasswod,signupname,signupemail,signuppassword)))
*/
}

func (m* Repository)Contact(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"contact.page.tmpl",&Models.TemplateData{
	})
}
func (m* Repository)Request(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"request.page.tmpl",&Models.TemplateData{
	})
}