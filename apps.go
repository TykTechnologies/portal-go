package portal

import (
	"context"
	"fmt"
	"net/http"
)

type AppsService service

type App struct {
	ID            int64  `json:"ID,omitempty"`
	Name          string `json:"Name,omitempty"`
	Description   string `json:"Description,omitempty"`
	RedirectURLs  string `json:"RedirectURLs,omitempty"`
	UserID        int64  `json:"UserID,omitempty"`
	AccessRequest []AR   `json:"AccessRequests,omitempty"`
	CreatedAt     string `json:"CreatedAt,omitempty"`
}

type appInput struct {
	Name         string `json:"Name,omitempty"`
	Description  string `json:"Description,omitempty"`
	RedirectURLs string `json:"RedirectURLs,omitempty"`
	UserID       int64  `json:"UserID,omitempty"`
}

func (u *AppsService) ListApps(ctx context.Context, opts *ListOptions) ([]*App, *Response, error) {
	ulrPath := "/apps"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, ulrPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var apps []*App

	resp, err := u.client.Do(ctx, req, &apps)
	if err != nil {
		return nil, resp, err
	}

	return apps, resp, nil
}

func (u *AppsService) CreateApp(ctx context.Context, input *App) (*App, *Response, error) {
	urlPath := "/apps"

	appReq := &appInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, appReq)
	if err != nil {
		return nil, nil, err
	}

	app := new(App)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

func (u *AppsService) GetApp(ctx context.Context, appID int64) (*App, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v", appID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	app := new(App)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

func (u *AppsService) UpdateApp(ctx context.Context, appID int64, input *App) (*App, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v", appID)

	appReq := &appInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, appReq)
	if err != nil {
		return nil, nil, err
	}

	app := new(App)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

func (u *AppsService) DeleteApp(ctx context.Context, appID string) (*Response, error) {
	urlPath := fmt.Sprintf("/apps/%v", appID)

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

func (u *AppsService) ProvisionApp(ctx context.Context, appID string, input *App) (*App, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v", appID)

	appReq := &appInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, appReq)
	if err != nil {
		return nil, nil, err
	}

	app := new(App)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

func (u *AppsService) GetAR(ctx context.Context, appID, arID string) (*AR, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests/%v", appID, arID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AR)

	resp, err := u.client.Do(ctx, req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, nil
}

func (u *AppsService) DeleteAR(ctx context.Context, appID, arID string) (*Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests/%v", appID, arID)

	appReq := &appInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, appReq)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *AppsService) ListARs(ctx context.Context, appID string, input *AR, opts *ListOptions) (*AR, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests", appID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	ars := new(AR)

	resp, err := u.client.Do(ctx, req, ars)
	if err != nil {
		return nil, resp, err
	}

	return ars, resp, nil
}

func (u *AppsService) ListARCredentials(ctx context.Context, appID, arID string, opts *ListOptions) ([]*Credentials, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests/%v/credentials", appID, arID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var creds []*Credentials

	resp, err := u.client.Do(ctx, req, creds)
	if err != nil {
		return nil, resp, err
	}

	return creds, resp, nil
}

func (u *AppsService) GetARCredential(ctx context.Context, appID, arID, credID string) (*Credentials, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests/%v/credentials/%v", appID, arID, credID)

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	app := new(Credentials)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}

func (u *AppsService) DeleteARCredential(ctx context.Context, appID, arID, credID string) (*Credentials, *Response, error) {
	urlPath := fmt.Sprintf("/apps/%v/access-requests/%v/credentials/%v", appID, arID, credID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	app := new(Credentials)

	resp, err := u.client.Do(ctx, req, app)
	if err != nil {
		return nil, resp, err
	}

	return app, resp, nil
}
