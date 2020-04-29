package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type codeHandler struct {
	snippet string
}

type codeSnippet struct {
	Code string
}
type codeRun struct {
	Result string `json:"result"`
}

var path = "/Users/satriyo/dev/golang/entergolang/code/test.php"

func startServer() {
	mux := http.NewServeMux()
	ch := &codeHandler{}

	mux.Handle("/exec", ch)

	// mux.HandleFunc("/cookie", createCookie)

	// mux.HandleFunc("/session", createSession)

	// mux.HandleFunc("/tes", tesSession)

	log.Println("Listening....")
	http.ListenAndServe(":3000", mux)
}

func (ch *codeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// start session

	// TODO : check lang ID

	// TODO : get code snippet from client
	var cs codeSnippet
	body, err := ioutil.ReadAll(r.Body)
	check(err)

	err = json.Unmarshal(body, &cs)
	check(err)
	fmt.Println(cs)

	// TODO : save to file (and how to identify unique user to file map)
	createFile(cs.Code)

	// trigger docker start
	out := runDockerPhp()
	fmt.Println(out)

	// prepare response body as json
	var resp = &codeRun{
		Result: out,
	}
	jsonData, _ := json.Marshal(resp)

	// enable cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// send json to client
	w.Write([]byte(jsonData))
}

func createFile(cs string) {
	_, err := os.Stat(path)

	file, err := os.Create(path)
	check(err)
	defer file.Close()

	fmt.Println("File created")
	writeFile(cs)
}

func writeFile(cs string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	check(err)
	defer file.Close()

	_, err = file.WriteString(cs)
	check(err)

	err = file.Sync()
	check(err)
	fmt.Println("Write to file success")
}
