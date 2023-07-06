package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathApps = "/portal-api/users"
	pathApp  = "/portal-api/apps/%d"
)

type appsService struct {
	client *Client
}

func (p appsService) CreateApp(input CreateAppInput) (*CreateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathApps, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateAppOutput{}, nil
}

func (p appsService) GetApp(id uint64) (*GetAppOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathApp, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetAppOutput{}, nil
}

func (p appsService) ListApps(options *ListAppsOptions) (*ListAppsOutput, error) {
	req, err := p.client.newGetRequest(pathApps, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListAppsOutput{}, nil
}

func (p appsService) UpdateApp(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) CreateAccessRequest(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) ListAccessRequest(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) GetAccessRequest(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) UpdateAccessRequest(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) DeleteAccessRequest(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) Provision(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

func (p appsService) PSProvision(id uint64, input UpdateAppInput) (*UpdateAppOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathApp, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAppOutput{}, nil
}

type UpdateAppInput struct {
	Catalogues []uint64
}

type CreateAppInput struct{}

type ListAppsOptions struct{}

type ListAppsOutput struct{}

type App struct{}

type AppOutput struct{}

type UpdateAppOutput = AppOutput

type GetAppOutput = AppOutput

type CreateAppOutput = AppOutput
