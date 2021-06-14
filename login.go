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
--------------------------------------Login Page ---------------------------------------------
----------------------------------------------------------------------------------------------*/
func login(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now() // Print the time of execution of the function

	password := r.FormValue("user_password")
	email := r.FormValue("user_mail") // get the values of the password and email from user

	//---------------------------On vérififie si l'adresse email et le mdp sont dans la base de donnée et on connecte l'utilisateur--------------------

	user := User{}

	tab := ReadUsertoDB()

	for i := range tab {

		if tab[i].Email == email && CheckPasswordHash(password, tab[i].Password) == true {
			fmt.Println("connexion effectuée") // Cheack if the email and password are in the db
			user = tab[i]
		}

	}

	output := ""

	if user.ID != 0 {

		// Log the user with cookies

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
