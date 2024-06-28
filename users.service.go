package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UserService struct {
	store Store
}

func NewUsersService(store Store) *UserService {
	return &UserService{store: store}
}

func (usersService *UserService) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/users", usersService.handleCreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", usersService.handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/users", usersService.handleGetUser).Methods("GET")
}


func (usersService *UserService) handleCreateUser(w http.ResponseWriter, r *http.Request){}

func (userService *UserService) handleDeleteUser(w http.ResponseWriter, r *http.Request){}



func (userService *UserService) handleGetUser(w http.ResponseWriter, r *http.Request){}