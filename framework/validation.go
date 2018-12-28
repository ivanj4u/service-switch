/*
 * Copyright (c) 2018.
 */

package framework

import (
	"log"
	"errors"
	"strconv"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/database"
)

func validateField(urlName string, req map[string] string) (bool, restUrl, error) {
	url := UrlLoader[urlName]
	if url == (restUrl{}) {
		return false, url, errors.New(constant.MSG_ERR_URL_NOT_FOUND)
	}

	fields := FieldLoader[url.Url_id]

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		// Validate Null
		if req[field.Field] == "" {
			return true, url, errors.New("Field " + field.Field + " is Wrong")
		}

		length := len(req[field.Field])
		// Validate Min Length
		if length < field.Min_length {
			return true, url, errors.New("Field " + field.Field + " Must be Greater Than " + strconv.Itoa(field.Min_length))
		}
		if length > field.Max_length {
			return true, url, errors.New("Field " + field.Field + " Must be Less Than " + strconv.Itoa(field.Max_length))
		}
	}

	return true, url, nil
}

func validateRole(clientId, role, requestType string) (bool, error) {
	valid := true

	rows, err := database.DBCon.Query("SELECT * FROM tbl_rest_ca_role WHERE username = ? AND role = ? AND requestType = ?",
		clientId, role, requestType)
	if err != nil {
		log.Println(err.Error())
		return false, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	if rows == nil {
		valid = false
	}
	return valid, nil
}

func validateProduct(channelId, clientId, productCode string) (bool, error) {
	valid := true

	if productCode == "" {
		return valid, nil
	}

	rows, err := database.DBCon.Query("SELECT * FROM tbl_rest_ca_product " +
		"WHERE channel_id = ? AND client_id = ? AND productCode = ?", channelId, clientId, productCode)
	if err != nil {
		log.Println(err.Error())
		return false, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	if rows == nil {
		valid = false
	}

	return valid, nil
}