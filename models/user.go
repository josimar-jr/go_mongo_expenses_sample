package models

import (
	"github.com/josimar-jr/go_mongo_expenses_sample/models/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

// User model
type User struct {
	Id        bson.ObjectId `bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	Email     string        `json:"email" bson:"email"`
}

// newUserCollection
func newUserCollection() *db.Collection {
	db := db.NewCollectionSession("users")
	
	db.Session.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
		Sparse: true})
	
	return db
}

// StartUserCollection...
func StartUserCollection(){
	db := newUserCollection()
	db.Close()
	return
}

// CreateUser...
func CreateUser(user User) (User, error) {
	var err error
	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// set default mongodb ID
	user.Id = bson.NewObjectId()

	// Insert user to mongodb
	err = c.Session.Insert(&user)
	if err != nil {
		return user, err
	}
	return user, err
}

// UpdateUser...
func (user User) UpdateUser(userParam User) (User, error) {
	var err error

	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// update post
	err = c.Session.Update(bson.M{
		"_id": user.Id,
	}, bson.M{
		"$set": bson.M{
			"firstName": userParam.FirstName,
			"email":     userParam.Email,
		},
	})

	if err != nil {
		return user, err
	}

	return user, err
}

// FindUsers...
func FindUsers() ([]User, error) {
	var err error
	var users []User

	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// get users
	err = c.Session.Find(nil).All(&users)

	if err != nil {
		return users, err
	}

	return users, err
}

// GetUserById...
func GetUserById(id bson.ObjectId) (User, error) {
	var err error
	var user User

	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// get post
	err = c.Session.FindId(id).One(&user)

	if err != nil {
		return user, err
	}

	return user, err
}

// GetUserByEmail...
func GetUserByEmail(email string) (User, error) {
	var err error
	var user User

	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// get user
	err = c.Session.Find(bson.M{"email": email}).One(&user)

	if err != nil {
		return user, err
	}

	return user, err
}

// DeleteUser...
func DeleteUser(user User) error {
	var err error

	// Get user collection connection
	c := newUserCollection()
	defer c.Close()

	// remove account
	err = c.Session.Remove(bson.M{"_id": user.Id})

	if err != nil {
		return err
	}

	return err
}
