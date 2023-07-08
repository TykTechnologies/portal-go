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

//go:generate mockery --name ProductsService
type ProductsService interface {
	CreateProduct(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error)
	GetProduct(ctx context.Context, id uint64) (*GetProductOutput, error)
	ListProducts(ctx context.Context, options *ListProductsOptions) (*ListProductsOutput, error)
	UpdateProduct(ctx context.Context, id uint64, input UpdateProductInput) (*UpdateProductOutput, error)
}

type productsService struct {
	client *Client
}

func (p productsService) CreateProduct(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathProducts, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var product Product

	if err := resp.Parse(&product); err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		Product: &product,
	}, nil
}

func (p productsService) GetProduct(ctx context.Context, id uint64) (*GetProductOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathProduct, id), nil)
	if err != nil {
		return nil, err
	}

	var product Product
	if err := resp.Parse(&product); err != nil {
		return nil, err
	}

	return &GetProductOutput{
		Product: &product,
	}, nil
}

func (p productsService) ListProducts(ctx context.Context, options *ListProductsOptions) (*ListProductsOutput, error) {
	resp, err := p.client.doGet(pathProducts, nil)
	if err != nil {
		return nil, err
	}

	var products []Product

	if err := resp.Parse(&products); err != nil {
		return nil, err
	}

	return &ListProductsOutput{
		Products: products,
	}, nil
}

func (p productsService) UpdateProduct(ctx context.Context, id uint64, input UpdateProductInput) (*UpdateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathProduct, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var product Product

	if err := resp.Parse(&product); err != nil {
		return nil, err
	}

	return &UpdateProductOutput{
		Product: &product,
	}, nil
}

type ProductInput struct {
	ID   *uint64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateProductInput = ProductInput

type CreateProductInput = ProductInput

type ListProductsOptions struct{}

type ListProductsOutput struct {
	Products []Product
}

type Product struct {
	ID          uint64
	Name        string
	DisplayName string
	ReferenceID string
	Feature     bool
	DCREnabled  bool
	AuthType    string
	Scopes      string
	Path        string
}

type ProductOutput struct {
	Product *Product
}

type UpdateProductOutput = ProductOutput

type GetProductOutput = ProductOutput

type CreateProductOutput = ProductOutput
