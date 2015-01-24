package user

type User struct {
	Name      string
	Location  Location
	Interests []*Interest
}

type Location struct {
	Longitute float64
	Latitude  float64
}

type Interest struct {
	Name   string
	Rating float64
}

type UserRepository interface {
	GetUsers() ([]User, error)
}

func NewUser(name string, interests []*Interest) *User {
	return &User{Name: name, Interests: interests}
}

func NewInterest(name string, rating float64) *Interest {
	return &Interest{Name: name, Rating: rating}
}
