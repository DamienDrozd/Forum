package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------------Structs--------------------------------------------
--------------------------------------------------------------------------------------------*/

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
	Category        []string
	PostDate        time.Time
	UserName        string
	PostDescription string
	Avatar          string
	Likes           int
	Dislikes        int
}

type Out struct {
	TabList      []Post
	CategoryList []Category
}

type Post struct {
	UserID          int
	UserName        string
	UserAvatar      string
	TabComment      []Comment
	PostName        string
	PostCategory    string
	PostDate        time.Time
	PostDateString  string
	PostDescription string
	PostLikes       int
	PostDislikes    int
}

type Category struct {
	CategoryName   string
	CategoryNumber int
}

/*--------------------------------------------------------------------------------------------
------------------------------ Func Handler Index and MainPage -------------------------------
----------------------------------------------------------------------------------------------*/
// index.html

func newPost(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var newpost Post
	var ID string
	var Avatar string
	var UserName string

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

	newpost.UserID, _ = strconv.Atoi(ID)
	newpost.UserName = UserName
	newpost.UserAvatar = Avatar
	newpost.PostName = r.FormValue("Titre_sujet")
	newpost.PostCategory = r.FormValue("categorie")
	newpost.PostDescription = r.FormValue("message_newpost")
	newpost.PostDate = time.Now()
	newpost.PostDateString = newpost.PostDate.Format("2006-01-02 15:04:05")

	addPost(newpost)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/newpost.html"))
	err := templates.ExecuteTemplate(w, "newpost", nil)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var out Out

	tablist := postlist()

	postmap := make(map[string]int)

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

	if tricategories == "run" {

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostCategory < tablist[j].PostCategory {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	} else if tricreatedposts == "run" {

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostDate.Format("2006-01-02 15:04:05") < tablist[j].PostDate.Format("2006-01-02 15:04:05") {
					tablist[i], tablist[j] = tablist[j], tablist[i]
				}
			}
		}

	} else if trilikedposts == "run" {

		for i := range tablist {
			for j := range tablist {
				if tablist[i].PostLikes < tablist[j].PostLikes {
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
	x := 0

	for i := range postmap {
		CategoryList[x].CategoryName = i
		CategoryList[x].CategoryNumber = postmap[i]
		x++
	}

	out.CategoryList = CategoryList
	out.TabList = tablist

	fmt.Println(postmap)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/index.html"))
	err := templates.ExecuteTemplate(w, "index", out)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
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

func register(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("psw")

	user := User{
		Username: pseudo,
		Email:    email,
		Password: password,
	}

	fmt.Println(pseudo, email, password)
	InsertIntoDB(user)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/register.html"))
	err := templates.ExecuteTemplate(w, "register", nil)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

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

		addComment(postName, Comment)
	}

	CommentTab := readComment(postName)

	var output OutputPost

	output.PostName = postName
	output.Category = []string{"categorie1", "categorie2"}
	output.TabComment = CommentTab

	//------------------------------------Provisoire-------------------

	output.PostName = "ceci est un post"
	output.Category = []string{"categorie1", "categorie2"}
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
	db, _ = sql.Open("sqlite3", "data.db")

	createDB(UserTab)
	createDB(CommentTab)
	createDB(PostTab)
	// test User Tab
	// user1 := User{
	// 	Username: "Nick",
	// 	Password: "Hello",
	// 	Email:    "azerty@hotmail.fr",
	// }
	// InsertIntoDB(user1)
	// Serving templates files
	filesServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesServer))

	// Index handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/post", post)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/newpost", newPost)

	fmt.Println("Server is starting...\n")
	fmt.Println("Go on http://localhost:8080/\n")
	fmt.Println("To shut down the server press CTRL + C")

	// Starting serveur
	http.ListenAndServe(":8080", nil)
}

/*--------------------------------------------------------------------------------------------
-------------------------------------------A finir--------------------------------------------
--------------------------------------------------------------------------------------------*/

//----------------------écriture----------------------------

func addComment(postName string, comment Comment) {

	// ID       int
	// UserID   int
	// PostID   int
	// UserName string
	// Message  string
	// Likes    int
	// Dislikes int
	// Date     string
	// Avatar   string

}

func addLogin(user User) {

	// user.Username
	// user.Email
	// user.Password

}

func addPost(newpost Post) {

	// UserID          int
	// UserName        string
	// UserAvatar      string
	// TabComment      []Comment
	// PostName        string
	// PostCategory    string
	// PostDate        time.Time
	// PostDateString  string
	// PostDescription string
	// PostLikes       int
	// PostDislikes    int

}

//----------------------Lecture----------------------------

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

func testLogin(email, password string) User {

	var user User

	return user
}

func postlist() []Post { //Get a listof all posts

	var post1 Post
	post1.PostName = "Post1"
	post1.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post1.UserName = "toto1"
	post1.PostDescription = "Description du post"
	post1.PostCategory = "categorie2"
	post1.PostDate = time.Now()
	post1.PostLikes = 15

	var post2 Post
	post2.PostName = "Post2"
	post2.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post2.UserName = "toto1"
	post2.PostDescription = "Description du post"
	post2.PostCategory = "categorie1"
	post2.PostDate = time.Now()
	post2.PostLikes = 15

	var post3 Post
	post3.PostName = "Post3"
	post3.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post3.UserName = "toto1"
	post3.PostDescription = "Description du post"
	post3.PostCategory = "categorie2"
	post3.PostDate = time.Now()
	post3.PostLikes = 15

	var post4 Post
	post4.PostName = "Post4"
	post4.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post4.UserName = "toto1"
	post4.PostDescription = "Description du post"
	post4.PostCategory = "categorie2"
	post4.PostDate = time.Now()
	post4.PostLikes = 15

	var post5 Post
	post5.PostName = "Post5"
	post5.UserAvatar = "https://tse4.mm.bing.net/th?id=OIP.YdkNhFNLUQ_NN3gZir70pQHaHZ&pid=Api"
	post5.UserName = "toto1"
	post5.PostDescription = "Description du post"
	post5.PostCategory = "categorie1"
	post5.PostDate = time.Now()
	post5.PostLikes = 15

	return []Post{post1, post2, post3, post4, post5}

}

// func FindUser(ID int) User {

// }

var db *sql.DB

const UserTab = `
	CREATE TABLE IF NOT EXISTS users (
		id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
		username	TEXT UNIQUE NOT NULL, 
		password	TEXT NOT NULL, 
		email		TEXT UNIQUE NOT NULL,
		avatar		TEXT
	)`

const CommentTab = `
	CREATE	TABLE IF NOT EXISTS comments (
		id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		user_id 	INTEGER,
		post_id		INTEGER,
		message		TEXT,
		date		DATETIME
		)`

const PostTab = `
		CREATE TABLE IF NOT EXISTS post(
			post_id		INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			user_id		INTEGER NOT NULL,
			title		TEXT NOT NULL,
			post_description		TEXT NOT NULL,
			date		DATETIME,
			image		TEXT
		)`

func createDB(tab string) error {
	stmt, err := db.Prepare(tab)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}

// !! A récupérer à partir de la requete http !!
func InsertIntoDB(user User) error {
	add, err := db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
	defer add.Close()
	if err != nil {
		return err
	}

	add.Exec(user.Username, user.Password, user.Email)
	return nil
}
