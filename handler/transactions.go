package handler

import (
	"app/dto/http"
	"app/dto/model"
	"app/repository"
	"math"

	"fmt"
)

func contains(denom []float64, amount float64) bool {
	for _, d := range denom {
		if d == amount {
			return true
		}
	}
	return false
}

func CheckedTransaction(paymentRequest *http.CreatePaymentRequest, client *model.Client) map[string]interface{} {
	var chargingPrice float64

	fmt.Printf("Processing payment: %+v\n", paymentRequest)
	fmt.Printf("Client details: %+v\n", client)

	if len(paymentRequest.UserID) > 50 || len(paymentRequest.MerchantTransactionID) > 36 || len(paymentRequest.ItemName) > 25 {
		fmt.Printf("Too long parameters: %+v\n", paymentRequest)
		return map[string]interface{}{
			"success": false,
			"retcode": "E0021",
		}
	}

	paymentMethod := paymentRequest.PaymentMethod
	switch paymentMethod {
	case "xl_gcpay", "xl_gcpay2":
		paymentMethod = "xl_airtime"
	case "smartfren":
		paymentMethod = "smartfren_airtime"
	case "three":
		paymentMethod = "three_airtime"
	case "telkomsel_airtime_sms", "telkomsel_airtime_ussd", "telkomsel_airtime_mdm":
		paymentMethod = "telkomsel_airtime"
	case "indosat_huawei", "indosat_mimopay", "indosat_simplepayment":
		paymentMethod = "indosat_airtime"
	}

	arrPaymentMethod, err := repository.FindPaymentMethodBySlug(paymentMethod, "")
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"retcode": "E0005",
		}
	}

	var route string

	// Loop through client's payment methods
	for _, arrayPayments := range client.PaymentMethods {
		if arrayPayments.Name == paymentMethod {
			arrRoutes := arrayPayments.Route
			if !arrPaymentMethod.Value.Flexible {
				// Non-flexible payment method
				for routename, arrayDenom := range arrRoutes {
					// Type assertion to []int
					// if denom, ok := arrayDenom.([]string); ok {
					// 	// Check if the amount is in the denominated range
					// 	if contains(denom, paymentRequest.Amount) {
					// 		route = routename
					// 		break
					// 	}
					// } else {
					// 	fmt.Printf("Invalid type for arrayDenom: %T\n", arrayDenom)
					// }
					denom := arrayDenom
					// Check if the amount is in the denominated range
					if containsString(denom, fmt.Sprintf("%.0f", paymentRequest.Amount)) {
						route = routename
						break
					} else {
						fmt.Printf("Invalid type for arrayDenom: %T\n", arrayDenom)
					}
				}
			} else {
				// Flexible payment method
				for routename, value := range arrRoutes {
					if containsString(value, "1") {
						route = routename
						break
					}
				}
			}
		}
	}

	if route == "" {
		return map[string]interface{}{
			"success": false,
			"retcode": "E0007",
		}
	}

	// TODO
	// perlu checksupported di repository, check code legacy
	// check func search_sub_array di code legacy

	arrPaymentMethodRoute, err := repository.FindPaymentMethodBySlug(route, "")

	if arrPaymentMethodRoute.Value.Flexible {
		if paymentRequest.Amount < arrPaymentMethodRoute.MinimumDenom {
			if client.Testing == 0 {
				return map[string]interface{}{
					"success": false,
					"retcode": "E0020",
				}
			}
			chargingPrice = float64(paymentRequest.Amount)
		} else {
			switch {
			case paymentMethod == "indosat_airtime2" && route == "indosat_triyakom4":
				chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
			case paymentMethod == "gopay" && route == "gopay_midtrans":
				clientName := client.ClientName
				if clientName == "Topfun 2 New Qiuqiu" || clientName == "Topfun" || clientName == "SPOLIVE" || clientName == "Tricklet (Hong Kong) Limited" {
					chargingPrice = paymentRequest.Amount
				} else {
					chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
				}
			case paymentMethod == "shopeepay" && route == "shopeepay_midtrans":
				clientName := client.ClientName
				if clientName == "Tricklet (Hong Kong) Limited" {
					chargingPrice = paymentRequest.Amount
				} else {
					chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
				}
			case paymentMethod == "alfamart_otc" && route == "alfamart_faspay":
				clientName := client.ClientName
				if clientName == "Tricklet (Hong Kong) Limited" {
					chargingPrice = paymentRequest.Amount
				} else if clientName == "Redigame" {
					chargingPrice = paymentRequest.Amount + 6000
				} else {
					chargingPrice = paymentRequest.Amount + 6500
				}
			case paymentMethod == "indomaret_otc" && route == "indomaret_otc_mst":
				clientName := client.ClientName
				subtotal := paymentRequest.Amount
				var adminFee, totalAmount float64
				switch clientName {
				case "Tricklet (Hong Kong) Limited":
					adminFee = 0
				case "Redigame":
					adminFee = ((0.06 * paymentRequest.Amount) + 1000) / (1 - 0.06)
				case "Higo Game PTE LTD":
					adminFee = ((0.07 * paymentRequest.Amount) + 1000) / (1 - 0.07)
				default:
					adminFee = ((0.075 * paymentRequest.Amount) + 1000) / (1 - 0.075)
				}
				totalAmount = subtotal + adminFee
				chargingPrice = math.Round(totalAmount/100) * 100
			case paymentMethod == "smartfren_airtime2" && (route == "smartfren_triyakom_flex" || route == "smartfren_triyakom_flex2"):
				chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
			case paymentMethod == "three_airtime2" && route == "three_triyakom_flex2":
				chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
			default:
				chargingPrice = paymentRequest.Amount
			}

			if client.ClientName == "Zingplay International PTE,. LTD" && paymentMethod == "ovo_wallet" && route == "ovo" {
				chargingPrice = paymentRequest.Amount + math.Round(0.11*paymentRequest.Amount)
			}
		}

	} else {

		// TODO
		// recheck jika perlu pengecekan denom saat request charging, bisa check code legacy

		// statusDenoms := arrPaymentMethodRoute.Value.Status.(map[float64]string)
		// status := false

		// if denom == paymentRequest.Amount && val == "1" {
		// 	status = true
		// 	break
		// }

		// if !status {
		// 	return map[string]interface{}{
		// 		"success": false,
		// 		"retcode": "E0014",
		// 	}
		// }

		chargingPrice, err = repository.GetPrice(route, paymentRequest.Amount)
	}

	return map[string]interface{}{
		"success":        true,
		"retcode":        "E0020",
		"mobile":         client.Mobile,
		"testing":        client.Testing,
		"charging_price": chargingPrice,
		"route":          route,
		"payment_method": paymentMethod,
	}
}

func containsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
