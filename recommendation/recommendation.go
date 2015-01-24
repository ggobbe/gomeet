package recommendation

import "gomeet/user"

type Recommendation struct {
	User  *user.User
	Ratio float64
}

type Recommender interface {
	GetRecommendations(user *user.User) []Recommendation
}

func GetRecommendations(user *user.User) []Recommendation {
	return make([]Recommendation, 0)
}
