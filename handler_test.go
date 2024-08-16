package handler

import (
	"bytes"
	"conn/internal/connections"
	"conn/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var id int
var timest string
var have models.Task

func TestCreate(t *testing.T) {
	a := connections.NewHandler()
	reqBody := `{"title":"Test Task","description":"Test Description"}`
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var have1 models.Task
	if err := json.NewDecoder(rr.Body).Decode(&have1); err != nil {
		t.Error(err)
	}
	have = have1
	id = have.ID
	timest = have.Created_at
	want := models.Task{
		ID:          have.ID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "created",
		Created_at:  timest,
	}
	log.Println(have)
	if have != want {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}

func TestGet(t *testing.T) {
	a := connections.NewHandler()
	want := models.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "created",
		Created_at:  timest,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("/tasks/%v", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Get)
	handler.ServeHTTP(rr, req)

	if have != want {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}

func TestGetAll(t *testing.T) {
	a := connections.NewHandler()
	want := []models.Task{
		{
			ID:          id,
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "created",
			Created_at:  timest,
		},
	}

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Getall)
	handler.ServeHTTP(rr, req)

	var have []models.Task
	if err := json.NewDecoder(rr.Body).Decode(&have); err != nil {
		t.Error(err)
	}

	if len(have) != len(want) || have[0] != want[0] {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}

func TestUpdate(t *testing.T) {
	a := connections.NewHandler()
	want := models.Task{
		ID:          id,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "updated",
		Created_at:  timest,
	}

	taskUpdate := models.TaskUpdate{
		ID:          id,
		Title:       "Updated Task",
		Description: "Updated Description",
	}
	have.Title = taskUpdate.Title
	have.Description = taskUpdate.Description
	body, err := json.Marshal(taskUpdate)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("/tasks/%v", id), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Update)
	handler.ServeHTTP(rr, req)

	have.Status = "updated"
	if have != want {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}

func TestDelete(t *testing.T) {
	a := connections.NewHandler()
	want := ""

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/tasks/%v", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Delete)
	handler.ServeHTTP(rr, req)

	have1 := ""
	if have1 != want {
		t.Errorf("handler returned unexpected body: got %v want %v", have, want)
	}
}
