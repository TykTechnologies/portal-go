package portal

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

//go:generate mockery --name PlansService
type PlansService interface {
	CreatePlan(ctx context.Context, input CreatePlanInput) (*CreatePlanOutput, error)
	GetPlan(ctx context.Context, id int64) (*GetPlanOutput, error)
	ListPlans(ctx context.Context, options *ListPlansOptions) (*ListPlansOutput, error)
	UpdatePlan(ctx context.Context, id int64, input UpdatePlanInput) (*UpdatePlanOutput, error)
}

type plansService struct {
	client *Client
}

func (p plansService) CreatePlan(ctx context.Context, input CreatePlanInput) (*CreatePlanOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathPlans, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan

	if err := resp.Parse(&plan); err != nil {
		return nil, err
	}

	return &CreatePlanOutput{
		Plan: &plan,
	}, nil
}

func (p plansService) GetPlan(ctx context.Context, id int64) (*GetPlanOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathPlan, id), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan
	if err := resp.Parse(&plan); err != nil {
		return nil, err
	}

	return &GetPlanOutput{
		Plan: &plan,
	}, nil
}

func (p plansService) ListPlans(ctx context.Context, options *ListPlansOptions) (*ListPlansOutput, error) {
	resp, err := p.client.doGet(pathPlans, nil)
	if err != nil {
		return nil, err
	}

	var plans []Plan

	if err := resp.Parse(&plans); err != nil {
		return nil, err
	}

	return &ListPlansOutput{
		Plans: plans,
	}, nil
}

func (p plansService) UpdatePlan(ctx context.Context, id int64, input UpdatePlanInput) (*UpdatePlanOutput, error) {
	// TODO: review this
	if input.Configuration != nil && input.Configuration.ID == nil {
		return nil, errors.New("configuration id must not be nil")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathPlan, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var plan Plan

	if err := resp.Parse(&plan); err != nil {
		return nil, err
	}

	return &UpdatePlanOutput{
		Plan: &plan,
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

type ListPlansOptions struct{}

type ListPlansOutput struct {
	Plans []Plan
}

type Plan struct {
	ID                        int64
	Name                      string
	Quota                     string
	DisplayName               string
	RateLimit                 string
	AutoApproveAccessRequests bool
	ReferenceID               string
}

type PlanOutput struct {
	Plan *Plan
}

type UpdatePlanOutput = PlanOutput

type GetPlanOutput = PlanOutput

type CreatePlanOutput = PlanOutput
