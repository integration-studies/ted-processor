package ted

import (
	"ted-processor/pkg/domain"
	"time"
)

type PreTransaction struct {
	FromAccount string
	ToAccount   string
	Value       float64
	Time        string
	DeviceType  string
	Status      string
	StartedAt   time.Time
	Metadata    map[string]string
}

func NewPreTransaction(t *domain.Transaction, metadata map[string]string) *PreTransaction {
	return &PreTransaction{
		FromAccount: t.FromAccount,
		ToAccount:   t.ToAccount,
		Value:       t.Value,
		Time:        t.Time,
		DeviceType:  t.DeviceType,
		Status:      "created",
		StartedAt:   time.Now(),
		Metadata:    metadata,
	}
}
