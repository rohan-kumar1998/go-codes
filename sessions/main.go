package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	name  string
	email string
	pass  string
}
type logins struct {
	pass   string
	userID string
}

var userData = map[string]user{}
var sessions = map[string]string{}
var userLogins = map[string]logins{}

func alreadyLoggedIn(req *http.Request) bool {
	co, err := req.Cookie("my-cookie")
	if err != nil {
		return false
	}
	if _, ok := sessions[co.Value]; ok {
		return true
	}
	return false
}

func login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if alreadyLoggedIn(req) {
			w.Header().Set("Location", "/profile")
			w.WriteHeader(http.StatusSeeOther)
			return
		}
		sessionId := uuid.Must(uuid.NewRandom())
		req.ParseForm()
		if _, ok := userLogins[req.FormValue("email")]; !ok {
			tpl.ExecuteTemplate(w, "login.gohtml", nil)
			fmt.Fprintln(w, "USERNAME DOESNT MATCH")
			return
		}

		if userLogins[req.FormValue("email")].pass != req.FormValue("password") {
			tpl.ExecuteTemplate(w, "login.gohtml", nil)
			fmt.Fprintln(w, "PASSWORD DOESNT MATCH")
			return
		}
		sessions[sessionId.String()] = userLogins[req.FormValue("email")].userID
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: sessionId.String(),
		})

		w.Header().Set("Location", "/profile")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	if alreadyLoggedIn(req) {
		w.Header().Set("Location", "/profile")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		userId := uuid.Must(uuid.NewRandom())
		sessionId := uuid.Must(uuid.NewRandom())

		co := &http.Cookie{
			Name:  "my-cookie",
			Value: sessionId.String()}
		http.SetCookie(w, co)
		sessions[sessionId.String()] = userId.String()

		userData[userId.String()] = user{
			name:  req.FormValue("name"),
			email: req.FormValue("email"),
			pass:  req.FormValue("password")}

		userLogins[req.FormValue("email")] = logins{
			pass:   req.FormValue("password"),
			userID: userId.String()}

		w.Header().Set("Location", "/profile")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	if alreadyLoggedIn(req) {
		w.Header().Set("Location", "/profile")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	co, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	co.MaxAge = -1
	http.SetCookie(w, co)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", nil)
}

func profile(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "profile.gohtml", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
