package main

import (
	"Users/grpc_cont/GRPC"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Grpc_GetUsers()
	//Grpc_GetUser(2)
	//Grpc_CreateUser(5, "abc", "d")
	Grpc_DeleteUser(5)
}

func Grpc_GetUsers() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.GetUsersRequest{}

	response, err := b.GetUsers(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling GetUsers: %s", err)
	}

	for _, user := range response.User {
		fmt.Println(user)
	}
}

func Grpc_GetUser(id int64) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.GetUserRequest{Id: id}

	response, err := b.GetUser(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	fmt.Println(response.User)
}

func Grpc_CreateUser(id int, first_name string, last_name string) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	newUser := GRPC.GRPCUser{Id: int64(id), FirstName: first_name, LastName: last_name}
	request := GRPC.CreateUserRequest{User: &newUser}

	response, err := b.CreateUser(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}

	fmt.Println(response.User)
}

func Grpc_DeleteUser(id int64) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.DeleteUserRequest{Id: id}

	response, err := b.DeleteUser(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling DeleteUser: %s", err)
	}

	fmt.Println(response.User)
}

func Grpc_CreateAccount(user_id int, account_number int) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	newAccount := GRPC.GRPCAccount{UserId: int64(user_id), AccountNumber: int64(account_number)}
	request := GRPC.CreateAccountRequest{Account: &newAccount}

	response, err := b.CreateAccount(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}

	fmt.Println(response.User)
}
