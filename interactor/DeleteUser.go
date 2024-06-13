package interactor

import (
	"Users/models"
	"fmt"
)

func (i *Interactor) DeleteUser(id int) *models.User {
	fmt.Println("Into interactor")
	return i.My_gateway.DeleteUser(id)
}
