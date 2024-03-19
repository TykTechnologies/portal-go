package portal

import (
	"context"
	"fmt"
	"net/http"
)

type OrgsService service

type Org struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
}

type Team struct {
	ID        int64       `json:"ID,omitempty"`
	Name      string      `json:"Name,omitempty"`
	Teams     interface{} `json:"Teams,omitempty"`
	Users     interface{} `json:"Users,omitempty"`
	UpdatedAt string      `json:"UpdatedAt,omitempty"`
	CreatedAt string      `json:"CreatedAt,omitempty"`
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

type OrgCart struct {
	CatalogOrders string `json:"CatalogOrders,omitempty"`
	ID            int64  `json:"ID,omitempty"`
	ProviderID    any    `json:"ProviderID,omitempty"`
}

type orgInput struct {
	Name string `json:"Name"`
}

type teamInput struct {
	Name string `json:"Name"`
}

func (u *OrgsService) ListOrgs(ctx context.Context, opts *ListOptions) ([]*Org, *Response, error) {
	urlPath := "/organisations"

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*Org

	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

func (u *OrgsService) CreateOrg(ctx context.Context, input *Org) (*Org, *Response, error) {
	urlPath := "/organisations"

	userReq := &orgInput{
		Name: input.Name,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(Org)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *OrgsService) GetOrg(ctx context.Context, orgID int64) (*Org, *Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v", orgID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(Org)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *OrgsService) UpdateOrg(ctx context.Context, orgID int64, input *Org) (*Org, *Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v", orgID)

	userReq := &orgInput{
		Name: input.Name,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	user := new(Org)

	resp, err := u.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

func (u *OrgsService) DeleteOrg(ctx context.Context, orgID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v", orgID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *OrgsService) ListTeams(ctx context.Context, orgID int64, opts *ListOptions) ([]*Team, *Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v/teams", orgID)

	req, err := u.client.NewRequestWithOptions(ctx, http.MethodGet, urlPath, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var teams []*Team

	resp, err := u.client.Do(ctx, req, &teams)
	if err != nil {
		return nil, resp, err
	}

	return teams, resp, nil
}

func (u *OrgsService) CreateTeam(ctx context.Context, orgID int64, input *Team) (*Team, *Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v/teams", orgID)

	teamBody := &teamInput{
		Name: input.Name,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPost, urlPath, teamBody)
	if err != nil {
		return nil, nil, err
	}

	team := new(Team)

	resp, err := u.client.Do(ctx, req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, nil
}

func (u *OrgsService) UpdateTeam(ctx context.Context, orgID int64, teamID int64, input *Team) (*Team, *Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v/teams/%v", orgID, teamID)

	userReq := &orgInput{
		Name: input.Name,
	}

	req, err := u.client.NewRequest(ctx, http.MethodPut, urlPath, userReq)
	if err != nil {
		return nil, nil, err
	}

	team := new(Team)

	resp, err := u.client.Do(ctx, req, team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, nil
}

func (u *OrgsService) DeleteTeam(ctx context.Context, orgID int64, teamID int64) (*Response, error) {
	urlPath := fmt.Sprintf("/organisations/%v/teams/%v", orgID, teamID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
