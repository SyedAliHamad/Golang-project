package handlers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/SyedAliHamad/internproject/internal/driver"
	"github.com/SyedAliHamad/internproject/internal/forms"
	"github.com/SyedAliHamad/internproject/internal/repository"
	"github.com/SyedAliHamad/internproject/internal/repository/dbrepo"
	"github.com/SyedAliHamad/internproject/pkg/Models"
	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/SyedAliHamad/internproject/pkg/render"
	"tawesoft.co.uk/go/dialog"

	"net/http"
)

//Repository: creates a new repository
type Repository struct{
	App *config.AppConfig
	DB repository.DatabaseRepo
}



//Repo: The repository used by the handlers
var Repo *Repository


//NewRepo: creates a new repository and assigns values to
//repository structure
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


//logs our user out
func (m*Repository)Logout(w http.ResponseWriter,r *http.Request){
	_ =m.App.Session.Destroy(r.Context())
	_= m.App.Session.RenewToken(r.Context())

	http.Redirect(w,r,"/login",http.StatusSeeOther)
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
		log.Println("error pasing login form")
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
		http.Redirect(w,r,"/home",http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(),"user_id",id)
	http.Redirect(w,r,"/",http.StatusSeeOther)
	dialog.Alert("Successfully Logged in")

}
var dropuniversities[]string
var dropcourse[]string
var dropdept[]string

func (m* Repository)filldata(w *http.ResponseWriter){
	dropuni,err :=m.DB.Getuniversities()
	if err !=nil{
		log.Println("error getting university data")
	} 
	
	dropuniversities=dropuni	

	dropc,err:=m.DB.GetCourses()
	if err !=nil{
		log.Println("error getting course data")
	} 
	dropcourse=dropc

	dropd,err:=m.DB.Getdepartment()
	if err !=nil{
		log.Println("error getting department data")
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

		log.Println("error parsing signup form")
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
		log.Println("Error inserting students info")
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

	m.filldata(&w)
	render.Template(w,r,"upload.page.tmpl",&Models.TemplateData{
		Dropuni: dropuniversities,
		DropCourse: dropcourse,
		DropDept: dropdept,
	})
}

func (m* Repository)PostUpload(w http.ResponseWriter, r *http.Request){

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dialog.Alert("File submitted")
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
		log.Println("error pasing the contact form")
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
		log.Println("error inserting contact information")
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
		log.Println("error parsing request form")
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
			log.Println("error inserting request information")
		} 
}

//we can send data from handers to
//goline templates