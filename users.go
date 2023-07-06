package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathUsers = "/portal-api/users"
	pathUser  = "/portal-api/users/%d"
)

type usersService struct {
	client *Client
}

func (p usersService) CreateUser(input CreateUserInput) (*CreateUserOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathUsers, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{}, nil
}

func (p usersService) GetUser(id uint64) (*GetUserOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathUser, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{}, nil
}

func (p usersService) ListUsers(options *ListUsersOptions) (*ListUsersOutput, error) {
	req, err := p.client.newGetRequest(pathUsers, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListUsersOutput{}, nil
}

func (p usersService) UpdateUser(id uint64, input UpdateUserInput) (*UpdateUserOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathUser, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateUserOutput{}, nil
}

type UpdateUserInput struct {
	Catalogues []uint64
}

type CreateUserInput struct{}

type ListUsersOptions struct{}

type ListUsersOutput struct{}

type User struct{}

type UserOutput struct{}

type UpdateUserOutput = UserOutput

type GetUserOutput = UserOutput

type CreateUserOutput = UserOutput
