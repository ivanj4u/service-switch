package framework

import "database/sql"

type object interface{}

type restUrl struct {
	urlId           string
	isTransaction   string
	isReffSwitching sql.NullString
	amountReqField  sql.NullString
	amountResField  sql.NullString
	keyField        sql.NullString
	routingField    string
	requestType     sql.NullString
	urlName         string
	isExisting      sql.NullString
}

type restField struct {
	fieldName string
	minLength int
	maxLength int
}

type restSurcharge struct {
	channelId string
	clientId string
	jenisTransaksi string
	productCode string
	paymentMethod sql.NullString
	kodeBank sql.NullString
	amount float64
}