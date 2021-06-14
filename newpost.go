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
-------------------------------------- NewPost Page -------------------------------------------
----------------------------------------------------------------------------------------------*/
func newPost(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now() // Print the time of execution of the function

	var newpost Post
	var ID string
	var Avatar string
	var UserName string

	for _, cookie := range r.Cookies() { // Read the cookies of the user
		if cookie.Name == "Username" {
			UserName = cookie.Value
		}
		if cookie.Name == "Avatar" {
			Avatar = cookie.Value
		}
		if cookie.Name == "ID" {
			ID = cookie.Value
		}
	}
	var user User
	user.Username = UserName
	user.Avatar = Avatar
	user.ID, _ = strconv.Atoi(ID)

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

	erroutput := ""

	if UserName == "" || ID == "" { // The user must be connected to create a post
		erroutput += "Vous devez être connecté pour ajouter un post"
	}

	// file, fileHeader, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println(file)
	// fmt.Println(fileHeader)

	newpost.PostName = r.FormValue("Titre_sujet")
	newpost.PostCategory = r.FormValue("categorie")
	newpost.PostDate = time.Now()
	newpost.PostDateString = newpost.PostDate.Format("2006-01-02 15:04:05")
	newpost.PostDescription = r.FormValue("message_newpost")
	newpost.PostURL = "/post?name=" + newpost.PostName
	newpost.UserID, _ = strconv.Atoi(ID)
	newpost.UserName = UserName
	newpost.UserAvatar = Avatar

	tab := ReadPosttoDB()

	for i := range tab {
		if tab[i].PostName == newpost.PostName { // Test if the post not already exist
			erroutput += "Un post avec le même nom existe déja"
			break
		}
	}

	if erroutput == "" && newpost.PostName != "" && newpost.PostCategory != "" && newpost.PostDescription != "" {

		err1 := InsertPosttoDB(newpost) // Insert to post in the database

		if err1 != nil {
			log.Fatalf("DataBase execution: %s", err1)
		}

	}

	if r.Method != "POST" {
		erroutput = ""
	}

	var output Error
	output.Error = erroutput
	output.User = user
	output.CategoryList = ReadCategorytoDB()

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/newpost.html"))
	erro := templates.ExecuteTemplate(w, "newpost", output)

	if erro != nil {
		log.Fatalf("Template execution: %s", erro) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))
}
