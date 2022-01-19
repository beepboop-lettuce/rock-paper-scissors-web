package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello, world</strong>`
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting wb server on port 8080")
	http.ListenAndServe(":8080", nil)
}
