package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Products struct {
	client *Client
}

func (p Products) CreateProduct(input CreateProductInput) (error, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := u.client.newPostRequest("/portal-api/products", bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p Products) GetProduct(id uint64) (error, error) {
	req, err := u.client.newGetRequest(fmt.Sprintf("/portal-api/products/%d", id), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p Products) ListProducts(options *ListProductsOptions) (error, error) {
	req, err := u.client.newGetRequest("/portal-api/products", nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p Products) UpdateProduct(id uint64, input UpdateProductInput) (error, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := u.client.newPutRequest(fmt.Sprintf("/portal-api/products/%d", id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type UpdateProductInput struct {
	Catalogues []uint64
}

type CreateProductInput struct{}

type ListProductsOptions struct{}
