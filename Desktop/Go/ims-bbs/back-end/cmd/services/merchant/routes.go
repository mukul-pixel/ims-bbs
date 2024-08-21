package merchant

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"

	"github.com/mukul-pixel/ims-bbs/cmd/types"
	"github.com/mukul-pixel/ims-bbs/cmd/utils"

)

type Handler struct {
	//import an interface for productStore
	merchantStore types.MerchantStore
}

func NewHandler(merchantStore types.MerchantStore) *Handler {
	return &Handler{merchantStore: merchantStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/addMerchant", h.handleAddMerchant).Methods("POST")
	router.HandleFunc("/viewMerchant", h.viewMerchant).Methods("GET")
}

func (h *Handler) handleAddMerchant(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
		return
	}

	var merchant []types.MerchantPayload
	if merchant, err = ReadFromFile(file); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not able to read the file"))
		fmt.Println("checkpoint")
		return
	}

	for _, row := range merchant {
		if err := utils.Validate.Struct(row);err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload for %s, %v", row.Merchant_Name, errors))
		}

		err = h.merchantStore.CreateMerchant(types.Merchant{
			Merchant_Name: row.Merchant_Name,
	        Merchant_Address: row.Merchant_Address,
	        In_Contact:row.In_Contact,
	        Contact_Info:row.Contact_Info,
	        Category:row.Category,
		})
		if err!=nil{
			utils.WriteError(w,http.StatusInternalServerError,fmt.Errorf("error adding merchant: %v",err))
		}
		
	}
	utils.WriteJSON(w,http.StatusCreated,nil)
}

func (h *Handler) viewMerchant(w http.ResponseWriter, r *http.Request) {
	merchants,err:= h.merchantStore.GetMerchants()
	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
	}

	utils.WriteJSON(w,http.StatusOK,merchants)
}

func ReadFromFile(file multipart.File) ([]types.MerchantPayload, error) {
    if file == nil {
        return nil, fmt.Errorf("file is nil")
    }

    f, err := excelize.OpenReader(file)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }

    rows, err := f.GetRows(f.GetSheetName(0))
    if err != nil{
        return nil, fmt.Errorf("failed to get rows or no rows found: %w", err)
    }

    var merchants []types.MerchantPayload
    for i, row := range rows {
        if i == 0 {
            continue // skipping the header row
        }
        if len(row) < 5 {
            continue // skipping rows with insufficient columns
        }

        merchant := types.MerchantPayload{
            Merchant_Name:    row[0],
            Merchant_Address: row[1],
            In_Contact:       row[2],
            Contact_Info:     row[3],
            Category:         row[4],
        }

        merchants = append(merchants, merchant)
    }

    if len(merchants) == 0 {
        return nil, fmt.Errorf("no valid merchants found in the file")
    }

    return merchants, nil
}

