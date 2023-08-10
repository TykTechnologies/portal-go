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
	pathProducts = "/portal-api/products"
	pathProduct  = "/portal-api/products/%d"
)

//go:generate mockery --name Products --filename products.go
type Products interface {
	CreateProduct(ctx context.Context, input *CreateProductInput, opts ...Option) (*CreateProductOutput, error)
	GetProduct(ctx context.Context, id int64, opts ...Option) (*GetProductOutput, error)
	ListProducts(ctx context.Context, options *ListProductsInput, opts ...Option) (*ListProductsOutput, error)
	UpdateProduct(ctx context.Context, id int64, input *UpdateProductInput, opts ...Option) (*UpdateProductOutput, error)
}

type products struct {
	client *Client
}

func (p products) CreateProduct(ctx context.Context, input *CreateProductInput, opts ...Option) (*CreateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathProducts, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var product Product

	if err := resp.Unmarshal(&product); err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		Data: &product,
	}, nil
}

func (p products) GetProduct(ctx context.Context, id int64, opts ...Option) (*GetProductOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathProduct, id), nil)
	if err != nil {
		return nil, err
	}

	var product Product
	if err := resp.Unmarshal(&product); err != nil {
		return nil, err
	}

	return &GetProductOutput{
		Data: &product,
	}, nil
}

func (p products) ListProducts(ctx context.Context, options *ListProductsInput, opts ...Option) (*ListProductsOutput, error) {
	resp, err := p.client.doGet(ctx, pathProducts, nil)
	if err != nil {
		return nil, err
	}

	var products []Product

	if err := resp.Unmarshal(&products); err != nil {
		return nil, err
	}

	return &ListProductsOutput{
		Data: products,
	}, nil
}

func (p products) UpdateProduct(ctx context.Context, id int64, input *UpdateProductInput, opts ...Option) (*UpdateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathProduct, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var product Product

	if err := resp.Unmarshal(&product); err != nil {
		return nil, err
	}

	return &UpdateProductOutput{
		Data: &product,
	}, nil
}

type ProductInput struct {
	Content     string `json:"Content,omitempty"`
	Description string `json:"Description,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	Feature     *bool  `json:"Feature,omitempty"`
	DCREnabled  *bool  `json:"DCREnabled,omitempty"`
}

type UpdateProductInput = ProductInput

type CreateProductInput = ProductInput

type ListProductsInput struct{}

type ListProductsOutput struct {
	Data []Product
}

type Product struct {
	ID          int          `json:"ID,omitempty"`
	APIDetails  []APIDetails `json:"APIDetails,omitempty"`
	AuthType    string       `json:"AuthType,omitempty"`
	Catalogues  []string     `json:"Catalogues,omitempty"`
	Content     string       `json:"Content,omitempty"`
	DCREnabled  bool         `json:"DCREnabled,omitempty"`
	Description string       `json:"Description,omitempty"`
	DisplayName string       `json:"DisplayName,omitempty"`
	Feature     bool         `json:"Feature,omitempty"`
	Name        string       `json:"Name,omitempty"`
	Path        string       `json:"Path,omitempty"`
	Logo        string       `json:"Logo,omitempty"`
	ReferenceID string       `json:"ReferenceID,omitempty"`
	Scopes      string       `json:"Scopes,omitempty"`
	Tags        []string     `json:"Tags,omitempty"`
	Templates   []string     `json:"Templates,omitempty"`
}

type APIDetails struct {
	APIID       string `json:"APIID,omitempty"`
	APIType     string `json:"APIType,omitempty"`
	Description string `json:"Description,omitempty"`
	ListenPath  string `json:"ListenPath,omitempty"`
	Name        string `json:"Name,omitempty"`
	OASURL      string `json:"OASUrl,omitempty"`
	Status      bool   `json:"Status,omitempty"`
	TargetURL   string `json:"TargetURL,omitempty"`
}

type ProductOutput struct {
	Data *Product
}

type UpdateProductOutput = ProductOutput

type GetProductOutput = ProductOutput

type CreateProductOutput = ProductOutput
