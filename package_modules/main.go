package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web !!")
}

func main() {
	fmt.Println("Welcome to packages and moudles")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
