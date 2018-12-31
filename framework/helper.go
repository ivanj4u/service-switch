package framework

import (
	"strconv"
	"strings"
	"github.com/ivanj4u/service-switch/constant"
	"fmt"
)

func getJenisTransaksi(url string) (string, string) {
	sizePath := len(url)
	url = url[1:sizePath]
	index := strings.Index(url, "/")
	return url[:index], url[index+1:]
}

func getProductCode(tipeTransaksi string, req map[string] object) string {
	productCode := ""

	if req["norek"] != nil {
		return req["norek"].(string)[7:9]
	}
	if req["productCode"] != nil {
		return req["productCode"].(string)
	}

	if req["jenisTransaksi"] == "OP" {
		if tipeTransaksi == constant.TABUNGAN {
			productCode = "62"
		}
	}

	return productCode
}

func createOldJson(jenisTransaksi string, req map[string] object) (map[string] object) {
	old := map[string] object {}
	var kewajiban float64

	// Master Request
	old["channelId"] = req["channelId"]
	old["clientId"] = req["clientId"]
	old["jenisTransaksi"] = req["jenisTransaksi"]

	// Pembayaran
	if req["paymentMethod"] != nil {
		old["paymentMethod"] = req["paymentMethod"]
	}
	if req["agentId"] != nil {
		old["agentId"] = req["agentId"]
	}
	if req["walletId"] != nil {
		old["walletId"] = req["walletId"]
	}
	if req["norekWallet"] != nil {
		old["norekWallet"] = req["norekWallet"]
	}
	if req["kodeBankPembayar"] != nil {
		old["kodeBankPembayar"] = req["kodeBankPembayar"]
	}
	if req["reffSwitching"] != nil {
		old["reffIdSwitching"] = req["reffSwitching"]
	}

	amount := "0"
	if jenisTransaksi == constant.GADAI {
		if req["amount"] != nil {
			amount = req["amount"].(string)
		}
		kewajiban, _ = strconv.ParseFloat(amount, 64)
		old["kewajiban"] = kewajiban
		if req["minimalUpCicil"] != nil && req["minimalUpCicil"] != "0" {
			amount = req["minimalUpCicil"].(string)
			kewajiban, _ = strconv.ParseFloat(amount, 64)
			old["minimalUpCicil"] = kewajiban
		}
		old["nilaiTransaksi"] = kewajiban
		old["noKredit"] = req["norek"]
	} else if jenisTransaksi == constant.MIKRO {
		if req["amount"] != nil {
			amount = req["amount"].(string)
			kewajiban, _ = strconv.ParseFloat(amount, 64)
		}
		old["kewajiban"] = kewajiban
		old["nilaiTransaksi"] = kewajiban
		old["noKredit"] = req["norek"]
	} else if jenisTransaksi == constant.TABUNGAN {
		if req["norek"] != nil {
			old["norek"] = req["norek"]
		}

		if req["kodeCabang"] != nil {
			old["kodeCabang"] = req["kodeCabang"]
		}

		if req["amount"] != nil {
			amount = req["amount"].(string)
		} else {
			amount = req["nilaiTransaksi"].(string)
		}
		kewajiban, _ = strconv.ParseFloat(amount, 64)
		old["nilaiTransaksi"] = amount

		if req["administrasi"] != nil {
			kewajiban, _ = strconv.ParseFloat(req["administrasi"].(string), 64)
			old["admin"] = kewajiban
		}
		if req["surcharge"] != nil {
			kewajiban, _ = strconv.ParseFloat(req["surcharge"].(string), 64)
			old["charges"] = kewajiban
		}
		if req["gram"] != nil {
			kewajiban, _ = strconv.ParseFloat(req["gram"].(string), 64)
			old["gram"] = kewajiban
		}
		if req["harga"] != nil {
			kewajiban, _ = strconv.ParseFloat(req["harga"].(string), 64)
			old["norek"] = kewajiban
		}
		if req["totalKewajiban"] != nil {
			kewajiban, _ = strconv.ParseFloat(req["totalKewajiban"].(string), 64)
			old["kewajiban"] = kewajiban
		}
	}
	return old
}

