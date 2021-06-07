package main

import (
	"database/sql"
	"fmt"
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

//----------------------Lecture----------------------------

func readComment(postName string) []Comment {

	//----------------------------------------------------Provisoire---------------------------------------

	var Comment1 Comment

	Comment1.CommentMessage = "Message test"
	Comment1.UserName = "toto"
	Comment1.UserAvatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment1.CommentLikes = 5
	Comment1.CommentDislikes = 2
	Comment1.CommentDate = time.Now()

	var Comment2 Comment

	Comment2.CommentMessage = "Message test"
	Comment2.UserName = "toto"
	Comment2.UserAvatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment2.CommentLikes = 5
	Comment2.CommentDislikes = 2
	Comment2.CommentDate = time.Now()

	return []Comment{Comment1, Comment2}

}

func testLogin(email, password string) User {

	var user User

	return user
}

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
	CREATE TABLE IF NOT EXISTS comments (
		commentid			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		commentmessage    	TEXT NOT NULL,
		commentlikes      	INTEGER NOT NULL,
		commentdislikes   	INTEGER NOT NULL,
		commentdate       	DATETIME,
		commentdatestring 	TEXT NOT NULL,
		postid            	INTEGER NOT NULL,
		userid            	INTEGER NOT NULL,
		username          	INTEGER NOT NULL,
		useravatar        	INTEGER NOT NULL
		)`

const PostTab = `
	CREATE TABLE IF NOT EXISTS post (
		postid				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		postname			TEXT NOT NULL UNIQUE,
		postcategory		TEXT NOT NULL,
		postdate			DATETIME,
		postdatestring		TEXT NOT NULL,
		postdescription		TEXT NOT NULL,
		posturl				TEXT NOT NULL,
		postlikes			INTEGER NOT NULL,
		postdislikes		INTEGER NOT NULL,
		userid				INTEGER NOT NULL,
		username			TEXT NOT NULL,
		useravatar			TEXT NOT NULL
	)`

func createDB(tab string) error {

	// fmt.Println("test")
	stmt, err := db.Prepare(tab)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

//----------------------Lecture----------------------------

func ReadCommenttoDB(PostID int) []Comment {

	//----------------------------------------------------Provisoire---------------------------------------

	rows := selectAllFromTable(db, "comments")
	// fmt.Println(rows)

	var tab []Comment

	for rows.Next() {
		var u Comment
		err := rows.Scan(&u.CommentID, &u.CommentMessage, &u.CommentLikes, &u.CommentDislikes, &u.CommentDate, &u.CommentDateString, &u.PostID, &u.UserID, &u.UserName, &u.UserAvatar)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	var tabcomment []Comment

	for i := range tab {
		// fmt.Println(tab[i].PostID, PostID)
		if tab[i].PostID == PostID {
			tabcomment = append(tabcomment, tab[i])
		}
	}

	return tabcomment

	// return tab

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

func ReadPosttoDB() []Post {

	rows := selectAllFromTable(db, "post")
	// fmt.Println(rows)

	var tab []Post

	for rows.Next() {
		var u Post
		err := rows.Scan(&u.PostID, &u.PostName, &u.PostCategory, &u.PostDate, &u.PostDateString, &u.PostDescription, &u.PostURL, &u.PostLikes, &u.PostDislikes, &u.UserID, &u.UserName, &u.UserAvatar)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab

}

//----------------------écriture----------------------------

func InsertPosttoDB(newpost Post) error {

	// fmt.Println(newpost)

	add, err := db.Prepare("INSERT INTO post (postname, postcategory, postdate, postdatestring, postdescription, posturl, postlikes, postdislikes, userid, username, useravatar) VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,?)")
	defer add.Close()
	if err != nil {
		return err
	}

	fmt.Println("test", add)

	add.Exec(newpost.PostName, newpost.PostCategory, newpost.PostDate, newpost.PostDateString, newpost.PostDescription, newpost.PostURL, newpost.PostLikes, newpost.PostDislikes, newpost.UserID, newpost.UserName, newpost.UserAvatar)
	return nil

}

// !! A récupérer à partir de la requete http !!
func InsertUsertoDB(user User) error {

	user.Avatar = "https://i.redd.it/wellr7jjiv011.jpg"
	add, err := db.Prepare("INSERT INTO users (username, password, email, avatar) VALUES (?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(user.Username, user.Password, user.Email, user.Avatar)
	return nil
}

func InsertCommenttoDB(comment Comment) error {

	add, err := db.Prepare("INSERT INTO comments (commentmessage, commentlikes, commentdislikes, commentdate, commentdatestring, postid, userid, username, useravatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(comment.CommentMessage, comment.CommentLikes, comment.CommentDislikes, comment.CommentDate, comment.CommentDateString, comment.PostID, comment.UserID, comment.UserName, comment.UserAvatar)
	return nil

}

func AddLiketoPosttoDB(typeadd string, nb int, PostID int) {

	if typeadd == "likes" {

		stmt, _ := db.Prepare("update post set postlikes=? where postid=?")

		stmt.Exec(nb, PostID)
	}

	if typeadd == "dislikes" {

		stmt, _ := db.Prepare("update post set postdislikes=? where postid=?")

		stmt.Exec(nb, PostID)
	}

}

func AddLiketoCommenttoDB(typeadd string, nb int, CommentID int) {

}
