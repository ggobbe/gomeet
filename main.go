package main

import (
	"fmt"
	"gomeet/user"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/profile", profileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

var templates = template.Must(template.ParseFiles("static/tpl/header.html", "static/tpl/footer.html", "static/tpl/home.html", "static/tpl/login.html"))

type Page struct {
	Title string
	User  *user.User
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-session")
	checkError(err)
	username, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	user := &user.User{Name: username.(string)}
	display(w, "home", &Page{Title: "Home", User: user})
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "login", &Page{Title: "Login"})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-session")
	checkError(err)
	username := r.FormValue("username")
	if strings.Trim(username, " ") == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	session.Values["username"] = username
	session.Save(r, w)
	fmt.Printf("Username %s is logged", username)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-session")
	checkError(err)
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Profile")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
