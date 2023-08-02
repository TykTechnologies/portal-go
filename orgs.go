// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	pathOrgs     = "/portal-api/organisations"
	pathOrg      = "/portal-api/organisations/%d"
	pathOrgTeams = "/portal-api/organisations/%d/teams"
	pathOrgTeam  = "/portal-api/organisations/%d/teams/%d"
)

//go:generate mockery --name OrgsService --filename orgs_service.go
type OrgsService interface {
	CreateOrg(ctx context.Context, input *CreateOrgInput, opts ...Option) (*CreateOrgOutput, error)
	GetOrg(ctx context.Context, id int64, opts ...Option) (*GetOrgOutput, error)
	ListOrgs(ctx context.Context, options *ListOrgsInput, opts ...Option) (*ListOrgsOutput, error)
	UpdateOrg(ctx context.Context, id int64, input *UpdateOrgInput, opts ...Option) (*UpdateOrgOutput, error)
	DeleteOrg(ctx context.Context, id int64, opts ...Option) (*DeleteOrgOutput, error)
	CreateTeam(ctx context.Context, orgId int64, input *TeamInput, opts ...Option) (*TeamOutput, error)
	GetTeam(ctx context.Context, orgId, teamId int64, opts ...Option) (*TeamOutput, error)
	ListTeams(ctx context.Context, orgId int64, options *ListTeamsInput, opts ...Option) (*ListTeamsOutput, error)
	UpdateTeam(ctx context.Context, orgId, teamId int64, input *TeamInput, opts ...Option) (*TeamOutput, error)
	DeleteTeam(ctx context.Context, orgId, teamId int64, opts ...Option) (*TeamOutput, error)
}

type orgsService struct {
	client *Client
}

func (p orgsService) CreateOrg(ctx context.Context, input *CreateOrgInput, opts ...Option) (*CreateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	if err := input.validate(); err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, pathOrgs, bytes.NewReader(payload), nil)
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
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathOrg, id), nil)
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
	resp, err := p.client.doGet(ctx, pathOrgs, nil)
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

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathOrg, id), bytes.NewReader(payload), nil)
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

func (p orgsService) DeleteOrg(ctx context.Context, id int64, opts ...Option) (*DeleteOrgOutput, error) {
	_, err := p.client.doDelete(ctx, fmt.Sprintf(pathOrg, id), nil, nil)
	if err != nil {
		return nil, err
	}

	return &GetOrgOutput{}, nil
}

func (p orgsService) CreateTeam(ctx context.Context, orgId int64, input *TeamInput, opts ...Option) (*TeamOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	if err := input.validate(); err != nil {
		return nil, err
	}

	resp, err := p.client.doPost(ctx, fmt.Sprintf(pathOrgTeams, orgId), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Team

	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &TeamOutput{
		Data: &org,
	}, nil
}

func (p orgsService) GetTeam(ctx context.Context, orgId, teamId int64, opts ...Option) (*TeamOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathOrgTeam, orgId, teamId), nil)
	if err != nil {
		return nil, err
	}

	var org Team
	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &TeamOutput{
		Data: &org,
	}, nil
}

func (p orgsService) ListTeams(ctx context.Context, orgId int64, options *ListTeamsInput, opts ...Option) (*ListTeamsOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathOrgTeams, orgId), nil)
	if err != nil {
		return nil, err
	}

	var orgs []Team

	if err := resp.Unmarshal(&orgs); err != nil {
		return nil, err
	}

	return &ListTeamsOutput{
		Data: orgs,
	}, nil
}

