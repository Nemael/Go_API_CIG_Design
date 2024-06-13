package interactor

import (
	"Users/models"
	"fmt"
)

func (i *Interactor) GetUser(id int) *models.User {
	fmt.Println("Into interactor")
	return i.My_gateway.GetUser(id)
}
