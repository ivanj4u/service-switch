package framework

import (
	"log"
	"database/sql"

	"github.com/globalsign/mgo"
)

func openDatabaseConnection() (*sql.DB) {
	log.Println("Start Open Connection to Database")
	var db *sql.DB
	var err error

	driver := databaseProperties["db.driver"]
	schema := databaseProperties["db.dbname"]
	host := databaseProperties["db.url"]
	port := databaseProperties["db.port"]
	user := databaseProperties["db.user"]
	pass := databaseProperties["db.pass"]

	if driver == "mysql" {
		dbSourceName := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + schema
		log.Println("datasource :", dbSourceName)

		db, err = sql.Open(driver, dbSourceName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("MySQL Driver only!")
	}

	log.Println("Success Open Connection to Database")

	return db
}

func openMongoDBConnection() (*mgo.Database) {
	log.Println("Start Open Connection to MongoDB")

	url := databaseProperties["mongo.url"]
	port := databaseProperties["mongo.port"]
	dbname := databaseProperties["mongo.dbname"]

	mongoDial := "mongodb://"+url+":"+port
	log.Println("mongoDial :", mongoDial)

	session, err := mgo.Dial(mongoDial)
	if err != nil {
		log.Fatal(err)
	}

	db := session.DB(dbname)

	log.Println("Success Open Connection to MongoDB")
	return db
}
