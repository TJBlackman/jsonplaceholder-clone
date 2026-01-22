package middleware

import (
	"net/http"
	"strconv"
	"time"
)

func Delay(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if delayStr := r.URL.Query().Get("_delay"); delayStr != "" {
			if delay, err := strconv.Atoi(delayStr); err == nil && delay > 0 {
				if delay > 10000 {
					delay = 10000
				}
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
		}
		next.ServeHTTP(w, r)
	})
}
