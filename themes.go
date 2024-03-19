// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
	"net/http"
)

type ThemesService service

type Theme struct {
	Author  string `json:"Author,omitempty"`
	ID      string `json:"ID,omitempty"`
	Name    string `json:"Name,omitempty"`
	Path    string `json:"Path,omitempty"`
	Status  string `json:"Status,omitempty"`
	Version string `json:"Version,omitempty"`
}

type themeInput struct {
	Active        bool   `json:"Active,omitempty"`
	Email         string `json:"Email,omitempty"`
	First         string `json:"First,omitempty"`
	Last          string `json:"Last,omitempty"`
	OrgID         int64  `json:"OrganisationID,omitempty"`
	Role          string `json:"Role,omitempty"`
	Provider      string `json:"Provider,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}

func (u *ThemesService) ListThemes(ctx context.Context, opts *ListOptions) ([]*Theme, *Response, error) {
	urlPath := "/themes"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var themes []*Theme

	resp, err := u.client.Do(ctx, req, &themes)
	if err != nil {
		return nil, resp, err
	}

	return themes, resp, nil
}

func (u *ThemesService) CreateTheme(ctx context.Context, input *Theme) (*Theme, *Response, error) {
	urlPath := "/themes"

	themeReq := &themeInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, themeReq)
	if err != nil {
		return nil, nil, err
	}

	theme := new(Theme)

	resp, err := u.client.Do(ctx, req, theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

func (u *ThemesService) GetTheme(ctx context.Context, themeID int64) (*Theme, *Response, error) {
	urlPath := fmt.Sprintf("/themes/%v", themeID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	theme := new(Theme)

	resp, err := u.client.Do(ctx, req, theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

func (u *ThemesService) DownloadTheme(ctx context.Context, themeID int64) (*Theme, *Response, error) {
	urlPath := fmt.Sprintf("/themes/%v", themeID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	theme := new(Theme)

	resp, err := u.client.Do(ctx, req, theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

func (u *ThemesService) ActivateTheme(ctx context.Context, themeID int64) (*Theme, *Response, error) {
	urlPath := fmt.Sprintf("/themes/%v/activate", themeID)

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	theme := new(Theme)

	resp, err := u.client.Do(ctx, req, theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

func (u *ThemesService) UploadTheme(ctx context.Context, themeID int64, input *Theme) (*Theme, *Response, error) {
	urlPath := fmt.Sprintf("/themes/upload", themeID)

	themeReq := &themeInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, themeReq)
	if err != nil {
		return nil, nil, err
	}

	theme := new(Theme)

	resp, err := u.client.Do(ctx, req, theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

func (u *ThemesService) DeleteTheme(ctx context.Context, themeID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/themes/%v", themeID)

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
