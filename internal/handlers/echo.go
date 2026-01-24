package handlers

import (
	"net/http"
	"time"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "ok"
	}

	response := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"message":   message,
	}

	respondJSON(w, http.StatusOK, response)
}
