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
