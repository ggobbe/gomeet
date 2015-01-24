package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/profile", profileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

var templates = template.Must(template.ParseFiles("static/tpl/header.html", "static/tpl/footer.html", "static/tpl/home.html", "static/tpl/login.html"))

type Page struct {
	Title string
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "session-name")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it.
	session.Save(r, w)

	display(w, "home", &Page{Title: "Home"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "login", &Page{Title: "Login"})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Profile")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
