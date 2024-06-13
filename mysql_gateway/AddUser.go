package mysql_gateway

import (
	"Users/models"
	"fmt"
)

func (g *Gateway) AddUser(user *models.User) *models.User {
	fmt.Println("Into gateway")
	db, err := getDB()
	if err != nil {
		fmt.Println("Error getting database")
		return nil
	}
	_, err = db.Exec("INSERT INTO users (id, first_name, last_name) VALUES (?, ?, ?)", user.Id, user.First_name, user.Last_name)
	if err != nil {
		fmt.Println("Error inserting data", err)
		return nil
	}
	return user
}
