package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathProducts = "/portal-api/products"
	pathProduct  = "/portal-api/products/%d"
)

type productsService struct {
	client *Client
}

func (p productsService) CreateProduct(input CreateProductInput) (*CreateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathProducts, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateProductOutput{}, nil
}

func (p productsService) GetProduct(id uint64) (*GetProductOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathProduct, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetProductOutput{}, nil
}

func (p productsService) ListProducts(options *ListProductsOptions) (*ListProductsOutput, error) {
	req, err := p.client.newGetRequest(pathProducts, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListProductsOutput{}, nil
}

func (p productsService) UpdateProduct(id uint64, input UpdateProductInput) (*UpdateProductOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathProduct, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateProductOutput{}, nil
}

type UpdateProductInput struct {
	Catalogues []uint64
}

type CreateProductInput struct{}

type ListProductsOptions struct{}

type ListProductsOutput struct{}

type Product struct{}

type ProductOutput struct{}

type UpdateProductOutput = ProductOutput

type GetProductOutput = ProductOutput

type CreateProductOutput = ProductOutput
