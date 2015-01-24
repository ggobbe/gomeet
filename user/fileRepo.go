package user

import (
	"edigophers/common"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type FileRepo struct {
	usersMap map[string]*User
	users    []User
}

// GetUser gets a user per his username
func (r FileRepo) GetUser(name string) (*User, error) {

	usr, ok := r.usersMap[name]
	if !ok {
		return nil, errors.New("User doesn't exists")
	}
	return usr, nil
}

//GetUsers is a function returning a list of users
func (r FileRepo) GetUsers() ([]User, error) {
	return r.users, nil
}

func loadFile(filepath string) ([]User, error) {

	body, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var usersJSON []User
	err = json.Unmarshal(body, &usersJSON)
	if err != nil {
		return nil, err
	}

	return usersJSON, nil
}

func NewRepo(filepath string) *FileRepo {

	u, err := loadFile(filepath)
	common.CheckError(err)

	usersMap := make(map[string]*User)
	for _, u := range u {
		usersMap[u.Name] = &u
	}
	return &FileRepo{users: u, usersMap: usersMap}
}
