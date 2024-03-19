// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type ARsService service

type AR struct {
	ID                   int64         `json:"ID,omitempty"`
	AuthType             string        `json:"AuthType,omitempty"`
	Catalogue            string        `json:"Catalogue,omitempty"`
	Client               string        `json:"Client,omitempty"`
	CreatedAt            string        `json:"CreatedAt,omitempty"`
	Credentials          []Credentials `json:"Credentials,omitempty"`
	DCREnabled           bool          `json:"DCREnabled,omitempty"`
	DeletedAt            string        `json:"DeletedAt,omitempty"`
	Plan                 string        `json:"Plan,omitempty"`
	Products             []string      `json:"Products,omitempty"`
	ProvisionImmediately bool          `json:"ProvisionImmediately,omitempty"`
	Status               string        `json:"Status,omitempty"`
	UpdatedAt            string        `json:"UpdatedAt,omitempty"`
	User                 string        `json:"User,omitempty"`
}

type Credentials struct {
	AccessRequest              string  `json:"AccessRequest,omitempty"`
	Credential                 string  `json:"Credential,omitempty"`
	CredentialHash             string  `json:"CredentialHash,omitempty"`
	DCRRegistrationAccessToken string  `json:"DCRRegistrationAccessToken,omitempty"`
	DCRRegistrationClientURI   string  `json:"DCRRegistrationClientURI,omitempty"`
	DCRResponse                string  `json:"DCRResponse,omitempty"`
	Expires                    Time    `json:"Expires,omitempty"`
	OAuthClientID              string  `json:"OAuthClientID,omitempty"`
	OAuthClientSecret          string  `json:"OAuthClientSecret,omitempty"`
	RedirectURI                string  `json:"RedirectURI,omitempty"`
	ResponseType               string  `json:"ResponseType,omitempty"`
	Scope                      string  `json:"Scope,omitempty"`
	TokenEndpoints             string  `json:"TokenEndpoints,omitempty"`
	GrantType                  *string `json:"GrantType,omitempty"`
	ID                         *int64  `json:"ID,omitempty"`
}

type arInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Role          string `json:"Role,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}

func (u *ARsService) ListARs(ctx context.Context, opts *ListOptions) ([]*AR, *Response, error) {
	urlPath := "/access_requests"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var ars []*AR

	resp, err := u.client.Do(ctx, req, &ars)
	if err != nil {
		return nil, resp, err
	}

	return ars, resp, nil
}

func (u *ARsService) CreateAR(ctx context.Context, input *AR) (*AR, *Response, error) {
	urlPath := "/access_requests"

	arReq := &arInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, arReq)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AR)

	resp, err := u.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}

func (u *ARsService) GetAR(ctx context.Context, arID int64) (*AR, *Response, error) {
	urlPath := fmt.Sprintf("/access_requests/%v", arID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AR)

	resp, err := u.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}

func (u *ARsService) DeleteAR(ctx context.Context, arID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/access_requests/%v", arID)

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

func (u *ARsService) ApproveAR(ctx context.Context, arID int64) (*AR, *Response, error) {
	urlPath := fmt.Sprintf("/access_requests/%v/approve", arID)

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AR)

	resp, err := u.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}

func (u *ARsService) RejectAR(ctx context.Context, arID int64) (*AR, *Response, error) {
	urlPath := fmt.Sprintf("/access_requests/%v/reject", arID)

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AR)

	resp, err := u.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}
