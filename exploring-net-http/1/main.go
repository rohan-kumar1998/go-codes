package main

import (
	"fmt"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is dog")
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index")
}
func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is me")
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", index)
	r.HandleFunc("/dog/", dog)
	r.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", r)
}
