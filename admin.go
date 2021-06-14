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
-------------------------------------- Admin Page----------------------------------------------
----------------------------------------------------------------------------------------------*/

func admin(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now() // Print the time of execution of the function

	var user User
	var output Out

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

	if user.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// If the user isn't a admin, redirect the user in the page index
	}

	TabUser := ReadUsertoDB()

	AddCategory := (r.FormValue("category")) // Add a new category for the forum
	if AddCategory != "" {
		var NewCategory Category
		NewCategory.CategoryName = AddCategory
		InsertCategorytoDB(NewCategory)

	}
	DeleteCategory, _ := strconv.Atoi((r.FormValue("deletecategory"))) // Delete a category

	if DeleteCategory != 0 {
		DeleteCategorytoDB(DeleteCategory)
	}

	Promote, _ := strconv.Atoi(r.FormValue("promote")) // Promote a user
	Demote, _ := strconv.Atoi(r.FormValue("demote"))   // Demote a user

	var userpromote User

	for i := range TabUser {

		if TabUser[i].ID == Promote || TabUser[i].ID == Demote {
			userpromote = TabUser[i]
		}
	}

	if Promote != 0 { // Promote a user
		PromoteUsertoDB(userpromote, "promote")
	}
	if Demote != 0 { // Demote a user
		PromoteUsertoDB(userpromote, "demote")
	}

	output.TabUser = ReadUsertoDB()
	output.CategoryList = ReadCategorytoDB()

	//--------------------------------------------------------------------

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/admin.html"))
	err := templates.ExecuteTemplate(w, "admin", output)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}
