package user

import (
	"edigophers/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//FileRepo is a file based user repository
type FileRepo struct {
	filepath string
	usersMap map[string]User
	users    []User
}

// GetUser gets a user per his username
func (r FileRepo) GetUser(name string) (*User, error) {

	usr, ok := r.usersMap[name]
	if !ok {
		return nil, errors.New("User doesn't exists")
	}

	return &usr, nil
}

//GetUsers is a function returning a list of users
func (r FileRepo) GetUsers() ([]User, error) {
	return r.users, nil
}

//SaveUser is a function to persist user changes
func (r FileRepo) SaveUser(usr User) error {
	_, ok := r.usersMap[usr.Name]

	if !ok {
		r.usersMap[usr.Name] = usr
		r.users = append(r.users, usr)
	} else {
		r.usersMap[usr.Name] = usr
		for i, val := range r.users {
			if val.Name == usr.Name {
				r.users[i] = usr
			}
		}

	}
	log.Printf("(%v)", r)
	err := r.saveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (r FileRepo) saveToFile() error {

	body, err := json.Marshal(r.users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filepath, body, 0600)
	if err != nil {
		return err
	}

	return nil
}

func loadFile(filepath string) ([]User, error) {

	body, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var usersJSON []User
	if len(body) > 0 {
		err = json.Unmarshal(body, &usersJSON)
		if err != nil {
			return nil, err
		}
	}

	return usersJSON, nil
}

//NewRepo creates a new File base repository
func NewRepo(filepath string) (*FileRepo, error) {

	u, err := loadFile(filepath)
	if err != nil {
		return nil, err
	}
	utils.CheckError(err)

	usersMap := make(map[string]User)
	for _, usr := range u {
		localUsr := usr //iteration variable usr is reused in the loop
		usersMap[localUsr.Name] = localUsr
	}

	return &FileRepo{users: u, usersMap: usersMap, filepath: filepath}, nil
}
