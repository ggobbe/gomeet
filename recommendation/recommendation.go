package recommendation

import "gomeet/user"

type Recommendation struct {
	User  *user.User
	Ratio float64
}

type Recommender interface {
	GetRecommendations(user *user.User) ([]Recommendation, error)
}

type SimpleRecommender struct {
	userRepository user.UserRepository
}

func (sr SimpleRecommender) GetRecommendations(user *user.User) ([]Recommendation, error) {

	return make([]Recommendation, 0), nil
}
