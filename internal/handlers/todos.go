package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/data"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	if userID := r.URL.Query().Get("userId"); userID != "" {
		filters["userId"] = userID
	}

	todos := data.FilterTodos(filters)
	todos = applyPagination(todos, r)

	respondJSON(w, http.StatusOK, todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	todo := data.GetTodoByID(id)
	if todo == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	todo.ID = data.GetNextID("todos")
	respondJSON(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetTodoByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	todo.ID = id
	respondJSON(w, http.StatusOK, todo)
}

func PatchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing := data.GetTodoByID(id)
	if existing == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	// Merge patch into existing
	result := *existing
	if title, ok := patch["title"].(string); ok {
		result.Title = title
	}
	if completed, ok := patch["completed"].(bool); ok {
		result.Completed = completed
	}
	if userID, ok := patch["userId"].(float64); ok {
		result.UserID = int(userID)
	}

	respondJSON(w, http.StatusOK, result)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetTodoByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{})
}
