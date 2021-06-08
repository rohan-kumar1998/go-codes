package main

import (
	"fmt"
	"net/http"
)

func redir(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "/redirected")
	w.WriteHeader(http.StatusSeeOther)
	fmt.Fprintln(w, "Redirecting...")
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Redirected")
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { fmt.Fprintln(w, "HOME") })
	http.HandleFunc("/invalid", redir)
	http.HandleFunc("/redirected", foo)
	http.ListenAndServe(":8080", nil)
}
