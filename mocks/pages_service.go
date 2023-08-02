// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	edp "github.com/edsonmichaque/edp-go"
	mock "github.com/stretchr/testify/mock"
)

// PagesService is an autogenerated mock type for the PagesService type
type PagesService struct {
	mock.Mock
}

// CreatePage provides a mock function with given fields: ctx, input, opts
func (_m *PagesService) CreatePage(ctx context.Context, input *edp.PageInput, opts ...edp.Option) (*edp.PageOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *edp.PageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *edp.PageInput, ...edp.Option) (*edp.PageOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *edp.PageInput, ...edp.Option) *edp.PageOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*edp.PageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *edp.PageInput, ...edp.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPage provides a mock function with given fields: ctx, id, opts
func (_m *PagesService) GetPage(ctx context.Context, id int64, opts ...edp.Option) (*edp.PageOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *edp.PageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...edp.Option) (*edp.PageOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...edp.Option) *edp.PageOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*edp.PageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...edp.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPages provides a mock function with given fields: ctx, options, opts
func (_m *PagesService) ListPages(ctx context.Context, options *edp.ListPagesInput, opts ...edp.Option) (*edp.ListPagesOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *edp.ListPagesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *edp.ListPagesInput, ...edp.Option) (*edp.ListPagesOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *edp.ListPagesInput, ...edp.Option) *edp.ListPagesOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*edp.ListPagesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *edp.ListPagesInput, ...edp.Option) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePage provides a mock function with given fields: ctx, id, input, opts
func (_m *PagesService) UpdatePage(ctx context.Context, id int64, input *edp.PageInput, opts ...edp.Option) (*edp.PageOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *edp.PageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *edp.PageInput, ...edp.Option) (*edp.PageOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *edp.PageInput, ...edp.Option) *edp.PageOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*edp.PageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *edp.PageInput, ...edp.Option) error); ok {
		r1 = rf(ctx, id, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPagesService creates a new instance of PagesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPagesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PagesService {
	mock := &PagesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}