func parseOldResponse(jenisTransaksi string, old, responseData map[string] object) map[string] object {
	if jenisTransaksi == constant.GADAI {
		if old["biayaAdmBjpdl"] != nil {
			responseData["biayaAdmBjpdl"] = old["biayaAdmBjpdl"]
		}
		if old["biayaLelang"] != nil {
			responseData["biayaLelang"] = old["biayaLelang"]
		}
		if old["biayaProsesLelang"] != nil {
			responseData["biayaProsesLelang"] = old["biayaProsesLelang"]
		}
		responseData["denda"] = fmt.Sprintf("%.0f", old["denda"].(float64))
		responseData["golongan"] = old["golonganRubrik"]
		responseData["jumlahHariReal"]= old["jumlahHariReal"]
		responseData["jumlahHariTarif"] = old["jumlahHariTarif"]
		responseData["minimalUpCicil"] = fmt.Sprintf("%.0f", old["minimalUpCicil"].(float64))
		responseData["norek"] = old["noKredit"]
		responseData["sewaModal"] = fmt.Sprintf("%.0f", old["sewaModal"].(float64))
		if old["sewaModalBaru"] != nil {
			responseData["sewaModalBaru"] = fmt.Sprintf("%.0f", old["sewaModalBaru"].(float64))
		}
		if old["status"] != nil {
			responseData["status"] = old["status"]
		}
		if old["taksiranBaru"] != nil {
			responseData["taksiranBaru"] = fmt.Sprintf("%.0f", old["taksiranBaru"].(float64))
		}
		responseData["tenor"] = old["tenor"]
		responseData["tglJatuhTempo"] = old["tglJatuhTempo"]
		responseData["tglKredit"] = old["tglKredit"]
		responseData["tglLelang"] = old["tglLelang"]
		responseData["up"] = fmt.Sprintf("%.0f", old["up"].(float64))
		if old["upLama"] != nil {
			responseData["upLama"] = fmt.Sprintf("%.0f", old["upLama"].(float64))
		}
	} else if jenisTransaksi == constant.MIKRO {
		responseData["angsuran"] = fmt.Sprintf("%.0f", old["angsuran"].(float64))
		responseData["angsuranKe"] = old["angsuranKe"]
		responseData["denda"] = fmt.Sprintf("%.0f", old["denda"].(float64))
		responseData["jumlahAngsuran"] = old["jumlahAngsuran"]
		responseData["norek"] = old["noKredit"]
		responseData["noRekeningPendamping"] = old["rekeningPendamping"]
		responseData["saldoRekeningPendamping"] = fmt.Sprintf("%.0f", old["saldoRekeningPendamping"].(float64))
		responseData["sisaTenor"] = old["sisaTenor"]
		responseData["sewaModal"] = fmt.Sprintf("%.0f", old["sewaModal"].(float64))
		if old["status"] != nil {
			responseData["status"] = old["status"]
		}
		responseData["tenor"] = old["tenor"]
		responseData["tglJatuhTempo"] = old["tglJatuhTempo"]
		responseData["tglKredit"] = old["tglKredit"]
		responseData["tunggakan"] = fmt.Sprintf("%.0f", old["tunggakan"].(float64))
		responseData["up"] = fmt.Sprintf("%.0f", old["up"].(float64))
	} else if jenisTransaksi == constant.TABUNGAN {
		responseData["biayaTitip"] = fmt.Sprintf("%.0f", old["biayaTitip"].(float64))
		responseData["gram"] = fmt.Sprintf("%.4f", old["gram"].(float64))
		responseData["hakNasabah"]= fmt.Sprintf("%.0f", old["hakNasabah"].(float64))
		responseData["harga"] = fmt.Sprintf("%.0f", old["harga"].(float64))
		responseData["saldoEmas"] = fmt.Sprintf("%.4f", old["saldoEmas"].(float64))
		responseData["saldoNominal"] = fmt.Sprintf("%.0f", old["saldoNominal"].(float64))
		responseData["satuan"] = fmt.Sprintf("%.4f", old["satuan"].(float64))
		responseData["tglBuka"] = old["tglBuka"]
	}
	responseData["administrasi"] = fmt.Sprintf("%.0f", old["administrasi"].(float64))
	responseData["namaNasabah"] = old["customerName"]
	responseData["namaProduk"] = old["namaProduk"]
	responseData["nilaiTransaksi"] = fmt.Sprintf("%.0f", old["nilaiTransaksi"].(float64))
	responseData["jenisTransaksi"] = old["jenisTransaksi"]

	if old["reffIdClient"] != nil {
		responseData["reffIdClient"] = old["reffIdClient"]
	}
	if old["idJurnal"] != nil {
		responseData["idJurnal"] = old["idJurnal"]
	}
	if old["reffIdSwitching"] != nil {
		responseData["reffIdSwitching"] = old["reffIdSwitching"]
	}
	responseData["charges"] = fmt.Sprintf("%.0f", old["charges"].(float64))
	responseData["tglTransaksi"] = old["tglTransaksi"]
	responseData["totalKewajiban"] = fmt.Sprintf("%.0f", old["totalKewajiban"].(float64))

	return responseData
}