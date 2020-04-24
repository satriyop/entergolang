package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


type codeHandler struct {
	snippet string
}
var path = "/Users/satriyo/dev/golang/entergolang/code/test.php"

func startServer() {
	mux := http.NewServeMux()
	ch := &codeHandler{}
	mux.Handle("/exec", ch)
	log.Println("Listening....")
	http.ListenAndServe(":3000", mux)
}



func (ch *codeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO : check lang ID

	// TODO : get code snippet from client
	ch.snippet = `<?php echo "test for php succeed";`
	// ch.snippet = "tes"

	// TODO : save to file (and how to identify unique user to file map)
	createFile(ch.snippet)

	// trigger docker start
	resp := runDockerPhp()
	fmt.Println(resp)

	// send response body
	w.Write([]byte(resp))
}

func createFile(cs string) {
	_, err := os.Stat(path)

	// if os.IsNotExist(err) {
	// 	file, err := os.Create(path)
	// 	check(err)
	// 	defer file.Close()
	// }

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