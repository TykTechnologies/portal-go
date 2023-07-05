package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Catalogues struct {
	client *Client
}

func (c Catalogues) CreateCatalogue(input CreateCatalogueInput) (*CreateCatalogueOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := u.client.newPostRequest("/portal-api/products/%d", bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateCatalogueOutput{}, nil
}

func (c Catalogues) GetCatalogue(id uint64) (*GetCatalogueOutput, error) {
	req, err := u.client.newGetRequest(fmt.Sprintf("/portal-api/catalogues/%d", id), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{}, nil
}

func (c Catalogues) UpdateCatalogue(id uint64, params UpdateCatalogueInput) (*UpdateCatalogueOutput, error) {
	req, err := u.client.newGetRequest(fmt.Sprintf("/portal-api/catalogues/%d", id), nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetCatalogueOutput{}, nil
}

func (c Catalogues) DeleteCatalogue(id uint64) (*DeleteCatalogueOutput, error) {
	req, err := u.client.newDeleteRequest(fmt.Sprintf("/portal-api/catalogues/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &DeleteCatalogueOutput{}, nil
}

func (c Catalogues) ListCatalogues(opts *ListCataloguesOptions) (*ListCataloguesOutput, error) {
	req, err := u.client.newGetRequest("/portal-api/catalogues", nil)
	if err != nil {
		return nil, err
	}

	_, err = u.client.performRequest(req)
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
