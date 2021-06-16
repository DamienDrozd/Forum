package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/time/rate"
)

func main() {
	// open database and sqlite3 module
	db, _ = sql.Open("sqlite3", "data.db")

	createDB(PostTab)
	createDB(UserTab)
	createDB(CommentTab)
	createDB(CategoryTab)
	createDB(LikePostTab)
	createDB(LikeCommentTab)
	// Serving templates files
	mux := http.NewServeMux()
	filesServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", filesServer))

	// Index handler
	mux.HandleFunc("/", indexHandler)
	// Login Handler
	mux.HandleFunc("/login", login)
	// Register Handler
	mux.HandleFunc("/register", register)
	// Post Handler
	mux.HandleFunc("/post", post)
	// Index Handler
	mux.HandleFunc("/index", indexHandler)
	// Newpost Handler
	mux.HandleFunc("/newpost", newPost)
	// user Handler
	mux.HandleFunc("/user", user)
	// admin Handler
	mux.HandleFunc("/admin", admin)
	// moderator Handler
	mux.HandleFunc("/moderator", moderator)

	fmt.Println("Server is starting...")
	fmt.Print("\n")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Print("\n")
	fmt.Println("To shut down the server press CTRL + C")
	fmt.Print("\n")

	// Starting serveur
	http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", limit(mux))
}

var limiter = rate.NewLimiter(3, 7)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
