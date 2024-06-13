package interactor

import (
	"Users/models"
	"fmt"
)

func (i *Interactor) GetUsers() []models.User {
	fmt.Println("Into interactor")
	return i.My_gateway.GetUsers()
}
