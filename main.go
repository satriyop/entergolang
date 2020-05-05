package main

import (
	"fmt"
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
	FrontMatter
	Body string
	File string
}

// FrontMatter is representation of meta data of post
type FrontMatter struct {
	Title      string
	Date       string
	Draft      bool
	Tags       []string
	Categories []string
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

	files, err := filepath.Glob("posts/*.md")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// get the filename without dir and ext .md
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fmt.Println(f)
		content, fm := parseContent(f)
		body := string(blackfriday.MarkdownCommon([]byte(content)))

		posts = append(posts, Post{*fm, body, file})
	}

	return posts
}

func getPost(p string) Post {
	post := Post{}
	f := "posts/" + p + ".md"
	// fileread, _ := ioutil.ReadFile(f)
	// lines := strings.Split(string(fileread), "\n")
	// title := string(lines[0])
	// date := string(lines[1])
	// summary := string(lines[2])
	// body := strings.Join(lines[3:], "\n")
	// body = string(blackfriday.MarkdownCommon([]byte(body)))
	// post = Post{title, date, summary, body, p}

	content, fm := parseContent(f)
	body := string(blackfriday.MarkdownCommon([]byte(content)))

	post = Post{*fm, body, p}

	return post
}
