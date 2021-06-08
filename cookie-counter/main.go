package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func setcookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "0"})
	tpl.ExecuteTemplate(w, "home.gohtml", nil)

	fmt.Fprintln(w, "Cookie was set")
}

func readcookie(w http.ResponseWriter, req *http.Request) {
	co, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "home.gohtml", nil)
	fmt.Fprintln(w, "The number of times the button was pressed was is :", (*co).Value)
}

func buttonpress(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	co, err := req.Cookie("my-cookie")
	if err != nil {
		tpl.ExecuteTemplate(w, "home.gohtml", nil)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	i, err := strconv.Atoi(co.Value)
	if err != nil {
		log.Fatalln(err)
	}
	if _, ok := req.Form["submit-button"]; ok {
		fmt.Println("counted")
		i++
		co.Value = strconv.Itoa(i)
		http.SetCookie(w, co)
	}
	tpl.ExecuteTemplate(w, "home.gohtml", nil)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/setcookie", setcookie)
	http.HandleFunc("/readcookie", readcookie)
	http.HandleFunc("/", buttonpress)
	http.ListenAndServe(":8080", nil)
}
