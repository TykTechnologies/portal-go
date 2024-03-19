package portal

import (
	"context"
	"fmt"
	"net/http"
)

type PlansService service

type Plan struct {
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

type planInput struct {
	AllowFormSubmission bool   `json:"AllowFormSubmission"`
	PageTypeID          int64  `json:"PageTypeID"`
	Path                string `json:"Path"`
	Status              string `json:"Status"`
	Template            string `json:"Template"`
	Title               string `json:"Title"`
}

func (u *PlansService) ListPlans(ctx context.Context, opts *ListOptions) ([]*Plan, *Response, error) {
	urlPath := "/plans"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var plans []*Plan

	resp, err := u.client.Do(ctx, req, &plans)
	if err != nil {
		return nil, resp, err
	}

	return plans, resp, nil
}

func (u *PlansService) GetPlan(ctx context.Context, planID int64) (*Plan, *Response, error) {
	urlPath := fmt.Sprintf("/plans/%v", planID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	plan := new(Plan)

	resp, err := u.client.Do(ctx, req, plan)
	if err != nil {
		return nil, resp, err
	}

	return plan, resp, nil
}

func (u *PlansService) UpdatePlan(ctx context.Context, planID int64, input *Plan) (*Plan, *Response, error) {
	urlPath := fmt.Sprintf("/plans/%v", planID)

	planReq := planInput{}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, planReq)
	if err != nil {
		return nil, nil, err
	}

	plan := new(Plan)

	resp, err := u.client.Do(ctx, req, plan)
	if err != nil {
		return nil, resp, err
	}

	return plan, resp, nil
}
