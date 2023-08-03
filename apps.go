// Copyright 2023 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"context"
	"fmt"
)

const (
	pathAppARs = "/portal-api/apps/%v/access-requests"
	pathAppAR  = "/portal-api/apps/%v/access-requests/%d"
)

//go:generate mockery --name Apps --filename apps.go
type Apps interface {
	ListARs(ctx context.Context, id int64, opts ...Option) (*ListARsOutput, error)
}

type apps struct {
	client *Client
}

func (p apps) ListARs(ctx context.Context, id int64, opts ...Option) (*ListARsOutput, error) {
	resp, err := p.client.doGet(ctx, fmt.Sprintf(pathAppARs, id), nil, opts...)
	if err != nil {
		return nil, err
	}

	var ars []ARSummary

	if err := resp.Unmarshal(&ars); err != nil {
		return nil, err
	}

	return &ListARsOutput{
		Data: ars,
	}, nil
}
