package user

import (
	"errors"
	"net/http"
	"strings"

	"gomeet/common"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("gomeet-for-gopher-gala-by-gg-and-mk"))

// User type
type User struct {
	Name      string
	Location  Location
	Interests Interests
}

// Location type
type Location struct {
	Longitute float64
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
	user := &User{Name: username.(string)}
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
