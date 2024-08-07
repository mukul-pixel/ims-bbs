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
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
		return
	}
	fmt.Println("checkpoint-1")

	//register type payload
	var user []types.AdminPayload
	if user,err = utils.ReadFromFile(file); err != nil {
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("not able to read the file"))
		return
	}

	fmt.Println("checkpoint 2 :",user);

	
	for _,row := range user{
		// //first i will get the data from the req and parse it to the json format
		// if err := utils.ParseJSON(r, &row); err != nil {
		// 	utils.WriteError(w, http.StatusBadRequest, err)
		// 	return
	    // }
		// fmt.Println("checkpoint 3 :");

		//validating the fields
	    if err := utils.Validate.Struct(row); err != nil {
		    errors := err.(validator.ValidationErrors)
		    utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
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
	    if err !=nil {
		    utils.WriteError(w,http.StatusInternalServerError,err)
		    return
	    }
		// fmt.Println("checkpoint 6 :");

	    //create the user by inserting the values to sql
	    err = h.store.CreateUser(types.Admin{
		    FirstName: row.FirstName,
		    LastName: row.LastName,
		    Email: row.Email,
		    Password: hashedPassword,
		    Contact: row.Contact,
		    Address: row.Address,
		    Age: row.Age,
		    JoiningDate: row.JoiningDate,
	    })
	    if err != nil {
		    utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error registering the user: %s",row.Email))
		    return
	    }
    }

	utils.WriteJSON(w,http.StatusCreated,nil)
}
