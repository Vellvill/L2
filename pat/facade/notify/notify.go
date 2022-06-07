package notify

type Notify struct {
}

func NewNotify() *Notify {
	return &Notify{}
}

func (n *Notify) NotifyAccount(accountID string) string {
	return "Account notified "
}
