package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/mukul-pixel/ims-bbs/cmd/auth"
	"github.com/mukul-pixel/ims-bbs/cmd/types"
	"github.com/mukul-pixel/ims-bbs/cmd/utils"

)

type Handler struct {
	store types.AdminStore
}

func NewHandler(store types.AdminStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/addAdmin", h.handleAdmin).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//need to test the handleAdmin and then read the data from the excel/file and create a user in the db
	//create a migrate to add the table in db
}

// this will create user but first read the data from the file and then create the user by hashing the password and storing in db
func (h *Handler) handleAdmin(w http.ResponseWriter, r *http.Request) {

	//register type payload
	var user types.AdminPayload

	//first i will get the data from the req and parse it to the json format
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validating the fields
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//after parsing the data i'll check if the user already exists in the db or not
	_, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	//if the user not exists i'll convert hash the password
	hashedPassword, err := auth.HashThePassword(user.Password)
	if err !=nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	//create the user by inserting the values to sql
	err = h.store.CreateUser(types.Admin{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: hashedPassword,
		Contact: user.Contact,
		Address: user.Address,
		Age: user.Age,
		JoiningDate: user.JoiningDate,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error registering the user"))
		return
	}

	utils.WriteJSON(w,http.StatusCreated,nil)

}
