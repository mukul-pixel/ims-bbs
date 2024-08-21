package product

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"

	"github.com/mukul-pixel/ims-bbs/cmd/types"
	"github.com/mukul-pixel/ims-bbs/cmd/utils"
)

type Handler struct {
	productStore types.ProductStore
}

func NewHandler(productStore types.ProductStore) *Handler {
	return &Handler{productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/addProduct", h.handleAddProduct).Methods("POST")
	router.HandleFunc("/getProducts", h.getProducts).Methods("GET")
	router.HandleFunc("/findProductByUpc", h.findProductByUpc).Methods("GET")
}

func (h *Handler) findProductByUpc(w http.ResponseWriter, r *http.Request) {

	upc := r.URL.Query().Get("upc")

	// Validate that the UPC is present
	if upc == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("upc is missing"))
		return
	}

	// Check if the product exists using the GetProductByUpcBool method
	if !h.productStore.GetProductByUpcBool(upc) {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}

func (h *Handler) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
		return
	}

	//register type payload
	var product []types.ProductPayload
	if product, err = ReadFromFile(file); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not able to read the file"))
		return
	}

	for _, row := range product {
		//validating the fields
		if err := utils.Validate.Struct(row); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload for %s, %v", row.Name, errors))
			return
		}

		err = h.productStore.CreateProduct(types.Product{
			Name:     row.Name,
			Upc:      row.Upc,
			Category: row.Category,
			Image:    row.Image,
			Quantity: row.Quantity,
			Location: row.Location,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding product: %v", err))
		}
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productStore.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func ReadFromFile(file multipart.File) ([]types.ProductPayload, error) {
	if file == nil {
		return nil, fmt.Errorf("file is nil")
	}

	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("failed to get rows or no rows found: %w", err)
	}

	var products []types.ProductPayload

	for i, row := range rows {
		if i == 0 {
			continue //skipping heading row
		}
		if len(row) < 6 {
			continue
		}

		quantity, err := strconv.Atoi(row[4])
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value in row %d: %w", i+1, err)
		}

		product := types.ProductPayload{
			Name:     row[0],
			Upc:      row[1],
			Category: row[2],
			Image:    row[3],
			Quantity: quantity,
			Location: row[5],
		}

		products = append(products, product)
	}
	return products, nil
}
