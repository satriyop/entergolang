package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/russross/blackfriday"
)

// Post is representation of indivial post in dir posts
type Post struct {
	Title   string
	Date    string
	Summary string
	Body    string
	File    string
}

func main() {
	fmt.Println("Entergolang Start")
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", routeHandler)

	log.Fatal(http.ListenAndServe(":9900", nil))
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		posts := getPosts()

		viewsPath := path.Join("views", "index.html")

		t := template.New("index.html")

		t = template.Must(t.ParseFiles(viewsPath))

		t.Execute(w, posts)

	default:
		p := r.URL.Path[1:]

		post := getPost(p)

		viewsPath := path.Join("views", "post.html")

		t := template.New("post.html")

		t = template.Must(t.ParseFiles(viewsPath))

		t.Execute(w, post)
	}
}

func getPosts() []Post {
	posts := []Post{}
	files, _ := filepath.Glob("posts/*.md")

	for _, f := range files {
		// get the filename without dir and ext .md
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)

		// read
		fileread, _ := ioutil.ReadFile(f)

		// break per line
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		posts = append(posts, Post{title, date, summary, body, file})
	}

	return posts
}

func getPost(p string) Post {
	post := Post{}
	f := "posts/" + p + ".md"
	fileread, _ := ioutil.ReadFile(f)
	lines := strings.Split(string(fileread), "\n")
	title := string(lines[0])
	date := string(lines[1])
	summary := string(lines[2])
	body := strings.Join(lines[3:], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post = Post{title, date, summary, body, p}

	return post
}
