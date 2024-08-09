package types

import "time"

type Staff struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Contact     int `json:"contact"`
	Address     string `json:"address"`
	Age         int `json:"age"`
	JoiningDate string `json:"joiningDate"`
	CreatedAt time.Time `json:"createdAt"`
}

//Admin Service
type AdminStore interface{
	GetUserByEmail(email string) (*Admin,error) 
	CreateUser(user Admin)error
}

type Admin struct {
	ID int `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	Contact     string `json:"contact"`
	Address     string `json:"address"`
	Age         int `json:"age"`
	JoiningDate string `json:"joiningDate"`
	CreatedAt time.Time `json:"createdAt"`
}

type AdminPayload struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=3,max=130"`
	Contact     string `json:"contact"  validate:"required,min=10,max=10"`
	Address     string `json:"address" validate:"required"`
	Age         int `json:"age" validate:"required"`
	JoiningDate string `json:"joiningDate" validate:"required"`
}

type LoginPayload struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

