package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("listening on 3000")
	http.ListenAndServe(":3000", nil)
}