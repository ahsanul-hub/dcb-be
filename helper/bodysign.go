package helper

import (
	"app/dto/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// createBodySign menerima input berupa data yang akan diproses menjadi JSON
// dan secretKey yang digunakan untuk proses HMAC.
func CreateBodySign(data model.InputPaymentRequest, secretKey string) (string, error) {
	// Ubah data menjadi JSON
	inputJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Membuat HMAC SHA256
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write(inputJSON)
	signature := h.Sum(nil)

	// Encode hasil signature ke base64
	bodySign := base64.StdEncoding.EncodeToString(signature)

	// Mengganti karakter '+' dengan '-' dan '/' dengan '_' untuk menyesuaikan format
	bodySign = strings.ReplaceAll(bodySign, "+", "-")
	bodySign = strings.ReplaceAll(bodySign, "/", "_")

	return bodySign, nil
}
