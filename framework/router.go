package framework

import (
	"bytes"
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ivanj4u/service-switch/constant"
	"github.com/ivanj4u/service-switch/database"
	"github.com/ivanj4u/service-switch/dto"
	"github.com/ivanj4u/service-switch/util"
	"github.com/globalsign/mgo/bson"
)

func handlerUniversal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Go Switching Services")
}

func reload(w http.ResponseWriter, r *http.Request) {
	load()
	responseService := responseApproved("")
	writeResponse(w, responseService)
	return
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var responseService, responseData dto.Json
	var param map[string] string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_RES_BODY)
		writeResponse(w, responseService)
		return
	}
	log.Println("Request Body :", string(body))

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_JSON_PARSING_RES)
		writeResponse(w, responseService)
		return
	}

	path := r.URL.Path

	err = json.Unmarshal(body, &param)
	if err != nil {
		responseService = generalError(constant.MSG_ERR_JSON_PARSING_RES)
		writeResponse(w, responseService)
		return
	}

	valid, restUrl, err := validateField(path, param)
	if err != nil {
		if valid {
			responseService = wrongFormat(err.Error())
		} else {
			responseService = generalError(err.Error())
		}
		writeResponse(w, responseService)
		return
	}
	if !valid {
		responseService = wrongFormat(constant.MSG_ERR_WRONG_FORMAT)
		writeResponse(w, responseService)
		return
	}

	jenisTransaksi, endPoint := getJenisTransaksi(path)

	valid, err = validateRole(responseData.ClientId, jenisTransaksi, responseData.RequestType)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}
	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	productCode := getProductCode(jenisTransaksi, responseData)
	valid, err = validateProduct(responseData.ChannelId, responseData.ClientId, productCode)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}

	if !valid {
		responseService = caNotRegistered(constant.MSG_ERR_CLIENT_UNAUTHORIZE)
		writeResponse(w, responseService)
		return
	}

	data := strings.Split(endPoint, "/")
	if len(data) > 0 {
		endPoint = data[0]
	}

	isParam := restUrl.Routing_field == "flag"
	key := param[restUrl.Routing_field]
	isTransaction := restUrl.Is_transaction == "1"
	isReffSwitching := restUrl.Is_reffswitching == "1"
	if isReffSwitching && responseData.ReffSwitching == "" {
		reffSwitching := responseData.ClientId + strconv.Itoa(util.GetCurrentTimeMilis())
		responseData.ReffSwitching = reffSwitching
	}

	// Create Old Request
	if restUrl.Is_Existing == "1" || (jenisTransaksi == constant.TABUNGAN && responseData.JenisTransaksi == "SL") {
		old, err := createOldJson(jenisTransaksi, responseData)
		if err != nil {
			log.Println(err.Error())
			responseService = generalError(err.Error())
			writeResponse(w, responseService)
			return
		}
		old, err = postOld(old, path, key, isParam, isTransaction)
		if err != nil {
			log.Println(err.Error())
			responseService = generalError(err.Error())
			writeResponse(w, responseService)
			return
		}

		responseService.ResponseCode = old.ResponseCode
		responseService.ResponseDesc = old.ResponseDesc

		if old.ResponseCode == constant.CODE_APPROVED {
			_ = json.Unmarshal([]byte(old.Data), &old)

			if jenisTransaksi == constant.GADAI {
				responseData.BiayaAdmBjpdl = old.BiayaAdmBjdpl
				responseData.BiayaLelang = old.BiayaLelang
				responseData.BiayaProsesLelang = old.BiayaProsesLelang
				responseData.Denda = fmt.Sprintf("%.0f", old.Denda)
				responseData.Golongan = old.GolonganRubrik
				responseData.JumlahHariReal = old.JumlahHariReal
				responseData.JumlahHariTarif = old.JumlahHariTarif
				responseData.MinimalUpCicil = fmt.Sprintf("%.0f", old.MinimalUpCicil)
				responseData.Norek = old.NoKredit
				responseData.SewaModal = fmt.Sprintf("%.0f", old.SewaModal)
				responseData.SewaModalBaru = fmt.Sprintf("%.0f", old.SewaModalBaru)
				responseData.Status = old.Status
				responseData.TaksiranBaru = fmt.Sprintf("%.0f", old.TaksiranBaru)
				responseData.Tenor = old.Tenor
				responseData.TglJatuhTempo = old.TglJatuhTempo
				responseData.TglKredit = old.TglKredit
				responseData.TglLelang = old.TglLelang
				responseData.Up = fmt.Sprintf("%.0f", old.Up)
				responseData.UpLama = fmt.Sprintf("%.0f", old.UpLama)
			} else if jenisTransaksi == constant.MIKRO {
				responseData.Angsuran = fmt.Sprintf("%.0f", old.Angsuran)
				responseData.AngsuranKe = strconv.Itoa(old.AngsuranKe)
				responseData.Denda = fmt.Sprintf("%.0f", old.Denda)
				responseData.JumlahAngsuran = strconv.Itoa(old.JumlahAngsuran)
				responseData.Norek = old.NoKredit
				responseData.NoRekeningPendamping = old.RekeningPendamping
				responseData.SaldoRekeningPendamping = fmt.Sprintf("%.0f", old.SaldoRekeningPendamping)
				responseData.SisaTenor = strconv.Itoa(old.SisaTenor)
				responseData.SewaModal = fmt.Sprintf("%.0f", old.SewaModal)
				responseData.Status = old.Status
				responseData.Tenor = old.Tenor
				responseData.TglJatuhTempo = old.TglJatuhTempo
				responseData.TglKredit = old.TglKredit
				responseData.Tunggakan = fmt.Sprintf("%.0f", old.Tunggakan)
				responseData.Up = fmt.Sprintf("%.0f", old.Up)
			} else if jenisTransaksi == constant.TABUNGAN {
				if old.BiayaTitip > 0 {
					responseData.BiayaTitip = fmt.Sprintf("%.0f", old.BiayaTitip)
				}
				responseData.Gram = fmt.Sprintf("%.4f", old.Gram)
				responseData.HakNasabah = fmt.Sprintf("%.0f", old.HakNasabah)
				responseData.Harga = fmt.Sprintf("%.0f", old.Harga)
				responseData.SaldoEmas = fmt.Sprintf("%.4f", old.SaldoEmas)
				responseData.SaldoNominal = fmt.Sprintf("%.0f", old.SaldoNominal)
				responseData.Satuan = fmt.Sprintf("%.4f", old.Satuan)
				responseData.TglBuka = old.TglBuka
			}
			responseData.Administrasi = fmt.Sprintf("%.0f", old.Administrasi)
			responseData.NamaNasabah = old.CustomerName
			responseData.NamaProduk = old.NamaProduk
			responseData.NilaiTransaksi = fmt.Sprintf("%.0f", old.NilaiTransaksi)
			responseData.JenisTransaksi = old.JenisTransaksi
			responseData.ReffBiller = old.ReffIdClient
			responseData.ReffCore = old.IdJurnal
			responseData.ReffSwitching = old.ReffIdSwitching
			responseData.Surcharge = fmt.Sprintf("%.0f", old.Charges)
			responseData.TglTransaksi = old.TglTransaksi
			responseData.TotalKewajiban = fmt.Sprintf("%.0f", old.TotalKewajiban)

			body, _ = json.Marshal(responseData)
			responseService.Data = string(body)
		}
	} else {
		responseService = post(responseData, path, key, isParam, isTransaction)
	}

	//Create Log
	err = saveLog(responseData)
	if err != nil {
		responseService = generalError(err.Error())
		writeResponse(w, responseService)
		return
	}
	log.Println("Response Body :", responseService)

	writeResponse(w, responseService)
}

