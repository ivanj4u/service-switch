package framework

import (
	"log"
	"net/http"
	"strings"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/dto"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request)

func responseApproved(data string) dto.Json {
	var r dto.Json

	r.ResponseCode = constant.CODE_APPROVED
	r.ResponseDesc = constant.DESC_APPROVED

	if data != "" {
		r.Data = data
	}

	return r
}

func routingByBranch(key string, isTransaction bool) string {
	log.Println("Routing by Branch")
	if strings.HasPrefix(key, constant.PREFIX_BRANCH_PUSAT_KONVEN) || strings.HasPrefix(key, constant.PREFIX_BRANCH_KONVEN) || strings.HasPrefix(key, constant.PREFIX_MIGRASI_KONVEN) {
		if isTransaction {
			return ParamLoader[constant.SERVER_KONVEN]
		} else {
			return ParamLoader[constant.SERVER_KONVEN_DRC]
		}
	} else if strings.HasPrefix(key, constant.PREFIX_BRANCH_PUSAT_SYAR) || strings.HasPrefix(key, constant.PREFIX_BRANCH_SYAR) || strings.HasPrefix(key, constant.PREFIX_MIGRASI_SYAR) {
		if isTransaction {
			return ParamLoader[constant.SERVER_SYARIAH]
		} else {
			return ParamLoader[constant.SERVER_SYARIAH_DRC]
		}
	} else if strings.HasPrefix(key, constant.PREFIX_BRANCH_PUSAT_GALERI) || strings.HasPrefix(key, constant.PREFIX_BRANCH_GALERI) {
		if isTransaction {
			return ParamLoader[constant.SERVER_GALERI]
		} else {
			return ParamLoader[constant.SERVER_GALERI_DRC]
		}
	}
	return ""
}

func routingByFlag(key string, isTransaction bool) string {
	log.Println("Routing by Flag")
	if key == constant.FLAG_KONVEN {
		if isTransaction {
			return ParamLoader[constant.SERVER_KONVEN]
		} else {
			return ParamLoader[constant.SERVER_KONVEN_DRC]
		}
	} else if key == constant.FLAG_SYAR {
		if isTransaction {
			return ParamLoader[constant.SERVER_SYARIAH]
		} else {
			return ParamLoader[constant.SERVER_SYARIAH_DRC]
		}
	} else if key == constant.FLAG_GALERI {
		if isTransaction {
			return ParamLoader[constant.SERVER_GALERI]
		} else {
			return ParamLoader[constant.SERVER_GALERI_DRC]
		}
	}
	return ""
}
