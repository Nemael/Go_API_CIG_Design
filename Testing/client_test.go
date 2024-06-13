package main

import (
	"Users/grpc_cont/GRPC"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Grpc_GetUsers()
	//Grpc_GetUser(2)
	//Grpc_CreateUser(5, "abc", "d")
	//Grpc_DeleteUser(5)
	//Grpc_CreateAccount(2, 1111)
	//Grpc_DeleteAccount(2, 1111)
	//Rest_GetUsers()
	//Rest_GetUser(2)
	//Rest_CreateUser(5, "abc", "d")
	//Rest_DeleteUser(5)
	//Rest_CreateAccount(2, 1111)
	//Rest_DeleteAccount(2, 1111)
}

func Grpc_GetUsers() string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.GetUsersRequest{}

	response, err := b.GetUsers(context.Background(), &request)
	if err != nil {
		return ("Error when calling GetUsers: " + err.Error())
	}
	return response.User[0].LastName
}

func Grpc_GetUser(id int64) string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.GetUserRequest{Id: id}

	response, err := b.GetUser(context.Background(), &request)
	if err != nil {
		return ("Error when calling GetUser: " + err.Error())
	}

	return (response.User.LastName)
}

func Grpc_CreateUser(id int, first_name string, last_name string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()
	b := GRPC.NewUsersClient(conn)
	newUser := GRPC.GRPCUser{Id: int64(id), FirstName: first_name, LastName: last_name}
	request := GRPC.CreateUserRequest{User: &newUser}

	response, err := b.CreateUser(context.Background(), &request)
	if err != nil {
		return ("Error when calling CreateUsert: " + err.Error())
	}

	return (response.User.LastName)
}

func Grpc_DeleteUser(id int64) string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.DeleteUserRequest{Id: id}

	response, err := b.DeleteUser(context.Background(), &request)
	if err != nil {
		return ("Error when calling DeleteUser: " + err.Error())
	}

	return (response.User.LastName)
}

func Grpc_CreateAccount(user_id int, account_number int) string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	newAccount := GRPC.GRPCAccount{UserId: int64(user_id), AccountNumber: int64(account_number)}
	request := GRPC.CreateAccountRequest{Account: &newAccount}

	response, err := b.CreateAccount(context.Background(), &request)
	if err != nil {
		return ("Error when calling CreateUser:" + err.Error())
	}

	return (strconv.Itoa(int(response.Account.AccountNumber)))
}

func Grpc_DeleteAccount(user_id int64, account_number int64) string {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ("Could not connect: " + err.Error())
	}
	defer conn.Close()

	b := GRPC.NewUsersClient(conn)
	request := GRPC.DeleteAccountRequest{UserId: user_id, AccountNumber: account_number}

	_, err = b.DeleteAccount(context.Background(), &request)
	if err != nil {
		return ("Error when calling DeleteUser: " + err.Error())
	}

	return ("Account deleted succesfully")
}

