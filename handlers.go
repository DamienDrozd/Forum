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
	var user User
	user.Username = UserName
	user.Avatar = Avatar
	user.ID, _ = strconv.Atoi(ID)

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
		c, _ = r.Cookie("Role")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		user = User{}
	}

	erroutput := ""

	if UserName == "" || ID == "" {
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

/*--------------------------------------------------------------------------------------------
-------------------------------------- Index Page -------------------------------------------
----------------------------------------------------------------------------------------------*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var user User

	for _, cookie := range r.Cookies() {
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
		var tab []Post
		for i := range tablist {
			if tablist[i].PostCategory == categorie {
				tab = append(tab, tablist[i])
			}
		}
		tablist = tab
	}

	x := len(tablist)
	for i := 0; i < x; i++ {
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
		cookie = http.Cookie{Name: "Role", Value: user.Role} //	, Expires: expiration}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

	if password != "" && email != "" && user.ID == 0 {
		output += "L'email ou le mot de passe entré est incorrect"
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
		Role:     "user",
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
	var user User

	for _, cookie := range r.Cookies() {
		if cookie.Name == "Username" {
			UserName = cookie.Value
			user.Username = cookie.Value
		}
		if cookie.Name == "Avatar" {
			Avatar = cookie.Value
			user.Avatar = cookie.Value
		}
		if cookie.Name == "ID" {
			ID = cookie.Value
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
		c, _ = r.Cookie("Role")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		user = User{}
	}

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

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

	SupprimerPost := r.FormValue("SuprimerPost")

	if user.Role == "modo" || user.Role == "admin" {
		if SupprimerPost != "" {
			DeletePosttoDB(output.PostID)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	SupprimerComment, _ := strconv.Atoi(r.FormValue("SuprimerCommentaire"))

	if user.Role == "modo" || user.Role == "admin" {
		if SupprimerComment != 0 {
			DeleteCommenttoDB(SupprimerComment)
		}
	}

	var comment Comment

	if inputcomment != "" && ID != "" {
		comment.CommentMessage = inputcomment
		comment.CommentLikes = 0
		comment.CommentDislikes = 0
		comment.CommentDate = time.Now()
		comment.CommentDateString = comment.CommentDate.Format("2006-01-02 15:04:05")
		comment.PostID = output.PostID
		comment.UserID, _ = strconv.Atoi(ID)
		comment.UserName = UserName
		comment.UserAvatar = Avatar
		InsertCommenttoDB(comment)
		SendMail("Comment", comment, output)
	}

	output.TabComment = ReadCommenttoDB(output.PostID)

	if r.FormValue("likes") == "run" {
		AddLiketoPosttoDB("likes", output.PostLikes+1, output.PostID)
		output.PostLikes += 1
		SendMail("NewLike", Comment{}, output)
	}
	if r.FormValue("dislikes") == "run" {
		AddLiketoPosttoDB("dislikes", output.PostDislikes+1, output.PostID)
		output.PostDislikes += 1
		SendMail("NewDislike", Comment{}, output)
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

	var Error Error
	Error.User = user
	Error.Post = output
	//--------------------------------------------------------------------

	templates := template.New("Label de ma template")
	templates = template.Must(templates.ParseFiles("./templates/post.html"))
	err := templates.ExecuteTemplate(w, "post", Error)

	if err != nil {
		log.Fatalf("Template execution: %s", err) // If the executetemplate function cannot run, displays an error message
	}
	t := time.Now()
	fmt.Println("time1:", t.Sub(timestart))

}

/*--------------------------------------------------------------------------------------------
-------------------------------------- User Page----------------------------------------------
----------------------------------------------------------------------------------------------*/
func user(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var user User

	for _, cookie := range r.Cookies() {
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
		c, _ = r.Cookie("Role")
		if c != nil {
			c.MaxAge = -1 // delete cookie
			http.SetCookie(w, c)
		}
		user = User{}
	}

	if user.ID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	UserList := ReadUsertoDB()
	PostList := ReadPosttoDB()

	for i := range UserList {
		if UserList[i].ID == user.ID {
			user = UserList[i]
		}
	}

	for i := range PostList {
		if PostList[i].UserID == user.ID {
			PostList[i].TabComment = ReadCommenttoDB(PostList[i].PostID)
			user.PostList = append(user.PostList, PostList[i])
		}

	}

	DeletePost, _ := strconv.Atoi(r.FormValue("delete_post"))
	DeleteComment, _ := strconv.Atoi(r.FormValue("delete_comment"))

	if DeletePost != 0 {
		DeletePosttoDB(DeletePost)
		for i := range user.PostList {
			if user.PostList[i].PostID == DeletePost {

				user.PostList = append(user.PostList[:i], user.PostList[i+1:]...)
				break

			}
		}
	}
	if DeleteComment != 0 {
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

/*--------------------------------------------------------------------------------------------
-------------------------------------- Admin Page----------------------------------------------
----------------------------------------------------------------------------------------------*/

func admin(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var user User
	var output Out

	for _, cookie := range r.Cookies() {
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
	}

	TabUser := ReadUsertoDB()

	AddCategory := (r.FormValue("category"))
	if AddCategory != "" {
		var NewCategory Category
		NewCategory.CategoryName = AddCategory
		InsertCategorytoDB(NewCategory)

	}
	DeleteCategory, _ := strconv.Atoi((r.FormValue("deletecategory")))

	if DeleteCategory != 0 {
		DeleteCategorytoDB(DeleteCategory)
	}

	Promote, _ := strconv.Atoi(r.FormValue("promote"))
	Demote, _ := strconv.Atoi(r.FormValue("demote"))

	var userpromote User

	for i := range TabUser {

		if TabUser[i].ID == Promote || TabUser[i].ID == Demote {
			userpromote = TabUser[i]
		}
	}

	if Promote != 0 {
		PromoteUsertoDB(userpromote, "promote")
	}
	if Demote != 0 {
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

/*--------------------------------------------------------------------------------------------
-------------------------------------- Moderator Page----------------------------------------------
----------------------------------------------------------------------------------------------*/

func moderator(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	var user User

	for _, cookie := range r.Cookies() {
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
	}

	uploadpost, _ := strconv.Atoi(r.FormValue("uploadpost"))
	deletepost, _ := strconv.Atoi(r.FormValue("deletepost"))

	if uploadpost != 0 {
		ValidatePosttoDB(uploadpost)
	}
	if deletepost != 0 {
		DeletePosttoDB(deletepost)
	}

	var output Out
	tablist := ReadPosttoDB()

	x := len(tablist)
	for i := 0; i < x; i++ {
		if tablist[i].Validated == "true" {
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
