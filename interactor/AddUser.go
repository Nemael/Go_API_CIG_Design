package interactor

import (
	"Users/models"
	"errors"
	"fmt"
)

func (i *Interactor) AddUser(user *models.User) (*models.User, error) {
	fmt.Println("Into interactor")
	if len(user.First_name) > len(user.Last_name) {
		return nil, errors.New("user first_name must be shorter than user last_name")
	}
	return i.My_gateway.AddUser(user), nil
}
