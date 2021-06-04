package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// open database and sqlite3 module
	db, _ = sql.Open("sqlite3", "data.db")

	createDB(UserTab)
	createDB(CommentTab)
	createDB(PostTab)

	// Serving templates files
	filesServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	// Index handler
	http.HandleFunc("/", indexHandler)
	// Login Handler
	http.HandleFunc("/login", login)
	// Register Handler
	http.HandleFunc("/register", register)
	// Post Handler
	http.HandleFunc("/post", post)
	// Index Handler
	http.HandleFunc("/index", indexHandler)
	// Newpost Handler
	http.HandleFunc("/newpost", newPost)

	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")

	// Starting serveur
	http.ListenAndServe(":8080", nil)
}
