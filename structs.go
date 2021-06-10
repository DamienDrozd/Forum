package main

import "time"

/*--------------------------------------------------------------------------------------------
-------------------------------------- Type Struct -------------------------------------------
----------------------------------------------------------------------------------------------*/

type User struct {
	ID           int
	Username     string
	Password     string
	Email        string
	Avatar       string
	Role         string
	PostList     []Post
	CommentList  []Comment
	PostLiked    []Post
	PostDisliked []Post
}

type Error struct {
	Error        string
	User         User
	Post         Post
	CategoryList []Category
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
	TabUser      []User
	User         User
	NbPost       int
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
	CategoryID     int
	CategoryName   string
	CategoryNumber int
}
