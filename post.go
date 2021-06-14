package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------- Post Page----------------------------------------------
----------------------------------------------------------------------------------------------*/
func post(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()
	var UserName string
	var Avatar string
	var ID string
	var user User

	for _, cookie := range r.Cookies() { // Read the cookies of the user to log in the user
		if cookie.Name == "Username" {
			UserName = cookie.Value
			user.Username = cookie.Value
		}
		if cookie.Name == "Avatar" {
			Avatar = cookie.Value
			user.Avatar = cookie.Value
		}
		if cookie.Name == "ID" {
			ID = cookie.Value
			user.ID, _ = strconv.Atoi(cookie.Value)
		}
		if cookie.Name == "Email" {
			user.Email = cookie.Value
		}
		if cookie.Name == "Role" {
			user.Role = cookie.Value
		}
	}

	deco := r.FormValue("Deconnexion")

	if deco == "run" { // Deconnect the user by deleting all cookies
		c, _ := r.Cookie("Username")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		c, _ = r.Cookie("Email")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		c, _ = r.Cookie("Avatar")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		c, _ = r.Cookie("ID")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		c, _ = r.Cookie("Role")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		user = User{}
	}

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	var postName string

	postName = keys[0]

	inputcomment := r.FormValue("comment")

	var output Post

	tablist := ReadPosttoDB()

	for i := range tablist {
		if tablist[i].PostName == postName {
			output = tablist[i]
		}
	}

	SupprimerPost := r.FormValue("SuprimerPost")

	if user.Role == "modo" || user.Role == "admin" {
		// Delete a post by a moderator or an admin
		if SupprimerPost != "" {
			DeletePosttoDB(output.PostID)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	SupprimerComment, _ := strconv.Atoi(r.FormValue("SuprimerCommentaire"))

	if user.Role == "modo" || user.Role == "admin" {
		// Delete a comment by a moderator or an admin
		if SupprimerComment != 0 {
			DeleteCommenttoDB(SupprimerComment)
		}
	}

	var comment Comment

	if inputcomment != "" && ID != "" { // Add a comment to a post
		comment.CommentMessage = inputcomment
		comment.CommentLikes = 0
		comment.CommentDislikes = 0
		comment.CommentDate = time.Now()
		comment.CommentDateString = comment.CommentDate.Format("2006-01-02 15:04:05")
		comment.PostID = output.PostID
		comment.UserID, _ = strconv.Atoi(ID)
		comment.UserName = UserName
		comment.UserAvatar = Avatar
		InsertCommenttoDB(comment)

	}

	output.TabComment = ReadCommenttoDB(output.PostID)

	if r.FormValue("likes") == "run" { // Add a like to a post
		AddLiketoPosttoDB("likes", output.PostLikes+1, output.PostID)
		output.PostLikes += 1

	}
	if r.FormValue("dislikes") == "run" { // Add a dislike to a post
		AddLiketoPosttoDB("dislikes", output.PostDislikes+1, output.PostID)
		output.PostDislikes += 1

	}

	if r.FormValue("CommentLikes") != "" { // Add a like to a comment
		ID, _ := strconv.Atoi(r.FormValue("CommentLikes"))
		for i := range output.TabComment {
			if ID == output.TabComment[i].CommentID {
				AddLiketoCommenttoDB("likes", output.TabComment[i].CommentLikes+1, ID)
				output.TabComment[i].CommentLikes += 1
			}

		}

	}
	if r.FormValue("CommentDislikes") != "" { // Add a dislike to a comment
		ID, _ := strconv.Atoi(r.FormValue("CommentDislikes"))
		for i := range output.TabComment {
			if ID == output.TabComment[i].CommentID {
				AddLiketoCommenttoDB("dislikes", output.TabComment[i].CommentDislikes+1, ID)
				output.TabComment[i].CommentDislikes += 1
			}

		}
	}

	var Error Error
	Error.User = user
	Error.Post = output

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/post.html"))
	err := templates.ExecuteTemplate(w, "post", Error)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}
