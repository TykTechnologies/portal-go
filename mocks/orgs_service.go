// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/edsonmichaque/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// OrgsService is an autogenerated mock type for the OrgsService type
type OrgsService struct {
	mock.Mock
}

// CreateOrg provides a mock function with given fields: ctx, input, opts
func (_m *OrgsService) CreateOrg(ctx context.Context, input *portal.OrgInput, opts ...portal.Option) (*portal.OrgOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.OrgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.OrgInput, ...portal.Option) (*portal.OrgOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.OrgInput, ...portal.Option) *portal.OrgOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.OrgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.OrgInput, ...portal.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrg provides a mock function with given fields: ctx, id, opts
func (_m *OrgsService) GetOrg(ctx context.Context, id int64, opts ...portal.Option) (*portal.OrgOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.OrgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.OrgOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.OrgOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.OrgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrgs provides a mock function with given fields: ctx, options, opts
func (_m *OrgsService) ListOrgs(ctx context.Context, options *portal.ListOrgsInput, opts ...portal.Option) (*portal.ListOrgsOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListOrgsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListOrgsInput, ...portal.Option) (*portal.ListOrgsOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListOrgsInput, ...portal.Option) *portal.ListOrgsOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListOrgsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListOrgsInput, ...portal.Option) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrg provides a mock function with given fields: ctx, id, input, opts
func (_m *OrgsService) UpdateOrg(ctx context.Context, id int64, input *portal.OrgInput, opts ...portal.Option) (*portal.OrgOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.OrgOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.OrgInput, ...portal.Option) (*portal.OrgOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.OrgInput, ...portal.Option) *portal.OrgOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.OrgOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *portal.OrgInput, ...portal.Option) error); ok {
		r1 = rf(ctx, id, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrgsService creates a new instance of OrgsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrgsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrgsService {
	mock := &OrgsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
