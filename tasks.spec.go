package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {

	ms := &MockStore{}

	service := NewTaskService(ms)

	t.Run("Should return an error if name is not set", func(t *testing.T){

		payload := &Task{
			Name: "",
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))

		if err != nil{
			t.Fatal(err)
		}

		testingRecorder := httptest.NewRecorder()

		router := mux.NewRouter()

		router.HandleFunc("/tasks",service.handleCreateTasks )


		router.ServeHTTP(testingRecorder, req)

		if testingRecorder.Code != http.StatusBadRequest{
			t.Error("Invalid status code, it should fail")
		}
	})
}


func TestGetTaskById(t *testing.T) {

	ms := &MockStore{}

	service := NewTaskService(ms)

	t.Run("Should return an error if name is not set", func(t *testing.T){

		req, err := http.NewRequest(http.MethodPost, "/tasks/42", nil)
		
		if err != nil{
			t.Fatal(err)
		}
	
		testingRecorder := httptest.NewRecorder()
	
		router := mux.NewRouter()
	
		router.HandleFunc("/tasks/{id}",service.handleCreateTasks)
	
	
		router.ServeHTTP(testingRecorder, req)
	
		if testingRecorder.Code != http.StatusOK{
			t.Error("Invalid status code, it should fail")
		}
	})
}
