package ted

import (
	"ted-processor/pkg/domain"
	"ted-processor/pkg/domain/ted/data"
	"time"
)

type TransactionConfirmation struct {
	FromAccount string
	ToAccount   string
	Value       float64
	Time        string
	DeviceType  string
	Status      string
	EndAt       time.Time
	Metadata    map[string]string
	PaymentId   string
}

func NewTransactionConfirmation(metadata map[string]string, paymentData *data.PaymentData, t *domain.Transaction) *TransactionConfirmation {
	return &TransactionConfirmation{
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Value:       t.Value,
		Time:        t.Time,
		DeviceType:  t.DeviceType,
		Status:      "processed",
		EndAt:       time.Now(),
		Metadata:    metadata,
		PaymentId:   paymentData.Id,
	}
}

func NewTransactionError(metadata map[string]string, t *domain.Transaction) *TransactionConfirmation {
	return &TransactionConfirmation{
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Value:       t.Value,
		Time:        t.Time,
		DeviceType:  t.DeviceType,
		Status:      "failed",
		EndAt:       time.Now(),
		Metadata:    metadata,
		PaymentId:   "",
	}
}
