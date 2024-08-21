package grn

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mukul-pixel/ims-bbs/cmd/types"

)

type Handler struct {
	productStore types.ProductStore
	merchantStore types.MerchantStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/createGRN", h.CreateGRN).Methods("POST")
	router.HandleFunc("/showPendingGRNs", h.getPendingGRNs).Methods("GET")
}

func (h *Handler) CreateGRN(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getPendingGRNs(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) AddProductAtMerchant(w http.ResponseWriter, r *http.Request) {

}
