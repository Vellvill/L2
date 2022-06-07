package account

type Account struct {
	id string
}

func NewAccount(id string) *Account {
	return &Account{id: id}
}

func (a *Account) GetID() string {
	return a.id
}
