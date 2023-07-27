package edp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	pathAccessRequests = "/portal-api/access_requests"
	pathAccessRequest  = "/portal-api/access_requests/%d"
)

//go:generate mockery --name PlansService --filename plans_service.go
type AccessRequestsService interface {
	CreateAccessRequest(ctx context.Context, input *CreateAccessRequestInput, opts ...Option) (*CreateAccessRequestOutput, error)
	GetAccessRequest(ctx context.Context, id int64, opts ...Option) (*GetAccessRequestOutput, error)
	ListAccessRequests(ctx context.Context, options *ListAccessRequestsInput, opts ...Option) (*ListAccessRequestsOutput, error)
	UpdateAccessRequest(ctx context.Context, id int64, input *UpdateAccessRequestInput, opts ...Option) (*UpdateAccessRequestOutput, error)
}

type accessRequestsService struct {
	client *Client
}

// CreatePlan ...
func (p accessRequestsService) CreateAccessRequest(ctx context.Context, input *CreateAccessRequestInput, opts ...Option) (*CreateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathAccessRequests, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var ar AccessRequest

	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &CreateAccessRequestOutput{
		Data: &ar,
	}, nil
}

// GetPlan ...
func (p accessRequestsService) GetAccessRequest(ctx context.Context, id int64, opts ...Option) (*GetAccessRequestOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathAccessRequest, id), nil)
	if err != nil {
		return nil, err
	}

	var ar AccessRequest
	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &GetAccessRequestOutput{
		Data: &ar,
	}, nil
}

// ListPlans ...
func (p accessRequestsService) ListAccessRequests(ctx context.Context, options *ListAccessRequestsInput, opts ...Option) (*ListAccessRequestsOutput, error) {
	resp, err := p.client.doGet(ctx, pathAccessRequests, nil)
	if err != nil {
		return nil, err
	}

	var ars []AccessRequest

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListAccessRequestsOutput{
		Data: ars,
	}, nil
}

// UpdatePlan ...
func (p accessRequestsService) UpdateAccessRequest(ctx context.Context, id int64, input *UpdatePlanInput, opts ...Option) (*UpdateAccessRequestOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathAccessRequest, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var ar AccessRequest

	if err := resp.Unmarshal(&ar); err != nil {
		return nil, err
	}

	return &UpdateAccessRequestOutput{
		Data: &ar,
	}, nil
}

type AccessRequestInput struct {
	ID            *int64 `json:",omitempty"`
	Type          string
	Name          string
	Configuration *PlanConfiguration `json:",omitempty"`
}

type UpdateAccessRequestInput = PlanInput

type CreateAccessRequestInput = PlanInput

type ListAccessRequestsInput struct{}

type ListAccessRequestsOutput struct {
	Data []AccessRequest
}

type AccessRequest struct {
	AuthType             string `json:"AuthType"`
	Catalogue            string `json:"Catalogue"`
	Client               string `json:"Client"`
	CreatedAt            string `json:"CreatedAt"`
	Credentials          []any  `json:"Credentials"`
	DCREnabled           bool   `json:"DCREnabled"`
	DeletedAt            string `json:"DeletedAt"`
	ID                   int64  `json:"ID"`
	Plan                 string `json:"Plan"`
	Products             string `json:"Products"`
	ProvisionImmediately bool   `json:"ProvisionImmediately"`
	Status               string `json:"Status"`
	UpdatedAt            string `json:"UpdatedAt"`
	User                 string `json:"User"`
}

type AccessRequestOutput struct {
	Data *AccessRequest
}

type UpdateAccessRequestOutput = AccessRequestOutput

type GetAccessRequestOutput = AccessRequestOutput

type CreateAccessRequestOutput = AccessRequestOutput
