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
//sets the config fot the templates package
func NewRenderer(a *config.AppConfig){
	app=a
}

func AddDefaultData(td *Models.TemplateData,r *http.Request)*Models.TemplateData{

	td.CSRFToken=nosurf.Token(r)

	return td
}

//Render Template renders templates using hml/templates
func Template(w http.ResponseWriter,r *http.Request, tmpl string,td *Models.TemplateData) {
	
	var tc map[string]* template.Template
	
	if app.UseCache{
		//get the template cache from the app config
		tc=app.TemplateCache
	}else{
		//Testing: used to rebuild the cache on every request
		tc,_=CreateTemplateCashe()
	}


	t,ok:=tc[tmpl]
	if!ok{
		log.Fatal("Could not get template from template cache")
	}

	buf:=new(bytes.Buffer)


	td=AddDefaultData(td,r)
	
	_=t.Execute(buf,td)

	_,err:=buf.WriteTo(w)
	
	if err!=nil{
		fmt.Println("Error writing template to browser",err)
	}

}
var functions=template.FuncMap{


}
//creates a teamplates cache as map
func CreateTemplateCashe()(map[string]*template.Template,error){

	myCache :=map[string]*template.Template{}
	
	pages,err:=filepath.Glob("./Templates/*.page.tmpl")
	if err != nil{
		return myCache ,err
	}

	for _,page:=range pages{
		name := filepath.Base(page)

		ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache ,err
		}

		matches,err:= filepath.Glob("./Templates/*.layout.tmpl")
		if err != nil{
			return myCache,err
		}
		if len(matches)>0{
			ts,err=ts.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil{
				return myCache ,err
			}
		}
		myCache[name]=ts
	}
	return myCache ,nil
}