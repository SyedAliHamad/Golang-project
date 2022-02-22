package dbrepo

import (
	"context"
	"time"

	"github.com/SyedAliHamad/internproject/pkg/Models"
)

func(m *postgresDBrepo)Allusers()bool{
	return true
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