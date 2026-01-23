package data

import "github.com/TJBlackman/jsonplaceholder-clone/internal/models"

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
	{
		ID:       2,
		Name:     "Ervin Howell",
		Username: "Antonette",
		Email:    "Shanna@melissa.tv",
		Address: models.Address{
			Street:  "Victor Plains",
			Suite:   "Suite 879",
			City:    "Wisokyburgh",
			Zipcode: "90566-7771",
			Geo:     models.Geo{Lat: "-43.9509", Lng: "-34.4618"},
		},
		Phone:   "010-692-6593 x09125",
		Website: "anastasia.net",
		Company: models.Company{
			Name:        "Deckow-Crist",
			CatchPhrase: "Proactive didactic contingency",
			BS:          "synergize scalable supply-chains",
		},
	},
	{
		ID:       3,
		Name:     "Clementine Bauch",
		Username: "Samantha",
		Email:    "Nathan@yesenia.net",
		Address: models.Address{
			Street:  "Douglas Extension",
			Suite:   "Suite 847",
			City:    "McKenziehaven",
			Zipcode: "59590-4157",
			Geo:     models.Geo{Lat: "-68.6102", Lng: "-47.0653"},
		},
		Phone:   "1-463-123-4447",
		Website: "ramiro.info",
		Company: models.Company{
			Name:        "Romaguera-Jacobson",
			CatchPhrase: "Face to face bifurcated interface",
			BS:          "e-enable strategic applications",
		},
	},
	{
		ID:       4,
		Name:     "Patricia Lebsack",
		Username: "Karianne",
		Email:    "Julianne.OConner@kory.org",
		Address: models.Address{
			Street:  "Hoeger Mall",
			Suite:   "Apt. 692",
			City:    "South Elvis",
			Zipcode: "53919-4257",
			Geo:     models.Geo{Lat: "29.4572", Lng: "-164.2990"},
		},
		Phone:   "493-170-9623 x156",
		Website: "kale.biz",
		Company: models.Company{
			Name:        "Robel-Corkery",
			CatchPhrase: "Multi-tiered zero tolerance productivity",
			BS:          "transition cutting-edge web services",
		},
	},
	{
		ID:       5,
		Name:     "Chelsey Dietrich",
		Username: "Kamren",
		Email:    "Lucio_Hettinger@annie.ca",
		Address: models.Address{
			Street:  "Skiles Walks",
			Suite:   "Suite 351",
			City:    "Roscoeview",
			Zipcode: "33263",
			Geo:     models.Geo{Lat: "-31.8129", Lng: "62.5342"},
		},
		Phone:   "(254)954-1289",
		Website: "demarco.info",
		Company: models.Company{
			Name:        "Keebler LLC",
			CatchPhrase: "User-centric fault-tolerant solution",
			BS:          "revolutionize end-to-end systems",
		},
	},
}

var Posts = []models.Post{
	{ID: 1, UserID: 1, Title: "Migrating from REST to gRPC for Internal Services", Body: "We recently shifted our microservices communication to gRPC to leverage Protobuf serialization. The performance gains in high-throughput scenarios have been significant compared to our old JSON over HTTP setup."},
	{ID: 2, UserID: 1, Title: "Why We Use Go for High-Concurrency Backend Systems", Body: "Go's goroutines and channels make it incredibly easy to handle thousands of concurrent requests without the overhead of traditional threading. The runtime scheduler does a great job of managing CPU resources efficiently."},
	{ID: 3, UserID: 2, Title: "Optimizing TypeScript Monorepos with Nx", Body: "Managing multiple TypeScript packages can be a headache. Nx has helped us reduce build times via computation caching and improved our developer experience with better dependency graphing."},
	{ID: 4, UserID: 2, Title: "The Case for Functional Programming in Modern JS", Body: "Even without a fully functional language, using patterns like immutability and pure functions in TypeScript leads to more predictable and testable codebases. It drastically reduces side-effect related bugs."},
	{ID: 5, UserID: 3, Title: "Handling Database Migrations in Distributed Environments", Body: "Running migrations becomes tricky when you have multiple instances of a service scaling horizontally. We've adopted a pattern of ensuring migrations are always backward compatible to prevent downtime during deployments."},
}

