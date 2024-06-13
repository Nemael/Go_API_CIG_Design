package grpc_cont

import (
	GRPC "Users/grpc_cont/GRPC"
	"Users/interactor"
	"Users/models"
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

var My_interactor interactor.Interactor

type myUsersServer struct {
	GRPC.UnimplementedUsersServer
}

func (s myUsersServer) GetUsers(ctx context.Context, req *GRPC.GetUsersRequest) (*GRPC.GetUsersResponse, error) {
	my_users := My_interactor.GetUsers()
	if my_users == nil {
		err := errors.New("no users")
		return nil, err
	}
	var grpc_users []*GRPC.GRPCUser
	for _, my_user := range my_users {
		var grpc_user GRPC.GRPCUser
		grpc_user.Id = int64(my_user.Id)
		grpc_user.FirstName = my_user.First_name
		grpc_user.LastName = my_user.Last_name
		var grpc_accounts []*GRPC.GRPCAccount
		for _, my_account := range my_user.Accounts {
			var grpc_account GRPC.GRPCAccount
			grpc_account.UserId = int64(my_account.User_id)
			grpc_account.AccountNumber = int64(my_account.Account_number)
			grpc_accounts = append(grpc_accounts, &grpc_account)
		}
		grpc_user.Accounts = grpc_accounts
		grpc_users = append(grpc_users, &grpc_user)

	}

	return &GRPC.GetUsersResponse{
		User: grpc_users,
	}, nil
}

func (s myUsersServer) GetUser(ctx context.Context, req *GRPC.GetUserRequest) (*GRPC.GetUserResponse, error) {
	my_user := My_interactor.GetUser(int(req.Id))
	if my_user == nil {
		err := errors.New("no user")
		return nil, err
	}
	var grpc_user GRPC.GRPCUser
	grpc_user.Id = int64(my_user.Id)
	grpc_user.FirstName = my_user.First_name
	grpc_user.LastName = my_user.Last_name
	var grpc_accounts []*GRPC.GRPCAccount
	for _, my_account := range my_user.Accounts {
		var grpc_account GRPC.GRPCAccount
		grpc_account.UserId = int64(my_account.User_id)
		grpc_account.AccountNumber = int64(my_account.Account_number)
		grpc_accounts = append(grpc_accounts, &grpc_account)
	}
	grpc_user.Accounts = grpc_accounts

	return &GRPC.GetUserResponse{
		User: &grpc_user,
	}, nil
}

func (s myUsersServer) CreateUser(ctx context.Context, req *GRPC.CreateUserRequest) (*GRPC.CreateUserResponse, error) {
	newUser := models.User{Id: int(req.User.Id), First_name: req.User.FirstName, Last_name: req.User.LastName}
	returnedUser, err := My_interactor.AddUser(&newUser)
	if err != nil {
		return nil, err
	}
	grpc_user := GRPC.GRPCUser{Id: int64(returnedUser.Id), FirstName: newUser.First_name, LastName: newUser.Last_name}
	return &GRPC.CreateUserResponse{
		User: &grpc_user,
	}, nil
}

func (s myUsersServer) DeleteUser(ctx context.Context, req *GRPC.DeleteUserRequest) (*GRPC.DeleteUserResponse, error) {
	my_user := My_interactor.DeleteUser(int(req.Id))
	if my_user == nil {
		err := errors.New("no user")
		return nil, err
	}
	var grpc_user GRPC.GRPCUser
	grpc_user.Id = int64(my_user.Id)
	grpc_user.FirstName = my_user.First_name
	grpc_user.LastName = my_user.Last_name
	var grpc_accounts []*GRPC.GRPCAccount
	for _, my_account := range my_user.Accounts {
		var grpc_account GRPC.GRPCAccount
		grpc_account.UserId = int64(my_account.User_id)
		grpc_account.AccountNumber = int64(my_account.Account_number)
		grpc_accounts = append(grpc_accounts, &grpc_account)
	}
	grpc_user.Accounts = grpc_accounts

	return &GRPC.DeleteUserResponse{
		User: &grpc_user,
	}, nil
}

func (s myUsersServer) CreateAccount(ctx context.Context, req *GRPC.CreateAccountRequest) (*GRPC.CreateAccountResponse, error) {
	newAccount := models.Account{User_id: int(req.Account.UserId), Account_number: int(req.Account.AccountNumber)}
	returnedAccount, err := My_interactor.AddAccount(&newAccount)
	if err != nil {
		return nil, err
	}
	grpc_account := GRPC.GRPCAccount{UserId: int64(returnedAccount.User_id), AccountNumber: int64(returnedAccount.Account_number)}
	return &GRPC.CreateAccountResponse{
		Account: &grpc_account,
	}, nil
}

/*func (s myUsersServer) DeleteAccount(ctx context.Context, req *users.DeleteAccountRequest) (*users.DeleteAccountResponse, error) {
	var book books.Book
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	id := req.Id // Query parameter
	row := db.QueryRow("SELECT id, title, author, quantity FROM books where id = ?", id)
	err = row.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	book.Quantity += int64(1)

	db.QueryRow("UPDATE books SET quantity = ? WHERE id = ?", book.Quantity, book.Id)
	return &books.ReturnBookResponse{
		Book: &book,
	}, nil
}*/

func InitGRPC(interactor interactor.Interactor) {
	My_interactor = interactor
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cannot create listener %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myUsersServer{}

	GRPC.RegisterUsersServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
