package handler

import (
	"app/dto/http"
	"app/dto/model"
	"app/helper"
	"app/lib"
	"app/pkg/response"
	. "app/repository"
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {
	headers := map[string]string{
		"appkey":    c.Get("appkey"),
		"appid":     c.Get("appid"),
		"timestamp": c.Get("timestamp"),
		"nonce":     c.Get("nonce"),
		"secret":    c.Get("secret"),
		"bodysign":  c.Get("bodysign"),
	}

	arrClient, err := FindClient(c.Get("appkey"), c.Get("appid"))

	if err != nil {
		return response.Response(c, fiber.StatusInternalServerError, err.Error())
	}

	var req http.CreatePaymentRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validator.New().Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	arrayTransactionCheck := CheckedTransaction(&req, arrClient)
	if !arrayTransactionCheck["success"].(bool) {
		return response.Response(c, fiber.StatusBadRequest, arrayTransactionCheck["retcode"].(string))
	}

	inputReq := model.InputPaymentRequest{
		ClientAppKey:  headers["appkey"],
		StatusCode:    "1001",
		Status:        helper.GetStatusMessage("1001"), // Fungsi untuk mendapatkan deskripsi status
		Mobile:        arrayTransactionCheck["mobile"].(string),
		Testing:       arrayTransactionCheck["testing"].(bool),
		Route:         arrayTransactionCheck["route"].(string),
		PaymentMethod: arrayTransactionCheck["payment_method"].(string),
		Currency:      "IDR",
		Price:         arrayTransactionCheck["charging_price"].(float64),
	}

	// Beautify UserMDN
	if req.UserMDN != "" {
		req.UserMDN = helper.BeautifyIDNumber(strings.TrimSpace(inputReq.UserMDN), false)
	}

	// Create the transaction order
	transactionToken, err := CreateOrder(context.Background(), &inputReq, arrClient)
	if err != nil {
		return response.Response(c, fiber.StatusInternalServerError, "E4001: Database Error")
	}

	// Save timestamps for transaction
	// err = SaveTransactionTimestamp(transactionToken)
	// if err != nil {
	// 	return response.Response(c, fiber.StatusInternalServerError, "E4001: Failed to update transaction timestamps")
	// }

	// Return successful response

	return response.ResponseSuccess(c, fiber.StatusOK, fiber.Map{
		"token": transactionToken,
	})
}

func TestPayment(c *fiber.Ctx) error {
	// Mendapatkan data dari request body
	var requestData model.InputPaymentRequest
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input",
		})
	}

	// URL endpoint untuk pembayaran
	url := "http://3.1.41.116/api/v1/create"

	// AppKey dan BodySign (contoh)
	secretKey := "72Zwth2Dd75yuYzRhgKhGcsdf"
	appKey := "7d51a9a750575a294df94a78bde79628"
	bodySign, err := helper.CreateBodySign(requestData, secretKey) // Anda harus membuat ini berdasarkan data yang diinginkan
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create body sign: " + err.Error(),
		})
	}

	// Kirim permintaan pembayaran
	err = lib.SendPaymentRequest(url, requestData, appKey, bodySign)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Payment successful",
	})
}
