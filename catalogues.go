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

//go:generate mockery --name CataloguesService
type CataloguesService interface {
	CreateCatalogue(ctx context.Context, input CreateCatalogueInput) (*CreateCatalogueOutput, error)
	GetCatalogue(ctx context.Context, id uint64) (*GetCatalogueOutput, error)
	ListCatalogues(ctx context.Context, options *ListCataloguesOptions) (*ListCataloguesOutput, error)
	UpdateCatalogue(ctx context.Context, id uint64, input UpdateCatalogueInput) (*UpdateCatalogueOutput, error)
}

type cataloguesService struct {
	client *Client
}

func (p cataloguesService) CreateCatalogue(ctx context.Context, input CreateCatalogueInput) (*CreateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathCatalogues, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue

	if err := resp.Parse(&catalogue); err != nil {
		return nil, err
	}

	return &CreateCatalogueOutput{
		Catalogue: &catalogue,
	}, nil
}

func (p cataloguesService) GetCatalogue(ctx context.Context, id uint64) (*GetCatalogueOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathCatalogue, id), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue
	if err := resp.Parse(&catalogue); err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{
		Catalogue: &catalogue,
	}, nil
}

func (p cataloguesService) ListCatalogues(ctx context.Context, options *ListCataloguesOptions) (*ListCataloguesOutput, error) {
	resp, err := p.client.doGet(pathCatalogues, nil)
	if err != nil {
		return nil, err
	}

	var catalogues []Catalogue

	if err := resp.Parse(&catalogues); err != nil {
		return nil, err
	}

	return &ListCataloguesOutput{
		Catalogues: catalogues,
	}, nil
}

func (p cataloguesService) UpdateCatalogue(ctx context.Context, id uint64, input UpdateCatalogueInput) (*UpdateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathCatalogue, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue

	if err := resp.Parse(&catalogue); err != nil {
		return nil, err
	}

	return &UpdateCatalogueOutput{
		Catalogue: &catalogue,
	}, nil
}

type CatalogueInput struct {
	ID   *uint64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateCatalogueInput = CatalogueInput

type CreateCatalogueInput = CatalogueInput

type ListCataloguesOptions struct{}

type ListCataloguesOutput struct {
	Catalogues []Catalogue
}

type Catalogue struct {
	ID               uint64
	Name             string
	NameWithSlug     string
	VisibilityStatus string
}

type CatalogueOutput struct {
	Catalogue *Catalogue
}

type UpdateCatalogueOutput = CatalogueOutput

type GetCatalogueOutput = CatalogueOutput

type CreateCatalogueOutput = CatalogueOutput
