package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	pathProviders    = "/portal-api/providers"
	pathProvider     = "/portal-api/providers/%d"
	pathProviderSync = "/portal-api/providers/%v/synchronize"
)

//go:generate mockery --name ProvidersService --filename providers_service.go
type ProvidersService interface {
	CreateProvider(ctx context.Context, input *CreateProviderInput, opts ...func(*Options)) (*CreateProviderOutput, error)
	GetProvider(ctx context.Context, id int64, opts ...func(*Options)) (*GetProviderOutput, error)
	ListProviders(ctx context.Context, options *ListProvidersOptions, opts ...func(*Options)) (*ListProvidersOutput, error)
	UpdateProvider(ctx context.Context, id int64, input *UpdateProviderInput, opts ...func(*Options)) (*UpdateProviderOutput, error)
	SyncProvider(ctx context.Context, id int64, opts ...func(*Options)) (*SyncProviderOutput, error)
}

type providersService struct {
	client *Client
}

func (p providersService) CreateProvider(ctx context.Context, input *CreateProviderInput, opts ...func(*Options)) (*CreateProviderOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathProviders, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider

	if err := resp.Parse(&provider); err != nil {
		return nil, err
	}

	return &CreateProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) GetProvider(ctx context.Context, id int64, opts ...func(*Options)) (*GetProviderOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathProvider, id), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider
	if err := resp.Parse(&provider); err != nil {
		return nil, err
	}

	return &GetProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) ListProviders(ctx context.Context, options *ListProvidersOptions, opts ...func(*Options)) (*ListProvidersOutput, error) {
	resp, err := p.client.doGet(pathProviders, nil)
	if err != nil {
		return nil, err
	}

	var providers []Provider

	if err := resp.Parse(&providers); err != nil {
		return nil, err
	}

	return &ListProvidersOutput{
		Data: providers,
	}, nil
}

func (p providersService) UpdateProvider(ctx context.Context, id int64, input *UpdateProviderInput, opts ...func(*Options)) (*UpdateProviderOutput, error) {
	// TODO: review this
	if input.Configuration != nil && input.Configuration.ID == nil {
		return nil, errors.New("configuration id must not be nil")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathProvider, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider

	if err := resp.Parse(&provider); err != nil {
		return nil, err
	}

	return &UpdateProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) SyncProvider(ctx context.Context, id int64, opts ...func(*Options)) (*SyncProviderOutput, error) {
	resp, err := p.client.doPut(fmt.Sprintf(pathProviderSync, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var msg SyncStatus

	if err := resp.Parse(&msg); err != nil {
		return nil, err
	}

	return &SyncProviderOutput{
		Data: msg,
	}, nil
}

func (p providersService) SyncProviders(ctx context.Context, opts ...func(*Options)) (*SyncProviderOutput, error) {
	resp, err := p.client.doPut(fmt.Sprintf(pathProviderSync, "all"), nil, nil)
	if err != nil {
		return nil, err
	}

	var msg SyncStatus

	if err := resp.Parse(&msg); err != nil {
		return nil, err
	}

	return &SyncProviderOutput{
		Data: msg,
	}, nil
}

type ProviderInput struct {
	ID            *int64 `json:",omitempty"`
	Type          string
	Name          string
	Configuration *ProviderConfig `json:",omitempty"`
}

type UpdateProviderInput = ProviderInput

type CreateProviderInput = ProviderInput

type ProviderConfig struct {
	ProviderID *int64 `json:"ProviderID,omitempty"`
	MetaData   string
	ID         *int64 `json:"ID,omitempty"`
}

type ListProvidersOptions struct{}

type ListProvidersOutput struct {
	Data []Provider
}

type Provider struct {
	ID            int64
	Name          string
	CreatedAt     string
	UpdatedAt     string
	Type          string
	Status        string
	LastSynced    string
	Configuration ProviderConfig
}

type ProviderOutput struct {
	Provider *Provider
}

type SyncProviderOutput struct {
	Data SyncStatus
}

type UpdateProviderOutput = ProviderOutput

type GetProviderOutput = ProviderOutput

type CreateProviderOutput = ProviderOutput

type SyncStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
