package framework

import (
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ivanj4u/service-switch/constant"
	"strconv"
)

func saveLog(data map[string] object, requestType string) (error) {
	dbname := databaseProperties["mongo.dbname"]

	session := mongoSession.Clone()
	defer session.Close()

	col := session.DB(dbname).C("log_echannel")

	data["_id"] = string(bson.NewObjectId())
	data["timestamp"] = time.Now().Format("2006-01-02 15:04:05 ")
	data["requestType"] = requestType

	err := col.Insert(data)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Error cannot insert log")
	}

	return nil

}

func validateLog(url restUrl, req map[string] object, productCode string) (error) {
	dbname := databaseProperties["mongo.dbname"]

	session := mongoSession.Clone()
	defer session.Close()

	col := session.DB(dbname).C("log_echannel")

	amountReq := ""
	amountRes := ""
	if url.amountReqField.Valid {
		amountReq = url.amountReqField.String
		amountRes = url.amountResField.String
	}

	key := ""
	if url.keyField.Valid {
		key = url.keyField.String
	}

	clientId := req["clientId"].(string)
	jenisTransaksi := req["jenisTransaksi"].(string)
	reffSwitching := req["reffSwitching"].(string)

	log.Println("Validating Log :", clientId, jenisTransaksi, reffSwitching, constant.REQ_INQUIRY)

	result := map[string] string {}
	query := col.Find(bson.M{
		"clientId": clientId,
		"jenisTransaksi": jenisTransaksi,
		"reffSwitching": reffSwitching,
		"requestType": constant.REQ_INQUIRY,
	})

	var err error
	if key != "" {
		err = query.Select(bson.M{
		key : 1,
		amountRes : 1,
		}).One(&result)

	} else {
		err = query.Select(bson.M{
		amountRes : 1,
		}).One(&result)
	}
	if err != nil {
		log.Println(err.Error())
		return errors.New("ReffSwitching : " + clientId + "-" + jenisTransaksi + "-" + reffSwitching + "-" + constant.REQ_INQUIRY + " not Found")
	}

	if key != "" && req[key] != result[key] {
		log.Println(key, "ReffSwitching tidak sesuai", req[key], result[key])
		return errors.New("Key ReffSwitching tidak sesuai")
	}

	// Validating Amount with Surcharges (Khusus paymentMethod BANK)
	amount := req[amountReq].(string)

	paymentMethod := req["paymentMethod"]
	if paymentMethod != nil && paymentMethod.(string) == "BANK" {
		key := req["channelId"].(string) + clientId + jenisTransaksi + productCode + paymentMethod.(string) + req["kodeBankPembayar"].(string)
		surcharge := SurchargeLoader[key]

		if surcharge != "" {
			charge, _ := strconv.Atoi(surcharge)
			value, _ := strconv.Atoi(amount)

			value = value - charge
			amount = strconv.Itoa(value)
		} else {
			log.Println("Surcharge :", key, "not found")
			return errors.New("Surcharge not found")
		}
	}
	if amount != result[amountRes] {
		log.Println(amountRes, "ReffSwitching tidak sesuai", amount, result[amountRes])
		return errors.New("Amount ReffSwitching tidak sesuai")
	}

	return nil
}