package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/data"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListPosts(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	if userID := r.URL.Query().Get("userId"); userID != "" {
		filters["userId"] = userID
	}

	posts := data.FilterPosts(filters)
	posts = applyPagination(posts, r)

	respondJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	post := data.GetPostByID(id)
	if post == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	post.ID = data.GetNextID("posts")
	respondJSON(w, http.StatusCreated, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetPostByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	post.ID = id
	respondJSON(w, http.StatusOK, post)
}

func PatchPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing := data.GetPostByID(id)
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
	if body, ok := patch["body"].(string); ok {
		result.Body = body
	}
	if userID, ok := patch["userId"].(float64); ok {
		result.UserID = int(userID)
	}

	respondJSON(w, http.StatusOK, result)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetPostByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{})
}

func GetPostComments(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	comments := data.GetCommentsByPostID(id)
	respondJSON(w, http.StatusOK, comments)
}
