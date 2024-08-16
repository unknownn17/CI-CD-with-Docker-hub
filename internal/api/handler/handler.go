package handler

import (
	"conn/internal/database/service"
	"conn/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	S *service.Service
}

func (u *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.TaskCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %s", err), http.StatusNoContent)
		return
	}
	res, err := u.S.CreateTask(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %s", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNoContent)
		return
	}
	res, err := u.S.GetTask(&models.Task_Get_Delete{ID: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %s", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Getall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := u.S.GetTasks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %s", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNoContent)
		return
	}

	var req models.TaskUpdate

	req.ID = id
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %s", err), http.StatusNoContent)
		return
	}

	res, err := u.S.UpdateTask(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %s", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNoContent)
		return
	}
	if err := u.S.DeleteTask(&models.Task_Get_Delete{ID: id}); err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong %s", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Delted Succesfully")
}
