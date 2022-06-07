package main

import (
	"L2/pat/facade/account"
	"L2/pat/facade/notify"
	"L2/pat/facade/payment"
	"fmt"
)

type Facade struct {
	account *account.Account
	notify  *notify.Notify
	payment *payment.Payment
}

func NewFacade(accountID, paymentType string) *Facade {
	return &Facade{
		account: account.NewAccount(accountID),
		notify:  notify.NewNotify(),
		payment: payment.NewPayment(paymentType),
	}
}

func (f *Facade) ProceedPay() (string, error) {
	id := f.account.GetID()
	f.notify.NotifyAccount(id)
	return f.payment.Pay(id), nil
}

func main() {
	F := NewFacade("12345", "QIWI")
	message, err := F.ProceedPay()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
