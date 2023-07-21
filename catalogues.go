package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	pathCatalogues = "/portal-api/catalogues"
	pathCatalogue  = "/portal-api/catalogues/%d"
)

//go:generate mockery --name CataloguesService --filename catalogues_service.go
type CataloguesService interface {
	CreateCatalogue(ctx context.Context, input *CreateCatalogueInput, opts ...Option) (*CreateCatalogueOutput, error)
	GetCatalogue(ctx context.Context, id int64, opts ...Option) (*GetCatalogueOutput, error)
	ListCatalogues(ctx context.Context, options *ListCataloguesInput, opts ...Option) (*ListCataloguesOutput, error)
	UpdateCatalogue(ctx context.Context, id int64, input *UpdateCatalogueInput, opts ...Option) (*UpdateCatalogueOutput, error)
}

type cataloguesService struct {
	client *Client
}

func (p cataloguesService) CreateCatalogue(ctx context.Context, input *CreateCatalogueInput, opts ...Option) (*CreateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathCatalogues, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue

	if err := resp.Unmarshal(&catalogue); err != nil {
		return nil, err
	}

	return &CreateCatalogueOutput{
		Data: &catalogue,
	}, nil
}

func (p cataloguesService) GetCatalogue(ctx context.Context, id int64, opts ...Option) (*GetCatalogueOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathCatalogue, id), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue
	if err := resp.Unmarshal(&catalogue); err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{
		Data: &catalogue,
	}, nil
}

func (p cataloguesService) ListCatalogues(ctx context.Context, options *ListCataloguesInput, opts ...Option) (*ListCataloguesOutput, error) {
	resp, err := p.client.doGet(ctx, pathCatalogues, nil)
	if err != nil {
		return nil, err
	}

	var catalogues []Catalogue

	if err := resp.Unmarshal(&catalogues); err != nil {
		return nil, err
	}

	return &ListCataloguesOutput{
		Data: catalogues,
	}, nil
}

func (p cataloguesService) UpdateCatalogue(ctx context.Context, id int64, input *UpdateCatalogueInput, opts ...Option) (*UpdateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathCatalogue, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue

	if err := resp.Unmarshal(&catalogue); err != nil {
		return nil, err
	}

	return &UpdateCatalogueOutput{
		Data: &catalogue,
	}, nil
}

type CatalogueInput struct {
	ID   *int64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateCatalogueInput = CatalogueInput

type CreateCatalogueInput = CatalogueInput

type ListCataloguesInput struct{}

type ListCataloguesOutput struct {
	Data []Catalogue
}

type Catalogue struct {
	ID               int64
	Name             string
	NameWithSlug     string
	VisibilityStatus string
}

type CatalogueOutput struct {
	Data *Catalogue
}

type UpdateCatalogueOutput = CatalogueOutput

type GetCatalogueOutput = CatalogueOutput

type CreateCatalogueOutput = CatalogueOutput
