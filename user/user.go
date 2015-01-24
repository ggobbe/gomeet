package user

type User struct {
	UserName  string
	Location  Location
	Interests []Interest
}

type Location struct {
	Longitute float64
	Latitude  float64
}

type Interest struct {
	Name   string
	Rating float64
}

func Login(username string) {

}
