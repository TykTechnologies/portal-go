package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathCatalogues = "/portal-api/catalogues"
	pathCatalogue  = "/portal-api/catalogues/%d"
)

type cataloguesService struct {
	client *Client
}

func (c cataloguesService) CreateCatalogue(input CreateCatalogueInput) (*CreateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := c.client.newPostRequest(pathCatalogues, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateCatalogueOutput{}, nil
}

func (c cataloguesService) GetCatalogue(id uint64) (*GetCatalogueOutput, error) {
	req, err := c.client.newGetRequest(fmt.Sprintf(pathCatalogue, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{}, nil
}

func (c cataloguesService) UpdateCatalogue(id uint64, params UpdateCatalogueInput) (*UpdateCatalogueOutput, error) {
	req, err := c.client.newGetRequest(fmt.Sprintf(pathCatalogue, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{}, nil
}

func (c cataloguesService) DeleteCatalogue(id uint64) (*DeleteCatalogueOutput, error) {
	req, err := c.client.newDeleteRequest(fmt.Sprintf(pathCatalogue, id), nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &DeleteCatalogueOutput{}, nil
}

func (c cataloguesService) ListCatalogues(opts *ListCataloguesOptions) (*ListCataloguesOutput, error) {
	req, err := c.client.newGetRequest(pathCatalogues, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListCataloguesOutput{}, nil
}

type CreateCatalogueInput struct{}

type UpdateCatalogueInput struct{}

type CatalogueOutput struct {
	Cataogue *Catalogue
}

type CreateCatalogueOutput = CatalogueOutput

type GetCatalogueOutput = CatalogueOutput

type UpdateCatalogueOutput = CatalogueOutput

type DeleteCatalogueOutput = CatalogueOutput

type ListCataloguesOutput struct {
	Catalogues []Catalogue
}

type ListCataloguesOptions struct{}

type Catalogue struct{}
