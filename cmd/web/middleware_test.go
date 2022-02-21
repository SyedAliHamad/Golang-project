package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoServe(t *testing.T){
	var myh myHandler

	h:=Nosurf(&myh)

	switch v:=h.(type){
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handdler but is %T",v))
	}
}

func TestSessionLoad(t *testing.T){
	var myh myHandler

	h:=SessionLoad(&myh)

	switch v:=h.(type){
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handdler but is %T",v))
	}
}