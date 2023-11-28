// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KVCacheI is an autogenerated mock type for the KVCacheI type
type KVCacheI struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, key
func (_m *KVCacheI) Get(ctx context.Context, key []byte) ([]byte, error) {
	ret := _m.Called(ctx, key)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) ([]byte, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) []byte); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNotFoundErr provides a mock function with given fields: err
func (_m *KVCacheI) IsNotFoundErr(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value, expired
func (_m *KVCacheI) Set(ctx context.Context, key []byte, value []byte, expired int64) error {
	ret := _m.Called(ctx, key, value, expired)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, int64) error); ok {
		r0 = rf(ctx, key, value, expired)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKVCacheI creates a new instance of KVCacheI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKVCacheI(t interface {
	mock.TestingT
	Cleanup(func())
}) *KVCacheI {
	mock := &KVCacheI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
