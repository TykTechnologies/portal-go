package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	pathOrgs = "/portal-api/organisations"
	pathOrg  = "/portal-api/organitations/%d"
)

//go:generate mockery --name OrgsService
type OrgsService interface {
	CreateOrg(ctx context.Context, input CreateOrgInput) (*CreateOrgOutput, error)
	GetOrg(ctx context.Context, id uint64) (*GetOrgOutput, error)
	ListOrgs(ctx context.Context, options *ListOrgsOptions) (*ListOrgsOutput, error)
	UpdateOrg(ctx context.Context, id uint64, input UpdateOrgInput) (*UpdateOrgOutput, error)
}

type orgsService struct {
	client *Client
}

func (p orgsService) CreateOrg(ctx context.Context, input CreateOrgInput) (*CreateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathOrgs, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Org

	if err := resp.Parse(&org); err != nil {
		return nil, err
	}

	return &CreateOrgOutput{
		Org: &org,
	}, nil
}

func (p orgsService) GetOrg(ctx context.Context, id uint64) (*GetOrgOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathOrg, id), nil)
	if err != nil {
		return nil, err
	}

	var org Org
	if err := resp.Parse(&org); err != nil {
		return nil, err
	}

	return &GetOrgOutput{
		Org: &org,
	}, nil
}

func (p orgsService) ListOrgs(ctx context.Context, options *ListOrgsOptions) (*ListOrgsOutput, error) {
	resp, err := p.client.doGet(pathOrgs, nil)
	if err != nil {
		return nil, err
	}

	var orgs []Org

	if err := resp.Parse(&orgs); err != nil {
		return nil, err
	}

	return &ListOrgsOutput{
		Orgs: orgs,
	}, nil
}

func (p orgsService) UpdateOrg(ctx context.Context, id uint64, input UpdateOrgInput) (*UpdateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathOrg, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Org

	if err := resp.Parse(&org); err != nil {
		return nil, err
	}

	return &UpdateOrgOutput{
		Org: &org,
	}, nil
}

type OrgInput struct {
	ID   *uint64 `json:",omitempty"`
	Type string
	Name string
}

type UpdateOrgInput = OrgInput

type CreateOrgInput = OrgInput

type ListOrgsOptions struct{}

type ListOrgsOutput struct {
	Orgs []Org
}

type Org struct {
	ID   uint64
	Name string
}

type OrgOutput struct {
	Org *Org
}

type UpdateOrgOutput = OrgOutput

type GetOrgOutput = OrgOutput

type CreateOrgOutput = OrgOutput
