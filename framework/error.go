package framework

import (
	"log"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/dto"
)

func httpNotFound(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_HTTP_NOT_FOUND
	if s == "" {
		res.ResponseDesc = constant.DESC_HTTP_NOT_FOUND
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func httpInternalError(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_HTTP_INTERNAL_ERROR
	if s == "" {
		res.ResponseDesc = constant.DESC_HTTP_INTERNAL_ERROR
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func invalidTransaction(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_INVALID_TRANSACTION
	if s == "" {
		res.ResponseDesc = constant.DESC_INVALID_TRANSACTION
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func invalidAmount(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_INVALID_AMOUNT
	if s == "" {
		res.ResponseDesc = constant.DESC_INVALID_AMOUNT
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func accountNotFound(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_ACCOUNT_NOT_FOUND
	if s == "" {
		res.ResponseDesc = constant.DESC_ACCOUNT_NOT_FOUND
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func invalidAccount(s string) (dto.Json)  {
	var res dto.Json
	res.ResponseCode = constant.CODE_INVALID_ACCOUNT
	if s == "" {
		res.ResponseDesc = constant.DESC_INVALID_ACCOUNT
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func wrongFormat(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_WRONG_FORMAT
	if s == "" {
		res.ResponseDesc = constant.DESC_WRONG_FORMAT
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func caNotRegistered(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_INVALID_CLIENT
	if s == "" {
		res.ResponseDesc = constant.DESC_INVALID_CLIENT
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func transactionTimeOut(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_TIMEOUT
	if s == "" {
		res.ResponseDesc = constant.DESC_TIMEOUT
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func billAlreadyPaid(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_INVALID_BILL
	if s == "" {
		res.ResponseDesc = constant.DESC_INVALID_BILL
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func cutOffTime(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_CUT_OFF_TIME
	if s == "" {
		s = constant.DESC_CUT_OFF_TIME
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func systemMaintenance(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_SYSTEM_MAINTENANCE
	if s == "" {
		res.ResponseDesc = constant.DESC_SYSTEM_MAINTENANCE
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}

func generalError(s string) (dto.Json) {
	var res dto.Json
	res.ResponseCode = constant.CODE_GENERAL_ERROR
	if s == "" {
		res.ResponseDesc = constant.DESC_GENERAL_ERROR
	}
	res.ResponseDesc = s
	log.Println(s)
	return res
}
