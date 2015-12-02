package main

import (
	"fmt"
	"html"
	"net/http"
)

func httptest() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello,%q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	httptest()
}
