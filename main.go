package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------------Structs--------------------------------------------
--------------------------------------------------------------------------------------------*/

/*--------------------------------------------------------------------------------------------
------------------------------ Func Handler Index and MainPage -------------------------------
----------------------------------------------------------------------------------------------*/
// index.html
func indexHandler(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/index.html"))
	// err := templates.ExecuteTemplate(w, "home", nil)

	// if err != nil {
	// 	log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	// }
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

func login(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("user_name")
	email := r.FormValue("user_mail")

	fmt.Println(pseudo, email)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/login.html"))
	err := templates.ExecuteTemplate(w, "login", nil)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

func register(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("psw")

	fmt.Println(pseudo, email, password)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/register.html"))
	err := templates.ExecuteTemplate(w, "register", nil)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

type Comment struct {
	Message  string
	UserName string
	Likes    int
	Dislikes int
	Date     string
	Avatar   string
}

type OutputPost struct {
	TabComment      []Comment
	PostName        string
	Categories      []string
	PostDate        string
	UserName        string
	PostDescription string
	Avatar          string
	Likes           int
	Dislikes        int
}

func post(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	//----------------------------------------------------Provisoire---------------------------------------

	var Comment1 Comment

	Comment1.Message = "Message test"
	Comment1.UserName = "toto"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment1.Likes = 5
	Comment1.Dislikes = 2
	Comment1.Date = "26/05/2021-10:42"

	var Comment2 Comment

	Comment2.Message = "Message test 2"
	Comment2.UserName = "toto2"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment2.Likes = 5
	Comment2.Dislikes = 2
	Comment2.Date = "26/05/2021-11h43"

	var output OutputPost

	output.PostName = "ceci est un post"
	output.Categories = []string{"categorie1", "categorie2"}
	output.TabComment = []Comment{Comment1, Comment2}

	output.PostDate = "26/05/2021-9h30"
	output.UserName = "toto0"
	output.PostDescription = "Ici on peut y écrire la description de ce post"
	output.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	output.Likes = 50
	output.Dislikes = 3

	//----------------------------------------------------------------------------------------------------

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/post.html"))
	err := templates.ExecuteTemplate(w, "post", output)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

/*--------------------------------------------------------------------------------------------
--------------------------------------Main Func-----------------------------------------------
----------------------------------------------------------------------------------------------*/
func main() {
	// Serving templates files
	filesServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	// Index handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/post", post)

	fmt.Println("Server is starting...\n")
	fmt.Println("Go on http://localhost:8080/\n")
	fmt.Println("To shut down the server press CTRL + C")

	// Starting serveur
	http.ListenAndServe(":8080", nil)
}
