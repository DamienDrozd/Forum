package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const UserTab = `
	CREATE TABLE IF NOT EXISTS users (
		id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
		username	TEXT UNIQUE NOT NULL, 
		password	TEXT NOT NULL, 
		email		TEXT UNIQUE NOT NULL,
		avatar		TEXT,
		role		TEXT
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
		useravatar			TEXT NOT NULL,
		validated 			TEXT NOT NULL
	)`
const CategoryTab = `
	CREATE TABLE IF NOT EXISTS category (
		categoryid				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		categoryname			TEXT NOT NULL UNIQUE
	)`

const LikePostTab = `	
	CREATE TABLE IF NOT EXISTS postlikes (
		postid			INTEGER NOT NULL,
		userid			INTEGER NOT NULL,
		type			TEXT	NOT NULL
		)`

const LikeCommentTab = `
	CREATE TABLE IF NOT EXISTS commentlikes (
		commentid		INTEGER NOT NULL,
		userid			INTEGER NOT NULL,
		type			TEXT	NOT NULL
)`

func createDB(tab string) error { // Create a database if the database don't exist

	stmt, err := db.Prepare(tab)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {
	// Select all elements from a table and return them
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

//----------------------Lecture----------------------------

func ReadCommenttoDB(PostID int) []Comment {

	//return the list of all comment in the forum stored in the database
	rows := selectAllFromTable(db, "comments")

	var tab []Comment

	for rows.Next() {
		var u Comment // the comments are stored in u
		err := rows.Scan(&u.CommentID, &u.CommentMessage, &u.CommentLikes, &u.CommentDislikes, &u.CommentDate, &u.CommentDateString, &u.PostID, &u.UserID, &u.UserName, &u.UserAvatar)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u) // Append in tab eatch comment
	}

	var tabcomment []Comment

	for i := range tab {
		if tab[i].PostID == PostID {
			tabcomment = append(tabcomment, tab[i]) // put in tab comment eatch tab linked to the PostID
		}
	}

	return tabcomment

}

func ReadUsertoDB() []User { //return the list of all users in the forum stored in the database
	rows := selectAllFromTable(db, "users")
	// fmt.Println(rows)

	var tab []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Avatar, &u.Role)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab

}

