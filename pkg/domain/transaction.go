package domain

import (
	"ted-processor/pkg/domain/ted/data"
)

type Transaction struct {
	Type        string
	SubType     string
	FromAccount string
	ToAccount   string
	Value       float64
	Time        string
	DeviceType  string
}

func (t *Transaction) ToRequest() *data.PaymentRequest {
	return &data.PaymentRequest{
		Type:        t.Type,
		SubType:     t.SubType,
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Value:       t.Value,
		Time:        t.Time,
		DeviceType:  t.DeviceType,
	}
}
