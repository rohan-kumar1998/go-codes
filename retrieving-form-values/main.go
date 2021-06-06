package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type temp struct {
}

func (t temp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	if len(req.Form["fname"]) >= 1 {
		req.Form["fname"] = strings.Split(req.Form["fname"][0], ",")
		for key := range req.Form["fname"] {
			req.Form["fname"][key] = strings.TrimSpace(req.Form["fname"][key])
		}
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	fmt.Println(req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d temp
	http.ListenAndServe(":8080", d)
}
