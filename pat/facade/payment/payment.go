package payment

import "fmt"

type Payment struct {
	paymentType string
}

func NewPayment(paymentType string) *Payment {
	return &Payment{
		paymentType: paymentType,
	}
}

func (p *Payment) Pay(accountID string) string {
	return fmt.Sprintf("%s pay by %s", accountID, p.paymentType)
}
