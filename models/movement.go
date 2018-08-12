package models

import (
	"time"

	"github.com/josimar-jr/go_mongo_expenses_sample/models/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Movement model
type Movement struct {
	Id           bson.ObjectId `bson:"_id"`
	User         bson.ObjectId `json:"user_id" bson:"user_id"`
	MovementDate time.Time     `json:"movement_date" bson:"movement_date"`
	Title        string        `json:"title" bson:"title"`
	Value        float64       `json:"value" bson:"value"`
	Category     bson.ObjectId `json:"category_id" bson:"category_id"`
	AccountFrom  bson.ObjectId `json:"account_from" bson:"account_from"`
	MovementType string        `json:"movement_type" bson:"movement_type"`
	Description  string        `json:"description" bson:"description"`
}

// newMovementCollection
func newMovementCollection() *db.Collection {
	db := db.NewCollectionSession("movements")

	db.Session.EnsureIndex(mgo.Index{
		Key:    []string{"user_id", "movement_date"},
		Unique: false,
		Sparse: true})

	db.Session.EnsureIndex(mgo.Index{
		Key:    []string{"user_id", "account_from"},
		Unique: false,
		Sparse: true})

	db.Session.EnsureIndex(mgo.Index{
		Key:    []string{"user_id", "category_id"},
		Unique: false,
		Sparse: true})

	return db
}

// StartMovementCollection...
func StartMovementCollection() {
	db := newMovementCollection()
	db.Close()
	return
}

// CreateMovement...
func CreateMovement(movement Movement) (Movement, error) {
	var err error
	// Get Movement collection connection
	c := newMovementCollection()
	defer c.Close()

	// set default mongodb ID
	movement.Id = bson.NewObjectId()

	// Insert movement to mongodb
	err = c.Session.Insert(&movement)
	if err != nil {
		return movement, err
	}
	return movement, err
}

// UpdateMovement...
func (movement Movement) UpdateMovement(movParam Movement) (Movement, error) {
	var err error

	// Get Movement collection connection
	c := newMovementCollection()
	defer c.Close()

	// update Movement
	err = c.Session.Update(bson.M{
		"_id": movParam.Id,
	}, bson.M{
		"$set": bson.M{
			"user_id":       movParam.User,
			"movement_date": movParam.MovementDate,
			"title":         movParam.Title,
			"value":         movParam.Value,
			"category_id":   movParam.Category,
			"account_from":  movParam.AccountFrom,
			"description":   movParam.Description,
		},
	})

	if err != nil {
		return movement, err
	}

	return movement, err
}

// FindMovements...
func FindMovements() ([]Movement, error) {
	var err error
	var movements []Movement

	// Get Movement collection connection
	c := newMovementCollection()
	defer c.Close()

	// get movements
	err = c.Session.Find(nil).All(&movements)

	if err != nil {
		return movements, err
	}

	return movements, err
}

// GetMovementById...
func GetMovementById(id bson.ObjectId) (Movement, error) {
	var err error
	var movement Movement

	// Get Movement collection connection
	c := newMovementCollection()
	defer c.Close()

	// get post
	err = c.Session.FindId(id).One(&movement)

	if err != nil {
		return movement, err
	}

	return movement, err
}

// DeleteMovement...
func DeleteMovement(movement Movement) error {
	var err error

	// Get movement collection connection
	c := newMovementCollection()
	defer c.Close()

	// remove account
	err = c.Session.Remove(bson.M{"_id": movement.Id})

	if err != nil {
		return err
	}

	return err
}
