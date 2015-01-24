package user

type User struct {
	UserName string
	Location Location
}

type Location struct {
	longitute float64
	latitude  float64
}
