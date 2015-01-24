package recommendation

import (
	"gomeet/user"
	"log"

	"github.com/muesli/regommend"
)

type Recommendation struct {
	User  *user.User
	Score float64
}

type Recommender interface {
	GetRecommendations(user *user.User) ([]Recommendation, error)
}

type SimpleRecommender struct {
	uRepo user.UserRepository
}

func (sr SimpleRecommender) GetRecommendations(user *user.User) ([]Recommendation, error) {

	users, err := sr.uRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	interests := regommend.Table("interests")
	for _, u := range users {
		interests.Add(u.Name, u.Interests.AsMap())
	}
	interests.Add(user.Name, user.Interests.AsMap())

	log.Print("Table ", interests)
	log.Println("Looking for recommendations for user ", user.Name)
	recs, err := interests.Recommend(user.Name)
	if err != nil {
		return nil, err
	}

	for _, rec := range recs {
		log.Println("Recommending", rec.Key, "with score", rec.Distance)
	}

	neighbours, err := interests.Neighbors("Chris")
	if err != nil {
		return nil, err
	}
	for _, rec := range neighbours {
		log.Println("Neighbours", rec.Key, "with score", rec.Distance)
	}

	return make([]Recommendation, 0), nil
}
