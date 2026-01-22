package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/data"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/models"
	"github.com/go-chi/chi/v5"
)

func ListPhotos(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]string)
	if albumID := r.URL.Query().Get("albumId"); albumID != "" {
		filters["albumId"] = albumID
	}

	photos := data.FilterPhotos(filters)
	photos = applyPagination(photos, r)

	respondJSON(w, http.StatusOK, photos)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	photo := data.GetPhotoByID(id)
	if photo == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, photo)
}

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	photo.ID = data.GetNextID("photos")
	respondJSON(w, http.StatusCreated, photo)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetPhotoByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	var photo models.Photo
	if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	photo.ID = id
	respondJSON(w, http.StatusOK, photo)
}

func PatchPhoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing := data.GetPhotoByID(id)
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
	if url, ok := patch["url"].(string); ok {
		result.URL = url
	}
	if thumbnailURL, ok := patch["thumbnailUrl"].(string); ok {
		result.ThumbnailURL = thumbnailURL
	}
	if albumID, ok := patch["albumId"].(float64); ok {
		result.AlbumID = int(albumID)
	}

	respondJSON(w, http.StatusOK, result)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if data.GetPhotoByID(id) == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{})
}
