package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Entergolang Start")
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9900", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EnterGolang")
}
