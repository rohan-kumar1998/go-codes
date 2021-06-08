package main

import (
	"fmt"
	"net/http"
)

func setcookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "cookie", Value: "hello"})
	fmt.Fprintln(w, "The cookie has been set")
}

func readcookie(w http.ResponseWriter, req *http.Request) {
	co, err := req.Cookie("cookie")
	if err != nil {
		fmt.Fprintln(w, "no cookie found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "cookie value", (*co).Value)
}
func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Home")
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", homepage)
	http.HandleFunc("/setcookie", setcookie)
	http.HandleFunc("/readcookie", readcookie)
	http.ListenAndServe(":8080", nil)
}
