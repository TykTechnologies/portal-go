// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/TykTechnologies/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// Products is an autogenerated mock type for the Products type
type Products struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, input, opts
func (_m *Products) CreateProduct(ctx context.Context, input *portal.ProductInput, opts ...portal.Option) (*portal.ProductOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ProductOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ProductInput, ...portal.Option) (*portal.ProductOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ProductInput, ...portal.Option) *portal.ProductOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ProductInput, ...portal.Option) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: ctx, id, opts
func (_m *Products) GetProduct(ctx context.Context, id int64, opts ...portal.Option) (*portal.ProductOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ProductOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) (*portal.ProductOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...portal.Option) *portal.ProductOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...portal.Option) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: ctx, options, opts
func (_m *Products) ListProducts(ctx context.Context, options *portal.ListProductsInput, opts ...portal.Option) (*portal.ListProductsOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ListProductsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProductsInput, ...portal.Option) (*portal.ListProductsOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProductsInput, ...portal.Option) *portal.ListProductsOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListProductsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListProductsInput, ...portal.Option) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, id, input, opts
func (_m *Products) UpdateProduct(ctx context.Context, id int64, input *portal.ProductInput, opts ...portal.Option) (*portal.ProductOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *portal.ProductOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.ProductInput, ...portal.Option) (*portal.ProductOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.ProductInput, ...portal.Option) *portal.ProductOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *portal.ProductInput, ...portal.Option) error); ok {
		r1 = rf(ctx, id, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProducts creates a new instance of Products. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducts(t interface {
	mock.TestingT
	Cleanup(func())
}) *Products {
	mock := &Products{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
