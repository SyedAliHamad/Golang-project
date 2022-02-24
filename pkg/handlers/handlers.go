package handlers

import (
	"log"
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


	render.Template(w,r,"login.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
	})
}



//PostLogin: Handles the postin of the form
func (m* Repository)PostLogin(w http.ResponseWriter, r *http.Request){

	_=m.App.Session.RenewToken(r.Context())

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

	form.Required("login_email","login_password")
	form.Minlength("login_password",8,r)
	form.IsEmail("login_email")

	if !form.Valid(){

		render.Template(w,r,"login.page.tmpl",&Models.TemplateData{
			Form: form,
		})
	
	return
	}

	id,_,err:=m.DB.Authenticate(login.LoginEmail,login.LoginPassword)
	if err!=nil{
		log.Println(err)
		m.App.Session.Put(r.Context(),"error","Invalid login credentials")
		http.Redirect(w,r,"/login",http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(),"user_id",id)
	m.App.Session.Put(r.Context(),"flash","Logged in successfully")
	http.Redirect(w,r,"/view",http.StatusSeeOther)

}
var dropuniversities[]string
var dropcourse[]string
var dropdept[]string

func (m* Repository)filldata(w *http.ResponseWriter){
	dropuni,err :=m.DB.Getuniversities()
	if err !=nil{
		helpers.ServerError(*w,err)
	} 
	
	dropuniversities=dropuni	

	dropc,err:=m.DB.GetCourses()
	if err !=nil{
		helpers.ServerError(*w,err)
	} 
	dropcourse=dropc

	dropd,err:=m.DB.Getdepartment()
	if err !=nil{
		helpers.ServerError(*w,err)
	} 
	dropdept=dropd
}

func (m* Repository)Signup(w http.ResponseWriter, r *http.Request){

	m.filldata(&w)
	emptysignup:= Models.Student_info{
	}

	data:= make(map[string]interface{})
	data["signupform"] = emptysignup
	

	render.Template(w,r,"signup.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
		Data: data,
		Dropuni: dropuniversities,
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
	form.Required("signup_name","signup_email","signup_password")
	form.Minlength("signup_name",3,r)
	form.Minlength("signup_password",8,r)
	form.IsEmail("signup_email")
	form.IsEqual("signup_password","confirm_password",r)
	form.University("signup_university",r)

	if !form.Valid(){
		data:=make(map[string]interface{})
		data["signupform"]=signup
		render.Template(w,r,"signup.page.tmpl",&Models.TemplateData{
			Form: form,
			Data: data,
			Dropuni: dropuniversities,
		})
		
		return
	} 
	
	err =m.DB.InsertStudentinfo(signup)
	if err !=nil{
		helpers.ServerError(w,err)
	} 
}


func (m* Repository)View(w http.ResponseWriter, r *http.Request){

	m.filldata(&w)

	render.Template(w,r,"view.page.tmpl",&Models.TemplateData{
		Dropuni: dropuniversities,
		DropCourse: dropcourse,
		DropDept: dropdept,
	})
}
func (m* Repository)Upload(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"upload.page.tmpl",&Models.TemplateData{
	})
}

func (m* Repository)PostUpload(w http.ResponseWriter, r *http.Request){

	render.Template(w,r,"upload.page.tmpl",&Models.TemplateData{
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

	var emptycontact Models.Contact
	data:=make(map[string]interface{})
	data["contactform"]=emptycontact

	render.Template(w,r,"contact.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository)PostContact(w http.ResponseWriter, r*http.Request){

	err := r.ParseForm()
	if err!=nil{
		helpers.ServerError(w,err)
		return
	}

	contact:=Models.Contact{
		Email: r.Form.Get("Email"),
		Username:r.Form.Get("name"),
		University: r.Form.Get("University"),
		Message:r.Form.Get("Message"),
	}

	log.Println(contact.Email,contact.Username,contact.Message)


	form:=forms.New(r.PostForm)
	form.Required("name","Message","Email","University")
	form.IsEmail("Email")
	form.Minlength("name",3,r)
	form.Minlength("Message",20,r)

	if!form.Valid(){
		data:=make(map[string]interface{})
		data["contactform"]=contact
		render.Template(w,r,"contact.page.tmpl",&Models.TemplateData{
			Form :form,
			Data:data,
		})
		return
	}

	err =m.DB.InsertContact(contact)
	if err !=nil{
		helpers.ServerError(w,err)
	} 

}

func (m* Repository)Request(w http.ResponseWriter,r *http.Request){

	var emptyrequest Models.Req_course
	data:=make(map[string]interface{})
	data["requestform"]=emptyrequest

	render.Template(w,r,"request.page.tmpl",&Models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m* Repository)PostRequest(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err!=nil{
		helpers.ServerError(w,err)
		return
	}

		request:=Models.Req_course{
			University_name: r.Form.Get("University"),
			Course: r.Form.Get("Course"),
			Department: r.Form.Get("Department"),
		}

		form:=forms.New(r.PostForm)
		form.Required("University","Course","Department")

		if!form.Valid(){
			data:=make(map[string]interface{})
			data["requestform"]=request
			render.Template(w,r,"request.page.tmpl",&Models.TemplateData{
				Form :form,
				Data:data,
			})
			return
		}

		err =m.DB.InsertRequest(request)
		if err !=nil{
			helpers.ServerError(w,err)
		} 
}