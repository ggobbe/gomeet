package user

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnMarshalUserArray(t *testing.T) {
	var m []User

	body := []byte(`[{"Name":"Guillaume","Location":{"Longitude":-3.34246233,"Latitude":55.86122317},"Interests":[{"Name":"Rugby","Rating":5},{"Name":"Tennis","Rating":8}]},{"Name":"Martin","Location":{"Longitude":-2.8197334,"Latitude":56.017244},"Interests":[{"Name":"Skiing","Rating":7},{"Name":"Cinema","Rating":8.5},{"Name":"Salsa","Rating":3}]}]`)

	err := json.Unmarshal(body, &m)

	if err != nil {
		t.Errorf("(%v)", err)
	}

	log.Printf("(%v)", m)
}

func TestCanReadUserDataFromFile(t *testing.T) {
	repo := NewRepo("../data/users.json")
	users, err := repo.GetUsers()

	if err != nil {
		t.Errorf("GetRecommendations failed: %s", err)
	}

	if len(users) == 0 {
		t.Error("Failed to read users data expected not empty but was empty failed")
	}

	log.Printf("(%v)", users[0])
}

func TestCanSerializeToJson(t *testing.T) {
	var users = []User{
		User{Name: "Guillaume",
			Location: Location{Latitude: 55.86122317, Longitude: -3.34246233},
			Interests: Interests{
				Interest{Name: "Rugby", Rating: 5},
				Interest{Name: "Tennis", Rating: 8}}},
		User{Name: "Martin",
			Location: Location{Latitude: 56.017244, Longitude: -2.8197334},
			Interests: Interests{
				Interest{Name: "Skiing", Rating: 7},
				Interest{Name: "Cinema", Rating: 8.5},
				Interest{Name: "Salsa", Rating: 3}}}}

	_, err := json.Marshal(users)
	if err != nil {
		t.Error(err)
	}

}
