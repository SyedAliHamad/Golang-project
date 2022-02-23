package Models

import (
	"time"
)

//Loginform holds form data
type Loginform struct{
	LoginEmail string
	LoginPassword string
}

//users is the user table from the database
type Student_info struct{
	Username string
	Email string
	University string
	Password string
	Created time.Time
	Status bool
	Hash string
	Dropuni []string
}

//Data is the data table from the database
type Data struct{
	Exam_id int
	Course_name string
	Department string
	Campus string
	Year time.Time
	Semester string
	Path_pdf string
}


//contact is the contact table from the database
type Contact struct{
	Message_id int
	Username string
	Email string
	University string
	Message string
}

//admin is the admin table from the database
type Admin struct{
	Username string
	Password string
	Email string
	Excesslevel int
}

//req_course is the req_course table from the database
type Req_course struct{
	Request_id int
	Email string
	University_name string
	Course string
}
