package models

type Account struct {
	User_id        int `json:"id"`
	Account_number int `json:"account_number"`
}

type User struct {
	Id         int       `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Accounts   []Account `json:"accounts"`
}
