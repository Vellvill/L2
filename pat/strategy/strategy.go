package main

func main() {
	var (
		product       = "tv"
		paymentMethod = 1
		payment       Pay
	)
	switch paymentMethod {
	case 1:
		payment = newPayCard()
	case 2:
		payment = newQIWI()
	case 3:
		payment = newPayPal()
	}
	err := ProcessOrder(product, payment)
	if err != nil {
		panic(err)
	}
}

func ProcessOrder(productName string, payment Pay) error {
	err := payment.Pay()
	if err != nil {
		return err
	}
	return nil
}

type Pay interface {
	Pay() error
}

type Card struct {
}

func newPayCard() *Card {
	return &Card{}
}

func (p *Card) Pay() error {
	return nil
}

type PayPal struct {
}

func newPayPal() *PayPal {
	return &PayPal{}
}

func (p PayPal) Pay() error {
	return nil
}

type QIWI struct {
}

func newQIWI() *QIWI {
	return &QIWI{}
}

func (p *QIWI) Pay() error {
	return nil
}
