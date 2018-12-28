package framework

import (
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/dto"
)

func postCore(req dto.Json, url string) dto.Json {
	var res dto.Json
	log.Println("URL:>", url)

	req.RequestType = ""

	r, err := json.Marshal(req)
	if err != nil {
		res = generalError(constant.MSG_ERR_JSON_PARSING_REQ)
		return res
	}

	log.Println("Request Body :", string(r))

	coreRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(r))
	coreRequest.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(coreRequest)
	if err != nil {
		log.Println(err)
		return generalError(constant.MSG_ERR_POST_HTTP)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return generalError(constant.MSG_ERR_RES_BODY)
	}

	log.Println("Response Body :", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return generalError(constant.MSG_ERR_JSON_PARSING_RES)
	}

	return res
}
