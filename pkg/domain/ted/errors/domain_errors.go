package errors

import "fmt"

type PaymentDeclinedError struct {
	Id  string
	Err error
}

func (p *PaymentDeclinedError) Error() string {
	return fmt.Sprintf("payment %v was declined", p.Id)
}
