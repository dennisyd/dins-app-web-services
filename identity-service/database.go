package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

var (
	// DB is the package global db connection
	DB *mgo.Collection
)

// ConnectDB connects to mongodb
func ConnectDB() *mgo.Session {

	// connect to mongodb
	session, err := mgo.DialWithTimeout(os.Getenv("DB_URI"), 3*time.Second) // 3 second timeout
	if err != nil {
		log.Fatalf("Unable to open mongodb connection: %s", err)
	}

	// extract DB name from uri
	uriParts := strings.Split(os.Getenv("DB_URI"), "/")
	dbName := uriParts[len(uriParts)-1]

	// set DB to our db instance collection
	DB = session.DB(dbName).C("identities")

	// ensure our indexes exist and duplicates don't exist for indexed fields
	DB.EnsureIndex(mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true, // delete duplicate documents in case they somehow get put in
	})

	// return db session
	return session
}
