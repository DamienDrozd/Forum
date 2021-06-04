package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//----------------------écriture----------------------------

func addComment(postName string, comment Comment) {

	// ID       int
	// UserID   int
	// PostID   int
	// UserName string
	// Message  string
	// Likes    int
	// Dislikes int
	// Date     string
	// Avatar   string

}

func addLogin(user User) {

	// user.Username
	// user.Email
	// user.Password

}

func addPost(newpost Post) {

	// UserID          int
	// UserName        string
	// UserAvatar      string
	// TabComment      []Comment
	// PostName        string
	// PostCategory    string
	// PostDate        time.Time
	// PostDateString  string
	// PostDescription string
	// PostLikes       int
	// PostDislikes    int

}

//----------------------Lecture----------------------------

func readComment(postName string) []Comment {

	//----------------------------------------------------Provisoire---------------------------------------

	var Comment1 Comment

	Comment1.Message = "Message test"
	Comment1.UserName = "toto"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment1.Likes = 5
	Comment1.Dislikes = 2
	Comment1.Date = time.Now().Format("2006-01-02 15:04:05")

	var Comment2 Comment

	Comment2.Message = "Message test 2"
	Comment2.UserName = "toto2"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment2.Likes = 5
	Comment2.Dislikes = 2
	Comment2.Date = time.Now().Format("2006-01-02 15:04:05")

	return []Comment{Comment1, Comment2}

}

func testLogin(email, password string) User {

	var user User

	return user
}

func postlist() []Post { //Get a listof all posts

	var post1 Post
	post1.PostName = "Post1"
	post1.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post1.UserName = "toto1"
	post1.PostDescription = "Description du post"
	post1.PostCategory = "categorie2"
	post1.PostDate = time.Now()
	post1.PostLikes = 15

	var post2 Post
	post2.PostName = "Post2"
	post2.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post2.UserName = "toto1"
	post2.PostDescription = "Description du post"
	post2.PostCategory = "categorie1"
	post2.PostDate = time.Now()
	post2.PostLikes = 15

	var post3 Post
	post3.PostName = "Post3"
	post3.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post3.UserName = "toto1"
	post3.PostDescription = "Description du post"
	post3.PostCategory = "categorie2"
	post3.PostDate = time.Now()
	post3.PostLikes = 15

	var post4 Post
	post4.PostName = "Post4"
	post4.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post4.UserName = "toto1"
	post4.PostDescription = "Description du post"
	post4.PostCategory = "categorie2"
	post4.PostDate = time.Now()
	post4.PostLikes = 15

	var post5 Post
	post5.PostName = "Post5"
	post5.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post5.UserName = "toto1"
	post5.PostDescription = "Description du post"
	post5.PostCategory = "categorie1"
	post5.PostDate = time.Now()
	post5.PostLikes = 15

	return []Post{post1, post2, post3, post4, post5}

}

// func FindUser(ID int) User {

// }

var db *sql.DB

const UserTab = `
	CREATE TABLE IF NOT EXISTS users (
		id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
		username	TEXT UNIQUE NOT NULL, 
		password	TEXT NOT NULL, 
		email		TEXT UNIQUE NOT NULL,
		avatar		TEXT
	)`

const CommentTab = `
	CREATE	TABLE IF NOT EXISTS comments (
		id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		user_id 	INTEGER,
		post_id		INTEGER,
		message		TEXT,
		date		DATETIME
		)`

const PostTab = `
		CREATE TABLE IF NOT EXISTS post(
			post_id		INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			user_id		INTEGER NOT NULL,
			title		TEXT NOT NULL,
			post_description		TEXT NOT NULL,
			date		DATETIME,
			image		TEXT
		)`

func createDB(tab string) error {
	stmt, err := db.Prepare(tab)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}

// !! A récupérer à partir de la requete http !!
func InsertUsertoDB(user User) error {

	user.Avatar = "test"
	add, err := db.Prepare("INSERT INTO users (username, password, email, avatar) VALUES (?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(user.Username, user.Password, user.Email, user.Avatar)
	return nil
}

func ReadUsertoDB() []User {

	rows := selectAllFromTable(db, "users")
	// fmt.Println(rows)

	var tab []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Avatar)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab

}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}
