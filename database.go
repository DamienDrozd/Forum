package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
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

const ImageTab = `
	CREATE TABLE IF NOT EXISTS image (
		imageid				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		image				BLOB NOT NULL, 
		postid				INTEGER NOT NULL
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

func SendMail(MailType string, comment Comment, post Post) {
	if MailType == "NewLike" {

	}
	if MailType == "NewDislike" {

	}
	if MailType == "NewComment" {

	}

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
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Avatar, &u.Role)
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
		err := rows.Scan(&u.PostID, &u.PostName, &u.PostCategory, &u.PostDate, &u.PostDateString, &u.PostDescription, &u.PostURL, &u.PostLikes, &u.PostDislikes, &u.UserID, &u.UserName, &u.UserAvatar, &u.Validated)
		if err != nil {
			log.Fatal(err)
		}

		tab = append(tab, u)
	}

	return tab

}

func ReadCategorytoDB() []Category {

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

//----------------------écriture----------------------------

func InsertPosttoDB(newpost Post) error {

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

// !! A récupérer à partir de la requete http !!
func InsertUsertoDB(user User) error {

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

func InsertCommenttoDB(comment Comment) error {

	add, err := db.Prepare("INSERT INTO comments (commentmessage, commentlikes, commentdislikes, commentdate, commentdatestring, postid, userid, username, useravatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(comment.CommentMessage, comment.CommentLikes, comment.CommentDislikes, comment.CommentDate, comment.CommentDateString, comment.PostID, comment.UserID, comment.UserName, comment.UserAvatar)
	return nil

}

func InsertCategorytoDB(category Category) error {

	add, err := db.Prepare("INSERT INTO category (categoryname) VALUES (?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(category.CategoryName)
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

	if typeadd == "likes" {

		stmt, _ := db.Prepare("update comments set commentlikes=? where commentid=?")

		stmt.Exec(nb, CommentID)
	}

	if typeadd == "dislikes" {

		stmt, _ := db.Prepare("update comments set commentdislikes=? where commentid=?")

		stmt.Exec(nb, CommentID)
	}

}

func DeletePosttoDB(PostID int) {
	TabComment := ReadCommenttoDB(PostID)

	for i := range TabComment {
		DeleteCommenttoDB(TabComment[i].CommentID)
	}

	stmt, _ := db.Prepare("DELETE FROM post WHERE postid=?;")

	stmt.Exec(PostID)

}

func DeleteCommenttoDB(CommentID int) {

	stmt, _ := db.Prepare("DELETE FROM comments WHERE commentid=?;")

	stmt.Exec(CommentID)

}

func DeleteCategorytoDB(CategoryID int) {

	stmt, _ := db.Prepare("DELETE FROM category WHERE categoryid=?;")

	stmt.Exec(CategoryID)

}

func PromoteUsertoDB(user User, typeadd string) {

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

func ValidatePosttoDB(PostID int) {
	stmt, _ := db.Prepare("update post set validated=? where postid=?")

	stmt.Exec("true", PostID)
}

func InsertImagetoDB(image Image) error {

	add, err := db.Prepare("INSERT INTO image (image, postid) VALUES (?, ?)")
	defer add.Close()
	if err != nil {
		fmt.Println(err)
	}

	out, _ := ioutil.TempFile("temp-images", "upload-*.png")
	defer out.Close()
	io.Copy(out, image.Image)

	fmt.Println(out)
	fmt.Println(image.Image)

	add.Exec(out, image.PostID)
	return nil
}
