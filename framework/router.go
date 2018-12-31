package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/dto"
	"github.com/ivanj4u/service-switch/util"
)

func handlerUniversal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Go Switching Services")
}

func reload(w http.ResponseWriter, r *http.Request) {
	load()
	responseService := responseApproved("")
	writeResponse(w, responseService)
	return
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var key string
	var responseService dto.Json
	var responseData map[string] object

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_RES_BODY)
		writeResponse(w, responseService)
		return
	}
	log.Println("Request Body :", string(body))

	path := r.URL.Path

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_JSON_PARSING_RES)
		writeResponse(w, responseService)
		return
	}

	valid, restUrl, err := validateField(path, responseData, true)
	if err != nil {
		if valid {
			responseService = wrongFormat(err.Error())
		} else {
			responseService = generalError(err.Error())
		}
		writeResponse(w, responseService)
		return
	}
	if !valid {
		responseService = wrongFormat(constant.MSG_ERR_WRONG_FORMAT)
		writeResponse(w, responseService)
		return
	}

	requestType := ""
	if restUrl.requestType.Valid {
		requestType = restUrl.requestType.String
	}

	if requestType != "" && requestType != "INQUIRY" {
		valid = validatePaymentMethod(responseData["paymentMethod"])

		if !valid {
			responseService = wrongFormat(constant.MSG_ERR_WRONG_FORMAT + " paymentMethod")
			writeResponse(w, responseService)
			return
		}
		valid, _, err = validateField(responseData["paymentMethod"].(string), responseData, false)
		if err != nil {
			if valid {
				responseService = wrongFormat(err.Error())
			} else {
				responseService = generalError(err.Error())
			}
			writeResponse(w, responseService)
			return
		}

	}

	jenisTransaksi, endPoint := getJenisTransaksi(path)

	valid = validateRole(responseData["clientId"].(string), jenisTransaksi, requestType)
	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_ROLE_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	productCode := getProductCode(jenisTransaksi, responseData)
	valid = validateProduct(responseData["channelId"].(string), responseData["clientId"].(string), productCode)
	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_PRODUCT_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	data := strings.Split(endPoint, "/")
	if len(data) > 0 {
		endPoint = data[0]
	}

	isFlag := restUrl.routingField == "flag"

	// Tabungan OPEN
	if productCode == constant.PRODUCT_CODE_TABUNGAN && restUrl.requestType.String == constant.REQ_INQUIRY && responseData["jenisTransaksi"] == "OP" {
		key = responseData["flag"].(string)
		isFlag = true
	} else {
		key = responseData[restUrl.routingField].(string)
	}

	isTransaction := restUrl.isTransaction == "1"

	isReffSwitching := false
	if restUrl.isReffSwitching.Valid && restUrl.isReffSwitching.String == "1" {
		isReffSwitching = true
	}

	if isReffSwitching && (responseData["reffSwitching"] == nil || responseData["reffSwitching"].(string) == "") {
		reffSwitching := responseData["clientId"].(string) + strconv.Itoa(util.GetCurrentTimeMilis())
		responseData["reffSwitching"] = reffSwitching
	}

	if requestType != "" && requestType != constant.REQ_CREATE && requestType != constant.REQ_INQUIRY {
		err = validateLog(restUrl, responseData)
		if err != nil {
			responseService = invalidTransaction(err.Error())
			writeResponse(w, responseService)
			return
		}
	}

	// Create Old Request
	if (restUrl.isExisting.Valid && restUrl.isExisting.String == "1") || (jenisTransaksi == constant.TABUNGAN && requestType == constant.REQ_INQUIRY) {
		old:= createOldJson(jenisTransaksi, responseData)

		old, err = postOld(old, path, key, isFlag, isTransaction)
		if err != nil {
			responseService = generalError(err.Error())
			writeResponse(w, responseService)
			return
		}

		responseService.ResponseCode = old["responseCode"].(string)
		responseService.ResponseDesc = old["responseDesc"].(string)

		if old["responseCode"] == constant.CODE_APPROVED {
			_ = json.Unmarshal([]byte(old["data"].(string)), &old)
			responseData = parseOldResponse(jenisTransaksi, old, responseData)
			body, _ = json.Marshal(responseData)
			responseService.Data = string(body)
		}
	} else {
		responseService = post(responseData, path, key, isFlag, isTransaction)
		_ = json.Unmarshal([]byte(responseService.Data), &responseData)
	}

	//Create Log
	err = saveLog(responseData, requestType)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}
	log.Println("Response Body :", responseService)

	writeResponse(w, responseService)
}

func writeResponse(w http.ResponseWriter, r dto.Json) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r)

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, data, "", "  "); err != nil {
		log.Println(err.Error())
	}
	w.Write(dst.Bytes())
}