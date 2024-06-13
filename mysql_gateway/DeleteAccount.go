package mysql_gateway

import (
	"fmt"
)

func (g *Gateway) DeleteAccount(id int, account_number int) {
	fmt.Println("Into gateway")
	db, err := getDB()
	if err != nil {
		fmt.Println("Error getting database")
	}
	db.Exec("DELETE FROM Accounts WHERE user_id = ? AND account_number = ?", id, account_number)
	row, err := db.Query("DELETE FROM Accounts WHERE user_id = ? AND account_number = ?", id, account_number)
	if err != nil {
		fmt.Println("Error deleting account: ", err)
	}
	fmt.Println("Deleted row: ", row)
}
