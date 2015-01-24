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
	uRepo user.UserRepository
}

func (sr SimpleRecommender) GetRecommendations(user *user.User) ([]Recommendation, error) {

	// := sr.uRepo.GetUsers()

	return make([]Recommendation, 0), nil
}
