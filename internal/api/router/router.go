package router

import (
	"conn/internal/connections"
	"fmt"
	"log"
	"net/http"
)

func NewRouter() {
	r := http.NewServeMux()

	handler := connections.NewHandler()

	r.HandleFunc("POST /tasks", handler.Create)
	r.HandleFunc("GET /tasks/{id}", handler.Get)
	r.HandleFunc("GET /tasks", handler.Getall)
	r.HandleFunc("PUT /tasks/{id}", handler.Update)
	r.HandleFunc("DELETE /tasks/{id}", handler.Delete)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
