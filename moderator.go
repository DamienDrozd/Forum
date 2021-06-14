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
-------------------------------------- Moderator Page----------------------------------------------
----------------------------------------------------------------------------------------------*/

func moderator(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now() // Print the time of execution of the function

	var user User

	for _, cookie := range r.Cookies() {
		// Read the cookies of the user if the user is logged
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

	if user.Role != "admin" && user.Role != "modo" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// If the user isn't a moderator, redirect the user in the page index
	}

	uploadpost, _ := strconv.Atoi(r.FormValue("uploadpost"))
	deletepost, _ := strconv.Atoi(r.FormValue("deletepost"))

	if uploadpost != 0 {
		ValidatePosttoDB(uploadpost) // Validate a post from db
	}
	if deletepost != 0 {
		DeletePosttoDB(deletepost) // delete a post from db
	}

	var output Out
	tablist := ReadPosttoDB() // List of all post from db

	x := len(tablist)
	for i := 0; i < x; i++ {
		if tablist[i].Validated == "true" { // Remove validated posts from postlist
			tablist = append(tablist[:i], tablist[i+1:]...)
			x = len(tablist)
			i = -1
		}
	}

	output.TabList = tablist

	//--------------------------------------------------------------------

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/moderator.html"))
	err := templates.ExecuteTemplate(w, "moderator", output)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}
