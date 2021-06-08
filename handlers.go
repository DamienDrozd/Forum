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

	erroutput := ""

	if UserName == "" || ID == "" {
		erroutput += "Vous devez être connecté pour ajouter un post"
	}

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
		if tab[i].PostName == newpost.PostName {
			erroutput += "Un post avec le même nom existe déja"
			break
		}
	}

	if erroutput == "" && newpost.PostName != "" && newpost.PostCategory != "" && newpost.PostDescription != "" {

		err1 := InsertPosttoDB(newpost)

		if err1 != nil {
			log.Fatalf("DataBase execution: %s", err1)
		}

	}

	if r.Method != "POST" {
		erroutput = ""
	}

	var output Error
	output.Error = erroutput

	fmt.Println(erroutput)

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/newpost.html"))
	err := templates.ExecuteTemplate(w, "newpost", output)

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

	tablist := ReadPosttoDB()

	postmap := make(map[string]int)

	deco := r.FormValue("Deconnexion")

	if deco == "run" {
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
		CategoryList[x].CategoryID = x + 1
		x++
	}
	categorie := r.FormValue("Categorie")

	if categorie != "" {
		var tab []Post
		for i := range tablist {
			if tablist[i].PostCategory == categorie {
				tab = append(tab, tablist[i])
			}
		}
		tablist = tab
	}

	for i := range tablist {
		tablist[i].TabComment = ReadCommenttoDB(tablist[i].PostID)
		tablist[i].NbComment = len(tablist[i].TabComment)
		// fmt.Println(tablist[i].NbComment)
	}

	out.CategoryList = CategoryList
	out.TabList = tablist

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

/*--------------------------------------------------------------------------------------------
--------------------------------------Login Page ---------------------------------------------
----------------------------------------------------------------------------------------------*/
func login(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	password := r.FormValue("user_password")
	email := r.FormValue("user_mail")

	// fmt.Println(password, email)
	//---------------------------On vérififie si l'adresse email et le mdp sont dans la base de donnée et on connecte l'utilisateur--------------------

	user := User{}

	tab := ReadUsertoDB()

	for i := range tab {

		// fmt.Println(tab[i].Email, email)
		// fmt.Println(tab[i].Password, password)

		if tab[i].Email == email && CheckPasswordHash(password, tab[i].Password) == true {
			fmt.Println("connexion effectuée")
			user = tab[i]
		}

	}

	// fmt.Println(user)

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

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
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

	var Comment Comment

	if inputcomment != "" && ID != "" {
		Comment.CommentMessage = inputcomment
		Comment.CommentLikes = 0
		Comment.CommentDislikes = 0
		Comment.CommentDate = time.Now()
		Comment.CommentDateString = Comment.CommentDate.Format("2006-01-02 15:04:05")
		Comment.PostID = output.PostID
		Comment.UserID, _ = strconv.Atoi(ID)
		Comment.UserName = UserName
		Comment.UserAvatar = Avatar
		InsertCommenttoDB(Comment)
	}

	output.TabComment = ReadCommenttoDB(output.PostID)

	if r.FormValue("likes") == "run" {
		AddLiketoPosttoDB("likes", output.PostLikes+1, output.PostID)
		output.PostLikes += 1
	}
	if r.FormValue("dislikes") == "run" {
		AddLiketoPosttoDB("dislikes", output.PostDislikes+1, output.PostID)
		output.PostDislikes += 1
	}
	// fmt.Println(r.FormValue("CommentLikes"))
	// fmt.Println(r.FormValue("CommentDislikes"))

	if r.FormValue("CommentLikes") != "" {
		ID, _ := strconv.Atoi(r.FormValue("CommentLikes"))
		for i := range output.TabComment {
			if ID == output.TabComment[i].CommentID {
				// fmt.Println(ID)
				AddLiketoCommenttoDB("likes", output.TabComment[i].CommentLikes+1, ID)
				output.TabComment[i].CommentLikes += 1
			}

		}

	}
	if r.FormValue("CommentDislikes") != "" {
		ID, _ := strconv.Atoi(r.FormValue("CommentDislikes"))
		for i := range output.TabComment {
			if ID == output.TabComment[i].CommentID {
				// fmt.Println(ID)
				AddLiketoCommenttoDB("dislikes", output.TabComment[i].CommentDislikes+1, ID)
				output.TabComment[i].CommentDislikes += 1
			}

		}
	}
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
