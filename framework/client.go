package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/dto"
)

func post(req map[string] object, path string, key string, isParam bool, isTransaction bool) dto.Json {
	log.Println("Start Services Post ")
	var url string
	var res dto.Json

	if isParam {
		url = routingByFlag(key, isTransaction)
	} else {
		url = routingByBranch(key, isTransaction)
	}

	if url == "" {
		return generalError(constant.MSG_ERR_PARAM_NOT_FOUND + " " + key)
	}

	url = url + path

	log.Println("URL:>", url)

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

	log.Println("End Services Post ")
	return res
}

func postOld(req map[string] object, path string, key string, isFlag bool, isTransaction bool) (map[string] object, error) {
	log.Println("Start Services Post ")
	var url string
	var res map[string] object

	if isFlag {
		url = routingByFlag(key, isTransaction)
	} else {
		url = routingByBranch(key, isTransaction)
	}

	if url == "" {
		return res, errors.New(constant.MSG_ERR_PARAM_NOT_FOUND + " " + key)
	}
	url = url + path

	log.Println("URL:>", url)

	r, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return res, errors.New(constant.MSG_ERR_JSON_PARSING_REQ)
	}

	log.Println("Request Body Core :", string(r))

	coreRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(r))
	coreRequest.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(coreRequest)
	if err != nil {
		log.Println(err)
		return res, errors.New(constant.MSG_ERR_POST_HTTP)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, errors.New(constant.MSG_ERR_RES_BODY)
	}

	log.Println("Response Body Core :", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
		return res, errors.New(constant.MSG_ERR_JSON_PARSING_RES)
	}

	log.Println("End Services Post ")
	return res, nil
}

func checkToken(token string) error {
	urlCheckToken := applicationProperties["server.oauth.url"] + "/" + "check_token" + "?token=" + token

	log.Println(urlCheckToken)

	req, _ := http.NewRequest(http.MethodPost, urlCheckToken, nil)
	req.Header.Add("Authorization", "Basic YXBsaWthc2lqczphcGxpa2FzaTEyMw==")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Println(err)
		return errors.New(constant.MSG_ERR_POST_HTTP + " OAuth Server")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Unauthorize - " + resp.Status)
	}
	return nil
}