package user

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

func NewUser(name string, interests []*Interest) *User {
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
