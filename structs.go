package main

import "time"

/*--------------------------------------------------------------------------------------------
-------------------------------------- Type Struct -------------------------------------------
----------------------------------------------------------------------------------------------*/

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Avatar   string
}

type Error struct {
	Error string
}

type Comment struct {
	CommentID         int
	CommentMessage    string
	CommentLikes      int
	CommentDislikes   int
	CommentDate       time.Time
	CommentDateString string
	PostID            int
	UserID            int
	UserName          string
	UserAvatar        string
}

type Out struct {
	TabList      []Post
	CategoryList []Category
	User         User
}

type Post struct {
	PostID          int
	PostName        string
	PostCategory    string
	PostDate        time.Time
	PostDateString  string
	PostDescription string
	PostURL         string
	PostLikes       int
	PostDislikes    int
	UserID          int
	UserName        string
	UserAvatar      string
	TabComment      []Comment
	NbComment       int
}

type Category struct {
	CategoryName   string
	CategoryNumber int
	CategoryID     int
}
