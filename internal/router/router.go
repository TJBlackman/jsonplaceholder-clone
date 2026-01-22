package router

import (
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/handlers"
	"github.com/TJBlackman/jsonplaceholder-clone/internal/middleware"
)

func New() chi.Router {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)
	r.Use(middleware.Delay)

	// Static landing page
	r.Get("/", handlers.ServeIndex)

	// API routes with JSON content type
	r.Route("/posts", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListPosts)
		r.Post("/", handlers.CreatePost)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetPost)
			r.Put("/", handlers.UpdatePost)
			r.Patch("/", handlers.PatchPost)
			r.Delete("/", handlers.DeletePost)
			r.Get("/comments", handlers.GetPostComments)
		})
	})

	r.Route("/comments", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListComments)
		r.Post("/", handlers.CreateComment)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetComment)
			r.Put("/", handlers.UpdateComment)
			r.Patch("/", handlers.PatchComment)
			r.Delete("/", handlers.DeleteComment)
		})
	})

	r.Route("/albums", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListAlbums)
		r.Post("/", handlers.CreateAlbum)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetAlbum)
			r.Put("/", handlers.UpdateAlbum)
			r.Patch("/", handlers.PatchAlbum)
			r.Delete("/", handlers.DeleteAlbum)
			r.Get("/photos", handlers.GetAlbumPhotos)
		})
	})

	r.Route("/photos", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListPhotos)
		r.Post("/", handlers.CreatePhoto)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetPhoto)
			r.Put("/", handlers.UpdatePhoto)
			r.Patch("/", handlers.PatchPhoto)
			r.Delete("/", handlers.DeletePhoto)
		})
	})

	r.Route("/todos", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListTodos)
		r.Post("/", handlers.CreateTodo)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetTodo)
			r.Put("/", handlers.UpdateTodo)
			r.Patch("/", handlers.PatchTodo)
			r.Delete("/", handlers.DeleteTodo)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.JSONContentType)
		r.Get("/", handlers.ListUsers)
		r.Post("/", handlers.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Put("/", handlers.UpdateUser)
			r.Patch("/", handlers.PatchUser)
			r.Delete("/", handlers.DeleteUser)
			r.Get("/posts", handlers.GetUserPosts)
			r.Get("/albums", handlers.GetUserAlbums)
			r.Get("/todos", handlers.GetUserTodos)
		})
	})

	return r
}
