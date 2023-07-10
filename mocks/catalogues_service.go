// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/edsonmichaque/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// CataloguesService is an autogenerated mock type for the CataloguesService type
type CataloguesService struct {
	mock.Mock
}

// CreateCatalogue provides a mock function with given fields: ctx, input
func (_m *CataloguesService) CreateCatalogue(ctx context.Context, input portal.CatalogueInput) (*portal.CatalogueOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, portal.CatalogueInput) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, portal.CatalogueInput) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, portal.CatalogueInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatalogue provides a mock function with given fields: ctx, id
func (_m *CataloguesService) GetCatalogue(ctx context.Context, id int64) (*portal.CatalogueOutput, error) {
	ret := _m.Called(ctx, id)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCatalogues provides a mock function with given fields: ctx, options
func (_m *CataloguesService) ListCatalogues(ctx context.Context, options *portal.ListCataloguesOptions) (*portal.ListCataloguesOutput, error) {
	ret := _m.Called(ctx, options)

	var r0 *portal.ListCataloguesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListCataloguesOptions) (*portal.ListCataloguesOutput, error)); ok {
		return rf(ctx, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListCataloguesOptions) *portal.ListCataloguesOutput); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListCataloguesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListCataloguesOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCatalogue provides a mock function with given fields: ctx, id, input
func (_m *CataloguesService) UpdateCatalogue(ctx context.Context, id int64, input portal.CatalogueInput) (*portal.CatalogueOutput, error) {
	ret := _m.Called(ctx, id, input)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, portal.CatalogueInput) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, id, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, portal.CatalogueInput) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, id, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, portal.CatalogueInput) error); ok {
		r1 = rf(ctx, id, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCataloguesService creates a new instance of CataloguesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCataloguesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CataloguesService {
	mock := &CataloguesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}