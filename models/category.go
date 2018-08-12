package models

import (
	"github.com/josimar-jr/go_mongo_expenses_sample/models/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

// Category model
type Category struct {
	Id             bson.ObjectId `bson:"_id"`
	Description    string        `json:"description" bson:"description"`
	User           bson.ObjectId `json:"user_id" bson:"user_id"`
	ParentCategory *bson.ObjectId `json:"parent_category" bson:"parent_category"`
}

// newCategoryCollection
func newCategoryCollection() *db.Collection {
	db := db.NewCollectionSession("categories")
	
	db.Session.EnsureIndex( mgo.Index{
		Key:    []string{"user_id","description"},
		Unique: true,
		Sparse: true})

	db.Session.EnsureIndex( mgo.Index{
		Key:    []string{"user_id","parent_category","description"},
		Unique: true,
		Sparse: true})

	return db
}

// StartCategoryCollection...
func StartCategoryCollection(){
	db := newCategoryCollection()
	db.Close()
	return
}

// CreateCategory...
func CreateCategory(category Category) (Category, error) {
	var err error
	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// set default mongodb ID
	category.Id = bson.NewObjectId()

	// Insert category to mongodb
	err = c.Session.Insert(&category)
	if err != nil {
		return category, err
	}
	return category, err
}

// UpdateCategory...
func (category Category) UpdateCategory(categoryParam Category) (Category, error) {
	var err error

	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// update post
	err = c.Session.Update(bson.M{
		"_id": category.Id,
	}, bson.M{
		"$set": bson.M{
			"description":     categoryParam.Description,
			"user_id":         categoryParam.User,
			"parent_category": categoryParam.ParentCategory,
		},
	})

	if err != nil {
		return category, err
	}

	return category, err
}

// FindCategories...
func FindCategories() ([]Category, error) {
	var err error
	var categories []Category

	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// get categories
	err = c.Session.Find(nil).All(&categories)

	if err != nil {
		return categories, err
	}

	return categories, err
}

// GetCategoryById...
func GetCategoryById(id bson.ObjectId) (Category, error) {
	var err error
	var category Category

	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// get post
	err = c.Session.FindId(id).One(&category)

	if err != nil {
		return category, err
	}

	return category, err
}

// GetCategoryByDescription...
func GetCategoryByDescription(description string) (Category, error) {
	var err error
	var category Category

	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// get category
	err = c.Session.Find(bson.M{"description": description}).One(&category)

	if err != nil {
		return category, err
	}

	return category, err
}

// DeleteCategory...
func DeleteCategory(category Category) error {
	var err error

	// Get category collection connection
	c := newCategoryCollection()
	defer c.Close()

	// remove account
	err = c.Session.Remove(bson.M{"_id": category.Id})

	if err != nil {
		return err
	}

	return err
}
