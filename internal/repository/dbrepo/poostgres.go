package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SyedAliHamad/internproject/pkg/Models"
	"golang.org/x/crypto/bcrypt"
)

func(m *postgresDBrepo)Allusers()bool{
	return true
}
func(m*postgresDBrepo) Getdepartment() ([]string,error) {

	rows,err:=m.DB.Query("select dep_name from department;")
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	defer rows.Close()
	var name string
	var dropdept []string
	for rows.Next(){
		err:=rows.Scan(&name)
		if err!=nil{
			log.Println(err)
			return nil,err
		}
		dropdept=append(dropdept, name)
	
	}
	if err=rows.Err(); err !=nil{
		log.Fatal("error scanning rows",err)
	}

	return dropdept,err

}

func (m*postgresDBrepo) GetCourses()([]string,error){
	rows,err:=m.DB.Query("select coursename from course;")
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	defer rows.Close()
	var name string
	var dropcourse []string
	for rows.Next(){
		err:=rows.Scan(&name)
		if err!=nil{
			log.Println(err)
			return nil,err
		}
		dropcourse=append(dropcourse, name)
	
	}
	if err=rows.Err(); err !=nil{
		log.Fatal("error scanning rows",err)
	}
	return dropcourse,err
}

func (m*postgresDBrepo) Getuniversities() ([]string,error){

	rows,err:=m.DB.Query("select university_name from university;")
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	defer rows.Close()
	var name string
	var dropuni []string
	for rows.Next(){

		err:=rows.Scan(&name)
		if err!=nil{
			log.Println(err)
			return nil,err
		}
		//fmt.Println("Record is",name)
		dropuni=append(dropuni, name)
	}
	if err=rows.Err(); err !=nil{
		log.Fatal("error scanning rows",err)
	}

	return dropuni,err

}

//inserts signup student's data into the database
func (m *postgresDBrepo) InsertStudentinfo(reg Models.Student_info) error{
	
	ctx,cancel:=context.WithTimeout(context.Background(),50*time.Second)
	defer cancel()

	stmt := `insert into student_info
	(username,email,university,password,created,status,hash) 
	values
	($1,$2,$3,$4,$5,$6,$7);`

	_,err:= m.DB.ExecContext(ctx,stmt,
		reg.Username,
		reg.Email,
		reg.University,
		reg.Password,
		time.Now(),
		reg.Status,
		reg.Hash,
	)

	if err!=nil{
		return err
	}
	return nil

}

//returns student by email
func(m* postgresDBrepo) GetStudent(email string)(Models.Student_info,error){

	ctx,cancel:=context.WithTimeout(context.Background(),50*time.Second)
	defer cancel()

	stmt:= `select username,email,university,password,created,status,hash
	from student_info where email= $1;`

	row:=m.DB.QueryRowContext(ctx,stmt,email)

	var u Models.Student_info

	err:=row.Scan(
		&u.Username,
		&u.Email,
		&u.University,
		&u.Password,
		&u.Created,
		&u.Status,
		&u.Hash,
	)
	if err!=nil{
		return u,err
	}
	return u,nil


}

func (m *postgresDBrepo) InsertContact(reg Models.Contact)error{
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Minute)
	defer cancel()

	stmt := `insert into contact
	(username,email,university,message) 
	values
	($1,$2,$3,$4);`

	_,err:= m.DB.ExecContext(ctx,stmt,
		reg.Username,
		reg.Email,
		reg.University,
		reg.Message,
	)


	if err!=nil{
		return err
	}
	return nil


}
func (m *postgresDBrepo) InsertRequest(req Models.Req_course)error{
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Minute)
	defer cancel()

	stmt := `insert into 
	Req_course(University_name,course,department) 
	values($1,$2,$3);
	`

	_,err:= m.DB.ExecContext(ctx,stmt,
		req.University_name,
		req.Course,
		req.Department,
	)

	if err!=nil{
		return err
	}
	return nil
}

//Authenticate: Authenticates a user
func (m*postgresDBrepo) Authenticate(email,testPassword string)(int,string,error){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Minute)
	defer cancel()


	var id int
	var hashedPassword string

	row:=m.DB.QueryRowContext(ctx,"select student_id ,password from student_info where email=$1;",email)
	err:=row.Scan(&id,&hashedPassword)
	if err!=nil{
		return id,"",err
	}

	err=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(testPassword))
	
	if err==bcrypt.ErrMismatchedHashAndPassword{
		return 0,"",errors.New("Incorrect password")
	}else if err!=nil{
		return 0,"",err
	}

	return id,hashedPassword,nil

}