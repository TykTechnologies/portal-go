// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type ProvidersService service

type Provider struct {
	Configuration *ProviderConfig `json:"Configuration,omitempty"`
	CreatedAt     string          `json:"CreatedAt,omitempty"`
	ID            int64           `json:"ID,omitempty"`
	LastSynced    string          `json:"LastSynced,omitempty"`
	Name          string          `json:"Name,omitempty"`
	Status        string          `json:"Status,omitempty"`
	Type          string          `json:"Type,omitempty"`
	UpdatedAt     string          `json:"UpdatedAt,omitempty"`
}

type ProviderConfig struct {
	ID       *int64 `json:"ID,omitempty"`
	MetaData string `json:"MetaData,omitempty"`
}

type providerInput struct {
	ID            *int64          `json:"ID,omitempty"`
	Type          string          `json:"Type,omitempty"`
	Name          string          `json:"Name,omitempty"`
	Configuration *ProviderConfig `json:"Configuration,omitempty"`
}

type SyncStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (u *ProvidersService) ListProviders(ctx context.Context, opts *ListOptions) ([]*Provider, *Response, error) {
	urlPath := "/providers"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var providers []*Provider

	resp, err := u.client.Do(ctx, req, &providers)
	if err != nil {
		return nil, resp, err
	}

	return providers, resp, nil
}

func (u *ProvidersService) CreateProvider(ctx context.Context, input *Provider) (*Provider, *Response, error) {
	urlPath := "/providers"

	providerReq := &providerInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, providerReq)
	if err != nil {
		return nil, nil, err
	}

	provider := new(Provider)

	resp, err := u.client.Do(ctx, req, provider)
	if err != nil {
		return nil, resp, err
	}

	return provider, resp, nil
}

func (u *ProvidersService) GetProvider(ctx context.Context, provID int64) (*Provider, *Response, error) {
	urlPath := fmt.Sprintf("/providers/%v", provID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	provider := new(Provider)

	resp, err := u.client.Do(ctx, req, provider)
	if err != nil {
		return nil, resp, err
	}

	return provider, resp, nil
}

func (u *ProvidersService) UpdateProvider(ctx context.Context, provID int64, input *Provider) (*Provider, *Response, error) {
	urlPath := fmt.Sprintf("/providers/%v", provID)

	providerReq := &providerInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, providerReq)
	if err != nil {
		return nil, nil, err
	}

	provider := new(Provider)

	resp, err := u.client.Do(ctx, req, provider)
	if err != nil {
		return nil, resp, err
	}

	return provider, resp, nil
}

func (u *ProvidersService) DeleteProvider(ctx context.Context, provID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/providers/%v", provID)

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

func (u *ProvidersService) SyncProvider(ctx context.Context, provID int64) (*SyncStatus, *Response, error) {
	urlPath := fmt.Sprintf("/providers/%v/synchronize", provID)

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(SyncStatus)

	resp, err := u.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}

func (u *ProvidersService) SyncAllProviders(ctx context.Context) (*SyncStatus, *Response, error) {
	urlPath := "/providers/all/synchronize"

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(SyncStatus)

	resp, err := u.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
