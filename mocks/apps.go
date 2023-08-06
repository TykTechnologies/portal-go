// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/TykTechnologies/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// Apps is an autogenerated mock type for the Apps type
type Apps struct {
	mock.Mock
}

// CreateApp provides a mock function with given fields: ctx, input, opts
func (_m *Apps) CreateApp(ctx context.Context, input *portal.AppInput, opts ...portal.Option) (*portal.AppOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.AppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.AppInput, ...portal.Option) (*portal.AppOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.AppInput, ...portal.Option) *portal.AppOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.AppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.AppInput, ...portal.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApp provides a mock function with given fields: ctx, id, opts
func (_m *Apps) GetApp(ctx context.Context, id int64, opts ...portal.Option) (*portal.AppOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.AppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.AppOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.AppOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.AppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListARs provides a mock function with given fields: ctx, id, opts
func (_m *Apps) ListARs(ctx context.Context, id int64, opts ...portal.Option) (*portal.ListARsOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListARsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.ListARsOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.ListARsOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListARsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApps provides a mock function with given fields: ctx, opts
func (_m *Apps) ListApps(ctx context.Context, opts ...portal.Option) (*portal.ListAppsOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListAppsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...portal.Option) (*portal.ListAppsOutput, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...portal.Option) *portal.ListAppsOutput); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListAppsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...portal.Option) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProvisionApp provides a mock function with given fields: ctx, id, opts
func (_m *Apps) ProvisionApp(ctx context.Context, id int64, opts ...portal.Option) (*portal.StatusOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.StatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.StatusOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.StatusOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.StatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewApps creates a new instance of Apps. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApps(t interface {
	mock.TestingT
	Cleanup(func())
}) *Apps {
	mock := &Apps{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
