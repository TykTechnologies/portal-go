// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type UsersService service

type User struct {
	Active            bool     `json:"Active,omitempty"`
	Email             string   `json:"Email,omitempty"`
	First             string   `json:"First,omitempty"`
	Last              string   `json:"Last,omitempty"`
	OrgID             int64    `json:"OrganisationID,omitempty"`
	Role              string   `json:"Role,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	JWTToken          string   `json:"JWTToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Org               string   `json:"Organisation,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	CreatedAt         string   `json:"CreatedAt,omitempty"`
	UpdatedAt         string   `json:"UpdatedAt,omitempty"`
}

type Validator interface {
	Validate() error
}

type userInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
	Role          string `json:"Role,omitempty"`
}

func (u userInput) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	return nil
}

func (u *UsersService) ListUsers(ctx context.Context, opts *ListOptions) ([]*User, *Response, error) {
	urlPath := "/users"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*User

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *UsersService) CreateUser(ctx context.Context, input *User) (*User, *Response, error) {
	urlPath := "/users"

	userReq := &userInput{
		Active:        input.Active,
		Email:         input.Email,
		First:         input.First,
		Last:          input.Last,
		OrgID:         input.OrgID,
		Provider:      input.Provider,
		ResetPassword: input.ResetPassword,
		Role:          input.Role,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *UsersService) GetUser(ctx context.Context, userID int64) (*User, *Response, error) {
	urlPath := fmt.Sprintf("/users/%v", userID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *UsersService) UpdateUser(ctx context.Context, userID int64, input *User) (*User, *Response, error) {
	urlPath := fmt.Sprintf("/users/%v", userID)

	userReq := &userInput{
		First: input.First,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *UsersService) DeleteUser(ctx context.Context, userID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/users/%v", userID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
