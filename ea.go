// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type EAsService service

type EA struct {
	Active            bool     `json:"Active,omitempty"`
	Email             string   `json:"Email,omitempty"`
	First             string   `json:"First,omitempty"`
	Last              string   `json:"Last,omitempty"`
	OrgID             int64    `json:"OrganisationID,omitempty"`
	Role              string   `json:"Role,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	JWTToken          string   `json:"JWTToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Org               string   `json:"Organisation,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	CreatedAt         string   `json:"CreatedAt,omitempty"`
	UpdatedAt         string   `json:"UpdatedAt,omitempty"`
}

type DA struct {
	Active            bool     `json:"Active,omitempty"`
	Email             string   `json:"Email,omitempty"`
	First             string   `json:"First,omitempty"`
	Last              string   `json:"Last,omitempty"`
	OrgID             int64    `json:"OrganisationID,omitempty"`
	Role              string   `json:"Role,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	JWTToken          string   `json:"JWTToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Org               string   `json:"Organisation,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	CreatedAt         string   `json:"CreatedAt,omitempty"`
	UpdatedAt         string   `json:"UpdatedAt,omitempty"`
}

type CA struct {
	Active            bool     `json:"Active,omitempty"`
	Email             string   `json:"Email,omitempty"`
	First             string   `json:"First,omitempty"`
	Last              string   `json:"Last,omitempty"`
	OrgID             int64    `json:"OrganisationID,omitempty"`
	Role              string   `json:"Role,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	JWTToken          string   `json:"JWTToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Org               string   `json:"Organisation,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	CreatedAt         string   `json:"CreatedAt,omitempty"`
	UpdatedAt         string   `json:"UpdatedAt,omitempty"`
}

type caInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Role          string `json:"Role,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}

type eaInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Role          string `json:"Role,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}

type daInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Role          string `json:"Role,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}

func (u *EAsService) ListEAs(ctx context.Context, opts *ListOptions) ([]*EA, *Response, error) {
	urlPath := "/extended_attributes"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var eas []*EA

	resp, err := u.client.Do(ctx, req, &eas)
	if err != nil {
		return nil, resp, err
	}

	return eas, resp, nil
}

func (u *EAsService) CreateEA(ctx context.Context, input *EA) (*EA, *Response, error) {
	urlPath := "/extended_attributes"

	eaReq := &eaInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(EA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) GetEA(ctx context.Context, id string) (*EA, *Response, error) {
	urlPath := fmt.Sprintf("extended_attributes/%v", id)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ea := new(EA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) UpdateEA(ctx context.Context, id string, input *EA) (*EA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v", id)

	eaReq := &eaInput{
		First: input.First,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(EA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) DeleteEA(ctx context.Context, extID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v", extID)

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

func (u *EAsService) ListCAs(ctx context.Context, extID int64, opts *ListOptions) ([]*CA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/custom-attributes", extID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var eas []*CA

	resp, err := u.client.Do(ctx, req, &eas)
	if err != nil {
		return nil, resp, err
	}

	return eas, resp, nil
}

func (u *EAsService) CreateCA(ctx context.Context, extID int64, input *CA) (*CA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/custom-attributes", extID)

	eaReq := &caInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(CA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) GetCA(ctx context.Context, extID, caID int64) (*CA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/custom-attributes/%v", extID, caID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ea := new(CA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) UpdateCA(ctx context.Context, extID, caID int64, input *CA) (*CA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/custom-attributes/%v", extID, caID)

	eaReq := &caInput{
		First: input.First,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(CA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) DeleteCA(ctx context.Context, extID, caID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/custom-attributes/%v", extID, caID)

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

func (u *EAsService) ListDAs(ctx context.Context, extID int64, opts *ListOptions) ([]*DA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/default-attributes", extID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var das []*DA

	resp, err := u.client.Do(ctx, req, &das)
	if err != nil {
		return nil, resp, err
	}

	return das, resp, nil
}

func (u *EAsService) CreateDA(ctx context.Context, extID int64, input *DA) (*DA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/default-attributes", extID)

	eaReq := &caInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(DA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) GetDA(ctx context.Context, extID, daID int64) (*DA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/default-attributes/%v", extID, daID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ea := new(DA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) UpdateDA(ctx context.Context, extID, daID int64, input *DA) (*DA, *Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/default-attributes/%v", extID, daID)

	eaReq := &caInput{
		First: input.First,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, eaReq)
	if err != nil {
		return nil, nil, err
	}

	ea := new(DA)

	resp, err := u.client.Do(ctx, req, ea)
	if err != nil {
		return nil, resp, err
	}

	return ea, resp, nil
}

func (u *EAsService) DeleteDA(ctx context.Context, extID, daID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/extended_attributes/%v/default-attributes/%v", extID, daID)

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
