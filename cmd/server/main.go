package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TJBlackman/jsonplaceholder-clone/internal/handlers"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/router"
)

//go:embed static
var staticFiles embed.FS

func main() {
	// Initialize static files for handlers
	handlers.StaticFiles = staticFiles

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := router.New()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
}
