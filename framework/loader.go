package framework

import (
	"log"
	"errors"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/database"
)

var (
	ParamLoader map[string] string
	UrlLoader map[string]restUrl
	FieldLoader map[int] []restField
)

func load() {
	ParamLoader, _ = loadParameter()
	UrlLoader, _ = loadUrl()
	FieldLoader, _ = loadField()
}

func loadParameter() (map[string] string, error) {
	h := map[string] string{}

	rows, err := database.DBCon.Query("SELECT a.key_, a.value_ FROM tbl_rest_param a")
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

	rows, err := database.DBCon.Query("SELECT a.url_id, a.url_name, a.routing_field, a.is_transaction FROM tbl_rest_url a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var r restUrl
		if err := rows.Scan(&r.Url_id, &r.Url_name, &r.Routing_field, &r.Is_transaction); err != nil {
			log.Panicln(err)
			return nil, errors.New(constant.ERR_ROWS_PARSING)
		}
		h[r.Url_name] = r
	}

	return h, nil
}

func loadField() (map[int] []restField, error) {
	h := map[int] []restField{}

	rows, err := database.DBCon.Query("SELECT a.url_id, a.field, a.min_length, a.max_length FROM tbl_rest_field a")
	if err != nil {
		log.Panicln(err)
		return nil, errors.New(constant.ERR_DATABASE)
	}

	defer rows.Close()

	for rows.Next() {
		var urlId int
		var r restField
		var list []restField
		if err := rows.Scan(&urlId, &r.Field, &r.Min_length, &r.Max_length); err != nil {
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
