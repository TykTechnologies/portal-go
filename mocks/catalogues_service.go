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

// CreateCatalogue provides a mock function with given fields: ctx, input, opts
func (_m *CataloguesService) CreateCatalogue(ctx context.Context, input *portal.CatalogueInput, opts ...portal.Option) (*portal.CatalogueOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.CatalogueInput, ...portal.Option) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.CatalogueInput, ...portal.Option) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.CatalogueInput, ...portal.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatalogue provides a mock function with given fields: ctx, id, opts
func (_m *CataloguesService) GetCatalogue(ctx context.Context, id int64, opts ...portal.Option) (*portal.CatalogueOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCatalogues provides a mock function with given fields: ctx, options, opts
func (_m *CataloguesService) ListCatalogues(ctx context.Context, options *portal.ListCataloguesInput, opts ...portal.Option) (*portal.ListCataloguesOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListCataloguesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListCataloguesInput, ...portal.Option) (*portal.ListCataloguesOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListCataloguesInput, ...portal.Option) *portal.ListCataloguesOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListCataloguesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListCataloguesInput, ...portal.Option) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCatalogue provides a mock function with given fields: ctx, id, input, opts
func (_m *CataloguesService) UpdateCatalogue(ctx context.Context, id int64, input *portal.CatalogueInput, opts ...portal.Option) (*portal.CatalogueOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.CatalogueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.CatalogueInput, ...portal.Option) (*portal.CatalogueOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.CatalogueInput, ...portal.Option) *portal.CatalogueOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.CatalogueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *portal.CatalogueInput, ...portal.Option) error); ok {
		r1 = rf(ctx, id, input, opts...)
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
