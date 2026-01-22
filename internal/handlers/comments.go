package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/data"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListComments(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	if postID := r.URL.Query().Get("postId"); postID != "" {
		filters["postId"] = postID
	}

	comments := data.FilterComments(filters)
	comments = applyPagination(comments, r)

	respondJSON(w, http.StatusOK, comments)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	comment := data.GetCommentByID(id)
	if comment == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, comment)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	comment.ID = data.GetNextID("comments")
	respondJSON(w, http.StatusCreated, comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetCommentByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	comment.ID = id
	respondJSON(w, http.StatusOK, comment)
}

func PatchComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing := data.GetCommentByID(id)
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
	if name, ok := patch["name"].(string); ok {
		result.Name = name
	}
	if email, ok := patch["email"].(string); ok {
		result.Email = email
	}
	if body, ok := patch["body"].(string); ok {
		result.Body = body
	}
	if postID, ok := patch["postId"].(float64); ok {
		result.PostID = int(postID)
	}

	respondJSON(w, http.StatusOK, result)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetCommentByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{})
}