func (p orgsService) UpdateTeam(ctx context.Context, orgId, teamId int64, input *TeamInput, opts ...Option) (*TeamOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.doPut(ctx, fmt.Sprintf(pathOrgTeam, orgId, teamId), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	var org Team

	if err := resp.Unmarshal(&org); err != nil {
		return nil, err
	}

	return &TeamOutput{
		Data: &org,
	}, nil
}

func (p orgsService) DeleteTeam(ctx context.Context, orgId, teamId int64, opts ...Option) (*TeamOutput, error) {
	_, err := p.client.doDelete(ctx, fmt.Sprintf(pathOrgTeam, orgId, teamId), nil, nil)
	if err != nil {
		return nil, err
	}

	return &TeamOutput{}, nil
}

type OrgInput struct {
	ID   *int64 `json:"ID,omitempty"`
	Type string `json:"Type,omitempty"`
	Name string `json:"Name,omitempty"`
}

func (o OrgInput) validate() error {
	if o.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

type (
	UpdateOrgInput = OrgInput
	CreateOrgInput = OrgInput
	ListOrgsInput  struct{}
)

type ListOrgsOutput struct {
	Data []Org
}

type OrgOutput struct {
	Data *Org
}

type (
	UpdateOrgOutput = OrgOutput
	GetOrgOutput    = OrgOutput
	CreateOrgOutput = OrgOutput
	DeleteOrgOutput = OrgOutput
)

type Org struct {
	ID        int64       `json:"ID,omitempty"`
	Name      string      `json:"Name,omitempty"`
	Teams     interface{} `json:"Teams,omitempty"`
	Users     interface{} `json:"Users,omitempty"`
	UpdatedAt string      `json:"UpdatedAt,omitempty"`
	CreatedAt string      `json:"CreatedAt,omitempty"`
}

type OrgTeam struct {
	Default        bool   `json:"Default,omitempty"`
	ID             int64  `json:"ID,omitempty"`
	Name           string `json:"Name,omitempty"`
	Organisation   string `json:"Organisation,omitempty"`
	OrganisationID string `json:"OrganisationID,omitempty"`
	Users          []User `json:"Users,omitempty"`
}

type OrgCart struct {
	CatalogueOrders string `json:"CatalogueOrders,omitempty"`
	ID              int64  `json:"ID,omitempty"`
	ProviderID      any    `json:"ProviderID,omitempty"`
}

type OrgUser struct {
	APIToken          string   `json:"APIToken,omitempty"`
	APITokenCreatedAt string   `json:"APITokenCreatedAt,omitempty"`
	Active            bool     `json:"Active,omitempty"`
	Cart              OrgCart  `json:"Cart,omitempty"`
	ConfirmedAt       string   `json:"ConfirmedAt,omitempty"`
	Email             string   `json:"Email,omitempty"`
	EncryptedPassword string   `json:"EncryptedPassword,omitempty"`
	First             string   `json:"First,omitempty"`
	ID                int64    `json:"ID,omitempty"`
	Joined            string   `json:"Joined,omitempty"`
	Last              string   `json:"Last,omitempty"`
	Organisation      string   `json:"Organisation,omitempty"`
	Password          string   `json:"Password,omitempty"`
	Provider          string   `json:"Provider,omitempty"`
	ProviderID        int      `json:"ProviderID,omitempty"`
	ResetPassword     bool     `json:"ResetPassword,omitempty"`
	Role              string   `json:"Role,omitempty"`
	SSOKey            string   `json:"SSOKey,omitempty"`
	Teams             []string `json:"Teams,omitempty"`
	UID               string   `json:"UID,omitempty"`
	UserID            string   `json:"UserID,omitempty"`
}

type Team struct {
	ID        int64    `json:"ID,omitempty"`
	Name      string   `json:"Name,omitempty"`
	Default   bool     `json:"Default,omitempty"`
	Users     []string `json:"Users,omitempty"`
	CreatedAt string   `json:"CreatedAt,omitempty"`
	UpdatedAt string   `json:"UpdatedAt,omitempty"`
}

type TeamOutput struct {
	Data *Team
}

type ListTeamsOutput struct {
	Data []Team
}

type (
	UpdateTeam = TeamOutput
	GetTeam    = TeamOutput
	CreateTeam = TeamOutput
	DeleteTeam = TeamOutput
)

type TeamInput struct {
	ID    *int64  `json:"ID,omitempty"`
	Name  string  `json:"Name,omitempty"`
	Users []int64 `json:"Users,omitempty"`
}

func (t TeamInput) validate() error {
	return nil
}

type ListTeamsInput struct{}
