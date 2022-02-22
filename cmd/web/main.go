package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SyedAliHamad/internproject/helpers"
	"github.com/SyedAliHamad/internproject/internal/driver"
	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/SyedAliHamad/internproject/pkg/handlers"
	"github.com/SyedAliHamad/internproject/pkg/render"

	"github.com/alexedwards/scs/v2"
)
const portNumber=": 8080"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main(){

	db,err:=run()
	if err!=nil{
		log.Fatal(err)
	}
	defer db.SQL.Close()     //connection won't be closed until main is stopped

	fmt.Println((fmt.Printf("Starting application on port %s ", portNumber)))


	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err=srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB,error){
	
	//change this to true when in production=> state for cookie
	app.InProduction=false

	infoLog =log.New(os.Stdout,"INFO \t",log.Ldate|log.Ltime)
	app.InfoLog=infoLog


	errorLog=log.New(os.Stdout,"ERROR \t",log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog=errorLog



	session= scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session=session

	//connnect to database

		log.Println("connecting to database")
		db,err:=driver.ConnectSQL("host=localhost port=5432 dbname=exams user=ali password=ali")
		if err!=nil{
			log.Fatal("cannot connect to database! dying ...")
		}
		log.Println("Connected to database")

	tc,err :=render.CreateTemplateCashe()
	if err!=nil{
		log.Fatal("cannot create template cache")
		return nil,err
	}

	app.TemplateCache=tc
	app.UseCache=false

	repo:=handlers.NewRepo(&app,db) 
	//we have application config avaliable to handlers along 
	//with db which is a pointer to a driver which can only 
	//handle postgres but can be changed to any driver with a new function

	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db,nil
}