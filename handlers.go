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

/*--------------------------------------------------------------------------------------------
-------------------------------------- Index Page -------------------------------------------
----------------------------------------------------------------------------------------------*/
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

/*--------------------------------------------------------------------------------------------
--------------------------------------Login Page ---------------------------------------------
----------------------------------------------------------------------------------------------*/
func login(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	password := r.FormValue("user_password")
	email := r.FormValue("user_mail")

	fmt.Println(password, email)
	//---------------------------On vérififie si l'adresse email et le mdp sont dans la base de donnée et on connecte l'utilisateur--------------------

	user := User{}

	tab := ReadUsertoDB()

	for i := range tab {

		fmt.Println(tab[i].Email, email)
		fmt.Println(tab[i].Password, password)

		if tab[i].Email == email && tab[i].Password == password {

			user = tab[i]
		}

	}

	fmt.Println(user)

	output := ""

	if user.ID != 0 {

		cookie := http.Cookie{Name: "ID", Value: strconv.Itoa(user.ID)} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Username", Value: user.Username} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Email", Value: user.Email} //, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Avatar", Value: user.Avatar} //	, Expires: expiration}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		output += "L'email ou le mot de passe n'existe pas"
	}

	var erroutput Error
	erroutput.Error = output

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/login.html"))
	err := templates.ExecuteTemplate(w, "login", erroutput)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

/*--------------------------------------------------------------------------------------------
-------------------------------------- Register Page------------------------------------------
----------------------------------------------------------------------------------------------*/
func register(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("psw")
	password_repeat := r.FormValue("psw_repeat")

	output := ""

	user := User{
		Username: pseudo,
		Email:    email,
		Password: password,
	}

	tab := ReadUsertoDB()

	if password_repeat != password {
		output += "Erreur, les mots de passe ne correspondent pas\n"
	}

	if len(password) == 0 || len(pseudo) == 0 || len(email) == 0 {

		output += "Erreur, vous devez remplir toutes les cases.\n"

	}

	for i := range tab {
		if tab[i].Username == user.Username {
			output += "Erreur, ce pseudo est déja pris\n"
		}
		if tab[i].Email == user.Email {
			output += "Erreur, cet email est déja pris\n"
		}

	}

	if output == "" {
		InsertUsertoDB(user)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var erroutput Error
	erroutput.Error = output

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/register.html"))
	err := templates.ExecuteTemplate(w, "register", erroutput)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

/*--------------------------------------------------------------------------------------------
-------------------------------------- Post Page----------------------------------------------
----------------------------------------------------------------------------------------------*/
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
