package mgox

import (
	"com.lee/fund/mgo"
)

type Database struct {
	db *mgo.Database
}

func (d *Database) Close() {
	d.db.Session.Close()
}

func (d *Database) Coll(name string) *mgo.Collection {
	return d.db.C(name)
}

func (d *Database) GridFS(prefix string) *mgo.GridFS {
	return d.db.GridFS(prefix)
}

func (d *Database) Run(cmd, result interface{}) error {
	return d.db.Run(cmd, result)
}
