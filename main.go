package main

import (
	"Users/grpc_cont"
	"Users/interactor"
	"Users/mysql_gateway"

	//"Users/rest_controller"
	"fmt"
)

func main() {
	gateway := mysql_gateway.Gateway{}
	interactor := interactor.Interactor{My_gateway: &gateway}
	//rest_controller.InitCont(interactor)
	grpc_cont.InitGRPC(interactor)
	fmt.Println("This is still going on")
}

/*
curl commands:
curl localhost:8080/users
curl localhost:8080/user?id=1
curl localhost:8080/user --header "Content-Type: application/json" -d @user_data.json --request "POST"
curl localhost:8080/user?id=1 --request "DELETE"
curl localhost:8080/account --header "Content-Type: application/json" -d @account_data.json --request "POST"
curl "localhost:8080/account?id=3&account_number=4321" --request "DELETE"

/*
TODO:
/Make sure that the error return values are received correctly to the client and the unit tests
/The controller doesn't need to know about the gateway
/I need to put some kind of verification on the interactor
/	-Maybe check for account number validity (exactly 4 numbers)?
/	-Maybe check that first_name is shorter than their last_name?
I need to use Rest and GRPC for the controllers
Faire un pseudo-client rest et grpc pour pouvoir faire tourner les tests
I need to do unit testing
	-Check getUsers
	-Check getUser
	-Check addUser
	-Check deleteUser
	-Check addAccount
	-Check deleteAccount
	-Check addUser with invalid name data (longer first_name than last_name)
	-Check addAccount with invalid data (Exactly 4 numbers)
	-Check addAccount with invalid data (Exactly letters and numbers)
*/
