package data

type PaymentData struct {
	Id               string `json:"id"`
	Reason           string `json:"reason"`
	OperationSuccess bool   `json:"operationSuccess"`
}
