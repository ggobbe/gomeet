package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/profile", profileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	UserName string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
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
