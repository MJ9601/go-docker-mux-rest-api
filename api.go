package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string

	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s*APIServer) Serve(){
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	TasksService := NewTaskService(s.store);
	TasksService.RegisterRoutes(subRouter)

	log.Println("Starting the API Server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subRouter))

}