package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	pathProviders    = "/portal-api/providers"
	pathProvider     = "/portal-api/providers/%d"
	pathProviderSync = "/portal-api/providers/%d/synchronize"
)

type ProvidersService interface {
	CreateProvider(ctx context.Context, input CreateProviderInput) (*CreateProviderOutput, error)
	GetProvider(ctx context.Context, id uint64) (*GetProviderOutput, error)
	ListProviders(ctx context.Context, options *ListProvidersOptions) (*ListProvidersOutput, error)
	UpdateProvider(ctx context.Context, id uint64, input UpdateProviderInput) (*UpdateProviderOutput, error)
	SynchronizeProvider(ctx context.Context, id uint64) (*SynchronizeProviderOutput, error)
}

type providersService struct {
	client *Client
}

func (p providersService) CreateProvider(ctx context.Context, input CreateProviderInput) (*CreateProviderOutput, error) {
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

func (p providersService) GetProvider(ctx context.Context, id uint64) (*GetProviderOutput, error) {
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

func (p providersService) ListProviders(ctx context.Context, options *ListProvidersOptions) (*ListProvidersOutput, error) {
	resp, err := p.client.doGet(pathProviders, nil)
	if err != nil {
		return nil, err
	}

	var providers []Provider

	if err := resp.Parse(&providers); err != nil {
		return nil, err
	}

	return &ListProvidersOutput{
		Providers: providers,
	}, nil
}

func (p providersService) UpdateProvider(ctx context.Context, id uint64, input UpdateProviderInput) (*UpdateProviderOutput, error) {
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

func (p providersService) SynchronizeProvider(ctx context.Context, id uint64) (*SynchronizeProviderOutput, error) {
	resp, err := p.client.doPut(fmt.Sprintf(pathProviderSync, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var msg SynchronizationStatus

	if err := resp.Parse(&msg); err != nil {
		return nil, err
	}

	return &SynchronizeProviderOutput{
		Synchronization: msg,
	}, nil
}

type UpdateProviderInput struct {
	Catalogues []uint64
}

type CreateProviderInput struct {
	Type          string
	Name          string
	Configuration *ProviderConfiguration `json:"Configuration,omitempty"`
}

type ProviderConfiguration struct {
	ProviderID *uint64 `json:"ProviderID,omitempty"`
	MetaData   string
	ID         uint64
}

type ListProvidersOptions struct{}

type ListProvidersOutput struct {
	Providers []Provider
}

type Provider struct {
	ID            uint64
	Name          string
	CreatedAt     string
	UpdatedAt     string
	Type          string
	Status        string
	LastSynched   string `json:"LastSynced"`
	Configuration ProviderConfiguration
}

type ProviderOutput struct {
	Provider *Provider
}

type SynchronizeProviderOutput struct {
	Synchronization SynchronizationStatus
}

type UpdateProviderOutput = ProviderOutput

type GetProviderOutput = ProviderOutput

type CreateProviderOutput = ProviderOutput

type SynchronizationStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
