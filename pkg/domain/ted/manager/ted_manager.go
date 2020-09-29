package manager

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	"ted-processor/pkg/domain"
	"ted-processor/pkg/domain/infra/errors"
	infra "ted-processor/pkg/domain/infra/logger"
	"ted-processor/pkg/domain/ted"
)

type TedManager struct {
	repo                ted.Repository
	paymentNotification *ted.PaymentNotification
}

func (tm *TedManager) Receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	infra.Logger.Infow("Event received", "event", event)
	data := &domain.Transaction{}
	if err := event.DataAs(data); err != nil {
		infra.Logger.Errorw("Error while extracting cloudevent Data", "error", errors.Wrap(err))
		return nil, cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
	}
	metadata := map[string]string{
		"eventId":   event.ID(),
		"subjectId": event.Subject(),
		"source":    event.Source(),
	}
	pt := ted.NewPreTransaction(data, metadata)
	pt, errPt := tm.repo.StorePreTransaction(pt)
	if errPt != nil {
		infra.Logger.Errorw("Error while storing pre-transaction", "error", errors.Wrap(errPt))
		return nil, cloudevents.NewHTTPResult(400, "failed to store pre-transaction: %s", errPt)
	}

	tc := &ted.TransactionConfirmation{}
	payment, errPay := tm.paymentNotification.CallPayment(data)
	if errPay != nil {
		infra.Logger.Errorw("Error in bank account service", "error", errors.Wrap(errPay))
		tc = ted.NewTransactionError(metadata, data)
	} else {
		tc = ted.NewTransactionConfirmation(metadata, payment, data)
	}
	td, errCon := tm.repo.StoreTransactionConfirmation(tc)
	if errCon != nil {
		infra.Logger.Errorw("Error while storing transaction confirmation", "error", errors.Wrap(errPay))
	}
	log.Printf("Transaction was decoded sucessfully!!! %q", data)
	return nil, cloudevents.NewHTTPResult(202, "event received", td)
}

func NewTedManager(repo ted.Repository, pm *ted.PaymentNotification) *TedManager {
	return &TedManager{repo: repo, paymentNotification: pm}
}
