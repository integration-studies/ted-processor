package data

type PaymentRequest struct {
	Type        string  `json:"type"`
	SubType     string  `json:"subType"`
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Value       float64 `json:"value"`
	Time        string  `json:"time"`
	DeviceType  string  `json:"deviceType"`
}
