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
-------------------------------------- Index Page -------------------------------------------
----------------------------------------------------------------------------------------------*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now() // Print the time of execution of the function

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
	}

	var out Out

	tablist := ReadPosttoDB() // Read all posts from the db

	postmap := make(map[string]int)

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

	for i := range tablist {

		if postmap[tablist[i].PostCategory] == 0 {

			postmap[tablist[i].PostCategory] = 1

		} else {

			postmap[tablist[i].PostCategory] += 1
		}

	}

	tricategories := r.FormValue("Categories")
	tricreatedposts := r.FormValue("CreatedPosts")
	trilikedposts := r.FormValue("LikedPosts")

	//-----------------------------Tri of all posts---------------------

	if tricategories == "run" { // Tri by categories

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostCategory < tablist[j].PostCategory {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	} else if tricreatedposts == "run" { // Tri by creation date

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostDate.Format("2006-01-02 15:04:05") < tablist[j].PostDate.Format("2006-01-02 15:04:05") {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	} else if trilikedposts == "run" { // Tri by liked posts

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostLikes > tablist[j].PostLikes {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	} else {

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostName < tablist[j].PostName {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	}

	CategoryList := make([]Category, len(postmap))

	CategoryList = ReadCategorytoDB()

	for j := range CategoryList {
		for i := range postmap {
			if CategoryList[j].CategoryName == i {
				CategoryList[j].CategoryNumber = postmap[i]
				break
			}
		}
	}

	categorie := r.FormValue("Categorie")

	if categorie != "" {
		var tab []Post //Return the post list by categories
		for i := range tablist {
			if tablist[i].PostCategory == categorie {
				tab = append(tab, tablist[i])
			}
		}
		tablist = tab
	}

	x := len(tablist)
	for i := 0; i < x; i++ { // Remove the non-validated categories of the list
		tablist[i].TabComment = ReadCommenttoDB(tablist[i].PostID)
		tablist[i].NbComment = len(tablist[i].TabComment)
		if tablist[i].Validated == "false" {
			tablist = append(tablist[:i], tablist[i+1:]...)
			x = len(tablist)
			i = -1
		}
	}
	out.NbPost = len(tablist)
	out.CategoryList = CategoryList
	out.TabList = tablist
	out.User = user

	// fmt.Println(postmap)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/index.html"))
	err := templates.ExecuteTemplate(w, "index", out)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}
