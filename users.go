package edp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	pathUsers = "/portal-api/users"
	pathUser  = "/portal-api/users/%d"
)

//go:generate mockery --name UsersService --filename users_service.go
type UsersService interface {
	CreateUser(ctx context.Context, input *CreateUserInput, opts ...Option) (*CreateUserOutput, error)
	GetUser(ctx context.Context, id int64, opts ...Option) (*GetUserOutput, error)
	ListUsers(ctx context.Context, options *ListUsersInput, opts ...Option) (*ListUsersOutput, error)
	UpdateUser(ctx context.Context, id int64, input *UpdateUserInput, opts ...Option) (*UpdateUserOutput, error)
	DeleteUser(ctx context.Context, id int64, opts ...Option) (*DeleteUserOutput, error)
}

type usersService struct {
	client *Client
}

func (p usersService) CreateUser(ctx context.Context, input *CreateUserInput, opts ...Option) (*CreateUserOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	if !p.client.skipValidation {
		if err := input.validate(); err != nil {
			return nil, err
		}
	}

	resp, err := p.client.doPost(ctx, pathUsers, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var user User

	if err := resp.Unmarshal(&user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		Data: &user,
	}, nil
}

func (p usersService) GetUser(ctx context.Context, id int64, opts ...Option) (*GetUserOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathUser, id), nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := resp.Unmarshal(&user); err != nil {
		return nil, err
	}

	return &GetUserOutput{
		Data: &user,
	}, nil
}

func (p usersService) ListUsers(ctx context.Context, options *ListUsersInput, opts ...Option) (*ListUsersOutput, error) {
	resp, err := p.client.doGet(ctx, pathUsers, nil)
	if err != nil {
		return nil, err
	}

	var users []User

	if err := resp.Unmarshal(&users); err != nil {
		return nil, err
	}

	return &ListUsersOutput{
		Users: users,
	}, nil
}

func (p usersService) UpdateUser(ctx context.Context, id int64, input *UpdateUserInput, opts ...Option) (*UpdateUserOutput, error) {
	input.ID = nil

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathUser, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var user User

	if err := resp.Unmarshal(&user); err != nil {
		return nil, err
	}

	return &UpdateUserOutput{
		Data: &user,
	}, nil
}

func (p usersService) DeleteUser(ctx context.Context, id int64, opts ...Option) (*DeleteUserOutput, error) {
	_, err := p.client.doDelete(ctx, fmt.Sprintf(pathUser, id), nil, nil)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{}, nil
}

type UserInput struct {
	ID             *int64 `json:"ID,omitempty"`
	Active         bool   `json:"Active,omitempty"`
	Email          string `json:"Email,omitempty"`
	First          string `json:"First,omitempty"`
	Last           string `json:"Last,omitempty"`
	OrganisationID int64  `json:"OrganisationID,omitempty"`
	Role           string `json:"Role,omitempty"`
	Provider       string `json:"Provider,omitempty"`
	ResetPassword  bool   `json:"ResetPassword,omitempty"`
}

func (u UserInput) validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.First == "" {
		return errors.New("first is required")
	}

	return nil
}

type CreateUserInput = UserInput

type UpdateUserInput = UserInput

type ListUsersInput struct{}

type ListUsersOutput struct {
	Users []User
}

type User struct {
	Active            bool     `json:"Active,omitempty"`
	Email             string   `json:"Email,omitempty"`
	First             string   `json:"First,omitempty"`
	Last              string   `json:"Last,omitempty"`
	OrganisationID    int64    `json:"OrganisationID,omitempty"`
	Role              string   `json:"Role,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	JWTToken          string   `json:"JWTToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Organisation      string   `json:"Organisation,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	CreatedAt         string   `json:"CreatedAt,omitempty"`
	UpdatedAt         string   `json:"UpdatedAt,omitempty"`
}

type UserOutput struct {
	Data *User
}

type UpdateUserOutput = UserOutput

type GetUserOutput = UserOutput

type CreateUserOutput = UserOutput

type DeleteUserOutput = UserOutput
