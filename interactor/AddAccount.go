package interactor

import (
	"Users/models"
	"errors"
	"fmt"
)

func (i *Interactor) AddAccount(account *models.Account) (*models.Account, error) {
	fmt.Println("Into interactor")
	fmt.Println(account.User_id)
	if account.Account_number < 1000 || account.Account_number > 9999 {
		return nil, errors.New("account number must be 4 digits long")
	}

	return i.My_gateway.AddAccount(account), nil
}
