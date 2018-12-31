package framework

import (
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

func saveLog(data map[string] object) (error) {
	dbname := databaseProperties["mongo.dbname"]

	session := mongoSession.Clone()
	defer session.Close()

	col := session.DB(dbname).C("log_echannel")

	data["_id"] = string(bson.NewObjectId())
	data["timestamp"] = time.Now().Format("2006-01-02 15:04:05 ")

	err := col.Insert(data)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Error cannot insert log ")
	}

	return nil

}