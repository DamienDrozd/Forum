package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	password := r.FormValue("user_password")
	email := r.FormValue("user_mail")

	fmt.Println(password, email)
	//---------------------------On vérififie si l'adresse email et le mdp sont dans la base de donnée et on connecte l'utilisateur--------------------

	var user User

	user = testLogin(email, password)

	if user.ID != 0 {

		// ID
		// Username
		// Email
		// Avatar

		// expiration := time.Now().Add(365 * 24 * time.Hour)

		cookie := http.Cookie{Name: "ID", Value: strconv.Itoa(user.ID)} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Username", Value: user.Username} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Email", Value: user.Email} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Avatar", Value: user.Avatar} //	, Expires: expiration}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/post", http.StatusSeeOther)

	}

	user.Password = password
	user.Email = email

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/login.html"))
	err := templates.ExecuteTemplate(w, "login", nil)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

func testLogin(email, password string) User {

	var user User

	return user
}

func register(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("psw")

	var user User

	user.Username = pseudo
	user.Email = email
	user.Password = password

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

// type Comment struct {
// 	Message  string
// 	UserName string
// 	Likes    int
// 	Dislikes int
// 	Date     string
// 	Avatar   string
// }

// type OutputPost struct {
// 	TabComment      []Comment
// 	PostName        string
// 	Categories      []string
// 	PostDate        string
// 	UserName        string
// 	PostDescription string
// 	Avatar          string
// 	Likes           int
// 	Dislikes        int
// }

func addComment(postName string) {

}

func readComment(postName string) []Comment {

	//----------------------------------------------------Provisoire---------------------------------------

	var Comment1 Comment

	Comment1.Message = "Message test"
	Comment1.UserName = "toto"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment1.Likes = 5
	Comment1.Dislikes = 2
	Comment1.Date = time.Now().Format("2006-01-02 15:04:05")

	var Comment2 Comment

	Comment2.Message = "Message test 2"
	Comment2.UserName = "toto2"
	Comment1.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	Comment2.Likes = 5
	Comment2.Dislikes = 2
	Comment2.Date = time.Now().Format("2006-01-02 15:04:05")

	return []Comment{Comment1, Comment2}

}

func post(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()
	var UserName string
	var Avatar string
	var ID string
	var postName string
	postName = "test"

	for _, cookie := range r.Cookies() {
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

	inputcomment := r.FormValue("comment")

	fmt.Println(inputcomment)

	var Comment Comment

	if inputcomment != "" && ID != "" {
		Comment.Message = inputcomment
		Comment.UserName = UserName
		Comment.Avatar = Avatar
		Comment.Likes = 0
		Comment.Dislikes = 0
		Comment.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	addComment(postName)

	CommentTab := readComment(postName)

	var output OutputPost

	output.PostName = postName
	output.Categories = []string{"categorie1", "categorie2"}
	output.TabComment = CommentTab

	//------------------------------------Provisoire-------------------

	output.PostName = "ceci est un post"
	output.Categories = []string{"categorie1", "categorie2"}
	output.PostDate = time.Now()
	output.UserName = "toto0"
	output.PostDescription = "Ici on peut y écrire la description de ce post"
	output.Avatar = "https://tse3.mm.bing.net/th?id=OIP.vzUhlFJFR5akQnwy8tWSvAHaF7&pid=Api"
	output.Likes = 50
	output.Dislikes = 3

	//--------------------------------------------------------------------

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

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Avatar   string
}

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	UserName string
	Message  string
	Likes    int
	Dislikes int
	Date     string
	Avatar   string
}

type OutputPost struct {
	TabComment      []Comment
	ID              int
	UserID          int
	title           string
	PostName        string
	Categories      []string
	PostDate        time.Time
	UserName        string
	PostDescription string
	Avatar          string
	Likes           int
	Dislikes        int
}
