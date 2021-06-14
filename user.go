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
-------------------------------------- User Page----------------------------------------------
----------------------------------------------------------------------------------------------*/
func user(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var user User

	for _, cookie := range r.Cookies() { // Read the cookies of the user
		if cookie.Name == "Username" {
			user.Username = cookie.Value
		}
		if cookie.Name == "Avatar" {
			user.Avatar = cookie.Value
		}
		if cookie.Name == "ID" {
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

	if user.ID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// If the user isn't logged, redirect the user in the page index
	}

	UserList := ReadUsertoDB()
	PostList := ReadPosttoDB()

	for i := range UserList {
		if UserList[i].ID == user.ID {
			user = UserList[i] // get the informations from the user
		}
	}

	for i := range PostList { // Get the posts from the username
		if PostList[i].UserID == user.ID {
			PostList[i].TabComment = ReadCommenttoDB(PostList[i].PostID)
			user.PostList = append(user.PostList, PostList[i])
		}

	}

	DeletePost, _ := strconv.Atoi(r.FormValue("delete_post"))
	DeleteComment, _ := strconv.Atoi(r.FormValue("delete_comment"))

	if DeletePost != 0 {
		DeletePosttoDB(DeletePost) // Delete a post from the user
		for i := range user.PostList {
			if user.PostList[i].PostID == DeletePost {

				user.PostList = append(user.PostList[:i], user.PostList[i+1:]...)
				break

			}
		}
	}
	if DeleteComment != 0 { // Delete a comment of a user post
		DeleteCommenttoDB(DeleteComment)
		for i := range user.PostList {
			for j := range user.PostList[i].TabComment {
				if user.PostList[i].TabComment[j].CommentID == DeleteComment {
					user.PostList[i].TabComment = append(user.PostList[i].TabComment[:j], user.PostList[i].TabComment[j+1:]...)
					break
				}
			}
		}
	}

	//--------------------------------------------------------------------

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/account.html"))
	err := templates.ExecuteTemplate(w, "user", user)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}
