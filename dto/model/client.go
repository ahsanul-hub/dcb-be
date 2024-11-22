package model

// PaymentMethod represents an individual payment method for a client.
// type PaymentMethodClient struct {
// 	Name   string                 `json:"name"`
// 	Status int                    `json:"status"`
// 	Msisdn int                    `json:"msisdn"`
// 	Route  map[string]interface{} `json:"route"`
// }

// // Client represents a client in the system.
// type Client struct {
// 	ID                string                `json:"_id"`
// 	ClientAppKey      string                `json:"client_appkey"`
// 	ClientSecret      string                `json:"client_secret"`
// 	ClientName        string                `json:"client_name"`
// 	AppName           string                `json:"app_name"`
// 	ClientStatus      int                   `json:"client_status"`
// 	Mobile            int                   `json:"mobile"`
// 	Testing           int                   `json:"testing"`
// 	PaymentMethods    []PaymentMethodClient `json:"payment_methods"`
// 	Lang              string                `json:"lang"`
// 	CallbackURL       string                `json:"callback_url"`
// 	FailCallback      string                `json:"fail_callback"`
// 	IsImportAvailable bool                  `json:"is_import_available"`
// 	// Route             string                `json:"route"`
// }

type Client struct {
	ID             string                `json:"_id"`
	UID            string                `json:"u_id"`
	ClientName     string                `json:"client_name"`
	ClientAppkey   string                `json:"client_appkey"`
	ClientSecret   string                `json:"client_secret"`
	ClientAppid    string                `json:"client_appid"`
	AppName        string                `json:"app_name"`
	Mobile         string                `json:"mobile"`
	ClientStatus   int                   `json:"client_status"`
	Testing        int                   `json:"testing"`
	Lang           string                `json:"lang"`
	CallbackURL    string                `json:"callback_url"`
	PaymentMethods []PaymentMethodClient `json:"payment_methods"`
	FailCallback   string                `json:"fail_callback"`
	Isdcb          string                `json:"isdcb"`
	UpdatedAt      string                `json:"updated_at"`
	CreatedAt      string                `json:"created_at"`
}

type PaymentMethodClient struct {
	Name   string              `json:"name"`
	Route  map[string][]string `json:"route"`
	Status int                 `json:"status"`
	Msisdn int                 `json:"msisdn"`
}
