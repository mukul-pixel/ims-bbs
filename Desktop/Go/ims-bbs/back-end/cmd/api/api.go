package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mukul-pixel/ims-bbs/cmd/services/merchant"
	"github.com/mukul-pixel/ims-bbs/cmd/services/product"
	"github.com/mukul-pixel/ims-bbs/cmd/services/user"

)

type APIServer struct {
	db   *sql.DB
	Addr string
}

func NewAPIServer(db *sql.DB, addr string) *APIServer {
	return &APIServer{
		db:   db,
		Addr: addr,
	}
}

func (s *APIServer) Run() error {
	//need router to run the server
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	merchantStore := merchant.NewStore(s.db)
	merchantHandler := merchant.NewHandler(merchantStore)
	merchantHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Server running on", s.Addr)

	return http.ListenAndServe(s.Addr, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Update with your frontend's URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(router))
}
