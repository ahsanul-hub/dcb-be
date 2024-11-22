package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type XlChargeRequest struct {
// 	Cache *cache.Cache
// }

type ChargeResponse struct {
	Success          bool        `json:"success"`
	Version          string      `json:"version"`
	Response         interface{} `json:"response"`
	PhoneNumber      string      `json:"phone_number"`
	HTTPCode         string      `json:"httpcode"`
	HTTPResponse     string      `json:"httpresponse"`
	StatusToken      string      `json:"status_token"`
	StatusInquiry    string      `json:"status_inquiry"`
	SubscriberNo     string      `json:"subscriberNo"`
	SubscriberStatus string      `json:"subscriberStatus"`
	ErrorMessage     string      `json:"error_message"`
}

func SendData(data map[string]interface{}) (ChargeResponse, error) {
	// Initialize response data
	responseData := ChargeResponse{
		Success:       false,
		Version:       data["version"].(string),
		Response:      nil,
		PhoneNumber:   data["msisdn"].(string),
		HTTPCode:      "",
		HTTPResponse:  "",
		StatusToken:   data["status_token"].(string),
		StatusInquiry: data["status_inquiry"].(string),
	}

	// Check required fields
	if body, ok := data["body"]; ok {
		if headers, ok := data["headers"]; ok {
			if subscriberNo, ok := data["subscriberNo"]; ok {
				if subscriberStatus, ok := data["subscriberStatus"]; ok {
					if accessToken, ok := data["access_token"]; ok {
						if data["status_token"] == "1" && data["status_inquiry"] == "1" && data["error_message"] == "" {
							// Proceed with sending the request
							// Log the request (optional)
							log.Printf("Sending request to %s", data["requestUrl"])

							// Prepare the HTTP request
							client := &http.Client{}
							req, err := http.NewRequest("POST", data["requestUrl"].(string), bytes.NewBuffer([]byte(body.(string))))
							if err != nil {
								return responseData, err
							}
							req.Header.Set("Content-Type", "application/json")
							req.Header.Set("Accept", "application/json")
							req.Header.Set("access-token", accessToken.(string))
							req.Header.Set("Connection", "keep-alive")
							req.Header.Set("cache-control", "no-cache")

							// Send the request
							resp, err := client.Do(req)
							if err != nil {
								return responseData, err
							}
							defer resp.Body.Close()

							// Read the response
							bodyBytes, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								return responseData, err
							}

							// Unmarshal the response
							var arrCharge map[string]interface{}
							if err := json.Unmarshal(bodyBytes, &arrCharge); err != nil {
								return responseData, err
							}

							// Get HTTP status code
							httpCode := resp.StatusCode
							responseData.HTTPCode = fmt.Sprintf("%d", httpCode)
							responseData.HTTPResponse = string(bodyBytes)

							if httpCode == http.StatusOK {
								responseCode := arrCharge["chargingResponse"].(map[string]interface{})["TransactionStatus"].(map[string]interface{})["responseCode"].(string)
								responseDesc := arrCharge["chargingResponse"].(map[string]interface{})["TransactionStatus"].(map[string]interface{})["responseDesc"].(string)

								if responseCode == "50" {
									// Success case
									responseData.Success = true
									responseData.Response = arrCharge
									responseData.SubscriberNo = subscriberNo.(string)
									responseData.SubscriberStatus = subscriberStatus.(string)
									responseData.ErrorMessage = ""
								} else {
									// Failure case
									responseData.ErrorMessage = "Request failed! (Err: ERR_CHARGE4)"
								}
							} else {
								// HTTP code not 200
								responseData.ErrorMessage = "Request failed! (Err: ERR_CHARGE2)"
							}
						}
					}
				}
			}
		}
	}

	// If we reach here, it means we didn't send the request
	if !responseData.Success {
		responseData.ErrorMessage = data["error_message"].(string)
	}

	return responseData, nil
}
