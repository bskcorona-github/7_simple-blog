// main.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/create", handleCreate)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Simple Blog!")
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Posts:")
	for _, post := range posts {
		fmt.Fprintf(w, "%d. %s - <a href=\"/post?id=%d\">Read More</a>\n", post.ID, post.Title, post.ID)
	}
	fmt.Fprintln(w, "<br/><a href=\"/create\">Create New Post</a>")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > len(posts) {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post := posts[id-1]
	fmt.Fprintf(w, "Post #%d - %s\n", post.ID, post.Title)
	fmt.Fprintf(w, "Author: %s\n", post.Author)
	fmt.Fprintf(w, "Content: %s\n", post.Content)
	fmt.Fprintf(w, "Created At: %s\n", post.CreatedAt)
}

// main.go
func handleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// フォームを表示
		http.ServeFile(w, r, "create.html")
	} else if r.Method == "POST" {
		// フォームから送信されたデータを受け取り、新しい投稿を作成
		title := r.FormValue("title")
		content := r.FormValue("content")
		author := r.FormValue("author")
		CreatePost(title, content, author)
		// 投稿が正常に作成されたら投稿一覧ページにリダイレクト
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}
