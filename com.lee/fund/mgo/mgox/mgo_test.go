package mgox

import (
	"testing"

	"com.lee/fund/config"
	"com.lee/fund/mgo/bson"
)

/********** Initialize **********/

func openDB(t *testing.T) *Database {
	config.ConfigDir = "."
	db, err := Open("Movie")
	if err != nil {
		t.Fatal(err)
	}

	return db
}

/********** Test Method **********/

func TestFind(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	result := make(map[string]interface{})
	err := db.Coll("MovieCount").Find(bson.M{"_id": 74830}).One(&result)
	if err != nil {
		t.Fatal(err)
	}

	// t.Fatal(result)
}
