# JSONPlaceholder Clone - Go Implementation Guide

## Project Overview

A RESTful API service in Go that mimics JSONPlaceholder. Uses hard-coded static data for reads and returns mocked responses for writes. No persistence layer.

---

## Stage 1: Project Initialization (COMPLETED)

### 1.1 Initialize Go Module (COMPLETED)

```bash
mkdir jsonplaceholder-go && cd jsonplaceholder-go
go mod init github.com/yourname/jsonplaceholder-go
```

### 1.2 Install Dependencies (COMPLETED)

```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```

### 1.3 Project Structure (COMPLETED)

```text
jsonplaceholder-go/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── internal/
│   ├── data/
│   │   └── data.go           # Hard-coded static data
│   ├── handlers/
│   │   ├── posts.go
│   │   ├── comments.go
│   │   ├── albums.go
│   │   ├── photos.go
│   │   ├── todos.go
│   │   ├── users.go
│   │   └── static.go         # Serve index.html
│   ├── middleware/
│   │   ├── cors.go
│   │   ├── delay.go
│   │   └── json.go           # Content-Type enforcement
│   ├── models/
│   │   └── models.go         # Struct definitions
│   └── router/
│       └── router.go         # Route definitions
├── static/
│   └── index.html
├── Dockerfile
└── go.mod
```

---

## Stage 2: Models (COMPLETED)

### 2.1 Define Structs (`internal/models/models.go`)

```go
package models

type User struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Username string  `json:"username"`
    Email    string  `json:"email"`
    Address  Address `json:"address"`
    Phone    string  `json:"phone"`
    Website  string  `json:"website"`
    Company  Company `json:"company"`
}

type Address struct {
    Street  string `json:"street"`
    Suite   string `json:"suite"`
    City    string `json:"city"`
    Zipcode string `json:"zipcode"`
    Geo     Geo    `json:"geo"`
}

type Geo struct {
    Lat string `json:"lat"`
    Lng string `json:"lng"`
}

type Company struct {
    Name        string `json:"name"`
    CatchPhrase string `json:"catchPhrase"`
    BS          string `json:"bs"`
}

type Post struct {
    ID     int    `json:"id"`
    UserID int    `json:"userId"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

