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
	pathPlans = "/portal-api/plans"
	pathPlan  = "/portal-api/plans/%d"
)

//go:generate mockery --name Plans --filename plans.go
type Plans interface {
	CreatePlan(ctx context.Context, input *CreatePlanInput, opts ...Option) (*CreatePlanOutput, error)
	GetPlan(ctx context.Context, id int64, opts ...Option) (*GetPlanOutput, error)
	ListPlans(ctx context.Context, options *ListPlansInput, opts ...Option) (*ListPlansOutput, error)
	UpdatePlan(ctx context.Context, id int64, input *UpdatePlanInput, opts ...Option) (*UpdatePlanOutput, error)
}

type plans struct {
	client *Client
}

// CreatePlan ...
func (p plans) CreatePlan(ctx context.Context, input *CreatePlanInput, opts ...Option) (*CreatePlanOutput, error) {
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
func (p plans) GetPlan(ctx context.Context, id int64, opts ...Option) (*GetPlanOutput, error) {
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
func (p plans) ListPlans(ctx context.Context, options *ListPlansInput, opts ...Option) (*ListPlansOutput, error) {
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
func (p plans) UpdatePlan(ctx context.Context, id int64, input *UpdatePlanInput, opts ...Option) (*UpdatePlanOutput, error) {
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
	AutoApproveAccessRequests *bool   `json:"AutoApproveAccessRequests,omitempty"`
	Catalogues                []int64 `json:"Catalogues,omitempty"`
	DisplayName               string  `json:"DisplayName,omitempty"`
	Description               string  `json:"Description,omitempty"`
	JWTScope                  string  `json:"JWTScope,omitempty"`
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
	Catalogues                []any  `json:"Catalogues"`
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
