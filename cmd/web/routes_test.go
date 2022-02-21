package main

import (
	"fmt"
	"testing"

	"github.com/SyedAliHamad/internproject/pkg/config"
	"github.com/go-chi/chi"
)

func TestRoutes (t *testing.T){
	var app config.AppConfig
	mux:=routes(&app)

	switch v:=mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not chi.mux type is %V",v))
		
	}
}