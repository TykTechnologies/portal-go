// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	portal "github.com/edsonmichaque/portal-go"
	mock "github.com/stretchr/testify/mock"
)

// ProductsService is an autogenerated mock type for the ProductsService type
type ProductsService struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, input, opts
func (_m *ProductsService) CreateProduct(ctx context.Context, input *portal.ProductInput, opts ...func(*portal.Options)) (*portal.ProductOutput, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ProductInput, ...func(*portal.Options)) (*portal.ProductOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ProductInput, ...func(*portal.Options)) *portal.ProductOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ProductInput, ...func(*portal.Options)) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: ctx, id, opts
func (_m *ProductsService) GetProduct(ctx context.Context, id int64, opts ...func(*portal.Options)) (*portal.ProductOutput, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...func(*portal.Options)) (*portal.ProductOutput, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...func(*portal.Options)) *portal.ProductOutput); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...func(*portal.Options)) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: ctx, options, opts
func (_m *ProductsService) ListProducts(ctx context.Context, options *portal.ListProductsOptions, opts ...func(*portal.Options)) (*portal.ListProductsOutput, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProductsOptions, ...func(*portal.Options)) (*portal.ListProductsOutput, error)); ok {
		return rf(ctx, options, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *portal.ListProductsOptions, ...func(*portal.Options)) *portal.ListProductsOutput); ok {
		r0 = rf(ctx, options, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ListProductsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *portal.ListProductsOptions, ...func(*portal.Options)) error); ok {
		r1 = rf(ctx, options, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, id, input, opts
func (_m *ProductsService) UpdateProduct(ctx context.Context, id int64, input *portal.ProductInput, opts ...func(*portal.Options)) (*portal.ProductOutput, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.ProductInput, ...func(*portal.Options)) (*portal.ProductOutput, error)); ok {
		return rf(ctx, id, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *portal.ProductInput, ...func(*portal.Options)) *portal.ProductOutput); ok {
		r0 = rf(ctx, id, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*portal.ProductOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *portal.ProductInput, ...func(*portal.Options)) error); ok {
		r1 = rf(ctx, id, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductsService creates a new instance of ProductsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductsService {
	mock := &ProductsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
