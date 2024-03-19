// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type PagesService service

type Page struct {
	AllowFormSubmission bool   `json:"AllowFormSubmission"`
	ID                  int64  `json:"ID"`
	PageTypeID          int64  `json:"PageTypeID"`
	Path                string `json:"Path"`
	Status              string `json:"Status"`
	Template            string `json:"Template"`
	Title               string `json:"Title"`
	CreatedAt           string `json:"CreatedAt"`
	UpdatedAt           string `json:"UpdatedAt"`
}

type ContentBlock struct {
	Content string `json:"Content,omitempty"`
	Name    string `json:"Name,omitempty"`
	ID      int    `json:"ID,omitempty"`
	PageID  int    `json:"PageID,omitempty"`
}

type pageInput struct {
	AllowFormSubmission bool   `json:"AllowFormSubmission"`
	PageTypeID          int64  `json:"PageTypeID"`
	Path                string `json:"Path"`
	Status              string `json:"Status"`
	Template            string `json:"Template"`
	Title               string `json:"Title"`
}

func (u *PagesService) ListPages(ctx context.Context, opts *ListOptions) ([]*Page, *Response, error) {
	urlPath := "/pages"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var pages []*Page

	resp, err := u.client.Do(ctx, req, &pages)
	if err != nil {
		return nil, resp, err
	}

	return pages, resp, nil
}

func (u *PagesService) CreatePage(ctx context.Context, input *Page) (*Page, *Response, error) {
	urlPath := "/pages"

	pageReq := &pageInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, pageReq)
	if err != nil {
		return nil, nil, err
	}

	page := new(Page)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) GetPage(ctx context.Context, pageID int64) (*Page, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v", pageID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(Page)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) UpdatePage(ctx context.Context, pageID int64, input *Page) (*Page, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v", pageID)

	pageReq := &pageInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, pageReq)
	if err != nil {
		return nil, nil, err
	}

	page := new(Page)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) DeletePage(ctx context.Context, pageID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/pages/%v", pageID)

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

func (u *PagesService) ListContentBlocks(ctx context.Context, pageID int64, opts *ListOptions) ([]*ContentBlock, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v/content-blocks", pageID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var pages []*ContentBlock

	resp, err := u.client.Do(ctx, req, &pages)
	if err != nil {
		return nil, resp, err
	}

	return pages, resp, nil
}

func (u *PagesService) CreateContentBlocks(ctx context.Context, pageID int64, input *Page) (*ContentBlock, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v/content-blocks", pageID)

	pageReq := &pageInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, pageReq)
	if err != nil {
		return nil, nil, err
	}

	page := new(ContentBlock)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) GetContentBlocks(ctx context.Context, pageID, cbID int64) (*ContentBlock, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v/content-blocks/%v", pageID, cbID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(ContentBlock)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) UpdateContentBlocks(ctx context.Context, pageID, cbID int64, input *Page) (*ContentBlock, *Response, error) {
	urlPath := fmt.Sprintf("/pages/%v/content-blocks/%v", pageID, cbID)

	pageReq := &pageInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, pageReq)
	if err != nil {
		return nil, nil, err
	}

	page := new(ContentBlock)

	resp, err := u.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

func (u *PagesService) DeleteContentBlocks(ctx context.Context, pageID, cbID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/pages/%v/content-blocks/%v", pageID, cbID)

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
