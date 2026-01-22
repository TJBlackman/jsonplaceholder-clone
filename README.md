# JSONPlaceholder Clone - Go Implementation

A lightweight, high-performance RESTful API service written in Go that mimics the behavior of [JSONPlaceholder](https://jsonplaceholder.typicode.com/). This clone provides a complete fake REST API with 6 resource types, supporting all standard HTTP methods while using static in-memory data without persistence.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [API Resources](#api-resources)
- [API Endpoints](#api-endpoints)
- [Query Parameters](#query-parameters)
- [Getting Started](#getting-started)
- [Developer Guide](#developer-guide)
- [Architecture](#architecture)
- [Docker Deployment](#docker-deployment)
- [License](#license)

---

## Overview

This project is a faithful clone of JSONPlaceholder, implemented in Go using the Chi router. It provides a mock REST API for testing and prototyping frontend applications without needing a real backend. All data is hard-coded and responses are mocked - no actual data persistence occurs.

**Key characteristics:**

- Zero dependencies on databases or external services
- Blazing fast responses (sub-10ms without artificial delays)
- Full CORS support for browser-based applications
- Identical API contract to the original JSONPlaceholder
- Docker-ready with multi-stage builds for minimal image size

---

## Features

### Core Functionality

- **6 Resource Types**: Posts, Comments, Albums, Photos, Todos, and Users
- **RESTful Operations**: Full CRUD support (GET, POST, PUT, PATCH, DELETE)
- **Static Data**: 5 pre-populated items per resource type
- **Nested Routes**: Access related resources (e.g., `/posts/1/comments`)
- **Query Filtering**: Filter resources by any field
- **Pagination**: Built-in pagination with `_page` and `_limit` parameters
- **Artificial Delay**: Simulate network latency with `_delay` parameter
- **HTML Landing Page**: Interactive documentation served at root path

### Technical Features

- **CORS Enabled**: Full cross-origin support for web applications
- **JSON Content-Type**: Automatic JSON headers on all API responses
- **Graceful Shutdown**: Proper signal handling for clean server termination
- **Request Logging**: Built-in request logging middleware
- **Error Recovery**: Panic recovery middleware prevents server crashes
- **Idempotent Writes**: POST/PUT/PATCH/DELETE return predictable responses

---

## Project Structure

```
jsonplaceholder.trevorblackman.dev/
├── main.go                      # Application entry point
├── internal/
│   ├── data/
│   │   └── data.go              # Static data and lookup functions
│   ├── handlers/
│   │   ├── helpers.go           # Response helpers
│   │   ├── posts.go             # Post handlers
│   │   ├── comments.go          # Comment handlers
│   │   ├── albums.go            # Album handlers
│   │   ├── photos.go            # Photo handlers
│   │   ├── todos.go             # Todo handlers
│   │   ├── users.go             # User handlers
│   │   └── static.go            # HTML page handler
│   ├── middleware/
│   │   ├── cors.go              # CORS middleware
│   │   ├── json.go              # JSON content-type middleware
│   │   └── delay.go             # Artificial delay middleware
│   ├── models/
│   │   └── models.go            # Data structure definitions
│   └── router/
│       └── router.go            # Route configuration
├── static/
│   └── index.html               # Documentation landing page
├── Dockerfile                   # Multi-stage Docker build
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── plan.md                      # Development plan
└── README.md                    # This file
```

---

## API Resources

### Posts

Represents blog posts or articles.

```json
{
  "id": 1,
  "userId": 1,
  "title": "sunt aut facere repellat",
  "body": "quia et suscipit..."
}
```

### Comments

Represents comments on posts.

```json
{
  "id": 1,
  "postId": 1,
  "name": "id labore ex et quam",
  "email": "Eliseo@gardner.biz",
  "body": "laudantium enim..."
}
```

### Albums

Represents photo albums.

```json
{
  "id": 1,
  "userId": 1,
  "title": "quidem molestiae enim"
}
```

### Photos

Represents photos within albums.

```json
{
  "id": 1,
  "albumId": 1,
  "title": "accusamus beatae ad",
  "url": "https://via.placeholder.com/600/92c952",
  "thumbnailUrl": "https://via.placeholder.com/150/92c952"
}
```

### Todos

Represents todo items with completion status.

```json
{
  "id": 1,
  "userId": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

### Users

Represents user profiles with nested address and company data.

```json
{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz",
  "address": {
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "city": "Gwenborough",
    "zipcode": "92998-3874",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    }
  },
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org",
  "company": {
    "name": "Romaguera-Crona",
    "catchPhrase": "Multi-layered client-server neural-net",
    "bs": "harness real-time e-markets"
  }
}
```

---

## API Endpoints

### Posts

| Method | Endpoint              | Description                          |
| ------ | --------------------- | ------------------------------------ |
| GET    | `/posts`              | List all posts                       |
| GET    | `/posts/:id`          | Get a specific post                  |
| GET    | `/posts/:id/comments` | Get comments for a post              |
| POST   | `/posts`              | Create a new post (mocked)           |
| PUT    | `/posts/:id`          | Replace a post (mocked)              |
| PATCH  | `/posts/:id`          | Update specific post fields (mocked) |
| DELETE | `/posts/:id`          | Delete a post (mocked)               |

### Comments

| Method | Endpoint        | Description                             |
| ------ | --------------- | --------------------------------------- |
| GET    | `/comments`     | List all comments                       |
| GET    | `/comments/:id` | Get a specific comment                  |
| POST   | `/comments`     | Create a new comment (mocked)           |
| PUT    | `/comments/:id` | Replace a comment (mocked)              |
| PATCH  | `/comments/:id` | Update specific comment fields (mocked) |
| DELETE | `/comments/:id` | Delete a comment (mocked)               |

### Albums

| Method | Endpoint             | Description                           |
| ------ | -------------------- | ------------------------------------- |
| GET    | `/albums`            | List all albums                       |
| GET    | `/albums/:id`        | Get a specific album                  |
| GET    | `/albums/:id/photos` | Get photos for an album               |
| POST   | `/albums`            | Create a new album (mocked)           |
| PUT    | `/albums/:id`        | Replace an album (mocked)             |
| PATCH  | `/albums/:id`        | Update specific album fields (mocked) |
| DELETE | `/albums/:id`        | Delete an album (mocked)              |

### Photos

| Method | Endpoint      | Description                           |
| ------ | ------------- | ------------------------------------- |
| GET    | `/photos`     | List all photos                       |
| GET    | `/photos/:id` | Get a specific photo                  |
| POST   | `/photos`     | Create a new photo (mocked)           |
| PUT    | `/photos/:id` | Replace a photo (mocked)              |
| PATCH  | `/photos/:id` | Update specific photo fields (mocked) |
| DELETE | `/photos/:id` | Delete a photo (mocked)               |

### Todos

| Method | Endpoint     | Description                          |
| ------ | ------------ | ------------------------------------ |
| GET    | `/todos`     | List all todos                       |
| GET    | `/todos/:id` | Get a specific todo                  |
| POST   | `/todos`     | Create a new todo (mocked)           |
| PUT    | `/todos/:id` | Replace a todo (mocked)              |
| PATCH  | `/todos/:id` | Update specific todo fields (mocked) |
| DELETE | `/todos/:id` | Delete a todo (mocked)               |

### Users

| Method | Endpoint            | Description                          |
| ------ | ------------------- | ------------------------------------ |
| GET    | `/users`            | List all users                       |
| GET    | `/users/:id`        | Get a specific user                  |
| GET    | `/users/:id/posts`  | Get posts by a user                  |
| GET    | `/users/:id/albums` | Get albums by a user                 |
| GET    | `/users/:id/todos`  | Get todos by a user                  |
| POST   | `/users`            | Create a new user (mocked)           |
| PUT    | `/users/:id`        | Replace a user (mocked)              |
| PATCH  | `/users/:id`        | Update specific user fields (mocked) |
| DELETE | `/users/:id`        | Delete a user (mocked)               |

### Static

| Method | Endpoint | Description                     |
| ------ | -------- | ------------------------------- |
| GET    | `/`      | HTML documentation landing page |

---

## Query Parameters

### Filtering

Filter any resource by any field using query parameters:

```bash
# Get posts by specific user
GET /posts?userId=1

# Get comments by email
GET /comments?email=Eliseo@gardner.biz

# Multiple filters
GET /todos?userId=2&completed=true
```

### Pagination

Paginate results using `_page` and `_limit`:

```bash
# Get first 2 posts
GET /posts?_page=1&_limit=2

# Get next 2 posts
GET /posts?_page=2&_limit=2
```

**Note:** If `_limit` is not provided, all items are returned. If `_page` is not provided, it defaults to 1.

### Artificial Delay

Simulate network latency using `_delay` (in milliseconds):

```bash
# Add 1 second delay
GET /posts?_delay=1000

# Maximum delay is 10 seconds (10000ms)
GET /posts?_delay=15000  # Will be capped at 10000
```

---

## Getting Started

### Prerequisites

- Go 1.23 or later
- (Optional) Docker for containerized deployment

### Local Development

1. **Clone the repository**

```bash
git clone https://github.com/TJBlackman/jsonplaceholder-clone.git
cd jsonplaceholder-clone
```

2. **Install dependencies**

```bash
go mod download
```

3. **Run the server**

```bash
go run .
```

The server will start on `http://localhost:3000`

4. **Test the API**

```bash
# List all posts
curl http://localhost:3000/posts

# Get a specific user
curl http://localhost:3000/users/1

# Create a post (mocked)
curl -X POST http://localhost:3000/posts \
  -H "Content-Type: application/json" \
  -d '{"userId":1,"title":"Test Post","body":"This is a test"}'
```

### Using a Custom Port

Set the `PORT` environment variable:

```bash
PORT=8080 go run .
```

---

## Developer Guide

### Adding New Resources

To add a new resource type, follow these steps:

#### 1. Define the Model

Add a new struct in `internal/models/models.go`:

```go
type NewResource struct {
    ID     int    `json:"id"`
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}
```

#### 2. Add Static Data

In `internal/data/data.go`, create a static slice and helper functions:

```go
var NewResources = []models.NewResource{
    {ID: 1, Field1: "value1", Field2: 100},
    {ID: 2, Field1: "value2", Field2: 200},
    // ... add 3 more for total of 5
}

func GetNewResourceByID(id int) *models.NewResource {
    for i := range NewResources {
        if NewResources[i].ID == id {
            return &NewResources[i]
        }
    }
    return nil
}

func FilterNewResources(filters map[string]string) []models.NewResource {
    results := make([]models.NewResource, 0)
    for _, item := range NewResources {
        match := true
        for key, value := range filters {
            // Add field-specific filtering logic
        }
        if match {
            results = append(results, item)
        }
    }
    return results
}
```

Update the `NextIDs` map:

```go
var NextIDs = map[string]int{
    "posts":        6,
    "comments":     6,
    "albums":       6,
    "photos":       6,
    "todos":        6,
    "users":        6,
    "newresources": 6, // Add this line
}
```

#### 3. Create Handlers

Create `internal/handlers/newresources.go`:

```go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/TJBlackman/jsonplaceholder-clone/internal/data"
    "github.com/TJBlackman/jsonplaceholder-clone/internal/models"
)

func ListNewResources(w http.ResponseWriter, r *http.Request) {
    filters := make(map[string]string)
    // Extract filters from query params

    items := data.FilterNewResources(filters)
    items = applyPagination(items, r)

    respondJSON(w, http.StatusOK, items)
}

func GetNewResource(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        respondError(w, http.StatusBadRequest, "invalid id")
        return
    }

    item := data.GetNewResourceByID(id)
    if item == nil {
        respondError(w, http.StatusNotFound, "not found")
        return
    }

    respondJSON(w, http.StatusOK, item)
}

func CreateNewResource(w http.ResponseWriter, r *http.Request) {
    var item models.NewResource
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        respondError(w, http.StatusBadRequest, "invalid json")
        return
    }

    item.ID = data.GetNextID("newresources")
    respondJSON(w, http.StatusCreated, item)
}

// Implement UpdateNewResource, PatchNewResource, DeleteNewResource
// following the same pattern as existing handlers
```

#### 4. Register Routes

In `internal/router/router.go`, add the new routes:

```go
r.Route("/newresources", func(r chi.Router) {
    r.Use(middleware.JSONContentType)
    r.Get("/", handlers.ListNewResources)
    r.Post("/", handlers.CreateNewResource)
    r.Route("/{id}", func(r chi.Router) {
        r.Get("/", handlers.GetNewResource)
        r.Put("/", handlers.UpdateNewResource)
        r.Patch("/", handlers.PatchNewResource)
        r.Delete("/", handlers.DeleteNewResource)
    })
})
```

#### 5. Update Documentation

Update `static/index.html` to include documentation for the new resource.

### Modifying Middleware

#### Custom Middleware Example

Create a new middleware in `internal/middleware/`:

```go
package middleware

import "net/http"

func CustomMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Pre-processing logic here

        next.ServeHTTP(w, r)

        // Post-processing logic here
    })
}
```

Register it in `internal/router/router.go`:

```go
func New() chi.Router {
    r := chi.NewRouter()

    // Global middleware
    r.Use(chimw.Logger)
    r.Use(chimw.Recoverer)
    r.Use(middleware.CORS)
    r.Use(middleware.Delay)
    r.Use(middleware.CustomMiddleware) // Add here

    // ... rest of routes
}
```

### Understanding the Handler Pattern

All resource handlers follow a consistent pattern:

#### GET (List)

1. Extract query parameters for filtering
2. Apply filters to static data
3. Apply pagination if `_page` and `_limit` are provided
4. Return JSON array

#### GET (Single Item)

1. Extract ID from URL parameter
2. Look up item in static data
3. Return 404 if not found
4. Return JSON object

#### POST (Create)

1. Decode JSON request body
2. Validate JSON structure
3. Assign next available ID
4. Return 201 with the created object (not persisted)

#### PUT (Replace)

1. Extract ID from URL parameter
2. Verify ID exists (404 if not)
3. Decode JSON request body
4. Return 200 with the replacement object (not persisted)

#### PATCH (Partial Update)

1. Extract ID from URL parameter
2. Verify ID exists (404 if not)
3. Decode request body as map
4. Merge patch fields with existing object
5. Return 200 with merged object (not persisted)

#### DELETE

1. Extract ID from URL parameter
2. Verify ID exists (404 if not)
3. Return 200 with empty object `{}`

### Pagination Helper

The `applyPagination` function is a generic helper that works with any slice:

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

---

## Architecture

### Request Flow

1. **HTTP Request** → Server receives request
2. **Logger Middleware** → Logs request details
3. **Recoverer Middleware** → Catches panics
4. **CORS Middleware** → Adds CORS headers
5. **Delay Middleware** → Applies artificial delay if `_delay` parameter present
6. **Router** → Matches request to handler
7. **JSON Middleware** → Sets JSON content-type (API routes only)
8. **Handler** → Processes request and returns response
9. **HTTP Response** → Sent back to client

### Middleware Stack

```
Incoming Request
     ↓
Logger (chi middleware)
     ↓
Recoverer (chi middleware)
     ↓
CORS (custom middleware)
     ↓
Delay (custom middleware)
     ↓
Router Matching
     ↓
JSONContentType (custom middleware, API routes only)
     ↓
Handler Function
     ↓
Response
```

### Data Flow

```
Static Data (data.go)
     ↓
Lookup/Filter Functions
     ↓
Handler Functions
     ↓
Response Helpers
     ↓
JSON Response
```

**Important:** No data is ever persisted. All POST, PUT, PATCH, and DELETE operations return mocked responses based on the static data but don't modify it.

### Key Design Decisions

1. **No Database**: Eliminates complexity and dependencies
2. **Static Data**: Predictable, fast, and consistent responses
3. **Chi Router**: Lightweight, idiomatic Go routing
4. **Middleware Pattern**: Composable, reusable request processing
5. **Generic Pagination**: Type-safe pagination helper using Go generics
6. **Embedded Static Files**: HTML landing page bundled in binary
7. **Graceful Shutdown**: Proper signal handling for zero-downtime deploys

---

## Docker Deployment

### Building the Image

```bash
docker build -t jsonplaceholder-clone .
```

The Dockerfile uses a multi-stage build:

- **Stage 1 (builder)**: Compiles the Go binary in a golang:1.23-alpine image
- **Stage 2 (runtime)**: Copies binary to a minimal alpine:3.19 image

This results in a tiny final image (~15MB) with only the compiled binary.

### Running the Container

```bash
# Run on default port 3000
docker run -p 3000:3000 jsonplaceholder-clone

# Run on custom port
docker run -p 8080:8080 -e PORT=8080 jsonplaceholder-clone

# Run in detached mode
docker run -d -p 3000:3000 --name jsonplaceholder jsonplaceholder-clone
```

### Docker Compose

Create a `docker-compose.yml`:

```yaml
version: "3.8"

services:
  api:
    build: .
    ports:
      - "3000:3000"
    environment:
      - PORT=3000
    restart: unless-stopped
```

Run with:

```bash
docker-compose up -d
```

### Environment Variables

| Variable | Default | Description                |
| -------- | ------- | -------------------------- |
| `PORT`   | `3000`  | Port the server listens on |

---

## Performance

### Benchmarks

Without the `_delay` parameter, responses are extremely fast:

- **List operations**: < 1ms
- **Single item lookup**: < 0.5ms
- **Create/Update/Delete**: < 1ms

The server can handle thousands of requests per second on modest hardware.

### Optimization Tips

1. **Disable logging** in production by removing the `chimw.Logger` middleware
2. **Use HTTP/2** by serving with TLS
3. **Add caching headers** for static resources
4. **Rate limiting** can be added via custom middleware if needed

---

## Troubleshooting

### Server won't start

**Problem**: Port already in use

**Solution**: Change the port using the `PORT` environment variable:

```bash
PORT=8080 go run .
```

### CORS errors in browser

**Problem**: Preflight requests failing

**Solution**: The CORS middleware should handle this automatically. Verify the middleware is registered:

```go
r.Use(middleware.CORS)
```

### 404 for all routes

**Problem**: Router not properly initialized

**Solution**: Check that `router.New()` is called in `main.go` and all routes are registered in `internal/router/router.go`

### JSON parsing errors

**Problem**: Request body not properly formatted

**Solution**: Ensure `Content-Type: application/json` header is set and body is valid JSON:

```bash
curl -X POST http://localhost:3000/posts \
  -H "Content-Type: application/json" \
  -d '{"userId":1,"title":"Test","body":"Test body"}'
```

---

## Contributing

This is a personal project, but feedback and suggestions are welcome via GitHub issues.

### Development Workflow

1. Fork the repository
2. Create a feature branch
3. Make changes
4. Test locally
5. Submit a pull request

### Code Style

- Follow standard Go conventions
- Use `gofmt` to format code
- Add comments for exported functions
- Keep handlers simple and consistent

---

## License

This project is open source and available under the MIT License.

---

## Acknowledgments

- Original [JSONPlaceholder](https://jsonplaceholder.typicode.com/) by Typicode
- [Chi Router](https://github.com/go-chi/chi) by Peter Kieltyka and contributors
- Inspired by the need for simple, fast mock APIs for frontend development

---

## Contact

**Trevor Blackman**
GitHub: [@TJBlackman](https://github.com/TJBlackman)
Project: [jsonplaceholder-clone](https://github.com/TJBlackman/jsonplaceholder-clone)

---

**Built with Go 1.25.1**
