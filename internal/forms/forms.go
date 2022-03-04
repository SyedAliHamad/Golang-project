package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Form createss a custom form struct, embeds a url.Values object
type Form struct{
	url.Values
	Errors errors
}


//returns true if no errors else false
func (f *Form) Valid() bool{
	return len(f.Errors)==0
}

func New(data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//checks for required fields
func (f *Form)Required(fields ...string){
	for _,field :=range fields{
		value :=f.Get(field)
		if strings.TrimSpace(value)==""{
			f.Errors.Add(field,"This field can not be blank")
		}else if strings.TrimSpace(value)=="select"{
			f.Errors.Add(field,"This field can not be blank")
		}
	}
}



func (f *Form) University(field string,r *http.Request) bool{
	x:=r.Form.Get(field)
	if x=="Select"{
		f.Errors.Add(field,"Please Select a University")
		return false
	}
	return true
}


//Has checksif form field is in p ost and not empty
func (f *Form) Has(field string, r *http.Request) bool{
	x:=r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field,"This field can't be blank")
		return false
	}
	return true
}
func (f *Form) IsEqual(field1 string,field2 string, r *http.Request)bool{
	x:=r.Form.Get(field1)
	y:=r.Form.Get(field2)

	if x!=y{
		f.Errors.Add(field2,"Passwords Don't match")
		return false
	}
	return true

}

//Min length: checks for string min length
func (f *Form) Minlength(field string,length int,r *http.Request)bool{
	x:=r.Form.Get(field)
	if len(x)<length{
		f.Errors.Add(field,fmt.Sprintf("This field must be atleast %d charactes long",length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string){
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field,"Invalid Email address")
	}
}