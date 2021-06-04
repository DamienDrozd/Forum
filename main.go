package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ = sql.Open("sqlite3", "data.db")

	createDB(UserTab)
	createDB(CommentTab)
	createDB(PostTab)

	filesServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	// Index handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/post", post)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/newpost", newPost)

	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")

	// Starting serveur
	http.ListenAndServe(":8080", nil)
}
