package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//holds the application config which can be accessed by any part of the
//application
type AppConfig struct{

	UseCache bool //
	TemplateCache map[string]*template.Template
	InProduction bool
	Session *scs.SessionManager
}