func writeResponse(w http.ResponseWriter, r dto.Json) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r)

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, data, "", "  "); err != nil {
		log.Println(err.Error())
	}
	w.Write(dst.Bytes())
}

func getJenisTransaksi(url string) (string, string) {
	sizePath := len(url)
	url = url[1:sizePath]
	index := strings.Index(url, "/")
	return url[:index], url[index+1:]
}

func getProductCode(tipeTransaksi string, responseService dto.Json) string {
	productCode := ""

	if responseService.Norek != "" {
		return responseService.Norek[5:7]
	}
	if responseService.ProductCode != "" {
		return responseService.ProductCode
	}

	if responseService.JenisTransaksi == "OP" {
		if tipeTransaksi == constant.TABUNGAN {
			productCode = "62"
		}
	}

	return productCode
}

func createOldJson(jenisTransaksi string, req dto.Json) (dto.JsonOld, error) {
	old := dto.JsonOld{}
	var kewajiban float64

	old.ChannelId = req.ChannelId
	old.ClientId = req.ClientId
	old.JenisTransaksi = req.JenisTransaksi

	// Pembayaran
	old.PaymentMethod = req.PaymentMethod
	old.AgentId = req.AgentId
	old.WalletId = req.WalletId
	old.NorekWallet = req.NorekWallet
	old.KodeBankPembayar = req.KodeBankPembayar
	old.ReffIdSwitching = req.ReffSwitching

	amount := "0"
	if jenisTransaksi == constant.GADAI {
		if req.Amount != "" {
			amount = req.Amount
		}
		kewajiban, _ = strconv.ParseFloat(amount, 64)
		old.Kewajiban = kewajiban
		if req.MinimalUpCicil != "" {
			kewajiban, _ = strconv.ParseFloat(req.MinimalUpCicil, 64)
			if kewajiban > 0 {
				old.MinimalUpCicil = kewajiban
			} else {
				old.MinimalUpCicil = 0
			}
			old.NilaiTransaksi = kewajiban
		} else {
			old.NilaiTransaksi = old.Kewajiban
		}
		old.NoKredit = req.Norek
	} else if jenisTransaksi == constant.MIKRO {
		if req.Amount != "" {
			amount = req.Amount
		}
		kewajiban, _ := strconv.ParseFloat(amount, 64)
		old.Kewajiban = kewajiban
		old.NilaiTransaksi = old.Kewajiban
		old.NoKredit = req.Norek
	} else if jenisTransaksi == constant.TABUNGAN {
		old.Norek = req.Norek
		old.KodeCabang = req.KodeCabang

		if req.Amount != "" {
			kewajiban, _ = strconv.ParseFloat(req.Amount, 64)
		} else {
			kewajiban, _ = strconv.ParseFloat(req.NilaiTransaksi, 64)
		}
		old.NilaiTransaksi = kewajiban

		if req.Administrasi != "" {
			kewajiban, _ = strconv.ParseFloat(req.Administrasi, 64)
			log.Println("ADMINISTRASI", kewajiban)
			old.Admin = kewajiban
		}
		if req.Surcharge != "" {
			kewajiban, _ = strconv.ParseFloat(req.Surcharge, 64)
			old.Charges = kewajiban
		}
		if req.Gram != "" {
			kewajiban, _ = strconv.ParseFloat(req.Gram, 64)
			log.Println("GRAM", kewajiban)
			old.Gram = kewajiban
		}
		if req.Harga != "" {
			kewajiban, _ = strconv.ParseFloat(req.Harga, 64)
			old.Harga = kewajiban
		}
		if req.TotalKewajiban != "" {
			kewajiban, _ = strconv.ParseFloat(req.TotalKewajiban, 64)
			old.Kewajiban = kewajiban
		}

	}

	return old, nil
}

func saveLog(data dto.Json) (error) {

	col := database.MBCon.C("log_echannel")

	data.Id = string(bson.NewObjectId())
	data.Timestamp = time.Now()

	err := col.Insert(data)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Error cannot insert log ")
	}

	return nil
}