package framework

import (
	"errors"
	"log"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/database"
	"github.com/ivanj4u/service-switch/dto"
	"time"
	"github.com/globalsign/mgo/bson"
)

func post(req dto.Json, path string, key string, isParam bool, isTransaction bool) dto.Json {
	log.Println("Start Services Post ")
	var url string

	if isParam {
		url = routingByFlag(key, isTransaction)
	} else {
		url = routingByBranch(key, isTransaction)
	}

	if url == "" {
		return generalError(constant.MSG_ERR_PARAM_NOT_FOUND + " " + key)
	}
	url = url + path

	res := postCore(req, url)

	log.Println("End Services Post ")
	return res
}

func saveLog(data dto.Json) (error) {

	col := database.MBCon.C("log_echannel")

	data.Id = string(bson.NewObjectId())
	data.Timestamp = time.Now()

	err := col.Insert(data)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Error cannot insert log ")
	}

	return nil
}