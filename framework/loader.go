package framework

import (
	"log"
	"errors"

	"github.com/ivanj4u/service-switch/constant"
)

var (
	ParamLoader map[string] string
	UrlLoader map[string]restUrl
	FieldLoader map[string] []restField
	CARole map[string] map[string] bool
	CAProduct map[string] bool
)

func load() {
	ParamLoader, _ = loadParameter()
	UrlLoader, _ = loadUrl()
	FieldLoader, _ = loadField()
	CARole, _ = loadCARole()
	CAProduct, _ = loadCAProduct()
}

func loadParameter() (map[string] string, error) {
	h := map[string] string{}

	rows, err := DBCon.Query("SELECT a.key_, a.value_ FROM tbl_rest_param a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		h[key] = value
	}

	return h, nil
}

func loadUrl() (map[string]restUrl, error) {
	h := map[string]restUrl{}

	rows, err := DBCon.Query("SELECT a.url_id, a.url_name, a.amount_field, a.key_field, a.routing_field, a.is_transaction, a.is_reff_switching, a.request_type, a.is_existing FROM tbl_rest_url a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var r restUrl
		if err := rows.Scan(&r.urlId, &r.urlName, &r.amountField, &r.keyField, &r.routingField, &r.isTransaction, &r.isReffSwitching, &r.requestType, &r.isExisting); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		h[r.urlName] = r
	}

	return h, nil
}

func loadField() (map[string] []restField, error) {
	h := map[string] []restField{}

	rows, err := DBCon.Query("SELECT a.url_id, a.field, a.min_length, a.max_length FROM tbl_rest_field a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var urlId string
		var r restField
		var list []restField
		if err := rows.Scan(&urlId, &r.fieldName, &r.minLength, &r.maxLength); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		if h[urlId] != nil {
			list = h[urlId]
		} else {
			list = []restField{}
		}
		list = append(list, r)

		h[urlId] = list
	}

	return h, nil
}

func loadCARole() (map[string] map[string] bool, error) {
	h := map[string] map[string] bool {}

	rows, err := DBCon.Query("SELECT a.username, a.role, a.request_type FROM tbl_rest_ca_role a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var data map[string] bool
		var username, requestType, role string
		if err := rows.Scan(&username, &role, &requestType); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		if h[username] != nil {
			data = h[username]
		} else {
			data = map[string] bool{}
		}
		data[role+requestType] = true
		h[username] = data
	}

	return h, nil
}

func loadCAProduct() (map[string] bool, error) {
	h := map[string] bool{}

	rows, err := DBCon.Query("SELECT a.channel_id, a.client_id, a.product_code FROM tbl_rest_ca_product a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var channelId, clientId, productCode string
		if err := rows.Scan(&channelId, &clientId, &productCode); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		h[channelId + clientId + productCode] = true
	}

	return h, nil
}