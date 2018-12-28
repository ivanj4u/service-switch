package dto

import "time"

type Json struct {
	Administrasi string `bson:"administrasi,omitempty" json:"administrasi,omitempty"`
	AgentId string `bson:"agentId,omitempty" json:"agentId,omitempty"`
	Amount string `bson:"amount,omitempty" json:"amount,omitempty"`
	Angsuran string `bson:"angsuran,omitempty" json:"angsuran,omitempty"`
	AngsuranKe string `bson:"angsuranKe,omitempty" json:"angsuranKe,omitempty"`
	BiayaAdmBjpdl string `bson:"biayaAdmBjpdl,omitempty" json:"biayaAdmBjpdl,omitempty"`
	BiayaLelang string `bson:"biayaLelang,omitempty" json:"biayaLelang,omitempty"`
	BiayaProsesLelang string `bson:"biayaProsesLelang,omitempty" json:"biayaProsesLelang,omitempty"`
	BiayaTitip string `bson:"biayaTitip,omitempty" json:"biayaTitip,omitempty"`
	ChannelId string `bson:"channelId,omitempty" json:"channelId,omitempty"`
	Cif string `bson:"cif,omitempty" json:"cif,omitempty"`
	ClientId string `bson:"clientId,omitempty" json:"clientId,omitempty"`
	Data string `bson:"data,omitempty" json:"data,omitempty"`
	Denda string `bson:"denda,omitempty" json:"denda,omitempty"`
	Flag string `bson:"flag,omitempty" json:"flag,omitempty"`
	Gram string `bson:"gram,omitempty" json:"gram,omitempty"`
	GramTransaksi string `bson:"gramTransaksi,omitempty" json:"gramTransaksi,omitempty"`
	Golongan string `bson:"golongan,omitempty" json:"golongan,omitempty"`
	HakNasabah string `bson:"hakNasabah,omitempty" json:"hakNasabah,omitempty"`
	Harga string `bson:"harga,omitempty" json:"harga,omitempty"`
	IbuKandung string `bson:"ibuKandung,omitempty" json:"ibuKandung,omitempty"`
	IdKelurahan string `bson:"idKelurahan,omitempty" json:"idKelurahan,omitempty"`
	Jalan string `bson:"jalan,omitempty" json:"jalan,omitempty"`
	JenisKelamin string `bson:"jenisKelamin,omitempty" json:"jenisKelamin,omitempty"`
	JenisTransaksi string `bson:"jenisTransaksi,omitempty" json:"jenisTransaksi,omitempty"`
	JumlahAngsuran string `bson:"jumlahAngsuran,omitempty" json:"jumlahAngsuran,omitempty"`
	JumlahHariTarif string `bson:"jumlahHariTarif,omitempty" json:"jumlahHariTarif,omitempty"`
	JumlahHariReal string `bson:"jumlahHariReal,omitempty" json:"jumlahHariReal,omitempty"`
	Kewarganegaraan string `bson:"kewarganegaraan,omitempty" json:"kewarganegaraan,omitempty"`
	KodeBankPembayar string `bson:"kodeBankPembayar,omitempty" json:"kodeBankPembayar,omitempty"`
	KodeBankTujuan string `bson:"kodeBankTujuan,omitempty" json:"kodeBankTujuan,omitempty"`
	KodeCabang string `bson:"kodeCabang,omitempty" json:"kodeCabang,omitempty"`
	KodeNamaCabang string `bson:"kodeNamaCabang,omitempty" json:"kodeNamaCabang,omitempty"`
	MinimalUpCicil string `bson:"minimalUpCicil,omitempty" json:"minimalUpCicil,omitempty"`
	NamaBankTujuan string `bson:"namaBankTujuan,omitempty" json:"namaBankTujuan,omitempty"`
	NamaNasabah string `bson:"namaNasabah,omitempty" json:"namaNasabah,omitempty"`
	NamaProduk string `bson:"namaProduk,omitempty" json:"namaProduk,omitempty"`
	NilaiTransaksi string `bson:"nilaiTransaksi,omitempty" json:"nilaiTransaksi,omitempty"`
	NoHp string `bson:"noHp,omitempty" json:"noHp,omitempty"`
	NoHpAgent string `bson:"noHpAgent,omitempty" json:"noHpAgent,omitempty"`
	NoIdentitas string `bson:"noIdentitas,omitempty" json:"noIdentitas,omitempty"`
	Norek string `bson:"norek,omitempty" json:"norek,omitempty"`
	NorekBankTujuan string `bson:"norekBankTujuan,omitempty" json:"norekBankTujuan,omitempty"`
	NoRekeningPendamping string `bson:"noRekeningPendamping,omitempty" json:"noRekeningPendamping,omitempty"`
	NorekWallet string `bson:"norekWallet,omitempty" json:"norekWallet,omitempty"`
	NorekWalletTujuan string `bson:"norekWalletTujuan,omitempty" json:"norekWalletTujuan,omitempty"`
	PaymentMethod string `bson:"paymentMethod,omitempty" json:"paymentMethod,omitempty"`
	PriceId string `bson:"priceId,omitempty" json:"priceId,omitempty"`
	ProductCode string `bson:"productCode,omitempty" json:"productCode,omitempty"`
	ReffBiller string `bson:"reffBiller,omitempty" json:"reffBiller,omitempty"`
	ReffCore string `bson:"reffCore,omitempty" json:"reffCore,omitempty"`
	ReffSwitching string `bson:"reffSwitching,omitempty" json:"reffSwitching,omitempty"`
	RequestType string `bson:"requestType,omitempty" json:"requestType,omitempty"`
	ResponseCode string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	ResponseDesc string `bson:"responseDesc,omitempty" json:"responseDesc,omitempty"`
	Rubrik string `bson:"rubrik,omitempty" json:"rubrik,omitempty"`
	SaldoEmas string `bson:"saldoEmas,omitempty" json:"saldoEmas,omitempty"`
	SaldoNominal string `bson:"saldoNominal,omitempty" json:"saldoNominal,omitempty"`
	SaldoRekeningPendamping string `bson:"saldoRekeningPendamping,omitempty" json:"saldoRekeningPendamping,omitempty"`
	Satuan string `bson:"satuan,omitempty" json:"satuan,omitempty"`
	SewaModal string `bson:"sewaModal,omitempty" json:"sewaModal,omitempty"`
	SewaModalBaru string `bson:"sewaModalBaru,omitempty" json:"sewaModalBaru,omitempty"`
	SisaTenor string `bson:"sisaTenor,omitempty" json:"sisaTenor,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	StatusKawin string `bson:"statusKawin,omitempty" json:"statusKawin,omitempty"`
	Surcharge string `bson:"surcharge,omitempty" json:"surcharge,omitempty"`
	TaksiranBaru string `bson:"taksiranBaru,omitempty" json:"taksiranBaru,omitempty"`
	TanggalExpiredId string `bson:"tanggalExpiredId,omitempty" json:"tanggalExpiredId,omitempty"`
	TanggalLahir string `bson:"tanggalLahir,omitempty" json:"tanggalLahir,omitempty"`
	TempatLahir string `bson:"tempatLahir,omitempty" json:"tempatLahir,omitempty"`
	Tenor string `bson:"tenor,omitempty" json:"tenor,omitempty"`
	TglBuka string `bson:"tglBuka,omitempty" json:"tglBuka,omitempty"`
	TglJatuhTempo string `bson:"tglJatuhTempo,omitempty" json:"tglJatuhTempo,omitempty"`
	TglKredit string `bson:"tglKredit,omitempty" json:"tglKredit,omitempty"`
	TglLelang string `bson:"tglLelang,omitempty" json:"tglLelang,omitempty"`
	TglTransaksi string `bson:"tglTransaksi,omitempty" json:"tglTransaksi,omitempty"`
	TipeIdentitas string `bson:"tipeIdentitas,omitempty" json:"tipeIdentitas,omitempty"`
	TipeTransaksi string `bson:"tipeTransaksi,omitempty" json:"tipeTransaksi,omitempty"`
	Token string `bson:"token,omitempty" json:"token,omitempty"`
	TotalKewajiban string `bson:"totalKewajiban,omitempty" json:"totalKewajiban,omitempty"`
	Tunggakan string `bson:"tunggakan,omitempty" json:"tunggakan,omitempty"`
	Up string `bson:"up,omitempty" json:"up,omitempty"`
	UpLama string `bson:"upLama,omitempty" json:"upLama,omitempty"`
	UserId string `bson:"userId,omitempty" json:"userId,omitempty"`
	WalletId string `bson:"walletId,omitempty" json:"walletId,omitempty"`

	// Untuk MongoDB
	Id string `bson:"_id,omitempty" json:"-"`
	Timestamp time.Time `json:"-"`
}

type JsonOld struct {
	// Old Request
	Admin float64 `json:"admin,omitempty"`
	Administrasi float64 `json:"administrasi,omitempty"`
	AgentId string `json:"agentId,omitempty"`
	Angsuran int `json:"angsuran,omitempty"`
	AngsuranKe int `json:"angsuranKe,omitempty"`
	BiayaAdmBjdpl string `json:"agentId,omitempty"`
	BiayaLelang string `json:"biayaLelang,omitempty"`
	BiayaProsesLelang string `json:"biayaProsesLelang,omitempty"`
	BiayaTitip float64 `json:"biayaTitip,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Charges float64 `json:"charges,omitempty"`
	Cif string `json:"cif,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	Data string `json:"data,omitempty"`
	Denda float64 `json:"denda,omitempty"`
	GolonganRubrik string `json:"golonganRubrik,omitempty"`
	Gram float64 `json:"gram,omitempty"`
	HakNasabah float64 `json:"hakNasabah,omitempty"`
	Harga float64 `json:"harga,omitempty"`
	IdJurnal string `json:"idJurnal,omitempty"`
	JenisTransaksi string `json:"jenisTransaksi,omitempty"`
	JumlahAngsuran int `json:"jumlahAngsuran,omitempty"`
	JumlahHariTarif string `json:"jumlahHariTarif,omitempty"`
	JumlahHariReal string `json:"jumlahHariReal,omitempty"`
	Kewajiban float64 `json:"kewajiban,omitempty"`
	KodeBankPembayar string `json:"kodeBankPembayar,omitempty"`
	KodeCabang string `json:"kodeCabang,omitempty"`
	KodeNamaCabang string `json:"kodeNamaCabang,omitempty"`
	MinimalUpCicil float64 `json:"minimalUpCicil,omitempty"`
	NamaNasabah string `json:"namaNasabah,omitempty"`
	NamaProduk string `json:"namaProduk,omitempty"`
	NilaiTransaksi float64 `json:"nilaiTransaksi,omitempty"`
	NoHpAgent string `json:"noHpAgent,omitempty"`
	NoKredit string `json:"noKredit,omitempty"`
	Norek string `json:"norek,omitempty"`
	NoRekBank string `json:"noRekBank,omitempty"`
	NorekWallet string `json:"norekWallet,omitempty"`
	PaymentMethod string `json:"paymentMethod,omitempty"`
	ReffIdClient string `json:"reffIdClient,omitempty"`
	ReffIdSwitching string `json:"reffIdSwitching,omitempty"`
	RekeningPendamping string `json:"rekeningPendamping,omitempty"`
	ResponseCode string `json:"responseCode,omitempty"`
	ResponseDesc string `json:"responseDesc,omitempty"`
	Rubrik string `json:"rubrik,omitempty"`
	SaldoEmas float64 `json:"saldoEmas,omitempty"`
	SaldoNominal float64 `json:"saldoNominal,omitempty"`
	SaldoRekeningPendamping int `json:"saldoRekeningPendamping,omitempty"`
	Satuan float64 `json:"satuan,omitempty"`
	SewaModal float64 `json:"sewaModal,omitempty"`
	SewaModalBaru float64 `json:"sewaModalBaru,omitempty"`
	SisaTenor int `json:"sisaTenor,omitempty"`
	Status string `json:"status,omitempty"`
	Struk string `json:"struk,omitempty"`
	TaksiranBaru float64 `json:"taksiranBaru,omitempty"`
	Tenor string `json:"tenor,omitempty"`
	TerminalId string `json:"terminalId,omitempty"`
	TglBuka string `json:"tglBuka,omitempty"`
	TglJatuhTempo string `json:"tglJatuhTempo,omitempty"`
	TglKredit string `json:"tglKredit,omitempty"`
	TglLelang string `json:"tglLelang,omitempty"`
	TglTransaksi string `json:"tglTransaksi,omitempty"`
	TotalKewajiban float64 `json:"totalKewajiban,omitempty"`
	Tunggakan float64 `json:"tunggakan,omitempty"`
	Up float64 `json:"up,omitempty"`
	UpLama float64 `json:"upLama,omitempty"`
	UserId string`json:"userId,omitempty"`
	WalletId string `json:"walletId,omitempty"`

}