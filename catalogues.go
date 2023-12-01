// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

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

//go:generate mockery --name Catalogs --filename catalogs.go
type Catalogues interface {
	CreateCatalogue(ctx context.Context, input *CreateCatalogueInput, opts ...Option) (*CreateCatalogueOutput, error)
	GetCatalogue(ctx context.Context, id int64, opts ...Option) (*GetCatalogueOutput, error)
	ListCatalogues(ctx context.Context, options *ListCataloguesInput, opts ...Option) (*ListCataloguesOutput, error)
	UpdateCatalogue(ctx context.Context, id int64, input *UpdateCatalogueInput, opts ...Option) (*UpdateCatalogueOutput, error)
	DeleteCatalogue(ctx context.Context, id int64, opts ...Option) (*CatalogueOutput, error)
}

type catalogues struct {
	client *Client
}

func (p catalogues) CreateCatalogue(ctx context.Context, input *CreateCatalogueInput, opts ...Option) (*CreateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathCatalogues, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalogue

	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &CreateCatalogueOutput{
		Data: &catalog,
	}, nil
}

func (p catalogues) GetCatalogue(ctx context.Context, id int64, opts ...Option) (*GetCatalogueOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathCatalogue, id), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalogue
	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{
		Data: &catalog,
	}, nil
}

func (p catalogues) ListCatalogues(ctx context.Context, options *ListCataloguesInput, opts ...Option) (*ListCataloguesOutput, error) {
	resp, err := p.client.doGet(ctx, pathCatalogues, nil)
	if err != nil {
		return nil, err
	}

	var catalogs []Catalogue

	if err := resp.Unmarshal(&catalogs); err != nil {
		return nil, err
	}

	return &ListCataloguesOutput{
		Data: catalogs,
	}, nil
}

func (p catalogues) UpdateCatalogue(ctx context.Context, id int64, input *UpdateCatalogueInput, opts ...Option) (*UpdateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathCatalogue, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalogue

	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &UpdateCatalogueOutput{
		Data: &catalog,
	}, nil
}

func (p catalogues) DeleteCatalogue(ctx context.Context, id int64, opts ...Option) (*CatalogueOutput, error) {
	_, err := p.client.doDelete(ctx, fmt.Sprintf(pathCatalogue, id), nil, nil)
	if err != nil {
		return nil, err
	}

	return &CatalogueOutput{}, nil
}

type CatalogueInput struct {
	ID               *int64  `json:"ID,omitempty"`
	Name             string  `json:"Name,omitempty"`
	NameWithSlug     string  `json:"NameWithSlug,omitempty"`
	Plans            []int64 `json:"Plans,omitempty"`
	Products         []int64 `json:"Products,omitempty"`
	VisibilityStatus string  `json:"VisibilityStatus,omitempty"`
	OrgCatalogs      []struct {
		OrganizationID int `json:"OrganizationID,omitempty"`
		TeamID         int `json:"TeamID,omitempty"`
	} `json:"OrgCatalogs,omitempty"`
}

type UpdateCatalogueInput = CatalogueInput

type CreateCatalogueInput = CatalogueInput

type ListCataloguesInput struct{}

type ListCataloguesOutput struct {
	Data []Catalogue
}

type Catalogue struct {
	CreatedAt        string   `json:"CreatedAt,omitempty"`
	ID               int64    `json:"ID,omitempty"`
	Name             string   `json:"Name,omitempty"`
	OrgCatalogs      []any    `json:"OrgCatalogs,omitempty"`
	Plans            []any    `json:"Plans,omitempty"`
	Products         []string `json:"Products,omitempty"`
	UpdatedAt        string   `json:"UpdatedAt,omitempty"`
	VisibilityStatus string   `json:"VisibilityStatus,omitempty"`
	NameWithSlug     string   `json:"NameWithSlug,omitempty"`
}

type CatalogueOutput struct {
	Data *Catalogue
}

type UpdateCatalogueOutput = CatalogueOutput

type GetCatalogueOutput = CatalogueOutput

type CreateCatalogueOutput = CatalogueOutput
