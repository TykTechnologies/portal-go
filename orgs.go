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

//go:generate mockery --name OrgsService --filename orgs_service.go
type OrgsService interface {
	CreateOrg(ctx context.Context, input *CreateOrgInput, opts ...Option) (*CreateOrgOutput, error)
	GetOrg(ctx context.Context, id int64, opts ...Option) (*GetOrgOutput, error)
	ListOrgs(ctx context.Context, options *ListOrgsInput, opts ...Option) (*ListOrgsOutput, error)
	UpdateOrg(ctx context.Context, id int64, input *UpdateOrgInput, opts ...Option) (*UpdateOrgOutput, error)
}

type orgsService struct {
	client *Client
}

func (p orgsService) CreateOrg(ctx context.Context, input *CreateOrgInput, opts ...Option) (*CreateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(pathOrgs, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Org

	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &CreateOrgOutput{
		Data: &org,
	}, nil
}

func (p orgsService) GetOrg(ctx context.Context, id int64, opts ...Option) (*GetOrgOutput, error) {
	resp, err := p.client.doGet(fmt.Sprintf(pathOrg, id), nil)
	if err != nil {
		return nil, err
	}

	var org Org
	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &GetOrgOutput{
		Data: &org,
	}, nil
}

func (p orgsService) ListOrgs(ctx context.Context, options *ListOrgsInput, opts ...Option) (*ListOrgsOutput, error) {
	resp, err := p.client.doGet(pathOrgs, nil)
	if err != nil {
		return nil, err
	}

	var orgs []Org

	if err := resp.Unmarshal(&orgs); err != nil {
		return nil, err
	}

	return &ListOrgsOutput{
		Data: orgs,
	}, nil
}

func (p orgsService) UpdateOrg(ctx context.Context, id int64, input *UpdateOrgInput, opts ...Option) (*UpdateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(fmt.Sprintf(pathOrg, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Org

	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &UpdateOrgOutput{
		Data: &org,
	}, nil
}

type OrgInput struct {
	ID   *int64 `json:"ID,omitempty"`
	Type string `json:"Type,omitempty"`
	Name string `json:"Name,omitempty"`
}

type UpdateOrgInput = OrgInput

type CreateOrgInput = OrgInput

type ListOrgsInput struct{}

type ListOrgsOutput struct {
	Data []Org
}

type OrgOutput struct {
	Data *Org
}

type UpdateOrgOutput = OrgOutput

type GetOrgOutput = OrgOutput

type CreateOrgOutput = OrgOutput

type Org struct {
	ID        int64     `json:"ID"`
	Name      string    `json:"Name"`
	Teams     []OrgTeam `json:"Teams,omitempty"`
	Users     []OrgUser `json:"Users,omitempty"`
	UpdatedAt string    `json:"UpdatedAt,omitempty"`
	CreatedAt string    `json:"CreatedAt,omitempty"`
}

type OrgTeam struct {
	Default        bool   `json:"Default"`
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	Organisation   string `json:"Organisation"`
	OrganisationID string `json:"OrganisationID"`
	Users          []User `json:"Users"`
}

type OrgCart struct {
	CatalogueOrders string `json:"CatalogueOrders"`
	ID              int64  `json:"ID"`
	ProviderID      any    `json:"ProviderID"`
}

type OrgUser struct {
	APIToken          string   `json:"APIToken"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt"`
	Active            bool     `json:"Active"`
	Cart              OrgCart  `json:"Cart"`
	ConfirmedAt       string   `json:"ConfirmedAt"`
	Email             string   `json:"Email"`
	EncryptedPassword string   `json:"EncryptedPassword"`
	First             string   `json:"First"`
	ID                int64    `json:"ID"`
	Joined            string   `json:"Joined"`
	Last              string   `json:"Last"`
	Organisation      string   `json:"Organisation"`
	Password          string   `json:"Password"`
	Provider          string   `json:"Provider"`
	ProviderID        int      `json:"ProviderID"`
	ResetPassword     bool     `json:"ResetPassword"`
	Role              string   `json:"Role"`
	SSOKey            string   `json:"SSOKey"`
	Teams             []string `json:"Teams"`
	UID               string   `json:"UID"`
	UserID            string   `json:"UserID"`
}
