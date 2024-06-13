package interactor

import (
	"fmt"
)

func (i *Interactor) DeleteAccount(id int, account_number int) {
	fmt.Println("Into interactor")
	i.My_gateway.DeleteAccount(id, account_number)
}
