package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	pathUsers = "/portal-api/users"
	pathUser  = "/portal-api/users/%d"
)

//go:generate mockery --name UsersService --filename catagogues_service.go
type UsersService interface {
	CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
	GetUser(ctx context.Context, id int64) (*GetUserOutput, error)
	ListUsers(ctx context.Context, options *ListUsersOptions) (*ListUsersOutput, error)
	UpdateUser(ctx context.Context, id int64, input UpdateUserInput) (*UpdateUserOutput, error)
}

type usersService struct {
	client *Client
}

func (p usersService) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathUsers, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var user User

	if err := resp.Parse(&user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		Data: &user,
	}, nil
}

func (p usersService) GetUser(ctx context.Context, id int64) (*GetUserOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathUser, id), nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := resp.Parse(&user); err != nil {
		return nil, err
	}

	return &GetUserOutput{
		Data: &user,
	}, nil
}

func (p usersService) ListUsers(ctx context.Context, options *ListUsersOptions) (*ListUsersOutput, error) {
	resp, err := p.client.doGet(pathUsers, nil)
	if err != nil {
		return nil, err
	}

	var users []User

	if err := resp.Parse(&users); err != nil {
		return nil, err
	}

	return &ListUsersOutput{
		Users: users,
	}, nil
}

func (p usersService) UpdateUser(ctx context.Context, id int64, input UpdateUserInput) (*UpdateUserOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathUser, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var user User

	if err := resp.Parse(&user); err != nil {
		return nil, err
	}

	return &UpdateUserOutput{
		Data: &user,
	}, nil
}

type UserInput struct {
	ID   *int64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateUserInput = UserInput

type CreateUserInput = UserInput

type ListUsersOptions struct{}

type ListUsersOutput struct {
	Users []User
}

type User struct {
	ID             int64
	Active         bool
	Email          string
	First          string
	Last           string
	OrganisationID int64
	Provider       string
	Role           string
}

type UserOutput struct {
	Data *User
}

type UpdateUserOutput = UserOutput

type GetUserOutput = UserOutput

type CreateUserOutput = UserOutput
