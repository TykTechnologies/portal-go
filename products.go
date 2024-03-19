// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type ProductsService service

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
	ReferenceID int64        `json:"ReferenceID,omitempty"`
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

type prodInput struct {
	Content     string `json:"Content,omitempty"`
	Description string `json:"Description,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	Feature     *bool  `json:"Feature,omitempty"`
	DCREnabled  *bool  `json:"DCREnabled,omitempty"`
}

type Tag struct{}

type ClientType struct{}

type APIDetail struct{}

type Tutorial struct{}

type clientTypeInput struct{}

type apiDetailsInput struct{}

type tagInput struct{}

type tutorialInput struct{}

func (u *ProductsService) ListProducts(ctx context.Context, opts *ListOptions) ([]*Product, *Response, error) {
	urlPath := "/products"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var products []*Product

	resp, err := u.client.Do(ctx, req, &products)
	if err != nil {
		return nil, resp, err
	}

	return products, resp, nil
}

func (u *ProductsService) GetProduct(ctx context.Context, id string) (*Product, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v", id)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	product := new(Product)

	resp, err := u.client.Do(ctx, req, product)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

func (u *ProductsService) UpdateProduct(ctx context.Context, prodID int64, input *Product) (*Product, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v", prodID)

	prodReq := prodInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, prodReq)
	if err != nil {
		return nil, nil, err
	}

	product := new(Product)

	resp, err := u.client.Do(ctx, req, product)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

func (u *ProductsService) ListTags(ctx context.Context, prodID int64, opts *ListOptions) ([]*Tag, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/tags", prodID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*Tag

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *ProductsService) CreateTag(ctx context.Context, prodID int64, input *Tag) (*Tag, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/tags", prodID)

	userReq := &tagInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(Tag)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) GetTag(ctx context.Context, prodID, tagID int64) (*Tag, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/tags/%v", prodID, tagID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(Tag)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) DeleteTag(ctx context.Context, prodID, tagID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/products/%v/tags/%v", prodID, tagID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *ProductsService) ListClientTypes(ctx context.Context, prodID int64, opts *ListOptions) ([]*ClientType, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/client_types", prodID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*ClientType

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *ProductsService) AttachClientType(ctx context.Context, prodID int64, input *ClientType) (*ClientType, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/client_types", prodID)

	userReq := &clientTypeInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(ClientType)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) GetClientType(ctx context.Context, prodID, ctID int64) (*ClientType, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/client_types/%v", prodID, ctID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(ClientType)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) UpdateClientType(ctx context.Context, prodID, ctID int64, input *ClientType) (*ClientType, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/client_types/%v", prodID, ctID)

	userReq := &clientTypeInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(ClientType)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) DetachClientType(ctx context.Context, prodID, ctID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/products/%v/client_types/%v", prodID, ctID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *ProductsService) ListAPIDetails(ctx context.Context, prodID int64, opts *ListOptions) ([]*APIDetail, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/api-details", prodID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*APIDetail

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *ProductsService) GetAPIDetails(ctx context.Context, prodID, apiID int64) (*APIDetail, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/api-details/%v", prodID, apiID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(APIDetail)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) UpdateAPIDetails(ctx context.Context, prodID, apiID int64, input *APIDetail) (*APIDetail, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/api-details/%v", prodID, apiID)

	userReq := &apiDetailsInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(APIDetail)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) DeleteAPIDetails(ctx context.Context, prodID, apiID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/products/%v/api-details/%v", prodID, apiID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *ProductsService) ListTutorials(ctx context.Context, prodID int64, opts *ListOptions) ([]*Tutorial, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/docs", prodID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*Tutorial

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *ProductsService) CreateTutorial(ctx context.Context, prodID int64, input *Tutorial) (*Tutorial, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/docs", prodID)

	userReq := &tutorialInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(Tutorial)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) GetTutorial(ctx context.Context, prodID, tutorialID int64) (*Tutorial, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/docs/%v", prodID, tutorialID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(Tutorial)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) UpdateTutorial(ctx context.Context, prodID, tutorialID int64, input *Tutorial) (*Tutorial, *Response, error) {
	urlPath := fmt.Sprintf("/products/%v/docs/%v", prodID, tutorialID)

	userReq := &tutorialInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(Tutorial)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *ProductsService) DeleteTutorial(ctx context.Context, prodID, tutorialID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/products/%v/docs/%v", prodID, tutorialID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
