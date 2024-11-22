package http

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