func Rest_GetUsers() string {
	client := &http.Client{}
	url := "http://localhost:8080/users"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

func Rest_GetUser(id int) string {
	client := &http.Client{}
	url := "http://localhost:8080/user?id=" + strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

func Rest_CreateUser(id int, first_name string, last_name string) string {
	client := &http.Client{}

	data := map[string]interface{}{
		"id":         id,
		"first_name": first_name,
		"last_name":  last_name,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		// handle error
		panic(err)
	}

	url := "http://localhost:8080/user"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

func Rest_DeleteUser(id int) string {
	client := &http.Client{}

	url := "http://localhost:8080/user?id=" + strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

func Rest_CreateAccount(user_id int, account_number int) string {
	client := &http.Client{}

	data := map[string]interface{}{
		"id":             user_id,
		"account_number": account_number,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		// handle error
		panic(err)
	}

	url := "http://localhost:8080/account"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

func Rest_DeleteAccount(id int, account_number int) string {
	client := &http.Client{}

	url := "http://localhost:8080/account?id=" + strconv.Itoa(id) + "&account_number=" + strconv.Itoa(account_number)
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return (err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return (err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(body))
}

// REST tests
func TestRestGetUsers(t *testing.T) {
	fmt.Println("Test 1")
	Rest_CreateUser(1, "Jeanne", "Kritzteker")
	res := Rest_GetUsers()
	if !strings.Contains(res, "Kritzteker") {
		t.Fatalf(`Error in getUser test`)
	}
	Rest_DeleteUser(1)
}

func TestRestCreateUser(t *testing.T) {
	fmt.Println("Test 2")
	Rest_CreateUser(1, "Jeanne", "Kritzteker")
	res := Rest_GetUsers()
	if !strings.Contains(res, "Kritzteker") {
		t.Fatalf(`Error in createUser test`)
	}
	Rest_DeleteUser(1)
}

func TestRestCreateInvalidUser(t *testing.T) {
	fmt.Println("Test 3")
	res := Rest_CreateUser(1, "this_is_a_longer_name", "shorter_name")
	if !strings.Contains(res, "User not created: user first_name must be shorter than user last_name") {
		t.Fatalf(`Error in createInvalidUser test`)
	}
}

func TestRestCreateAccount(t *testing.T) {
	fmt.Println("Test 4")
	Rest_CreateUser(1, "Jeanne", "Kritzteker")
	res := Rest_CreateAccount(1, 1234)
	if !strings.Contains(res, "1234") {
		t.Fatalf(`Error in createAccount test`)
	}
	Rest_DeleteAccount(1, 1234)
	Rest_DeleteUser(1)
}

func TestRestCreateInvalidAccount(t *testing.T) {
	fmt.Println("Test 5")
	Rest_CreateUser(1, "Jeanne", "Kritzteker")
	res := Rest_CreateAccount(1, 123)
	if !strings.Contains(res, "Account not created: account number must be 4 digits long") {
		t.Fatalf(`Error in createInvalidAccount test version 1`)
	}
	res = Rest_CreateAccount(1, 12345)
	if !strings.Contains(res, "Account not created: account number must be 4 digits long") {
		t.Fatalf(`Error in createInvalidAccount test version 1`)
	}
	Rest_DeleteUser(1)
}

// GRPC tests
func TestGrpcGetUsers(t *testing.T) {
	fmt.Println("Test 6")
	Grpc_CreateUser(1, "Jeanne", "Kritzteker")
	res := Grpc_GetUsers()
	if !strings.Contains(res, "Kritzteker") {
		t.Fatalf(`Error in getUser test`)
	}
	Grpc_DeleteUser(1)
}

func TestGrpcCreateUser(t *testing.T) {
	fmt.Println("Test 7")
	Grpc_CreateUser(1, "Jeanne", "Kritzteker")
	res := Grpc_GetUsers()
	if !strings.Contains(res, "Kritzteker") {
		t.Fatalf(`Error in createUser test`)
	}
	Grpc_DeleteUser(1)
}

func TestGrpcCreateInvalidUser(t *testing.T) {
	fmt.Println("Test 8")
	res := Grpc_CreateUser(1, "this_is_a_longer_name", "shorter_name")
	if !strings.Contains(res, "user first_name must be shorter than user last_name") {
		t.Fatalf(`Error in createInvalidUser test`)
	}
}

func TestGrpcCreateAccount(t *testing.T) {
	fmt.Println("Test 9")
	Grpc_CreateUser(1, "Jeanne", "Kritzteker")
	res := Grpc_CreateAccount(1, 1234)
	if !strings.Contains(res, "1234") {
		t.Fatalf(`Error in createAccount test`)
	}
	Grpc_DeleteAccount(1, 1234)
	Grpc_DeleteUser(1)
}

func TestGrpcCreateInvalidAccount(t *testing.T) {
	fmt.Println("Test 10")
	Grpc_CreateUser(1, "Jeanne", "Kritzteker")
	res := Grpc_CreateAccount(1, 123)
	if !strings.Contains(res, "account number must be 4 digits long") {
		t.Fatalf(`Error in createInvalidAccount test version 1`)
	}
	res = Rest_CreateAccount(1, 12345)
	if !strings.Contains(res, "account number must be 4 digits long") {
		t.Fatalf(`Error in createInvalidAccount test version 1`)
	}
	Rest_DeleteUser(1)
}

// Tests ideas
// 	- Create multiple accounts linked to one user
//  - DeleteUser
//  - GetUser
//  - DeleteAcount
