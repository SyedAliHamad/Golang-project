package dbrepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SyedAliHamad/internproject/pkg/Models"
)

func(m *postgresDBrepo)Allusers()bool{
	return true
}


func (m*postgresDBrepo) Getuniversities() ([]string,error){

	rows,err:=m.DB.Query("select university_name from university")
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	count:=0
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
		count++
	}
	if err=rows.Err(); err !=nil{
		log.Fatal("error scanning rows",err)
	}
	fmt.Println("------------------------------")
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