package repository

import "github.com/SyedAliHamad/internproject/pkg/Models"

type DatabaseRepo interface{
	Allusers()bool
	InsertStudentinfo(log Models.Student_info)error
	InsertContact(log Models.Contact)error
	InsertRequest(log Models.Req_course)error
	Getuniversities()([]string,error)
	GetCourses()([]string,error)
	Getdepartment()([]string,error)
}