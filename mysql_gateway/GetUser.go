package mysql_gateway

import (
	"Users/models"
	"fmt"
)

func (g *Gateway) GetUser(id int) *models.User {
	fmt.Println("Into gateway")
	user := models.User{}
	db, err := getDB()
	if err != nil {
		fmt.Println("Error getting database")
		return nil
	}
	//Get user data
	rows := db.QueryRow("SELECT id, first_name, last_name FROM users WHERE id = ?", id)
	fmt.Println("Got the rows")
	fmt.Println(rows)
	err = rows.Scan(&user.Id, &user.First_name, &user.Last_name)
	if err != nil {
		fmt.Println("Error parsing rows", err)
		return nil
	}

	//Get account data
	account_rows, _ := db.Query("SELECT user_id, account_number FROM Accounts WHERE user_id = ?", id)
	var accounts []models.Account
	fmt.Println("Got the rows")
	fmt.Println(id)
	for account_rows.Next() {
		var account models.Account
		err = account_rows.Scan(&account.User_id, &account.Account_number)
		if err != nil {
			fmt.Println("Error parsing rows", err)
			return nil
		}
		fmt.Println(accounts)
		accounts = append(accounts, account)
	}
	user.Accounts = accounts
	return &user
}
