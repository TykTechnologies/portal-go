// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	pathAccessRequests       = "/portal-api/access_requests"
	pathAccessRequest        = "/portal-api/access_requests/%d"
	pathAccessRequestApprove = "/portal-api/access_requests/%d/approve"
	pathAccessRequestReject  = "/portal-api/access_requests/%d/reject"
)

//go:generate mockery --name ARs --filename ars.go
type ARs interface {
	ListARs(ctx context.Context, opts ...Option) (*ListARsOutput, error)
	GetAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error)
	ApproveAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error)
	RejectAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error)
	DeleteAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error)
}

type ars struct {
	client *Client
}

// GetAccessRequest ...
func (p ars) GetAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathAccessRequest, id), nil, opts...)
	if err != nil {
		return nil, err
	}

	var ar ARDetails
	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &AROutput{
		Data: &ar,
	}, nil
}

// ListAccessRequests ...
func (p ars) ListARs(ctx context.Context, opts ...Option) (*ListARsOutput, error) {
	resp, err := p.client.doGet(ctx, pathAccessRequests, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ars []ARDetails

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListARsOutput{
		Data: ars,
	}, nil
}

type StatusOutput struct {
	Data     *Status
	Response *http.Response
}

// UpdateAccessRequest ...
func (p ars) ApproveAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequestApprove, id), nil, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ar Status

	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &StatusOutput{
		Data: &ar,
	}, nil
}

// UpdateAccessRequest ...
func (p ars) RejectAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequestReject, id), nil, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ar Status

	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &StatusOutput{
		Data: &ar,
	}, nil
}

// UpdateAccessRequest ...
func (p ars) DeleteAR(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error) {
	resp, err := p.client.doDelete(ctx, fmt.Sprintf(pathAccessRequest, id), nil, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ar Status
	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &StatusOutput{
		Data: &ar,
	}, nil
}

type ARDetails struct {
	AuthType             string        `json:"AuthType,omitempty"`
	Catalogue            string        `json:"Catalogue,omitempty"`
	Client               string        `json:"Client,omitempty"`
	CreatedAt            string        `json:"CreatedAt,omitempty"`
	Credentials          []Credentials `json:"Credentials,omitempty"`
	DCREnabled           bool          `json:"DCREnabled,omitempty"`
	DeletedAt            string        `json:"DeletedAt,omitempty"`
	ID                   int64         `json:"ID,omitempty"`
	Plan                 string        `json:"Plan,omitempty"`
	Products             []string      `json:"Products,omitempty"`
	ProvisionImmediately bool          `json:"ProvisionImmediately,omitempty"`
	Status               string        `json:"Status,omitempty"`
	UpdatedAt            string        `json:"UpdatedAt,omitempty"`
	User                 string        `json:"User,omitempty"`
}

func (a *ARDetails) UnmarshalJSON(b []byte) error {
	var customAR struct {
		AuthType             string        `json:"AuthType,omitempty"`
		Catalogue            string        `json:"Catalogue,omitempty"`
		Client               string        `json:"Client,omitempty"`
		CreatedAt            string        `json:"CreatedAt,omitempty"`
		Credentials          []Credentials `json:"Credentials,omitempty"`
		DCREnabled           bool          `json:"DCREnabled,omitempty"`
		DeletedAt            string        `json:"DeletedAt,omitempty"`
		ID                   int64         `json:"ID,omitempty"`
		Plan                 string        `json:"Plan,omitempty"`
		Products             interface{}   `json:"Products,omitempty"`
		ProvisionImmediately bool          `json:"ProvisionImmediately,omitempty"`
		Status               string        `json:"Status,omitempty"`
		UpdatedAt            string        `json:"UpdatedAt,omitempty"`
		User                 string        `json:"User,omitempty"`
	}

	if err := json.Unmarshal(b, &customAR); err != nil {
		return err
	}

	a.AuthType = customAR.AuthType
	a.Catalogue = customAR.Catalogue
	a.Client = customAR.Client
	a.CreatedAt = customAR.CreatedAt
	a.Credentials = customAR.Credentials
	a.DCREnabled = customAR.DCREnabled
	a.DeletedAt = customAR.DeletedAt
	a.ID = customAR.ID
	a.Plan = customAR.Plan
	a.ProvisionImmediately = customAR.ProvisionImmediately
	a.Status = customAR.Status
	a.UpdatedAt = customAR.UpdatedAt
	a.User = customAR.User

	switch k := customAR.Products.(type) {
	case []string:
		a.Products = k
	case string:
		a.Products = []string{k}
	}

	return nil
}

type Credentials struct {
	AccessRequest              string     `json:"AccessRequest,omitempty"`
	Credential                 string     `json:"Credential,omitempty"`
	CredentialHash             string     `json:"CredentialHash,omitempty"`
	DCRRegistrationAccessToken string     `json:"DCRRegistrationAccessToken,omitempty"`
	DCRRegistrationClientURI   string     `json:"DCRRegistrationClientURI,omitempty"`
	DCRResponse                string     `json:"DCRResponse,omitempty"`
	Expires                    CustomTime `json:"Expires,omitempty"`
	OAuthClientID              string     `json:"OAuthClientID,omitempty"`
	OAuthClientSecret          string     `json:"OAuthClientSecret,omitempty"`
	RedirectURI                string     `json:"RedirectURI,omitempty"`
	ResponseType               string     `json:"ResponseType,omitempty"`
	Scope                      string     `json:"Scope,omitempty"`
	TokenEndpoints             string     `json:"TokenEndpoints,omitempty"`
	GrantType                  *string    `json:"GrantType,omitempty"`
	ID                         *int64     `json:"ID,omitempty"`
}

type ListARsOutput struct {
	Data     []ARDetails
	Response *http.Response
}

type AROutput struct {
	Data     *ARDetails
	Response *http.Response
}
