package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/mukul-pixel/ims-bbs/cmd/auth"
	"github.com/mukul-pixel/ims-bbs/cmd/config"
	"github.com/mukul-pixel/ims-bbs/cmd/types"
	"github.com/mukul-pixel/ims-bbs/cmd/utils"

)

//------------------ NOTE ---------------------
//1. create front-end page then, login into the dashboard
//2.if everything is okay -then think of grn-which includes products

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
	//creating a session for the admin after it is added to db
	var admin types.LoginPayload

	if err := utils.ParseJSON(r, &admin); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(admin); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//checking if the user already exists
	u, err := h.store.GetUserByEmail(admin.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		fmt.Println("checkpoint1")
		return
	}

	//comparing the password
	if !auth.ComparePassword(u.Password, []byte(admin.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		fmt.Println("checkpoint2", u.Email)
		return
	}

	//token generation - need 2 things -> 1. Expiration time 2. secret-key
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusAccepted, map[string]string{"token": token})
}

// this will create user but first read the data from the file and then create the user by hashing the password and storing in db
func (h *Handler) handleAdmin(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
		return
	}

	//register type payload
	var user []types.AdminPayload
	if user, err = utils.ReadFromFile(file); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not able to read the file"))
		return
	}

	for _, row := range user {

		//validating the fields
		if err := utils.Validate.Struct(row); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload for %s, %v", row.Email, errors))
			return
		}
		// fmt.Println("checkpoint 4 :");

		//after parsing the data i'll check if the user already exists in the db or not
		_, err := h.store.GetUserByEmail(row.Email)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", row.Email))
			return
		}
		// fmt.Println("checkpoint 5 :");

		//if the user not exists i'll convert hash the password
		hashedPassword, err := auth.HashThePassword(row.Password)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		// fmt.Println("checkpoint 6 :");

		//create the user by inserting the values to sql
		err = h.store.CreateUser(types.Admin{
			FirstName:   row.FirstName,
			LastName:    row.LastName,
			Email:       row.Email,
			Password:    hashedPassword,
			Contact:     row.Contact,
			Address:     row.Address,
			Age:         row.Age,
			JoiningDate: row.JoiningDate,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error registering the user: %s", row.Email))
			return
		}
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
