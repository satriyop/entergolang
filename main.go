package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Entergolang Start")

	http.Handle("/", http.HandlerFunc(index))

	log.Fatal(http.ListenAndServe(":9900", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EnterGolang")
}