type Comment struct {
    ID     int    `json:"id"`
    PostID int    `json:"postId"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}

type Album struct {
    ID     int    `json:"id"`
    UserID int    `json:"userId"`
    Title  string `json:"title"`
}

type Photo struct {
    ID           int    `json:"id"`
    AlbumID      int    `json:"albumId"`
    Title        string `json:"title"`
    URL          string `json:"url"`
    ThumbnailURL string `json:"thumbnailUrl"`
}

type Todo struct {
    ID        int    `json:"id"`
    UserID    int    `json:"userId"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```

---

## Stage 3: Static Data (COMPLETED)

### 3.1 Hard-coded Data (`internal/data/data.go`)

Define 5 instances of each resource type as package-level variables:

```go
package data

import "github.com/yourname/jsonplaceholder-go/internal/models"

var Users = []models.User{
    {
        ID:       1,
        Name:     "Leanne Graham",
        Username: "Bret",
        Email:    "Sincere@april.biz",
        Address: models.Address{
            Street:  "Kulas Light",
            Suite:   "Apt. 556",
            City:    "Gwenborough",
            Zipcode: "92998-3874",
            Geo:     models.Geo{Lat: "-37.3159", Lng: "81.1496"},
        },
        Phone:   "1-770-736-8031 x56442",
        Website: "hildegard.org",
        Company: models.Company{
            Name:        "Romaguera-Crona",
            CatchPhrase: "Multi-layered client-server neural-net",
            BS:          "harness real-time e-markets",
        },
    },
    // ... 4 more users
}

var Posts = []models.Post{
    {ID: 1, UserID: 1, Title: "sunt aut facere repellat", Body: "quia et suscipit..."},
    {ID: 2, UserID: 1, Title: "qui est esse", Body: "est rerum tempore..."},
    {ID: 3, UserID: 2, Title: "ea molestias quasi", Body: "et iusto sed quo..."},
    {ID: 4, UserID: 2, Title: "eum et est occaecati", Body: "ullam et saepe..."},
    {ID: 5, UserID: 3, Title: "nesciunt quas odio", Body: "repudiandae veniam..."},
}

var Comments = []models.Comment{
    {ID: 1, PostID: 1, Name: "id labore ex et quam", Email: "Eliseo@gardner.biz", Body: "laudantium enim..."},
    {ID: 2, PostID: 1, Name: "quo vero reiciendis", Email: "Jayne_Kuhic@sydney.com", Body: "est natus enim..."},
    {ID: 3, PostID: 2, Name: "odio adipisci rerum", Email: "Nikita@garfield.biz", Body: "quia molestiae..."},
    {ID: 4, PostID: 2, Name: "alias odio sit", Email: "Lew@alysha.tv", Body: "non et atque..."},
    {ID: 5, PostID: 3, Name: "vero eaque aliquid", Email: "Hayden@althea.biz", Body: "harum non quasi..."},
}

var Albums = []models.Album{
    {ID: 1, UserID: 1, Title: "quidem molestiae enim"},
    {ID: 2, UserID: 1, Title: "sunt qui excepturi"},
    {ID: 3, UserID: 2, Title: "omnis laborum odio"},
    {ID: 4, UserID: 2, Title: "non esse culpa"},
    {ID: 5, UserID: 3, Title: "eaque aut omnis"},
}

var Photos = []models.Photo{
    {ID: 1, AlbumID: 1, Title: "accusamus beatae ad", URL: "https://via.placeholder.com/600/92c952", ThumbnailURL: "https://via.placeholder.com/150/92c952"},
    {ID: 2, AlbumID: 1, Title: "reprehenderit est deserunt", URL: "https://via.placeholder.com/600/771796", ThumbnailURL: "https://via.placeholder.com/150/771796"},
    {ID: 3, AlbumID: 2, Title: "officia porro iure", URL: "https://via.placeholder.com/600/24f355", ThumbnailURL: "https://via.placeholder.com/150/24f355"},
    {ID: 4, AlbumID: 2, Title: "culpa odio esse", URL: "https://via.placeholder.com/600/d32776", ThumbnailURL: "https://via.placeholder.com/150/d32776"},
    {ID: 5, AlbumID: 3, Title: "natus nisi omnis", URL: "https://via.placeholder.com/600/f66b97", ThumbnailURL: "https://via.placeholder.com/150/f66b97"},
}

var Todos = []models.Todo{
    {ID: 1, UserID: 1, Title: "delectus aut autem", Completed: false},
    {ID: 2, UserID: 1, Title: "quis ut nam facilis", Completed: true},
    {ID: 3, UserID: 2, Title: "fugiat veniam minus", Completed: false},
    {ID: 4, UserID: 2, Title: "et porro tempora", Completed: true},
    {ID: 5, UserID: 3, Title: "laboriosam mollitia", Completed: false},
}
```

### 3.2 Helper Functions

Add lookup and filter functions to the data package:

```go
func GetPostByID(id int) *models.Post
func GetPostsByUserID(userID int) []models.Post
func GetCommentByID(id int) *models.Comment
func GetCommentsByPostID(postID int) []models.Comment
func GetUserByID(id int) *models.User
func GetAlbumByID(id int) *models.Album
func GetAlbumsByUserID(userID int) []models.Album
func GetPhotoByID(id int) *models.Photo
func GetPhotosByAlbumID(albumID int) []models.Photo
func GetTodoByID(id int) *models.Todo
func GetTodosByUserID(userID int) []models.Todo
func FilterPosts(filters map[string]string) []models.Post
// ... similar filter functions for other resources
```

### 3.3 Next ID Tracking

Track the next available ID for mocked POST responses:

```go
var NextIDs = map[string]int{
    "posts":    6,
    "comments": 6,
    "albums":   6,
    "photos":   6,
    "todos":    6,
    "users":    6,
}

func GetNextID(resource string) int {
    return NextIDs[resource]
}
```

---

## Stage 4: HTTP Handlers (COMPLETED)

### 4.1 Response Helpers (`internal/handlers/helpers.go`)

```go
package handlers

import (
    "encoding/json"
    "net/http"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{"error": message})
}
```

### 4.2 Handler Pattern

Each resource handler follows this pattern:

**List (GET /resource)**:

- Parse query params for filters
- Parse `_page` and `_limit` for pagination
- Filter static data slice
- Apply pagination
- Return JSON array

**Get (GET /resource/:id)**:

- Extract ID from URL param
- Look up in static data
- Return 404 if not found
- Return JSON object

**Create (POST /resource)**:

- Decode request body
- Validate JSON is parseable
- Assign next ID to the object
- Return 201 with the object (not persisted)

**Update (PUT /resource/:id)**:

- Extract ID from URL param
- Check if ID exists in static data (return 404 if not)
- Decode request body
- Return 200 with merged object (not persisted)

**Patch (PATCH /resource/:id)**:

- Extract ID from URL param
- Check if ID exists in static data (return 404 if not)
- Decode request body as map
- Return 200 with original object merged with patch fields (not persisted)

**Delete (DELETE /resource/:id)**:

- Extract ID from URL param
- Check if ID exists in static data (return 404 if not)
- Return 200 with empty object `{}`

### 4.3 Example Handler (`internal/handlers/posts.go`)

```go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/yourname/jsonplaceholder-go/internal/data"
    "github.com/yourname/jsonplaceholder-go/internal/models"
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
```

### 4.4 Pagination Helper

```go
func applyPagination[T any](items []T, r *http.Request) []T {
    page, _ := strconv.Atoi(r.URL.Query().Get("_page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("_limit"))

    if limit <= 0 {
        return items
    }

    if page <= 0 {
        page = 1
    }

    start := (page - 1) * limit
    if start >= len(items) {
        return []T{}
    }

    end := start + limit
    if end > len(items) {
        end = len(items)
    }

    return items[start:end]
}
```

### 4.5 Implement Remaining Handlers

Create similar handlers for:

- `internal/handlers/comments.go`
- `internal/handlers/albums.go`
- `internal/handlers/photos.go`
- `internal/handlers/todos.go`
- `internal/handlers/users.go`

Each follows the same pattern with resource-specific fields and relationships.

---

## Stage 5: Static Landing Page (COMPLETED)

### 5.1 Create `static/index.html`

Include:

- Project title and description
- List of all 6 resource types with example JSON
- Route documentation table
- Example curl commands
- Query parameter documentation (`_page`, `_limit`, `_delay`, filters)

### 5.2 Static Handler (`internal/handlers/static.go`)

```go
package handlers

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed static
var staticFiles embed.FS

func StaticHandler() http.Handler {
    sub, _ := fs.Sub(staticFiles, "static")
    return http.FileServer(http.FS(sub))
}
```

Move embed directive to appropriate location or pass embedded FS from main.

---

## Stage 6: Middleware (COMPLETED)

### 6.1 CORS (`internal/middleware/cors.go`)

```go
package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

### 6.2 JSON Content-Type (`internal/middleware/json.go`)

```go
package middleware

import "net/http"

func JSONContentType(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        next.ServeHTTP(w, r)
    })
}
```

### 6.3 Delay (`internal/middleware/delay.go`)

```go
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
```

---

## Stage 7: Router (COMPLETED)

### 7.1 Route Definitions (`internal/router/router.go`)

```go
package router

import (
    "github.com/go-chi/chi/v5"
    chimw "github.com/go-chi/chi/v5/middleware"
    "github.com/yourname/jsonplaceholder-go/internal/handlers"
    "github.com/yourname/jsonplaceholder-go/internal/middleware"
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
```

---

## Stage 8: Entry Point (COMPLETED)

### 8.1 Main (`cmd/server/main.go`)

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/yourname/jsonplaceholder-go/internal/router"
)

func main() {
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
```

---

## Stage 9: Containerization (COMPLETED)

### 9.1 Dockerfile

```dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]
```

### 9.2 Build and Run

```bash
docker build -t jsonplaceholder-go .
docker run -p 3000:3000 jsonplaceholder-go
```

---

## Stage 10: Testing (COMPLETED)

### 10.1 Handler Tests

Test each handler with `httptest`:

```go
func TestListPosts(t *testing.T) {
    req := httptest.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()

    r := router.New()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", w.Code)
    }

    var posts []models.Post
    json.Unmarshal(w.Body.Bytes(), &posts)

    if len(posts) != 5 {
        t.Errorf("expected 5 posts, got %d", len(posts))
    }
}
```

### 10.2 Test Cases

For each resource:

- GET /resource returns 5 items
- GET /resource/1 returns single item
- GET /resource/999 returns 404
- POST /resource returns 201 with id=6
- PUT /resource/1 returns 200 with updated data
- PATCH /resource/1 returns 200 with merged data
- DELETE /resource/1 returns 200 with empty object
- GET /resource?filter=value returns filtered results
- GET /resource?\_page=1&\_limit=2 returns 2 items
- Nested routes return correct related resources

---

## Validation Checklist

### Functional

- [ ] GET / serves HTML landing page
- [ ] All 6 resources return 5 hard-coded items
- [ ] GET by ID returns single item or 404
- [ ] POST returns 201 with id=6 (not persisted)
- [ ] PUT returns 200 with sent data (not persisted)
- [ ] PATCH returns 200 with merged data (not persisted)
- [ ] DELETE returns 200 with empty object (not persisted)
- [ ] Query param filtering works
- [ ] Pagination with `_page` and `_limit` works
- [ ] `_delay` adds artificial latency
- [ ] Nested routes work (`/posts/1/comments`, `/users/1/posts`)

### Headers

- [ ] `Content-Type: application/json` on API responses
- [ ] `Content-Type: text/html` on landing page
- [ ] CORS headers present
- [ ] OPTIONS returns 204

### Performance

- [ ] Responses under 10ms (excluding `_delay`)

---

## Endpoint Reference

| Method | Path                | Description              |
| ------ | ------------------- | ------------------------ |
| GET    | /                   | Landing page (HTML)      |
| GET    | /posts              | List all posts           |
| GET    | /posts/:id          | Get post by ID           |
| GET    | /posts/:id/comments | Get comments for post    |
| POST   | /posts              | Create post (mocked)     |
| PUT    | /posts/:id          | Replace post (mocked)    |
| PATCH  | /posts/:id          | Update post (mocked)     |
| DELETE | /posts/:id          | Delete post (mocked)     |
| GET    | /comments           | List all comments        |
| GET    | /comments/:id       | Get comment by ID        |
| POST   | /comments           | Create comment (mocked)  |
| PUT    | /comments/:id       | Replace comment (mocked) |
| PATCH  | /comments/:id       | Update comment (mocked)  |
| DELETE | /comments/:id       | Delete comment (mocked)  |
| GET    | /albums             | List all albums          |
| GET    | /albums/:id         | Get album by ID          |
| GET    | /albums/:id/photos  | Get photos for album     |
| POST   | /albums             | Create album (mocked)    |
| PUT    | /albums/:id         | Replace album (mocked)   |
| PATCH  | /albums/:id         | Update album (mocked)    |
| DELETE | /albums/:id         | Delete album (mocked)    |
| GET    | /photos             | List all photos          |
| GET    | /photos/:id         | Get photo by ID          |
| POST   | /photos             | Create photo (mocked)    |
| PUT    | /photos/:id         | Replace photo (mocked)   |
| PATCH  | /photos/:id         | Update photo (mocked)    |
| DELETE | /photos/:id         | Delete photo (mocked)    |
| GET    | /todos              | List all todos           |
| GET    | /todos/:id          | Get todo by ID           |
| POST   | /todos              | Create todo (mocked)     |
| PUT    | /todos/:id          | Replace todo (mocked)    |
| PATCH  | /todos/:id          | Update todo (mocked)     |
| DELETE | /todos/:id          | Delete todo (mocked)     |
| GET    | /users              | List all users           |
| GET    | /users/:id          | Get user by ID           |
| GET    | /users/:id/posts    | Get posts by user        |
| GET    | /users/:id/albums   | Get albums by user       |
| GET    | /users/:id/todos    | Get todos by user        |
| POST   | /users              | Create user (mocked)     |
| PUT    | /users/:id          | Replace user (mocked)    |
| PATCH  | /users/:id          | Update user (mocked)     |
| DELETE | /users/:id          | Delete user (mocked)     |
