package user

import (
	"edigophers/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
)

//FileRepo is a file based user repository
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

//NewRepo creates a new File base repository
func NewRepo(filepath string) *FileRepo {

	u, err := loadFile(filepath)
	utils.CheckError(err)

	usersMap := make(map[string]*User)
	for _, usr := range u {
		localUsr := usr //iteration variable usr is reused in the loop
		usersMap[localUsr.Name] = &localUsr
	}

	return &FileRepo{users: u, usersMap: usersMap}
}
