package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathDocs = "/portal-api/users"
	pathDoc  = "/portal-api/docs/%d"
)

type docsService struct {
	client *Client
}

func (p docsService) CreateDoc(input CreateDocInput) (*CreateDocOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathDocs, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateDocOutput{}, nil
}

func (p docsService) GetDoc(id uint64) (*GetDocOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathDoc, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetDocOutput{}, nil
}

func (p docsService) ListDocs(options *ListDocsOptions) (*ListDocsOutput, error) {
	req, err := p.client.newGetRequest(pathDocs, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListDocsOutput{}, nil
}

func (p docsService) UpdateDoc(id uint64, input UpdateDocInput) (*UpdateDocOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathDoc, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateDocOutput{}, nil
}

func (p docsService) ReorderDocs(id uint64, input UpdateDocInput) (*UpdateDocOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathDoc, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateDocOutput{}, nil
}

type UpdateDocInput struct {
	Catalogues []uint64
}

type CreateDocInput struct{}

type ListDocsOptions struct{}

type ListDocsOutput struct{}

type Doc struct{}

type DocOutput struct{}

type UpdateDocOutput = DocOutput

type GetDocOutput = DocOutput

type CreateDocOutput = DocOutput
