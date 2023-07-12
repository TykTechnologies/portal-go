// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/edsonmichaque/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// PlansService is an autogenerated mock type for the PlansService type
type PlansService struct {
	mock.Mock
}

// CreatePlan provides a mock function with given fields: ctx, input, opts
func (_m *PlansService) CreatePlan(ctx context.Context, input *portal.PlanInput, opts ...portal.Option) (*portal.PlanOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.PlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.PlanInput, ...portal.Option) (*portal.PlanOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.PlanInput, ...portal.Option) *portal.PlanOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.PlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.PlanInput, ...portal.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlan provides a mock function with given fields: ctx, id, opts
func (_m *PlansService) GetPlan(ctx context.Context, id int64, opts ...portal.Option) (*portal.PlanOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.PlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.PlanOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.PlanOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.PlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlans provides a mock function with given fields: ctx, options, opts
func (_m *PlansService) ListPlans(ctx context.Context, options *portal.ListPlansInput, opts ...portal.Option) (*portal.ListPlansOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListPlansOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListPlansInput, ...portal.Option) (*portal.ListPlansOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListPlansInput, ...portal.Option) *portal.ListPlansOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListPlansOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListPlansInput, ...portal.Option) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlan provides a mock function with given fields: ctx, id, input, opts
func (_m *PlansService) UpdatePlan(ctx context.Context, id int64, input *portal.PlanInput, opts ...portal.Option) (*portal.PlanOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.PlanOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.PlanInput, ...portal.Option) (*portal.PlanOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.PlanInput, ...portal.Option) *portal.PlanOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.PlanOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *portal.PlanInput, ...portal.Option) error); ok {
		r1 = rf(ctx, id, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPlansService creates a new instance of PlansService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlansService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PlansService {
	mock := &PlansService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
