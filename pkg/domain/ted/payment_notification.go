package ted

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ted-processor/pkg/domain"
	"ted-processor/pkg/domain/infra/config"
	"ted-processor/pkg/domain/infra/errors"
	infra "ted-processor/pkg/domain/infra/logger"
	data2 "ted-processor/pkg/domain/ted/data"
	errors2 "ted-processor/pkg/domain/ted/errors"
)

type PaymentNotification struct {
	httpClient *http.Client
	config     *config.Config
}

func NewPaymentNotification(httpClient *http.Client, config *config.Config) *PaymentNotification {
	return &PaymentNotification{
		httpClient: httpClient,
		config:     config,
	}
}

func (pn *PaymentNotification) CallPayment(t *domain.Transaction) (*data2.PaymentData, error) {
	data := t.ToRequest()
	jc, err := json.Marshal(data)
	if err != nil {
		infra.Logger.Errorw("error to serialize to json", "error", errors.Wrap(err))
		return nil, err
	}
	pd := &data2.PaymentData{}
	rb, errHttp := http.Post(pn.config.BAHost, "application/json", bytes.NewBuffer(jc))
	if errHttp != nil {
		infra.Logger.Errorw("error call bank account service", "error", errors.Wrap(errHttp))
		return nil, errHttp
	}
	if err := json.NewDecoder(rb.Body).Decode(pd); err != nil {
		infra.Logger.Errorw("error to deserialize json", "error", errors.Wrap(err))
		return nil, err
	}
	if rb.StatusCode != 201 {
		infra.Logger.Errorw("payment was declined", "error", errors.Wrap(err))
		return pd, &errors2.PaymentDeclinedError{Id: pd.Id}
	}
	defer rb.Body.Close()
	return pd, nil
}
