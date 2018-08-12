package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

// Collection
type Collection struct {
	db      *Database
	name    string
	Session *mgo.Collection
}

// type Index struct {
// 	Key        []string // Index key fields; prefix name with dash (-) for descending order
// 	Unique     bool     // Prevent two documents from having the same index key
// 	Background bool     // Build index in background and return immediately
// 	Sparse     bool     // Only index documents containing the Key fields

// 	Name string // Index name, computed by EnsureIndex
// }

// Connect
func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	c.Session = &session
}

// NewCollectionSession
func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db:   newDBSession(beego.AppConfig.String("DBName")),
		name: name,
	}
	c.Connect()
	return &c
}

// Close
func (c *Collection) Close() {
	service.Close(c)
}
