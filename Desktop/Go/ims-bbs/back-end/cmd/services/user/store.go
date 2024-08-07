package user

import (
	"database/sql"

	"github.com/mukul-pixel/ims-bbs/cmd/types"

)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.Admin,error) {
	rows,err := s.db.Query("SELECT * FROM admins WHERE email=?",email)
	if err != nil {
		return nil,err
	}

	u:= new(types.Admin)
	for rows.Next(){
		u,err=scanRowsIntoUsers(rows)
		if err != nil {
			return nil,err
		}
	}

	if u.ID == 0{
		return nil,err
	}

	return u,nil
}

func(s *Store) CreateUser(user types.Admin)error{
	_,err := s.db.Exec("INSERT INTO admins (firstName,lastName,email,password,contact,address,age,joiningDate) VALUES (?,?,?,?,?,?,?,?)",user.FirstName,user.LastName,user.Email,user.Password,user.Contact,user.Address,user.Age,user.JoiningDate)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUsers(row *sql.Rows) (*types.Admin,error) {
	user := new(types.Admin)

	err:= row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Contact,
		&user.Address,
		&user.Age,
		&user.JoiningDate,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user,nil
}