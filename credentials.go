package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathCredentials = "/portal-api/users"
	pathCredential  = "/portal-api/credentials/%d"
)

type credentialsService struct {
	client *Client
}

func (p credentialsService) CreateCredential(input CreateCredentialInput) (*CreateCredentialOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathCredentials, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateCredentialOutput{}, nil
}

func (p credentialsService) GetCredential(id uint64) (*GetCredentialOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathCredential, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetCredentialOutput{}, nil
}

func (p credentialsService) ListCredentials(options *ListCredentialsOptions) (*ListCredentialsOutput, error) {
	req, err := p.client.newGetRequest(pathCredentials, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListCredentialsOutput{}, nil
}

func (p credentialsService) UpdateCredential(id uint64, input UpdateCredentialInput) (*UpdateCredentialOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathCredential, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateCredentialOutput{}, nil
}

type UpdateCredentialInput struct {
	Catalogues []uint64
}

type CreateCredentialInput struct{}

type ListCredentialsOptions struct{}

type ListCredentialsOutput struct{}

type Credential struct{}

type CredentialOutput struct{}

type UpdateCredentialOutput = CredentialOutput

type GetCredentialOutput = CredentialOutput

type CreateCredentialOutput = CredentialOutput
