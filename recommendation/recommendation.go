package recommendation

import "gomeet/domain"

type Recommendation struct {
	user  *user.User
	ratio float64
}

type Recommender interface {
	GetRecommendation(user *user.User) []Recommendation
}
