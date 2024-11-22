package model

// type Value struct {
// 	Status     int                `json:"status"`
// 	Denom      map[string]float64 `json:"denom"`
// 	DailyLimit int                `json:"daily_limit"`
// }

// type PaymentMethod struct {
// 	ID           string  `json:"_id"`
// 	Slug         string  `json:"slug"`
// 	Description  string  `json:"description"`
// 	Flexible     bool    `json:"flexible"`
// 	IsAirtime    bool    `json:"is_airtime"`
// 	MinimumDenom float64 `json:"minimum_denom"`
// 	Route        string  `json:"route"`

// 	Value Value `json:"value"`
// }

type PaymentMethod struct {
	ID           string   `json:"_id"`
	Slug         string   `json:"slug"`
	Description  string   `json:"description"`
	Value        Value    `json:"value"`
	Route        []string `json:"route"`
	Type         string   `json:"type"`
	Expired      string   `json:"expired"`
	Report       string   `json:"report"`
	JSONReturn   string   `json:"json_return"`
	Parent       string   `json:"parent"`
	IsAirtime    string   `json:"is_airtime"`
	MinimumDenom float64  `json:"minimum_denom"`
	Disabled     string   `json:"disabled"`
	UpdatedAt    string   `json:"updated_at"`
}

type Value struct {
	Flexible    bool              `json:"flexible"`
	Status      string            `json:"status"`
	Msisdn      string            `json:"msisdn"`
	StatusDenom map[string]string `json:"status_denom"`
	Denom       []float64         `json:"denom"`
	Prefix      []string          `json:"prefix"`
	DailyLimit  string            `json:"daily_limit"`
}
