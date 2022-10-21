package facade

import "fmt"

type Account struct {
	Name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