var Comments = []models.Comment{
	{ID: 1, PostID: 1, Name: "Technical Query on Protobuf", Email: "dev.lead@techcorp.io", Body: "Great write-up. Did you face any challenges with browser-based clients needing to communicate with these gRPC services, or are you using gRPC-Web?"},
	{ID: 2, PostID: 1, Name: "Performance Benchmarks", Email: "sre_eng@cloudscale.net", Body: "I'd love to see some actual latency numbers. We've been considering a similar move but the overhead of maintaining proto files is a concern for our team."},
	{ID: 3, PostID: 2, Name: "Goroutine Management", Email: "go_enthusiast@gopher.org", Body: "Completely agree. One thing to watch out for is goroutine leaks when using unbuffered channels without proper context cancellation."},
	{ID: 4, PostID: 2, Name: "Comparison with Node.js", Email: "node_fan@jsland.com", Body: "While Go is fast, the ecosystem in Node for quick prototyping is still hard to beat. Do you find the strictly typed nature of Go slows down your initial development phase?"},
	{ID: 5, PostID: 3, Name: "Nx vs Turbo", Email: "frontend_arch@webstack.sh", Body: "We just switched from Lerna to Nx and the caching is a game changer. Have you tried Turborepo as well, or did Nx just fit your requirements better?"},
}

var Albums = []models.Album{
	{ID: 1, UserID: 1, Title: "Echoes of the Silent Valley"},
	{ID: 2, UserID: 1, Title: "Fragments of a Lucid Dream"},
	{ID: 3, UserID: 2, Title: "The Persistence of Twilight"},
	{ID: 4, UserID: 2, Title: "Symphony of Rust and Gold"},
	{ID: 5, UserID: 3, Title: "Shadows on the Marble Floor"},
}

var Photos = []models.Photo{
	{ID: 1, AlbumID: 1, Title: "accusamus beatae ad facilis cum similique qui sunt", URL: "https://placehold.co/600x600/92c952/white", ThumbnailURL: "https://placehold.co/150x150/92c952/white"},
	{ID: 2, AlbumID: 1, Title: "reprehenderit est deserunt velit ipsam", URL: "https://placehold.co/600x600/771796/white", ThumbnailURL: "https://placehold.co/150x150/771796/white"},
	{ID: 3, AlbumID: 2, Title: "officia porro iure quia iusto qui ipsa ut modi", URL: "https://placehold.co/600x600/24f355/white", ThumbnailURL: "https://placehold.co/150x150/24f355/white"},
	{ID: 4, AlbumID: 2, Title: "culpa odio esse rerum omnis laboriosam voluptate repudiandae", URL: "https://placehold.co/600x600/d32776/white", ThumbnailURL: "https://placehold.co/150x150/d32776/white"},
	{ID: 5, AlbumID: 3, Title: "natus nisi omnis corporis facere molestiae rerum in", URL: "https://placehold.co/600x600/f66b97/white", ThumbnailURL: "https://placehold.co/150x150/f66b97/white"},
}

var Todos = []models.Todo{
	{ID: 1, UserID: 1, Title: "Implement JWT authentication middleware", Completed: false},
	{ID: 2, UserID: 1, Title: "Refactor database connection pool settings", Completed: true},
	{ID: 3, UserID: 2, Title: "Write unit tests for the Go service layer", Completed: false},
	{ID: 4, UserID: 2, Title: "Optimize TypeScript type definitions for API responses", Completed: true},
	{ID: 5, UserID: 3, Title: "Fix memory leak in background worker routine", Completed: false},
}

// GetPostByID returns a post by ID or nil if not found
func GetPostByID(id int) *models.Post {
	for i := range Posts {
		if Posts[i].ID == id {
			return &Posts[i]
		}
	}
	return nil
}

// GetCommentByID returns a comment by ID or nil if not found
func GetCommentByID(id int) *models.Comment {
	for i := range Comments {
		if Comments[i].ID == id {
			return &Comments[i]
		}
	}
	return nil
}

// GetUserByID returns a user by ID or nil if not found
func GetUserByID(id int) *models.User {
	for i := range Users {
		if Users[i].ID == id {
			return &Users[i]
		}
	}
	return nil
}

// GetAlbumByID returns an album by ID or nil if not found
func GetAlbumByID(id int) *models.Album {
	for i := range Albums {
		if Albums[i].ID == id {
			return &Albums[i]
		}
	}
	return nil
}

// GetPhotoByID returns a photo by ID or nil if not found
func GetPhotoByID(id int) *models.Photo {
	for i := range Photos {
		if Photos[i].ID == id {
			return &Photos[i]
		}
	}
	return nil
}

// GetTodoByID returns a todo by ID or nil if not found
func GetTodoByID(id int) *models.Todo {
	for i := range Todos {
		if Todos[i].ID == id {
			return &Todos[i]
		}
	}
	return nil
}

// GetPostsByUserID returns all posts for a given user ID
func GetPostsByUserID(userID int) []models.Post {
	var result []models.Post
	for _, post := range Posts {
		if post.UserID == userID {
			result = append(result, post)
		}
	}
	return result
}

