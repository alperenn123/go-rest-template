package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"myexample.com/todo/models"
)

func (h handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedTodo models.Todo
	json.Unmarshal(body, &updatedTodo)

	var todo models.Todo

	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	todo.Username = updatedTodo.Username
	todo.Desc = updatedTodo.Desc

	h.DB.Save(&todo)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
