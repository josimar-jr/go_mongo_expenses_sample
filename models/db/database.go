package db

import "gopkg.in/mgo.v2"

// Database...
type Database struct {
	s       *mgo.Session
	name    string
	session *mgo.Database
}

// Connect...
func (db *Database) Connect() {
	db.s = service.Session()
	session := *db.s.DB(db.name)
	db.session = &session
}

// newDBSession...
func newDBSession(name string) *Database {
	var db = Database{
		name: name,
	}
	db.Connect()
	return &db
}
