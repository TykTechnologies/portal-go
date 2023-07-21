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

//go:generate mockery --name ProductsService --filename products_service.go
type ProductsService interface {
	CreateProduct(ctx context.Context, input *CreateProductInput, opts ...Option) (*CreateProductOutput, error)
	GetProduct(ctx context.Context, id int64, opts ...Option) (*GetProductOutput, error)
	ListProducts(ctx context.Context, options *ListProductsInput, opts ...Option) (*ListProductsOutput, error)
	UpdateProduct(ctx context.Context, id int64, input *UpdateProductInput, opts ...Option) (*UpdateProductOutput, error)
}

type productsService struct {
	client *Client
}

func (p productsService) CreateProduct(ctx context.Context, input *CreateProductInput, opts ...Option) (*CreateProductOutput, error) {
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

func (p productsService) GetProduct(ctx context.Context, id int64, opts ...Option) (*GetProductOutput, error) {
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

func (p productsService) ListProducts(ctx context.Context, options *ListProductsInput, opts ...Option) (*ListProductsOutput, error) {
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

func (p productsService) UpdateProduct(ctx context.Context, id int64, input *UpdateProductInput, opts ...Option) (*UpdateProductOutput, error) {
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
	ID   *int64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateProductInput = ProductInput

type CreateProductInput = ProductInput

type ListProductsInput struct{}

type ListProductsOutput struct {
	Data []Product
}

type Product struct {
	APIDetails  []APIDetails `json:"APIDetails"`
	AuthType    string       `json:"AuthType"`
	Catalogues  any          `json:"Catalogues"`
	Content     string       `json:"Content"`
	DCREnabled  bool         `json:"DCREnabled"`
	Description string       `json:"Description"`
	DisplayName string       `json:"DisplayName"`
	Feature     bool         `json:"Feature"`
	ID          int          `json:"ID"`
	Logo        string       `json:"Logo"`
	Name        string       `json:"Name"`
	Path        string       `json:"Path"`
	ReferenceID string       `json:"ReferenceID"`
	Scopes      string       `json:"Scopes"`
	Tags        any          `json:"Tags"`
	Templates   any          `json:"Templates"`
}

type APIDetails struct {
	APIID       string `json:"APIID"`
	APIType     string `json:"APIType"`
	Description string `json:"Description"`
	ListenPath  string `json:"ListenPath"`
	Name        string `json:"Name"`
	OASDocument string `json:"OASDocument"`
	OASURL      string `json:"OASUrl"`
	Status      bool   `json:"Status"`
	TargetURL   string `json:"TargetURL"`
}

type ProductOutput struct {
	Data *Product
}

type UpdateProductOutput = ProductOutput

type GetProductOutput = ProductOutput

type CreateProductOutput = ProductOutput
