package main

import (
	"fmt"
	"net/http"
)
import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/profile", profileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("Home")
}

func registerHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("Register")
}

func profileHandler(http.ResponseWriter, *http.Request) {
	fmt.Println("Profile")
}
