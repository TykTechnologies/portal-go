// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/edsonmichaque/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// ProvidersService is an autogenerated mock type for the ProvidersService type
type ProvidersService struct {
	mock.Mock
}

// CreateProvider provides a mock function with given fields: ctx, input
func (_m *ProvidersService) CreateProvider(ctx context.Context, input portal.ProviderInput) (*portal.ProviderOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *portal.ProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, portal.ProviderInput) (*portal.ProviderOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, portal.ProviderInput) *portal.ProviderOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, portal.ProviderInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvider provides a mock function with given fields: ctx, id
func (_m *ProvidersService) GetProvider(ctx context.Context, id int64) (*portal.ProviderOutput, error) {
	ret := _m.Called(ctx, id)

	var r0 *portal.ProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*portal.ProviderOutput, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *portal.ProviderOutput); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviders provides a mock function with given fields: ctx, options
func (_m *ProvidersService) ListProviders(ctx context.Context, options *portal.ListProvidersOptions) (*portal.ListProvidersOutput, error) {
	ret := _m.Called(ctx, options)

	var r0 *portal.ListProvidersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProvidersOptions) (*portal.ListProvidersOutput, error)); ok {
		return rf(ctx, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProvidersOptions) *portal.ListProvidersOutput); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListProvidersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListProvidersOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncProvider provides a mock function with given fields: ctx, id
func (_m *ProvidersService) SyncProvider(ctx context.Context, id int64) (*portal.SyncProviderOutput, error) {
	ret := _m.Called(ctx, id)

	var r0 *portal.SyncProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*portal.SyncProviderOutput, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *portal.SyncProviderOutput); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.SyncProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProvider provides a mock function with given fields: ctx, id, input
func (_m *ProvidersService) UpdateProvider(ctx context.Context, id int64, input portal.ProviderInput) (*portal.ProviderOutput, error) {
	ret := _m.Called(ctx, id, input)

	var r0 *portal.ProviderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, portal.ProviderInput) (*portal.ProviderOutput, error)); ok {
		return rf(ctx, id, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, portal.ProviderInput) *portal.ProviderOutput); ok {
		r0 = rf(ctx, id, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProviderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, portal.ProviderInput) error); ok {
		r1 = rf(ctx, id, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProvidersService creates a new instance of ProvidersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvidersService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProvidersService {
	mock := &ProvidersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}