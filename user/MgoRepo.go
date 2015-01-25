package user

import (
	"edigophers/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoRepo is a repository based on a MongoDb
type MgoRepo struct {
	url        string
	database   string
	collection string
}

const collection = "users"

func (r MgoRepo) getUserCollection() *mgo.Collection {
	s, err := mgo.Dial(r.url)
	utils.CheckError(err)
	return s.DB(r.database).C(r.collection)
}

// GetUser gets a user per his username
func (r MgoRepo) GetUser(name string) (*User, error) {
	c := r.getUserCollection()
	user := User{}
	err := c.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUsers is a function returning a list of users
func (r MgoRepo) GetUsers() ([]User, error) {
	c := r.getUserCollection()
	users := []User{}
	err := c.Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// SaveUser saves a user to the database
func (r MgoRepo) SaveUser(usr User) error {
	c := r.getUserCollection()
	return c.Insert(usr)
}

// NewMgoRepo creates a new Mongo database repository
func NewMgoRepo(url, database string) Repository {
	repo := MgoRepo{url: url, database: database, collection: collection}
	return Repository(repo)
}
