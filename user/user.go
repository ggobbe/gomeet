package user

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"gomeet/common"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

var users = map[string]*User{
	"ggobbe": &User{Name: "Guillaume",
		Location: Location{Latitude: 55.86122317, Longitude: -3.34246233},
		Interests: Interests{
			&Interest{Name: "Rugby", Rating: 5},
			&Interest{Name: "Tennis", Rating: 8}}},
	"kotulamar": &User{Name: "Martin",
		Location: Location{Latitude: 56.017244, Longitude: -2.8197334},
		Interests: Interests{
			&Interest{Name: "Skiing", Rating: 7},
			&Interest{Name: "Cinema", Rating: 8.5},
			&Interest{Name: "Salsa", Rating: 3}}}}

// User type
type User struct {
	Name      string
	Location  Location
	Interests Interests
}

// Location type
type Location struct {
	Longitude float64
	Latitude  float64
}

// Interests type
type Interests []*Interest

// Interest type
type Interest struct {
	Name   string
	Rating float64
}

// Repository type
type Repository interface {
	GetUsers() ([]User, error)
}

// NewUser creates a new user
func NewUser(name string, interests Interests) *User {
	return &User{Name: name, Interests: interests}
}

// NewInterest creates a new interest
func NewInterest(name string, rating float64) *Interest {
	return &Interest{Name: name, Rating: rating}
}

// AsMap creates a map from the users interests
func (ui Interests) AsMap() map[interface{}]float64 {
	interestMap := make(map[interface{}]float64)

	for _, i := range ui {
		interestMap[i.Name] = i.Rating
	}
	return interestMap
}

// GetSessionUser gets the user stored in the session if there is one
func GetSessionUser(w http.ResponseWriter, r *http.Request) (*User, error) {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	username, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, errors.New("No user in the session")
	}
	user, err := GetUser(username.(string))
	if err != nil {
		if err := LogOutSessionUser(w, r); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, err
	}
	return user, nil
}

// GetUser gets a user per his username
func GetUser(username string) (*User, error) {
	user, ok := users[username]
	if !ok {
		return nil, errors.New("User doesn't exists")
	}
	return user, nil
}

// SetSessionUser sets the user in the session
func SetSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	if strings.Trim(username, " ") == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	session.Values["username"] = username
	return session.Save(r, w)
}

// LogOutSessionUser logs out the user
func LogOutSessionUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "user-session")
	common.CheckError(err)
	delete(session.Values, "username")
	return session.Save(r, w)
}
