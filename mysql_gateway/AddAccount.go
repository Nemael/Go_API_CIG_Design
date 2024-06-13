package mysql_gateway

import (
	"Users/models"
	"fmt"
)

func (g *Gateway) AddAccount(account *models.Account) *models.Account {
	fmt.Println("Into gateway")
	db, err := getDB()
	if err != nil {
		fmt.Println("Error getting database")
		return nil
	}
	_, err = db.Exec("INSERT INTO Accounts (user_id, account_number) VALUES (?, ?)", account.User_id, account.Account_number)
	if err != nil {
		fmt.Println("Error inserting data", err)
		return nil
	}
	return account
}
