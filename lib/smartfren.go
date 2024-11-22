package lib

import (
	"app/dto/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func SendPaymentSmartfren(data model.InputPaymentRequest, appKey string, bodySign string, config map[string]string) (error, map[string]interface{}) {
	config, err := config.GetGatewayConfig("smartfren")
	amount := data.Amount

	if !ok {
		return fmt.Errorf("keyword for amount %d not found in config", amount), nil

	}

	serviceNode := config["serviceNode"]
	// msisdn := data.UserMdn
	msgCoding := config["msgCoding"]
	sender := config["sender"]
	smscId := config["smscId"]
	bearerId := config["bearerId"]

	smsCode := generateSMSCode()

	hexMsg := keyword + " " + smsCode

	query := url.Values{}
	query.Add("serviceNode", serviceNode)
	// query.Add("msisdn", msisdn)
	query.Add("keyword", keyword)
	query.Add("msgCoding", msgCoding)
	query.Add("sender", sender)
	query.Add("hexMsg", hexMsg)
	query.Add("smscId", smscId)
	query.Add("bearerId", bearerId)

	gateway := config["ip"] + ":" + config["port"] + "/moReq"

	requestURL := gateway + "?" + query.Encode()

	// Encode request data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err), nil
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err), nil
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("appkey", appKey)
	req.Header.Set("bodysign", bodySign)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response status: %s", resp.Status), nil
	}

	responseData := map[string]interface{}{
		"requestURL": requestURL,
		"sms_code":   smsCode,
		"keyword":    keyword,
	}

	return nil, responseData

}
