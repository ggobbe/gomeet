package recommendation

import (
	"gomeet/user"
	"log"

	"github.com/muesli/regommend"
)

//Recommendation contains a user with similiar interests
type Recommendation struct {
	User  user.User
	Score float64
}

//Recommender returns recommendations of users with similiar interests
type Recommender interface {
	GetRecommendations(user *user.User) ([]Recommendation, error)
}

//SimpleRecommender is a simple implementation of the Recommender using regocommend neighbours
type SimpleRecommender struct {
	uRepo user.Repository
}

//GetRecommendations is a method for returning recommendations of users with similiar interests
func (sr SimpleRecommender) GetRecommendations(usr *user.User) ([]Recommendation, error) {
	minScore := 0.6

	userMap, interests, err := prepareData(sr, usr)

	neighbours, err := interests.Neighbors(usr.Name)
	if err != nil {
		return nil, err
	}

	result := make([]Recommendation, 0, 10)
	for _, rec := range neighbours {
		if rec.Key == "" {
			continue
		}

		if rec.Distance >= minScore {
			u, ok := userMap[rec.Key.(string)]
			if !ok {
				log.Printf("[WARN] User map does not contain user with id:(%s)", rec.Key.(string))
			}
			result = append(result, Recommendation{User: u, Score: rec.Distance})
		}
	}

	return result, nil
}

func prepareData(sr SimpleRecommender, usr *user.User) (map[string]user.User, *regommend.RegommendTable, error) {
	users, err := sr.uRepo.GetUsers()
	if err != nil {
		return nil, nil, err
	}

	interests := regommend.Table("interests")
	userMap := make(map[string]user.User)

	for _, u := range users {
		userMap[u.Name] = u
		interests.Add(u.Name, u.Interests.AsMap())
	}
	if _, ok := userMap[usr.Name]; !ok {
		interests.Add(usr.Name, usr.Interests.AsMap())
	}

	return userMap, interests, nil
}
