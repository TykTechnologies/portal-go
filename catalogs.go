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
	pathCatalogs = "/portal-api/catalogs"
	pathCatalog  = "/portal-api/catalogs/%d"
)

//go:generate mockery --name Catalogs --filename catalogs.go
type Catalogs interface {
	CreateCatalog(ctx context.Context, input *CreateCatalogInput, opts ...Option) (*CreateCatalogOutput, error)
	GetCatalog(ctx context.Context, id int64, opts ...Option) (*GetCatalogOutput, error)
	ListCatalogs(ctx context.Context, options *ListCatalogsInput, opts ...Option) (*ListCatalogsOutput, error)
	UpdateCatalog(ctx context.Context, id int64, input *UpdateCatalogInput, opts ...Option) (*UpdateCatalogOutput, error)
}

type catalogs struct {
	client *Client
}

func (p catalogs) CreateCatalog(ctx context.Context, input *CreateCatalogInput, opts ...Option) (*CreateCatalogOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathCatalogs, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalog

	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &CreateCatalogOutput{
		Data: &catalog,
	}, nil
}

func (p catalogs) GetCatalog(ctx context.Context, id int64, opts ...Option) (*GetCatalogOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathCatalog, id), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalog
	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &GetCatalogOutput{
		Data: &catalog,
	}, nil
}

func (p catalogs) ListCatalogs(ctx context.Context, options *ListCatalogsInput, opts ...Option) (*ListCatalogsOutput, error) {
	resp, err := p.client.doGet(ctx, pathCatalogs, nil)
	if err != nil {
		return nil, err
	}

	var catalogs []Catalog

	if err := resp.Unmarshal(&catalogs); err != nil {
		return nil, err
	}

	return &ListCatalogsOutput{
		Data: catalogs,
	}, nil
}

func (p catalogs) UpdateCatalog(ctx context.Context, id int64, input *UpdateCatalogInput, opts ...Option) (*UpdateCatalogOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathCatalog, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var catalog Catalog

	if err := resp.Unmarshal(&catalog); err != nil {
		return nil, err
	}

	return &UpdateCatalogOutput{
		Data: &catalog,
	}, nil
}

type CatalogInput struct {
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

type UpdateCatalogInput = CatalogInput

type CreateCatalogInput = CatalogInput

type ListCatalogsInput struct{}

type ListCatalogsOutput struct {
	Data []Catalog
}

type Catalog struct {
	CreatedAt        string   `json:"CreatedAt,omitempty"`
	ID               int64    `json:"ID,omitempty"`
	Name             string   `json:"Name,omitempty"`
	OrgCatalogs      []any    `json:"OrgCatalogs,omitempty"`
	Plans            any      `json:"Plans,omitempty"`
	Products         []string `json:"Products,omitempty"`
	UpdatedAt        string   `json:"UpdatedAt,omitempty"`
	VisibilityStatus string   `json:"VisibilityStatus,omitempty"`
	NameWithSlug     string   `json:"NameWithSlug,omitempty"`
}

type CatalogOutput struct {
	Data *Catalog
}

type UpdateCatalogOutput = CatalogOutput

type GetCatalogOutput = CatalogOutput

type CreateCatalogOutput = CatalogOutput
