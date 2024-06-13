package interactor

import (
	"Users/models"
	"errors"
	"fmt"
)

func (i *Interactor) AddAccount(user *models.Account) (*models.Account, error) {
	fmt.Println("Into interactor")
	if user.Account_number < 1000 || user.Account_number > 9999 {
		return nil, errors.New("account number must be 4 digits long")
	}

	return i.My_gateway.AddAccount(user), nil
}
