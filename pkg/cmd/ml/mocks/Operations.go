// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Operations is an autogenerated mock type for the Operations type
type Operations struct {
	mock.Mock
}

// IsValidPath provides a mock function with given fields:
func (_m *Operations) IsValidPath() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsValidPath")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReadJSONFile provides a mock function with given fields:
func (_m *Operations) ReadJSONFile() (map[string]interface{}, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadJSONFile")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOperations creates a new instance of Operations. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOperations(t interface {
	mock.TestingT
	Cleanup(func())
}) *Operations {
	mock := &Operations{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
