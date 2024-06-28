package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectService struct {
	store Store
}

func NewProjectsService(store Store) *ProjectService {
	return &ProjectService{store: store}
}

func (projectService *ProjectService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects", projectService.handleCreateProject).Methods("POST")
}


func (projectService *ProjectService) handleCreateProject(res http.ResponseWriter, req *http.Request){

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	defer req.Body.Close()

	var project *Project

	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(res, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := validateProjectPayload(project); err != nil {
		WriteJSON(res, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return 
	}

	p, err := projectService.store.CreateProject(project)

	if err != nil {
		WriteJSON(res, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	WriteJSON(res, http.StatusCreated, p)
}

func validateProjectPayload(project *Project) error {
	if project.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
