package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	pathTeams = "/portal-api/teams"
	pathTeam  = "/portal-api/teams/%d"
)

type teamsService struct {
	client *Client
}

func (p teamsService) CreateTeam(input CreateTeamInput) (*CreateTeamOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPostRequest(pathTeams, bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &CreateTeamOutput{}, nil
}

func (p teamsService) GetTeam(id uint64) (*GetTeamOutput, error) {
	req, err := p.client.newGetRequest(fmt.Sprintf(pathTeam, id), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &GetTeamOutput{}, nil
}

func (p teamsService) ListTeams(options *ListTeamsOptions) (*ListTeamsOutput, error) {
	req, err := p.client.newGetRequest(pathTeams, nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &ListTeamsOutput{}, nil
}

func (p teamsService) UpdateTeam(id uint64, input UpdateTeamInput) (*UpdateTeamOutput, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := p.client.newPutRequest(fmt.Sprintf(pathTeam, id), bytes.NewReader(payload), nil)
	if err != nil {
		return nil, err
	}

	_, err = p.client.performRequest(req)
	if err != nil {
		return nil, err
	}

	return &UpdateTeamOutput{}, nil
}

type UpdateTeamInput struct {
	Catalogues []uint64
}

type CreateTeamInput struct{}

type ListTeamsOptions struct{}

type ListTeamsOutput struct{}

type Team struct{}

type TeamOutput struct{}

type UpdateTeamOutput = TeamOutput

type GetTeamOutput = TeamOutput

type CreateTeamOutput = TeamOutput
