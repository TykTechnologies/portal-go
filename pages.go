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
	pathPages = "/portal-api/pages"
	pathPage  = "/portal-api/pages/%d"
)

//go:generate mockery --name Pages --filename pages.go
type Pages interface {
	CreatePage(ctx context.Context, input *CreatePageInput, opts ...Option) (*CreatePageOutput, error)
	GetPage(ctx context.Context, id int64, opts ...Option) (*GetPageOutput, error)
	ListPages(ctx context.Context, options *ListPagesInput, opts ...Option) (*ListPagesOutput, error)
	UpdatePage(ctx context.Context, id int64, input *UpdatePageInput, opts ...Option) (*UpdatePageOutput, error)
}

type pages struct {
	client *Client
}

func (p pages) CreatePage(ctx context.Context, input *CreatePageInput, opts ...Option) (*CreatePageOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathPages, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var page Page

	if err := resp.Unmarshal(&page); err != nil {
		return nil, err
	}

	return &CreatePageOutput{
		Data: &page,
	}, nil
}

func (p pages) GetPage(ctx context.Context, id int64, opts ...Option) (*GetPageOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathPage, id), nil)
	if err != nil {
		return nil, err
	}

	var page Page
	if err := resp.Unmarshal(&page); err != nil {
		return nil, err
	}

	return &GetPageOutput{
		Data: &page,
	}, nil
}

func (p pages) ListPages(ctx context.Context, options *ListPagesInput, opts ...Option) (*ListPagesOutput, error) {
	resp, err := p.client.doGet(ctx, pathPages, nil)
	if err != nil {
		return nil, err
	}

	var pages []Page

	if err := resp.Unmarshal(&pages); err != nil {
		return nil, err
	}

	return &ListPagesOutput{
		Pages: pages,
	}, nil
}

func (p pages) UpdatePage(ctx context.Context, id int64, input *UpdatePageInput, opts ...Option) (*UpdatePageOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathPage, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var page Page

	if err := resp.Unmarshal(&page); err != nil {
		return nil, err
	}

	return &UpdatePageOutput{
		Data: &page,
	}, nil
}

type PageInput struct {
	ID   *int64 `json:",omitempty"`
	Type string
	Name string
}

type UpdatePageInput = PageInput

type CreatePageInput = PageInput

type ListPagesInput struct{}

type ListPagesOutput struct {
	Pages []Page
}

type Page struct {
	AllowFormSubmission bool   `json:"AllowFormSubmission"`
	ID                  int    `json:"ID"`
	PageTypeID          int    `json:"PageTypeID"`
	Path                string `json:"Path"`
	Status              string `json:"Status"`
	Template            string `json:"Template"`
	Title               string `json:"Title"`
	CreatedAt           string `json:"CreatedAt"`
	UpdatedAt           string `json:"UpdatedAt"`
}

type PageOutput struct {
	Data *Page
}

type UpdatePageOutput = PageOutput

type GetPageOutput = PageOutput

type CreatePageOutput = PageOutput
