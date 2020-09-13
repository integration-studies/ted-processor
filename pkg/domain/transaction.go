package domain

import (
	"ted-processor/pkg/domain/ted/data"
	"time"
)

type Transaction struct {
	Type        string
	SubType     string
	FromAccount string
	ToAccount   string
	Value       float64
	Time        time.Time
	DeviceType  string
}

func (t *Transaction) ToRequest() *data.PaymentRequest  {
	return &data.PaymentRequest{
		Type:        t.Type,
		SubType:     t.SubType,
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Value:       t.Value,
		Time:        t.Time.Format(time.RFC1123),
		DeviceType:  t.DeviceType,
	}
}