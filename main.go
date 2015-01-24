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
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/profile", profileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Page struct {
	UserName string
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

	p := &Page{"ggobbe"}

	t, err := template.ParseFiles("static/tpl/home.html")
	checkError(err)
	t.Execute(w, p)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Profile")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
