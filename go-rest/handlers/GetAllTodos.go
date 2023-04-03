package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"myexample.com/todo/models"
)

func (h handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	if result := h.DB.Find(&todos); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