func ReadPosttoDB() []Post { //return the list of all Posts in the forum stored in the database

	rows := selectAllFromTable(db, "post")
	// fmt.Println(rows)

	var tab []Post

	for rows.Next() {
		var u Post
		err := rows.Scan(&u.PostID, &u.PostName, &u.PostCategory, &u.PostDate, &u.PostDateString, &u.PostDescription, &u.PostURL, &u.PostLikes, &u.PostDislikes, &u.UserID, &u.UserName, &u.UserAvatar, &u.Validated)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab

}

func ReadCategorytoDB() []Category { //return the list of all categories in the forum stored in the database

	rows := selectAllFromTable(db, "category")
	// fmt.Println(rows)

	var tab []Category

	for rows.Next() {
		var u Category
		err := rows.Scan(&u.CategoryID, &u.CategoryName)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab
}

//----------------------Insert an element in the tab----------------------------

func InsertCommenttoDB(comment Comment) error { // Insert a comment in the database

	add, err := db.Prepare("INSERT INTO comments (commentmessage, commentlikes, commentdislikes, commentdate, commentdatestring, postid, userid, username, useravatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	//add the elements from comment in the database

	add.Exec(comment.CommentMessage, comment.CommentLikes, comment.CommentDislikes, comment.CommentDate, comment.CommentDateString, comment.PostID, comment.UserID, comment.UserName, comment.UserAvatar)
	return nil

}

func InsertPosttoDB(newpost Post) error { // Insert a post in the database

	// fmt.Println(newpost)

	add, err := db.Prepare("INSERT INTO post (postname, postcategory, postdate, postdatestring, postdescription, posturl, postlikes, postdislikes, userid, username, useravatar, validated) VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,?)")
	defer add.Close()
	if err != nil {
		return err
	}

	fmt.Println("test", add)

	add.Exec(newpost.PostName, newpost.PostCategory, newpost.PostDate, newpost.PostDateString, newpost.PostDescription, newpost.PostURL, newpost.PostLikes, newpost.PostDislikes, newpost.UserID, newpost.UserName, newpost.UserAvatar, newpost.Validated)
	return nil

}

func InsertUsertoDB(user User) error { // Insert a user in the database

	user.Avatar = "https://i.redd.it/wellr7jjiv011.jpg"
	add, err := db.Prepare("INSERT INTO users (username, password, email, avatar, role) VALUES (?, ?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	user.Password = HashPassword(user.Password)

	add.Exec(user.Username, user.Password, user.Email, user.Avatar, user.Role)
	return nil
}

func InsertCategorytoDB(category Category) error { // Insert a category in the database

	add, err := db.Prepare("INSERT INTO category (categoryname) VALUES (?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(category.CategoryName)
	return nil
}

//--------------------------------Edit Tab in database------------------------------------------------

func AddLiketoPosttoDB(typeadd string, nb int, PostID int) { // Add a like to a post in the db

	if typeadd == "likes" {

		stmt, _ := db.Prepare("update post set postlikes=? where postid=?")

		stmt.Exec(nb, PostID)
	}

	if typeadd == "dislikes" {

		stmt, _ := db.Prepare("update post set postdislikes=? where postid=?")

		stmt.Exec(nb, PostID)
	}

}

func AddLiketoCommenttoDB(typeadd string, nb int, CommentID int) { // Add a like to a comment in the db

	if typeadd == "likes" {

		stmt, _ := db.Prepare("update comments set commentlikes=? where commentid=?")

		stmt.Exec(nb, CommentID)
	}

	if typeadd == "dislikes" {

		stmt, _ := db.Prepare("update comments set commentdislikes=? where commentid=?")

		stmt.Exec(nb, CommentID)
	}

}

func PromoteUsertoDB(user User, typeadd string) { // Promote a user in the db

	if typeadd == "promote" {

		var newrole string
		stmt, _ := db.Prepare("update users set role=? where id=?")

		if user.Role == "user" {
			newrole = "modo"
		}
		if user.Role == "modo" {
			newrole = "admin"
		}
		if user.Role == "admin" {
			newrole = "admin"
		}
		// fmt.Println(stmt, user.Role, newrole)
		stmt.Exec(newrole, user.ID)
	}

	if typeadd == "demote" {

		var newrole string
		stmt, _ := db.Prepare("update users set role=? where id=?")

		if user.Role == "modo" {
			newrole = "user"
		}
		if user.Role == "admin" {
			newrole = "modo"
		}
		if user.Role == "user" {
			newrole = "user"
		}
		stmt.Exec(newrole, user.ID)
	}

}

func ValidatePosttoDB(PostID int) { // Validate a post in the db
	stmt, _ := db.Prepare("update post set validated=? where postid=?")

	stmt.Exec("true", PostID)
}

//------------------------------Delete Element to database-----------------------------------------------

func DeletePosttoDB(PostID int) { // Delete a post in the db
	TabComment := ReadCommenttoDB(PostID)

	for i := range TabComment {
		DeleteCommenttoDB(TabComment[i].CommentID)
	}

	stmt, _ := db.Prepare("DELETE FROM post WHERE postid=?;")

	stmt.Exec(PostID)

}

func DeleteCommenttoDB(CommentID int) { // Delete a comment in the DB

	stmt, _ := db.Prepare("DELETE FROM comments WHERE commentid=?;")

	stmt.Exec(CommentID)

}

func DeleteCategorytoDB(CategoryID int) { // Delete a category in the db

	stmt, _ := db.Prepare("DELETE FROM category WHERE categoryid=?;")

	stmt.Exec(CategoryID)

}

// func InsertPostLiketoDB(like PostLike) error {
// 	add, err := db.Prepare("INSERT INTO postlikes (postid, userid, type) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}

// 	add.Exec(like.PostID, like.UserID, like.Type)
// 	add.Close()
// 	return nil
// }

// func postslikes() []PostLike {
// 	rows := selectAllFromTable(db, "postlikes")
// 	var result []PostLike
// 	for rows.Next() {
// 		var post PostLike
// 		err := rows.Scan(&post.PostID, &post.UserID, &post.Type)
// 		if err != nil {
// 			log.fatal(err)
// 		}
// 		result = append(result, post)
// 	}
// 	return result
// }
