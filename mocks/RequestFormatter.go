// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequestFormatter is an autogenerated mock type for the RequestFormatter type
type RequestFormatter struct {
	mock.Mock
}

// GetUniqKey provides a mock function with given fields: _a0
func (_m *RequestFormatter) GetUniqKey(_a0 ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(...interface{}) ([]byte, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...interface{}) []byte); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarshalWrapper provides a mock function with given fields: _a0
func (_m *RequestFormatter) MarshalWrapper(_a0 ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(...interface{}) ([]byte, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...interface{}) []byte); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnMarshalWrapper provides a mock function with given fields: _a0, _a1
func (_m *RequestFormatter) UnMarshalWrapper(_a0 []byte, _a1 ...interface{}) ([]interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, ...interface{}) ([]interface{}, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func([]byte, ...interface{}) []interface{}); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRequestFormatter creates a new instance of RequestFormatter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestFormatter(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestFormatter {
	mock := &RequestFormatter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
