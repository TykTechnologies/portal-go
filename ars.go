// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	pathAccessRequests       = "/portal-api/access_requests"
	pathAccessRequest        = "/portal-api/access_requests/%d"
	pathAccessRequestApprove = "/portal-api/access_requests/%d/approve"
	pathAccessRequestReject  = "/portal-api/access_requests/%d/reject"
)

//go:generate mockery --name ARsService --filename ars_service.go
type ARsService interface {
	ListARs(ctx context.Context, opts ...Option) (*ListARsOutput, error)
}

type arsService struct {
	client *Client
}

// GetAccessRequest ...
func (p arsService) GetAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error) {
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
func (p arsService) ListARs(ctx context.Context, opts ...Option) (*ListARsOutput, error) {
	resp, err := p.client.doGet(ctx, pathAccessRequests, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ars []ARSummary

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListARsOutput{
		Data: ars,
	}, nil
}

// UpdateAccessRequest ...
func (p arsService) ApproveAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequestApprove, id), nil, nil, opts...)
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

// UpdateAccessRequest ...
func (p arsService) RejectAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequestReject, id), nil, nil, opts...)
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

// UpdateAccessRequest ...
func (p arsService) DeleteAR(ctx context.Context, id int64, opts ...Option) (*AROutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequest, id), nil, nil, opts...)
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

type ARDetails struct {
	Catalogue            string        `json:"Catalogue,omitempty"`
	Client               string        `json:"Client,omitempty"`
	CreatedAt            string        `json:"CreatedAt,omitempty"`
	UpdatedAt            string        `json:"UpdatedAt,omitempty"`
	DeletedAt            string        `json:"DeletedAt,omitempty"`
	Plan                 string        `json:"Plan,omitempty"`
	User                 string        `json:"User,omitempty"`
	AuthType             string        `json:"AuthType,omitempty"`
	DCREnabled           bool          `json:"DCREnabled,omitempty"`
	ID                   int64         `json:"ID,omitempty"`
	ProvisionImmediately bool          `json:"ProvisionImmediately,omitempty"`
	Status               string        `json:"Status,omitempty"`
	Products             string        `json:"Products,omitempty"`
	Credentials          []Credentials `json:"Credentials,omitempty"`
}

type ARSummary struct {
	Catalogue            string `json:"Catalogue,omitempty"`
	Client               string `json:"Client,omitempty"`
	CreatedAt            string `json:"CreatedAt,omitempty"`
	UpdatedAt            string `json:"UpdatedAt,omitempty"`
	DeletedAt            string `json:"DeletedAt,omitempty"`
	Plan                 string `json:"Plan,omitempty"`
	User                 string `json:"User,omitempty"`
	AuthType             string `json:"AuthType,omitempty"`
	DCREnabled           bool   `json:"DCREnabled,omitempty"`
	ID                   int16  `json:"ID,omitempty"`
	ProvisionImmediately bool   `json:"ProvisionImmediately,omitempty"`
	Status               string `json:"Status,omitempty"`
	Products             string `json:"Products,omitempty"`
}

type Credentials struct {
	AccessRequest              string    `json:"AccessRequest,omitempty"`
	Credential                 string    `json:"Credential,omitempty"`
	CredentialHash             string    `json:"CredentialHash,omitempty"`
	DCRRegistrationAccessToken string    `json:"DCRRegistrationAccessToken,omitempty"`
	DCRRegistrationClientURI   string    `json:"DCRRegistrationClientURI,omitempty"`
	DCRResponse                string    `json:"DCRResponse,omitempty"`
	Expires                    time.Time `json:"Expires,omitempty"`
	OAuthClientID              string    `json:"OAuthClientID,omitempty"`
	OAuthClientSecret          string    `json:"OAuthClientSecret,omitempty"`
	RedirectURI                string    `json:"RedirectURI,omitempty"`
	ResponseType               string    `json:"ResponseType,omitempty"`
	Scope                      string    `json:"Scope,omitempty"`
	TokenEndpoints             string    `json:"TokenEndpoints,omitempty"`
	GrantType                  *string   `json:"GrantType,omitempty"`
	ID                         *int64    `json:"ID,omitempty"`
}

type ListARsOutput struct {
	Data     []ARSummary
	Response *http.Response
}

type AROutput struct {
	Data     *ARDetails
	Response *http.Response
}
