package models

import (
	"github.com/josimar-jr/go_mongo_expenses_sample/models/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

// Account model
type Account struct {
	Id      bson.ObjectId `bson:"_id"`
	Account string        `json:"acocunt" bson:"account"`
	User    bson.ObjectId `json:"user_id" bson:"user_id"`
}

// newAccountCollection
func newAccountCollection() *db.Collection {
	db := db.NewCollectionSession("accounts")

	db.Session.EnsureIndex(mgo.Index{
		Key:    []string{"user_id","account"},
		Unique: true,
		Sparse: true})
	
	return db
}

// StartAccountCollection...
func StartAccountCollection(){
	db := newAccountCollection()
	db.Close()
	return
}

// CreateAccount...
func CreateAccount(account Account) (Account, error) {
	var err error
	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// set default mongodb ID
	account.Id = bson.NewObjectId()

	// Insert account to mongodb
	err = c.Session.Insert(&account)
	if err != nil {
		return account, err
	}
	return account, err
}

// UpdateAccount...
func (account Account) UpdateAccount(accountParam Account) (Account, error) {
	var err error

	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// update post
	err = c.Session.Update(bson.M{
		"_id": account.Id,
	}, bson.M{
		"$set": bson.M{
			"account": accountParam.Account,
			"user_id": accountParam.User,
		},
	})

	if err != nil {
		return account, err
	}

	return account, err
}

// FindAccount...
func FindAccount() ([]Account, error) {
	var err error
	var accounts []Account

	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// get account
	err = c.Session.Find(nil).All(&accounts)

	if err != nil {
		return accounts, err
	}

	return accounts, err
}

// GetAccountById...
func GetAccountById(id bson.ObjectId) (Account, error) {
	var err error
	var account Account

	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// get post
	err = c.Session.FindId(id).One(&account)

	if err != nil {
		return account, err
	}

	return account, err
}

// GetAccountByDescription...
func GetAccountByDescription(description string) (Account, error) {
	var err error
	var account Account

	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// get account
	err = c.Session.Find(bson.M{"account": description}).One(&account)

	if err != nil {
		return account, err
	}

	return account, err
}

// DeleteAccount...
func DeleteAccount(account Account) error {
	var err error

	// Get account collection connection
	c := newAccountCollection()
	defer c.Close()

	// remove account
	err = c.Session.Remove(bson.M{"_id": account.Id})

	if err != nil {
		return err
	}

	return err
}
