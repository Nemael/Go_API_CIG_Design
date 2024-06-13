package rest_controller

import (
	"Users/interactor"
	"Users/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Controller struct {
}

func getUsers(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	users := interactor.GetUsers()
	if users == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "No users found"}) //Data sent is the books struct
		return
	}
	c.IndentedJSON(http.StatusOK, users) //Data sent is the books struct
}

func getUser(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	var user *models.User
	str_id, ok := c.GetQuery("id") //Get inline query
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id is not a valid number"}) //gin.H is a shortcut to write custom json
		return
	}
	user = interactor.GetUser(id)
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not found"}) //Data sent is the books struct
		return
	}
	c.IndentedJSON(http.StatusOK, user) //Data sent is the books struct
}

func addUser(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	var user *models.User
	if err := c.BindJSON(&user); err != nil { //Binds the received values to newBook
		return //In case we get an error, BindJSON gives a return response
	}
	user, err := interactor.AddUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not created: " + err.Error()}) //Data sent is the books struct
		return
	}
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not created"}) //Data sent is the books struct
		return
	}
	c.IndentedJSON(http.StatusOK, user) //Data sent is the books struct
}

func addAccount(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	var account *models.Account
	if err := c.BindJSON(&account); err != nil { //Binds the received values to newBook
		return //In case we get an error, BindJSON gives a return response
	}
	account, err := interactor.AddAccount(account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Account not created: " + err.Error()}) //Data sent is the books struct
		return
	}
	if account == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Account not created"}) //Data sent is the books struct
		return
	}
	c.IndentedJSON(http.StatusOK, account) //Data sent is the books struct
}

func deleteUser(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	var user *models.User
	str_id, ok := c.GetQuery("id") //Get inline query
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id is not a valid number"}) //gin.H is a shortcut to write custom json
		return
	}
	user = interactor.DeleteUser(id)
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not deleted"}) //Data sent is the books struct
		return
	}
	c.IndentedJSON(http.StatusOK, user) //Data sent is the books struct
}

func deleteAccount(c *gin.Context, interactor interactor.Interactor) {
	fmt.Println("Into controller")
	str_id, ok := c.GetQuery("id") //Get inline query
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id is not a valid number"}) //gin.H is a shortcut to write custom json
		return
	}
	str_account_number, ok := c.GetQuery("account_number") //Get inline query
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing account_number query parameter."})
		return
	}
	account_number, err := strconv.Atoi(str_account_number)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account_number is not a valid number"}) //gin.H is a shortcut to write custom json
		return
	}
	interactor.DeleteAccount(id, account_number)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Account deleted succesfully"}) //Data sent is the books struct
}

func InitCont(interactor interactor.Interactor) {
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) { getUsers(c, interactor) })
	router.GET("/user", func(c *gin.Context) { getUser(c, interactor) })
	router.POST("/user", func(c *gin.Context) { addUser(c, interactor) })
	router.DELETE("/user", func(c *gin.Context) { deleteUser(c, interactor) })
	router.POST("/account", func(c *gin.Context) { addAccount(c, interactor) })
	router.DELETE("/account", func(c *gin.Context) { deleteAccount(c, interactor) })
	router.Run("localhost:8080")
}
