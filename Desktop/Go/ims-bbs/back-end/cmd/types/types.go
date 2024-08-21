package types

import "time"


//product details

type ProductStore interface{
	CreateProduct(product Product) error
	GetProducts() ([]Product, error)
	GetProductByUpcBool(Upc string)bool
}

type Product struct{
	ID int `json:"id"`
	Name   string `json:"name"`
	Upc    string `json:"upc"`
	Category string `json:"category"`
	Image string `json:"image"`
	Quantity     int `json:"quantity"`
	Location string `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductPayload struct{
	Name   string `json:"name" validate:"required"`
	Upc    string `json:"upc" validate:"required"`
	Category string `json:"category" validate:"required"`
	Image string `json:"image" validate:"required"`
	Quantity     int `json:"quantity" validate:"required"`
	Location string `json:"location" validate:"required"`
}

// type UPCRequest struct {
// 	Upc string `json:"upc"`
// }

//merchant details

type MerchantStore interface{
	CreateMerchant(merchant Merchant) error
	GetMerchants() ([]Merchant, error)
}

type MerchantPayload struct{
	Merchant_Name string `json:"merchant_name" validate:"required"`
	Merchant_Address string `json:"merchant_address" validate:"required"`
	In_Contact string `json:"in_contact" validate:"required"`
	Contact_Info string `json:"contact_info" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type Merchant struct{
	ID int `json:"int"`
	Merchant_Name string `json:"merchant_name"`
	Merchant_Address string `json:"merchant_address"`
	In_Contact string `json:"in_contact"`
	Contact_Info string `json:"contact_info"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}

//worker details
type Staff struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Contact     string `json:"contact"`
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

//grn types

//1. grn details

type GRN struct{
	ID int `json:"id"`
	InvoiveId string `json:"invoiceId"`
	Merchant_id int `json:"merchant_id"`
	Total_skus int `json:"total_skus"`
	Total_quantity int `json:"total_quantity"`
	Transporter_name string `json:"transporter_name"`
	Driver_name string `json:"driver_name"`
	Driver_contact int `json:"driver_contact"`
	CreatedAt time.Time `json:"createdAt"`
}

type GRNPayload struct{
	InvoiveId string `json:"invoiceId" validate:"required"`
	Merchant_id int `json:"merchant_id" validate:"required"`
	Total_skus int `json:"total_skus" validate:"required"`
	Total_quantity int `json:"total_quantity" validate:"required"`
	Transporter_name string `json:"transporter_name" validate:"required"`
	Driver_name string `json:"driver_name" validate:"required"`
	Driver_contact int `json:"driver_contact" validate:"required"`
}

//2. grn item details

type GRNItem struct{
	ID        int       `json:"id"`
	GRNId   int64       `json:"grn_id"`
	ProductId int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}
