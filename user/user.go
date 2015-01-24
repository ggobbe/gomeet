package user

import (
	"errors"
	"net/http"
	"strings"

	"gomeet/common"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

type User struct {
	Name      string
	Location  Location
	Interests UserInterests
}

type Location struct {
	Longitute float64
	Latitude  float64
}

type UserInterests []*Interest

type Interest struct {
	Name   string
	Rating float64
}

type UserRepository interface {
	GetUsers() ([]User, error)
}

func NewUser(name string, interests UserInterests) *User {
	return &User{Name: name, Interests: interests}
}

func NewInterest(name string, rating float64) *Interest {
	return &Interest{Name: name, Rating: rating}
}

func (ui UserInterests) AsMap() map[interface{}]float64 {
	interestMap := make(map[interface{}]float64)

	for _, i := range ui {
		interestMap[i.Name] = i.Rating
	}
	return interestMap
}

func GetSessionUser(w http.ResponseWriter, r *http.Request) (*User, error) {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	username, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, errors.New("No user in the session")
	}
	user := &User{Name: username.(string)}
	return user, nil
}

func SetSessionUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	username := r.FormValue("username")
	if strings.Trim(username, " ") == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	session.Values["username"] = username
	return session.Save(r, w)
}

func LogOutSessionUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	delete(session.Values, "username")
	return session.Save(r, w)
}
