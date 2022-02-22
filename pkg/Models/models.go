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
}

//Data is the data table from the database
type Data struct{
	exam_id int
	course_name string
	department string
	campus string
	year time.Time
	semester string
	path_pdf string
}


//contact is the contact table from the database
type contact struct{
	message_id int
	email string
	username string
	message string
}

//admin is the admin table from the database
type admin struct{
	username string
	password string
	email string
	excesslevel int
}

//req_course is the req_course table from the database
type req_course struct{
	request_id int
	email string
	university_name string
	course string
}
