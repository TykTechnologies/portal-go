package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathAccessRequests = "/portal-api/users"
	pathAccessRequest  = "/portal-api/users/%d"
)

type accessRequestsService struct {
	client *Client
}

func (p accessRequestsService) CreateAccessRequest(input CreateAccessRequestInput) (*CreateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathAccessRequests, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateAccessRequestOutput{}, nil
}

func (p accessRequestsService) GetAccessRequest(id uint64) (*GetAccessRequestOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathAccessRequest, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetAccessRequestOutput{}, nil
}

func (p accessRequestsService) ListAccessRequests(options *ListAccessRequestsOptions) (*ListAccessRequestsOutput, error) {
	req, err := p.client.newGetRequest(pathAccessRequests, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListAccessRequestsOutput{}, nil
}

func (p accessRequestsService) UpdateAccessRequest(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

func (p accessRequestsService) ApproveAccessRequest(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

func (p accessRequestsService) RejectAccessRequest(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

func (p accessRequestsService) ListClients(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

func (p accessRequestsService) CreateClient(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

func (p accessRequestsService) UpdateClient(id uint64, input UpdateAccessRequestInput) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{}, nil
}

type UpdateAccessRequestInput struct {
	Catalogues []uint64
}

type CreateAccessRequestInput struct{}

type ListAccessRequestsOptions struct{}

type ListAccessRequestsOutput struct{}

type AccessRequest struct{}

type AccessRequestOutput struct{}

type UpdateAccessRequestOutput = AccessRequestOutput

type GetAccessRequestOutput = AccessRequestOutput

type CreateAccessRequestOutput = AccessRequestOutput
