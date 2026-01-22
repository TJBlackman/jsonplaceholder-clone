package handlers

import (
	"embed"
	"net/http"
	"path/filepath"
)

var StaticFiles embed.FS

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	// Read the index.html file from embedded filesystem
	content, err := StaticFiles.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func ServeStaticFile(w http.ResponseWriter, r *http.Request) {
	// Extract the filename from the URL path
	filename := filepath.Base(r.URL.Path)

	// Read the file from embedded filesystem
	content, err := StaticFiles.ReadFile("static/" + filename)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Set content type based on file extension
	ext := filepath.Ext(filename)
	contentType := "application/octet-stream"
	if ext == ".json" {
		contentType = "application/json"
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
