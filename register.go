package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------- Register Page------------------------------------------
----------------------------------------------------------------------------------------------*/
func register(w http.ResponseWriter, r *http.Request) {
	timestart := time.Now()

	pseudo := r.FormValue("username") // Get the values of user input
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

	if password_repeat != password { // Test the two password input
		output += "Erreur, les mots de passe ne correspondent pas\n"
	}

	if len(password) == 0 || len(pseudo) == 0 || len(email) == 0 {

		output += "Erreur, vous devez remplir toutes les cases.\n"

	}

	for i := range tab { // Test if the username and the password are already taken
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
