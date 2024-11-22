package model

import (
	"time"
)

type Transactions struct {
	ID                       uint `gorm:"primaryKey" json:"id"`
	Mt1Id                    string
	BersamaBookingId         string `json:"bersama_booking_id"`
	SmsCode                  string
	MerchantName             string `json:"merchant_name"`
	Keyword                  string
	Otp                      int       `json:"otp"`
	TcashId                  string    `json:"tcach_id"`
	VaBcadynamicFaspayBillno int       ` json:"va_bca"`
	MtTid                    int       `json:"mt_tid"`
	DisbursementId           string    `json:"disbursement_id"`
	PaymentMethod            string    `json:"payment_method"`
	StatusCode               int       `json:"status_code"`
	ItemName                 string    `json:"item_name"`
	Route                    string    `json:"route"`
	MdmTrxID                 string    `json:"mdm_trx_id"`
	Timestamp                time.Time `json:"timestamp"`
	Stan                     string    `json:"json"`
	Amount                   float64   `json:"amount" gorm:"amount"`
	ClientAppKey             string    `json:"client_appkey" gorm:"client_appkey"`
	AppID                    string    `json:"appid"`
	AppKey                   string    `json:"appkey"`
	Testing                  bool      `json:"testing"`
	Token                    string    `json:"token"`
	Currency                 string    `json:"currency"`
	Price                    float64   `json:"price"`
	BodySign                 string    `json:"bodysign"`
	UserMDN                  string    `json:"user_mdn"`
	RedirectURL              string    `json:"redirect_url"`
	RedirectTarget           string    `json:"redirect_target"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type InputPaymentRequest struct {
	ClientAppKey  string  `json:"client_appkey"`  // App Key dari klien
	StatusCode    string  `json:"status_code"`    // Status code dari transaksi
	Status        string  `json:"status"`         // Deskripsi status
	Mobile        string  `json:"mobile"`         // Nomor telepon pengguna
	Testing       bool    `json:"testing"`        // Indikator apakah dalam mode testing
	Route         string  `json:"route"`          // Rute transaksi
	PaymentMethod string  `json:"payment_method"` // Metode pembayaran yang digunakan
	Currency      string  `json:"currency"`       // Mata uang transaksi
	Price         float64 `json:"price"`          // Harga atau jumlah transaksi
	Amount        float64 `json:"amount"`         // Harga atau jumlah transaksi
	ItemName      float64 `json:"item_name"`      // Harga atau jumlah transaksi
	UserMDN       string  `json:"user_mdn"`       // Nomor telepon pengguna (format sudah dirapikan)
}