// GetCommentsByPostID returns all comments for a given post ID
func GetCommentsByPostID(postID int) []models.Comment {
	var result []models.Comment
	for _, comment := range Comments {
		if comment.PostID == postID {
			result = append(result, comment)
		}
	}
	return result
}

// GetAlbumsByUserID returns all albums for a given user ID
func GetAlbumsByUserID(userID int) []models.Album {
	var result []models.Album
	for _, album := range Albums {
		if album.UserID == userID {
			result = append(result, album)
		}
	}
	return result
}

// GetPhotosByAlbumID returns all photos for a given album ID
func GetPhotosByAlbumID(albumID int) []models.Photo {
	var result []models.Photo
	for _, photo := range Photos {
		if photo.AlbumID == albumID {
			result = append(result, photo)
		}
	}
	return result
}

// GetTodosByUserID returns all todos for a given user ID
func GetTodosByUserID(userID int) []models.Todo {
	var result []models.Todo
	for _, todo := range Todos {
		if todo.UserID == userID {
			result = append(result, todo)
		}
	}
	return result
}

// FilterPosts filters posts based on query parameters
func FilterPosts(filters map[string]string) []models.Post {
	result := make([]models.Post, 0)

	for _, post := range Posts {
		match := true

		if userID, ok := filters["userId"]; ok {
			if userID != "" {
				if post.UserID != atoi(userID) {
					match = false
				}
			}
		}

		if id, ok := filters["id"]; ok {
			if id != "" {
				if post.ID != atoi(id) {
					match = false
				}
			}
		}

		if match {
			result = append(result, post)
		}
	}

	return result
}

// FilterComments filters comments based on query parameters
func FilterComments(filters map[string]string) []models.Comment {
	result := make([]models.Comment, 0)

	for _, comment := range Comments {
		match := true

		if postID, ok := filters["postId"]; ok {
			if postID != "" {
				if comment.PostID != atoi(postID) {
					match = false
				}
			}
		}

		if id, ok := filters["id"]; ok {
			if id != "" {
				if comment.ID != atoi(id) {
					match = false
				}
			}
		}

		if match {
			result = append(result, comment)
		}
	}

	return result
}

// FilterAlbums filters albums based on query parameters
func FilterAlbums(filters map[string]string) []models.Album {
	result := make([]models.Album, 0)

	for _, album := range Albums {
		match := true

		if userID, ok := filters["userId"]; ok {
			if userID != "" {
				if album.UserID != atoi(userID) {
					match = false
				}
			}
		}

		if id, ok := filters["id"]; ok {
			if id != "" {
				if album.ID != atoi(id) {
					match = false
				}
			}
		}

		if match {
			result = append(result, album)
		}
	}

	return result
}

// FilterPhotos filters photos based on query parameters
func FilterPhotos(filters map[string]string) []models.Photo {
	result := make([]models.Photo, 0)

	for _, photo := range Photos {
		match := true

		if albumID, ok := filters["albumId"]; ok {
			if albumID != "" {
				if photo.AlbumID != atoi(albumID) {
					match = false
				}
			}
		}

		if id, ok := filters["id"]; ok {
			if id != "" {
				if photo.ID != atoi(id) {
					match = false
				}
			}
		}

		if match {
			result = append(result, photo)
		}
	}

	return result
}

// FilterTodos filters todos based on query parameters
func FilterTodos(filters map[string]string) []models.Todo {
	result := make([]models.Todo, 0)

	for _, todo := range Todos {
		match := true

		if userID, ok := filters["userId"]; ok {
			if userID != "" {
				if todo.UserID != atoi(userID) {
					match = false
				}
			}
		}

		if id, ok := filters["id"]; ok {
			if id != "" {
				if todo.ID != atoi(id) {
					match = false
				}
			}
		}

		if completed, ok := filters["completed"]; ok {
			if completed != "" {
				completedBool := completed == "true"
				if todo.Completed != completedBool {
					match = false
				}
			}
		}

		if match {
			result = append(result, todo)
		}
	}

	return result
}

// FilterUsers filters users based on query parameters
func FilterUsers(filters map[string]string) []models.User {
	result := make([]models.User, 0)

	for _, user := range Users {
		match := true

		if id, ok := filters["id"]; ok {
			if id != "" {
				if user.ID != atoi(id) {
					match = false
				}
			}
		}

		if match {
			result = append(result, user)
		}
	}

	return result
}

// atoi converts string to int, returns 0 if conversion fails
func atoi(s string) int {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result
}

// NextIDs tracks the next available ID for each resource type
var NextIDs = map[string]int{
	"posts":    6,
	"comments": 6,
	"albums":   6,
	"photos":   6,
	"todos":    6,
	"users":    6,
}

// GetNextID returns the next available ID for a resource type
func GetNextID(resource string) int {
	return NextIDs[resource]
}
