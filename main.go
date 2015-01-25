package main

import (
	"edigophers/recommendation"
	"edigophers/user"
	"edigophers/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var templates = template.Must(template.ParseFiles(
	"tpl/header.html", "tpl/footer.html",
	"tpl/home.html", "tpl/login.html", "tpl/profile.html", "tpl/list.html"))

var repository = user.GetRepo()
var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

type page struct {
	Title string
	User  *user.User
	Data  interface{}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/list", listHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/interest/add", interestAddHandler).Methods("POST")
	r.HandleFunc("/profile/{username}", profileHandler)
	r.HandleFunc("/profile", profileHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	usr, err := user.GetSessionUser(w, r, repository, store)
	if err != nil {
		return
	}

	_ = recommendation.New(repository)

	display(w, "home", &page{Title: "What would you like to do?", User: usr})
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "login", &page{Title: "Login"})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if user.SetSessionUser(w, r, username, store) != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.CheckError(user.LogOutSessionUser(w, r, store))
	http.Redirect(w, r, "/login", http.StatusFound)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	usr, err := user.GetSessionUser(w, r, repository, store)
	if err != nil {
		return
	}

	users, err := repository.GetUsers()
	if err != nil {
		return
	}

	display(w, "list", &page{Title: "List of users", User: usr, Data: users})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Display the profile of somebody else
	vars := mux.Vars(r)
	if username, ok := vars["username"]; ok {
		user, err := repository.GetUser(username)

		if err != nil {
			return
		}
		display(w, "profile", &page{Title: fmt.Sprintf("%s's Profile", user.Name), User: user})
		return
	}

	// Display the profile of the current user
	user, err := user.GetSessionUser(w, r, repository, store)
	utils.CheckError(err)
	display(w, "profile", &page{Title: "Your Profile", User: user})
}

func interestAddHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("interest")
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	utils.CheckError(err)
	usr, err := user.GetSessionUser(w, r, repository, store)
	if err != nil {
		return
	}
	interest := user.NewInterest(name, rating)
	usr.Interests = append(usr.Interests, *interest)

	err = repository.SaveUser(*usr)

	http.Redirect(w, r, "/profile", http.StatusFound)
}
