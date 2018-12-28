package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var responseService dto.Json
	var param map[string] string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_RES_BODY)
		writeResponse(w, responseService)
		return
	}
	err = json.Unmarshal(body, &responseService)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_JSON_PARSING_RES)
		writeResponse(w, responseService)
		return
	}

	path := r.URL.Path

	err = json.Unmarshal(body, &param)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_JSON_PARSING_RES)
		writeResponse(w, responseService)
		return
	}

	valid, restUrl, err := validateField(path, param)
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

	jenisTransaksi, endPoint := getJenisTransaksi(path)

	valid, err = validateRole(responseService.ClientId, jenisTransaksi, responseService.RequestType)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}
	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	productCode := getProductCode(jenisTransaksi, responseService)
	valid, err = validateProduct(responseService.ChannelId, responseService.ClientId, productCode)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}

	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	data := strings.Split(endPoint, "/")
	if len(data) > 0 {
		endPoint = data[0]
	}

	isParam := restUrl.Routing_field == "param"
	key := param[restUrl.Routing_field]
	isTransaction := restUrl.Is_transaction == "1"
	if isTransaction {
		reffSwitching := responseService.ClientId + strconv.Itoa(util.GetCurrentTimeMilis())
		responseService.ReffSwitching = reffSwitching
	}

	//Create Log
	err = saveLog(responseService)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}

	responseService = post(responseService, path, key, isParam, isTransaction)

	writeResponse(w, responseService)
}

func writeResponse(w http.ResponseWriter, r dto.Json) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r)
	w.Write(data)
}

func getJenisTransaksi(url string) (string, string) {
	sizePath := len(url)
	url = url[1:sizePath]
	index := strings.Index(url, "/")
	return url[:index], url[index+1:]
}

func getProductCode(tipeTransaksi string, responseService dto.Json) string {
	productCode := ""

	if responseService.Norek != "" {
		return responseService.Norek[5:7]
	}
	if responseService.ProductCode != "" {
		return responseService.ProductCode
	}

	if responseService.JenisTransaksi == "OP" {
		if tipeTransaksi == "tabunganemas" {
			productCode = "62"
		}
	}

	return productCode
}