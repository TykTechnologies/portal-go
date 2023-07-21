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
	CreateProvider(ctx context.Context, input *CreateProviderInput, opts ...Option) (*CreateProviderOutput, error)
	GetProvider(ctx context.Context, id int64, opts ...Option) (*GetProviderOutput, error)
	DeleteProvider(ctx context.Context, id int64, opts ...Option) (*DeleteProviderOutput, error)
	ListProviders(ctx context.Context, options *ListProvidersInput, opts ...Option) (*ListProvidersOutput, error)
	UpdateProvider(ctx context.Context, id int64, input *UpdateProviderInput, opts ...Option) (*UpdateProviderOutput, error)
	SyncProvider(ctx context.Context, id int64, opts ...Option) (*SyncProviderOutput, error)
}

type providersService struct {
	client *Client
}

func (p providersService) CreateProvider(ctx context.Context, input *CreateProviderInput, opts ...Option) (*CreateProviderOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathProviders, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider

	if err := resp.Unmarshal(&provider); err != nil {
		return nil, err
	}

	return &CreateProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) GetProvider(ctx context.Context, id int64, opts ...Option) (*GetProviderOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathProvider, id), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider
	if err := resp.Unmarshal(&provider); err != nil {
		return nil, err
	}

	return &GetProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) DeleteProvider(ctx context.Context, id int64, opts ...Option) (*DeleteProviderOutput, error) {
	_, err := p.client.doDelete(ctx, fmt.Sprintf(pathProvider, id), nil, nil)
	if err != nil {
		return nil, err
	}

	return &GetProviderOutput{}, nil
}

func (p providersService) ListProviders(ctx context.Context, options *ListProvidersInput, opts ...Option) (*ListProvidersOutput, error) {
	resp, err := p.client.doGet(ctx, pathProviders, nil)
	if err != nil {
		return nil, err
	}

	var providers []Provider

	if err := resp.Unmarshal(&providers); err != nil {
		return nil, err
	}

	return &ListProvidersOutput{
		Data: providers,
	}, nil
}

func (p providersService) UpdateProvider(ctx context.Context, id int64, input *UpdateProviderInput, opts ...Option) (*UpdateProviderOutput, error) {
	// TODO: review this
	if input.Configuration != nil && input.Configuration.ID == nil {
		return nil, errors.New("configuration id must not be nil")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathProvider, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var provider Provider

	if err := resp.Unmarshal(&provider); err != nil {
		return nil, err
	}

	return &UpdateProviderOutput{
		Provider: &provider,
	}, nil
}

func (p providersService) SyncProvider(ctx context.Context, id int64, opts ...Option) (*SyncProviderOutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathProviderSync, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var msg SyncStatus

	if err := resp.Unmarshal(&msg); err != nil {
		return nil, err
	}

	return &SyncProviderOutput{
		Data: msg,
	}, nil
}

func (p providersService) SyncProviders(ctx context.Context, opts ...Option) (*SyncProviderOutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathProviderSync, "all"), nil, nil)
	if err != nil {
		return nil, err
	}

	var msg SyncStatus

	if err := resp.Unmarshal(&msg); err != nil {
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
	Configuration *ProviderConfiguration `json:",omitempty"`
}

type UpdateProviderInput = ProviderInput

type CreateProviderInput = ProviderInput

type ListProvidersInput struct{}

type ListProvidersOutput struct {
	Data []Provider
}

type Provider struct {
	Configuration *ProviderConfiguration `json:"Configuration,omitempty"`
	CreatedAt     string                 `json:"CreatedAt"`
	ID            int64                  `json:"ID"`
	LastSynced    string                 `json:"LastSynced"`
	Name          string                 `json:"Name"`
	Status        string                 `json:"Status"`
	Type          string                 `json:"Type"`
	UpdatedAt     string                 `json:"UpdatedAt"`
}

type ProviderConfiguration struct {
	ID       *int64 `json:"ID,omitempty"`
	MetaData string `json:"MetaData"`
}

type ProviderOutput struct {
	Provider *Provider
}

type SyncProviderOutput struct {
	Data SyncStatus
}

type UpdateProviderOutput = ProviderOutput

type GetProviderOutput = ProviderOutput

type DeleteProviderOutput = ProviderOutput

type CreateProviderOutput = ProviderOutput

type SyncStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
