package handlers

import (
	"embed"
	"net/http"
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
