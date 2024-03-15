package main

import "time"

type Post struct {
	ID        int
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
}

var posts []Post

func CreatePost(title, content, author string) {
	post := Post{
		ID:        len(posts) + 1,
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}
	posts = append(posts, post)
}
