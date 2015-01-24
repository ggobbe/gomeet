package recommendation

import "os/user"

type Recommendation struct {
	User  *user.User
	Ratio float64
}

type Recommender interface {
	GetRecommendations(user *user.User) []Recommendation
}
