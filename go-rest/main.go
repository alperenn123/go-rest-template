package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"myexample.com/todo/db"
	"myexample.com/todo/handlers"
)

func main() {
	router := mux.NewRouter()
	DB := db.Init()
	h := handlers.New(DB)

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}).Methods("GET")

	router.HandleFunc("/todos", h.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", h.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todos", h.PostTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", h.UpdateTodo).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", h.DeleteTodo).Methods(http.MethodDelete)

	e := http.ListenAndServe(":8080", router)
	if e != nil {
		log.Fatalln("There's an error with the server,", e)
	}

	fmt.Println("Server is listening the port 8080")
}
