package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/data"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListAlbums(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	if userID := r.URL.Query().Get("userId"); userID != "" {
		filters["userId"] = userID
	}

	albums := data.FilterAlbums(filters)
	albums = applyPagination(albums, r)

	respondJSON(w, http.StatusOK, albums)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	album := data.GetAlbumByID(id)
	if album == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, album)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	album.ID = data.GetNextID("albums")
	respondJSON(w, http.StatusCreated, album)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetAlbumByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	album.ID = id
	respondJSON(w, http.StatusOK, album)
}

func PatchAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing := data.GetAlbumByID(id)
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
	if userID, ok := patch["userId"].(float64); ok {
		result.UserID = int(userID)
	}

	respondJSON(w, http.StatusOK, result)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetAlbumByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{})
}

func GetAlbumPhotos(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	photos := data.GetPhotosByAlbumID(id)
	respondJSON(w, http.StatusOK, photos)
}
