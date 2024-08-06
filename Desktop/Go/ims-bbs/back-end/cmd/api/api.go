package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

type APIServer struct {
	db      *sql.DB
	Addr string
}

func NewAPIServer(db *sql.DB, addr string) *APIServer {
	return &APIServer{
		db:      db,
		Addr: addr,
	}
}

func (s *APIServer) Run() error {
	//need router to run the server
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Server running on", s.Addr)

	return http.ListenAndServe(s.Addr, subrouter)
}
