package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathOrgs = "/portal-api/users"
	pathOrg  = "/portal-api/orgs/%d"
)

type orgsService struct {
	client *Client
}

func (p orgsService) CreateOrg(input CreateOrgInput) (*CreateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathOrgs, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateOrgOutput{}, nil
}

func (p orgsService) GetOrg(id uint64) (*GetOrgOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathOrg, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetOrgOutput{}, nil
}

func (p orgsService) ListOrgs(options *ListOrgsOptions) (*ListOrgsOutput, error) {
	req, err := p.client.newGetRequest(pathOrgs, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListOrgsOutput{}, nil
}

func (p orgsService) UpdateOrg(id uint64, input UpdateOrgInput) (*UpdateOrgOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathOrg, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateOrgOutput{}, nil
}

type UpdateOrgInput struct {
	Catalogues []uint64
}

type CreateOrgInput struct{}

type ListOrgsOptions struct{}

type ListOrgsOutput struct{}

type Org struct{}

type OrgOutput struct{}

type UpdateOrgOutput = OrgOutput

type GetOrgOutput = OrgOutput

type CreateOrgOutput = OrgOutput
