package mysql_gateway

import (
	"Users/models"
	"fmt"
)

func (g *Gateway) DeleteUser(id int) *models.User {
	fmt.Println("Into gateway")
	user := models.User{}
	db, err := getDB()
	if err != nil {
		fmt.Println("Error getting database")
		return nil
	}
	db.Exec("DELETE FROM users WHERE id = ?", id)
	db.Exec("DELETE FROM Accounts WHERE user_id = ?", id) //Deletes all the accounts of a user
	return &user
}
