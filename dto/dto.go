package dto

import "time"

type Json struct {
	AgentId string `bson:"agentId,omitempty" json:"agentId,omitempty"`
	Amount string `bson:"amount,omitempty" json:"amount,omitempty"`
	ChannelId string `bson:"channelId,omitempty" json:"channelId,omitempty"`
	Cif string `bson:"cif,omitempty" json:"cif,omitempty"`
	ClientId string `bson:"clientId,omitempty" json:"clientId,omitempty"`
	Data string `bson:"data,omitempty" json:"data,omitempty"`
	Flag string `bson:"flag,omitempty" json:"flag,omitempty"`
	GramTransaksi string `bson:"gramtransaksi,omitempty" json:"gramtransaksi,omitempty"`
	IbuKandung string `bson:"ibuKandung,omitempty" json:"ibuKandung,omitempty"`
	IdKelurahan string `bson:"idKelurahan,omitempty" json:"idKelurahan,omitempty"`
	Jalan string `bson:"jalan,omitempty" json:"jalan,omitempty"`
	JenisKelamin string `bson:"jenisKelamin,omitempty" json:"jenisKelamin,omitempty"`
	JenisTransaksi string `bson:"jenisTransaksi,omitempty" json:"jenisTransaksi,omitempty"`
	Kewarganegaraan string `bson:"kewarganegaraan,omitempty" json:"kewarganegaraan,omitempty"`
	KodeBankPembayar string `bson:"kodeBankPembayar,omitempty" json:"kodeBankPembayar,omitempty"`
	KodeBankTujuan string `bson:"kodeBankTujuan,omitempty" json:"kodeBankTujuan,omitempty"`
	KodeCabang string `bson:"kodeCabang,omitempty" json:"kodeCabang,omitempty"`
	NamaBankTujuan string `bson:"namaBankTujuan,omitempty" json:"namaBankTujuan,omitempty"`
	NamaNasabah string `bson:"namaNasabah,omitempty" json:"namaNasabah,omitempty"`
	NilaiTransaksi string `bson:"nilaiTransaksi,omitempty" json:"nilaiTransaksi,omitempty"`
	NoHp string `bson:"noHp,omitempty" json:"noHp,omitempty"`
	NoHpAgent string `bson:"noHpAgent,omitempty" json:"noHpAgent,omitempty"`
	NoIdentitas string `bson:"noIdentitas,omitempty" json:"noIdentitas,omitempty"`
	Norek string `bson:"norek,omitempty" json:"norek,omitempty"`
	NorekBankTujuan string `bson:"norekBankTujuan,omitempty" json:"norekBankTujuan,omitempty"`
	NorekWallet string `bson:"norekWallet,omitempty" json:"norekWallet,omitempty"`
	NorekWalletTujuan string `bson:"norekWalletTujuan,omitempty" json:"norekWalletTujuan,omitempty"`
	PaymentMethod string `bson:"paymentMethod,omitempty" json:"paymentMethod,omitempty"`
	PriceId string `bson:"priceId,omitempty" json:"priceId,omitempty"`
	ProductCode string `bson:"productCode,omitempty" json:"productCode,omitempty"`
	ReffBiller string `bson:"reffBiller,omitempty" json:"reffBiller,omitempty"`
	ReffSwitching string `bson:"reffSwitching,omitempty" json:"reffSwitching,omitempty"`
	RequestType string `bson:"requestType,omitempty" json:"requestType,omitempty"`
	ResponseCode string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	ResponseDesc string `bson:"responseDesc,omitempty" json:"responseDesc,omitempty"`
	StatusKawin string `bson:"statusKawin,omitempty" json:"statusKawin,omitempty"`
	TanggalExpiredId string `bson:"tanggalExpiredId,omitempty" json:"tanggalExpiredId,omitempty"`
	TanggalLahir string `bson:"tanggalLahir,omitempty" json:"tanggalLahir,omitempty"`
	TempatLahir string `bson:"tempatLahir,omitempty" json:"tempatLahir,omitempty"`
	TipeIdentitas string `bson:"tipeIdentitas,omitempty" json:"tipeIdentitas,omitempty"`
	TipeTransaksi string `bson:"tipeTransaksi,omitempty" json:"tipeTransaksi,omitempty"`
	Token string `bson:"token,omitempty" json:"token,omitempty"`
	UserId string `bson:"userId,omitempty" json:"userId,omitempty"`
	WalletId string `bson:"walletId,omitempty" json:"walletId,omitempty"`

	// Untuk MongoDB
	Id string `bson:"_id,omitempty" json:"-"`
	Timestamp time.Time `json:"-"`
}