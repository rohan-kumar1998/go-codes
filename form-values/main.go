package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name    string
	Surname string
	Active  bool
}

var tpl *template.Template

func homepage(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	user_data := user{
		Name:    req.FormValue("name"),
		Surname: req.FormValue("surname"),
		Active:  req.FormValue("active") == "on"}
	fmt.Println(req.Form)
	err = tpl.ExecuteTemplate(w, "index.gohtml", user_data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", homepage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
