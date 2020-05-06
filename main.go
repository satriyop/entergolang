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
	// TO DO better server config
	log.Fatal(http.ListenAndServe(":9900", nil))
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	// that r.URL.Path[1:] is a file at directory posts ?
	url := r.URL.Path

	viewsPath, err := filepath.Abs("views")
	printErr(err)

	switch url {
	case "/":
		posts := getPosts()

		indexPath := path.Join(viewsPath, "index.html")

		t := template.New("index.html")
		t = template.Must(t.ParseFiles(indexPath))
		t.Execute(w, posts)

	case url:
		// TODO better matching with existing MD file

		postPath := path.Join(viewsPath, "post.html")

		post, err := getPost(url[1:])
		// Render Error Page if Not Found
		if err != nil {
			errorPath := path.Join(viewsPath, "error.html")
			t := template.New("error.html")
			t = template.Must(t.ParseFiles(errorPath))
			t.Execute(w, "")
		}

		// Render Post Page
		t := template.New("post.html")
		t = template.Must(t.ParseFiles(postPath))
		t.Execute(w, post)

	default:
		errorPath := path.Join(viewsPath, "error.html")
		t := template.New("error.html")
		t = template.Must(t.ParseFiles(errorPath))
		t.Execute(w, "")
	}

}

func getPosts() []Post {
	posts := []Post{}
	postDir, err := filepath.Abs("posts")
	printErr(err)

	files, err := ioutil.ReadDir(postDir)
	printErr(err)

	for _, f := range files {
		filename := f.Name()
		// remove .md extension to be used
		trimmedFilename := strings.TrimSuffix(filename, ".md")
		if isMarkdown(filename) {
			content, fm, err := parseContent(path.Join(postDir, filename))
			printErr(err)

			body := string(blackfriday.MarkdownCommon([]byte(content)))

			posts = append(posts, Post{*fm, body[:200], trimmedFilename})
		}

	}

	return posts
}

func getPost(p string) (*Post, error) {

	// TO DO check for only listed post to serve
	// VALIDATE
	f := "posts/" + p + ".md"

	content, fm, err := parseContent(f)
	if err != nil {
		return nil, err
	}

	body := string(blackfriday.MarkdownCommon([]byte(content)))

	post := &Post{*fm, body, p}

	return post, nil
}

func isMarkdown(f string) bool {
	if !strings.HasSuffix(f, ".md") {
		return false
	}
	return true
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
