package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name required")
var errProjectIDRequired = errors.New("project ID required")
var errUserIDRequired = errors.New("user ID required")

type TasksService struct {
	store Store
}

func NewTaskService(store Store) *TasksService{
	return &TasksService{store: store}
}

func (tasksService *TasksService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", WithJWTAuth(tasksService.handleCreateTasks, tasksService.store)).Methods("POST")
	router.HandleFunc("/tasks/{id}", WithJWTAuth(tasksService.handleGetTasks, tasksService.store)).Methods("GET")
}


func (tasksService *TasksService) handleCreateTasks(res http.ResponseWriter, req *http.Request){

		body, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}

		defer req.Body.Close()

		var task *Task

		err = json.Unmarshal(body, &task)
		if err != nil {
			WriteJSON(res, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		if err := validateTaskPayload(task); err != nil {
			WriteJSON(res, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return 
		}

		t, err := tasksService.store.CreateTask(task)
		if err != nil {
			WriteJSON(res, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return 
		}

		WriteJSON(res, http.StatusCreated, t)
}

func (taskService *TasksService) handleGetTasks(res http.ResponseWriter, req *http.Request){

	vars := mux.Vars(req)
	id := vars["id"]

	if id == "" {
		WriteJSON(res, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return 
	}

	task, err := taskService.store.GetTaskById(id)

	if err != nil {
		WriteJSON(res, http.StatusInternalServerError, ErrorResponse{Error: "Something went wrong"})
		return
	}

	WriteJSON(res, http.StatusOK, task)
}


func validateTaskPayload(task *Task) error{
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}