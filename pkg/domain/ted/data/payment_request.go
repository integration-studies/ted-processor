package data

type PaymentRequest struct {
	Type        string
	SubType     string
	FromAccount string
	ToAccount   string
	Value       float64
	Time        string
	DeviceType  string
}
