package Data

import (
	"time"

	"database/sql"
	// "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// func PasswordHash(user User) string {
// 	H := sha1.New()
// 	H.Write([]byte(user.Password))
// 	result := H.Sum(nil)
// 	return string(result)
// }

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
	Date     time.Time
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

func data() {
	db, _ = sql.Open("sqlite3", "data.db")

	createDB(UserTab)
	createDB(CommentTab)
	createDB(PostTab)
}

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
