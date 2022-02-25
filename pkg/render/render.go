package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/SyedAliHamad/internproject/pkg/Models"
	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/justinas/nosurf"
)


var app *config.AppConfig

//sets the config for the templates package
func NewRenderer(a *config.AppConfig){
	app=a
}

func AddDefaultData(td *Models.TemplateData,r *http.Request)*Models.TemplateData{

	td.CSRFToken=nosurf.Token(r)

	if app.Session.Exists(r.Context(),"user_id"){
		td.IsAuthenticated=1
		log.Println("user logged in")
	}

	return td
}

//Render Template renders templates using hml/templates
func Template(w http.ResponseWriter,r *http.Request, tmpl string,td *Models.TemplateData) {
	
	//template cache holds all the templates
	var tc map[string]* template.Template
	
	//if I want my templates from  disk and not to be 
	//taken from cache this is just for developmening stage
	if app.UseCache{
		//get the template cache from the app config
		tc=app.TemplateCache
	}else{
		//Testing: used to rebuild the cache on every request
		tc,_=CreateTemplateCashe()
	} 

	//if template exists t has value else ok is false
	t,ok:=tc[tmpl]
	if!ok{
		log.Fatal("Could not get template from template cache")
	}
  
	//holds bytes
	buf:=new(bytes.Buffer)


	//addings additional data for security
	td=AddDefaultData(td,r)
	
	// takes template t executes it with data td and stores values in buf 
	_=t.Execute(buf,td)

	_,err:=buf.WriteTo(w)
	
	if err!=nil{ 
		fmt.Println("Error writing template to browser",err)
	}

}


//creates a teamplates cache as map
func CreateTemplateCashe()(map[string]*template.Template,error){

	//template cache holds all the templates
	myCache :=map[string]*template.Template{}
	

	//sarting from root folder go to templates and find all the files ending with
	//.page.tmpl
	pages,err:=filepath.Glob("./Templates/*.page.tmpl")
	if err != nil{
		return myCache ,err
	}

	for _,page:=range pages{

		//extracts the name of page without route
		name := filepath.Base(page)


		//template set: 
		//template.New(name) creates a template with name extracted in name
		//parsingfile parses page and saves it in ts
		ts,err := template.New(name).ParseFiles(page)
		if err != nil{
			return myCache ,err
		}

		//checks if template matches layout
		matches,err:= filepath.Glob("./Templates/*.layout.tmpl")
		if err != nil{ 
			return myCache,err
		}

		//If something is found
		if len(matches)>0{  
			ts,err=ts.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil{
				return myCache ,err
			}
		}
		//creating a map which takes string name and gives template
		myCache[name]=ts
	}
	return myCache ,nil
}