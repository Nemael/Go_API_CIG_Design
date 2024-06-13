package mysql_gateway

import (
	"Users/models"
	"fmt"
)

func (g *Gateway) GetUsers() []models.User {
	fmt.Println("Into gateway")
	users := []models.User{}
	db, _ := getDB()
	rows, _ := db.Query("SELECT id, first_name, last_name FROM users")
	fmt.Println("Got the rows")
	fmt.Println(rows)
	for rows.Next() {
		var user models.User
		_ = rows.Scan(&user.Id, &user.First_name, &user.Last_name)
		users = append(users, user)
	}
	var updated_users []models.User
	for _, user := range users {
		//Get account data
		account_rows, _ := db.Query("SELECT user_id, account_number FROM Accounts WHERE user_id = ?", user.Id)
		var accounts []models.Account
		fmt.Println("Got the rows")
		fmt.Println(user.Id)
		for account_rows.Next() {
			var account models.Account
			err := account_rows.Scan(&account.User_id, &account.Account_number)
			if err != nil {
				fmt.Println("Error parsing rows", err)
				return nil
			}
			fmt.Println(accounts)
			accounts = append(accounts, account)
		}
		user.Accounts = accounts
		updated_users = append(updated_users, user)
	}
	return updated_users
}
