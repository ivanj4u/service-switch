/*
 * Copyright (c) 2018.
 */

package framework

import (
	"errors"
	"strconv"

	"github.com/ivanj4u/service-switch/constant"
	"strings"
	"log"
)

func validateField(urlName string, req map[string] object, isInquiry bool) (bool, restUrl, error) {
	url := UrlLoader[urlName]
	if isInquiry && url == (restUrl{}) {
		return false, url, errors.New(constant.MSG_ERR_URL_NOT_FOUND)
	}

	fields := FieldLoader[url.urlId]

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		// Validate Null
		if req[field.fieldName] == nil || req[field.fieldName].(string) == "" {
			return true, url, errors.New("fieldName " + field.fieldName + " is Wrong")
		}

		length := len(req[field.fieldName].(string))
		// Validate Min Length
		if length < field.minLength {
			return true, url, errors.New("fieldName " + field.fieldName + " Must be Greater Than " + strconv.Itoa(field.minLength))
		}
		if length > field.maxLength {
			return true, url, errors.New("fieldName " + field.fieldName + " Must be Less Than " + strconv.Itoa(field.maxLength))
		}
	}

	return true, url, nil
}

func validateRole(clientId, role, requestType string) bool {
	data := CARole[clientId]
	if data == nil || len(data) == 0 {
		return false
	}
	key := role + requestType
	return data[key]
}

func validateProduct(channelId, clientId, productCode string) bool {
	key := channelId + clientId + productCode
	log.Println("Key :", key)
	log.Println("Value :", CAProduct[key])
	return CAProduct[key]
}

func validatePaymentMethod(paymentMethod object) bool {
	if paymentMethod == nil || paymentMethod == "" {
		return false
	}

	str := ParamLoader["REST.PAYMENT.METHOD"]
	index := strings.Split(str, ",")

	valid := false
	for i := 0; i < len(index); i++  {
		if paymentMethod.(string) == index[i] {
			valid = true
			break
		}
	}
	return valid
}