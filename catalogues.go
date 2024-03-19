package portal

import (
	"context"
	"fmt"
	"net/http"
)

type CataloguesService service

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

type catalogueInput struct {
	ID               *int64          `json:"ID,omitempty"`
	Name             string          `json:"Name,omitempty"`
	NameWithSlug     string          `json:"NameWithSlug,omitempty"`
	Plans            []int64         `json:"Plans,omitempty"`
	Products         []int64         `json:"Products,omitempty"`
	VisibilityStatus string          `json:"VisibilityStatus,omitempty"`
	OrgCatalogues    []OrgCatalogues `json:"OrgCatalogs,omitempty"`
}

type OrgCatalogues struct {
	OrganizationID int `json:"OrganizationID,omitempty"`
	TeamID         int `json:"TeamID,omitempty"`
}

type Audience struct {
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

type audienceInput struct{}

func (u *CataloguesService) ListCatalogues(ctx context.Context, opts *ListOptions) ([]*Catalogue, *Response, error) {
	urlPath := "/catalogues"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var catalogues []*Catalogue

	resp, err := u.client.Do(ctx, req, &catalogues)
	if err != nil {
		return nil, resp, err
	}

	return catalogues, resp, nil
}

func (u *CataloguesService) CreateCatalogue(ctx context.Context, input *Catalogue) (*Catalogue, *Response, error) {
	urlPath := "/catalogues"

	catalogueReq := &catalogueInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, catalogueReq)
	if err != nil {
		return nil, nil, err
	}

	catalogue := new(Catalogue)

	resp, err := u.client.Do(ctx, req, catalogue)
	if err != nil {
		return nil, resp, err
	}

	return catalogue, resp, nil
}

func (u *CataloguesService) GetCatalogue(ctx context.Context, catID int64) (*Catalogue, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v", catID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	catalogue := new(Catalogue)

	resp, err := u.client.Do(ctx, req, catalogue)
	if err != nil {
		return nil, resp, err
	}

	return catalogue, resp, nil
}

func (u *CataloguesService) UpdateCatalogue(ctx context.Context, catID int64, input *Catalogue) (*Catalogue, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v", catID)

	catalogueReq := &catalogueInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, catalogueReq)
	if err != nil {
		return nil, nil, err
	}

	catalogue := new(Catalogue)

	resp, err := u.client.Do(ctx, req, catalogue)
	if err != nil {
		return nil, resp, err
	}

	return catalogue, resp, nil
}

func (u *CataloguesService) DeleteCatalogue(ctx context.Context, catID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v", catID)

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

func (u *CataloguesService) ListAudiences(ctx context.Context, catID int64, opts *ListOptions) ([]*Audience, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v/audiences", catID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var audiences []*Audience

	resp, err := u.client.Do(ctx, req, &audiences)
	if err != nil {
		return nil, resp, err
	}

	return audiences, resp, nil
}

func (u *CataloguesService) CreateAudience(ctx context.Context, catalogueID int64, input *Audience) (*Audience, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v/audiences", catalogueID)

	catalogueReq := &audienceInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, catalogueReq)
	if err != nil {
		return nil, nil, err
	}

	audience := new(Audience)

	resp, err := u.client.Do(ctx, req, audience)
	if err != nil {
		return nil, resp, err
	}

	return audience, resp, nil
}

func (u *CataloguesService) GetAudience(ctx context.Context, catalogueID, audienceID int64) (*Audience, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v/audiences/%v", catalogueID, audienceID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	audience := new(Audience)

	resp, err := u.client.Do(ctx, req, audience)
	if err != nil {
		return nil, resp, err
	}

	return audience, resp, nil
}

func (u *CataloguesService) UpdateAudience(ctx context.Context, catalogueID, audienceID int64, input *Audience) (*Audience, *Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v/audiences/%v", catalogueID, audienceID)

	catalogueReq := &audienceInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, catalogueReq)
	if err != nil {
		return nil, nil, err
	}

	audience := new(Audience)

	resp, err := u.client.Do(ctx, req, audience)
	if err != nil {
		return nil, resp, err
	}

	return audience, resp, nil
}

func (u *CataloguesService) DeleteAudience(ctx context.Context, catalogueID, audienceID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/catalogues/%v/audiences/%v", catalogueID, audienceID)

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
