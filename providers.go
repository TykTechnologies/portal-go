package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathProviders = "/portal-api/providers"
	pathProvider  = "/portal-api/providers/%d"
)

type providersService struct {
	client *Client
}

func (p providersService) CreateProvider(input CreateProviderInput) (*CreateProviderOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathProviders, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateProviderOutput{}, nil
}

func (p providersService) GetProvider(id uint64) (*GetProviderOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathProvider, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetProviderOutput{}, nil
}

func (p providersService) ListProviders(options *ListProvidersOptions) (*ListProvidersOutput, error) {
	req, err := p.client.newGetRequest(pathProviders, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListProvidersOutput{}, nil
}

func (p providersService) UpdateProvider(id uint64, input UpdateProviderInput) (*UpdateProviderOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathProvider, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateProviderOutput{}, nil
}

type UpdateProviderInput struct {
	Catalogues []uint64
}

type CreateProviderInput struct{}

type ListProvidersOptions struct{}

type ListProvidersOutput struct{}

type Provider struct{}

type ProviderOutput struct{}

type UpdateProviderOutput = ProviderOutput

type GetProviderOutput = ProviderOutput

type CreateProviderOutput = ProviderOutput
