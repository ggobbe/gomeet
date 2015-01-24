package main

import (
	"gomeet/common"
	"gomeet/user"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/profile", profileHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

var templates = template.Must(template.ParseFiles(
	"tpl/header.html", "tpl/footer.html",
	"tpl/home.html", "tpl/login.html", "tpl/profile.html"))

type page struct {
	Title string
	User  *user.User
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetSessionUser(w, r)
	if err != nil {
		return
	}
	display(w, "home", &page{Title: "Home", User: user})
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "login", &page{Title: "Login"})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if user.SetSessionUser(w, r, username) != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	common.CheckError(user.LogOutSessionUser(w, r))
	http.Redirect(w, r, "/login", http.StatusFound)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetSessionUser(w, r)
	if err != nil {
		return
	}
	display(w, "profile", &page{Title: "Profile", User: user})
}
