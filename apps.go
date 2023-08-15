// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	pathAppAR        = "/portal-api/apps/%v/access-requests/%v"
	pathAppARs       = "/portal-api/apps/%v/access-requests"
	pathApps         = "/portal-api/apps"
	pathApp          = "/portal-api/apps/%v"
	pathAppProvision = "/portal-api/apps/%v/provision"
)

//go:generate mockery --name Apps --filename apps.go
type Apps interface {
	CreateApp(ctx context.Context, input *AppInput, opts ...Option) (*AppOutput, error)
	GetApp(ctx context.Context, id int64, opts ...Option) (*AppOutput, error)
	ListApps(ctx context.Context, opts ...Option) (*ListAppsOutput, error)
	ListARs(ctx context.Context, id int64, opts ...Option) (*ListARsOutput, error)
	ProvisionApp(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error)
	GetAR(ctx context.Context, appID int64, arID int64, opts ...Option) (*AROutput, error)
}

type apps struct {
	client *Client
}

func (p apps) CreateApp(ctx context.Context, input *AppInput, opts ...Option) (*AppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathApps, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var app App

	if err := resp.Unmarshal(&app); err != nil {
		return nil, err
	}

	return &AppOutput{
		Data: &app,
	}, nil
}

func (p apps) GetApp(ctx context.Context, id int64, opts ...Option) (*AppOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathApp, id), nil)
	if err != nil {
		return nil, err
	}

	var app App
	if err := resp.Unmarshal(&app); err != nil {
		return nil, err
	}

	return &AppOutput{
		Data: &app,
	}, nil
}

func (p apps) ProvisionApp(ctx context.Context, id int64, opts ...Option) (*StatusOutput, error) {
	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAppProvision, id), nil, nil)
	if err != nil {
		return nil, err
	}

	var status Status
	if err := resp.Unmarshal(&status); err != nil {
		return nil, err
	}

	return &StatusOutput{
		Data: &status,
	}, nil
}

// ListApps lists apps
func (p apps) ListApps(ctx context.Context, opts ...Option) (*ListAppsOutput, error) {
	resp, err := p.client.doGet(ctx, pathApps, nil, opts...)
	if err != nil {
		return nil, err
	}

	var ars []App

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListAppsOutput{
		Data: ars,
	}, nil
}

func (p apps) ListARs(ctx context.Context, id int64, opts ...Option) (*ListARsOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathApp, id), nil, opts...)
	if err != nil {
		return nil, err
	}

	var ars []ARSummary

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListARsOutput{
		Data: ars,
	}, nil
}

func (p apps) GetAR(ctx context.Context, appID int64, arID int64, opts ...Option) (*AROutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathAppAR, appID, arID), nil, opts...)
	if err != nil {
		return nil, err
	}

	var ar ARDetails

	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &AROutput{
		Data: &ar,
	}, nil
}

type ListAppsOutput struct {
	Response *http.Response
	Data     []App
}

type AppOutput struct {
	Response *http.Response
	Data     *App
}

type App struct {
	ID            int64       `json:"ID,omitempty"`
	Name          string      `json:"Name,omitempty"`
	Description   string      `json:"Description,omitempty"`
	RedirectURLs  string      `json:"RedirectURLs,omitempty"`
	UserID        int64       `json:"UserID,omitempty"`
	AccessRequest []ARDetails `json:"AccessRequests,omitempty"`
	CreatedAt     string      `json:"CreatedAt,omitempty"`
}

type AppInput struct {
	Name         string `json:"Name,omitempty"`
	Description  string `json:"Description,omitempty"`
	RedirectURLs string `json:"RedirectURLs,omitempty"`
	UserID       int64  `json:"UserID,omitempty"`
}
