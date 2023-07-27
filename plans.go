package edp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	pathPlans = "/portal-api/plans"
	pathPlan  = "/portal-api/plans/%d"
)

//go:generate mockery --name PlansService --filename plans_service.go
type PlansService interface {
	CreatePlan(ctx context.Context, input *CreatePlanInput, opts ...Option) (*CreatePlanOutput, error)
	GetPlan(ctx context.Context, id int64, opts ...Option) (*GetPlanOutput, error)
	ListPlans(ctx context.Context, options *ListPlansInput, opts ...Option) (*ListPlansOutput, error)
	UpdatePlan(ctx context.Context, id int64, input *UpdatePlanInput, opts ...Option) (*UpdatePlanOutput, error)
}

type plansService struct {
	client *Client
}

// CreatePlan ...
func (p plansService) CreatePlan(ctx context.Context, input *CreatePlanInput, opts ...Option) (*CreatePlanOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathPlans, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan

	if err := resp.Unmarshal(&plan); err != nil {
		return nil, err
	}

	return &CreatePlanOutput{
		Data: &plan,
	}, nil
}

// GetPlan ...
func (p plansService) GetPlan(ctx context.Context, id int64, opts ...Option) (*GetPlanOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathPlan, id), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan
	if err := resp.Unmarshal(&plan); err != nil {
		return nil, err
	}

	return &GetPlanOutput{
		Data: &plan,
	}, nil
}

// ListPlans ...
func (p plansService) ListPlans(ctx context.Context, options *ListPlansInput, opts ...Option) (*ListPlansOutput, error) {
	resp, err := p.client.doGet(ctx, pathPlans, nil)
	if err != nil {
		return nil, err
	}

	var plans []Plan

	if err := resp.Unmarshal(&plans); err != nil {
		return nil, err
	}

	return &ListPlansOutput{
		Data: plans,
	}, nil
}

// UpdatePlan ...
func (p plansService) UpdatePlan(ctx context.Context, id int64, input *UpdatePlanInput, opts ...Option) (*UpdatePlanOutput, error) {
	// TODO: review this
	if input.Configuration != nil && input.Configuration.ID == nil {
		return nil, errors.New("configuration id must not be nil")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathPlan, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan

	if err := resp.Unmarshal(&plan); err != nil {
		return nil, err
	}

	return &UpdatePlanOutput{
		Data: &plan,
	}, nil
}

type PlanInput struct {
	ID            *int64 `json:",omitempty"`
	Type          string
	Name          string
	Configuration *PlanConfiguration `json:",omitempty"`
}

type UpdatePlanInput = PlanInput

type CreatePlanInput = PlanInput

type PlanConfiguration struct {
	PlanID   *int64 `json:"PlanID,omitempty"`
	MetaData string
	ID       *int64 `json:"ID,omitempty"`
}

type ListPlansInput struct{}

type ListPlansOutput struct {
	Data []Plan
}

type Plan struct {
	AuthType                  string `json:"AuthType"`
	AutoApproveAccessRequests bool   `json:"AutoApproveAccessRequests"`
	Catalogues                any    `json:"Catalogues"`
	Description               string `json:"Description"`
	DisplayName               string `json:"DisplayName"`
	ID                        int64  `json:"ID"`
	JWTScope                  string `json:"JWTScope"`
	Name                      string `json:"Name"`
	Quota                     string `json:"Quota"`
	RateLimit                 string `json:"RateLimit"`
	ReferenceID               string `json:"ReferenceID"`
}

type PlanOutput struct {
	Data *Plan
}

type UpdatePlanOutput = PlanOutput

type GetPlanOutput = PlanOutput

type CreatePlanOutput = PlanOutput
