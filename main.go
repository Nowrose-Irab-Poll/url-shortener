package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("slash")
	fmt.Fprintf(w, "Hello, URL Shortener!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Hello")
	http.ListenAndServe(":8080", nil)
}
