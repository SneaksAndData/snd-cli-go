// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	algorithm "github.com/SneaksAndData/esd-services-api-client-go/algorithm"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CancelRun provides a mock function with given fields: algorithmName, requestId, initiator, reason
func (_m *Service) CancelRun(algorithmName string, requestId string, initiator string, reason string) (string, error) {
	ret := _m.Called(algorithmName, requestId, initiator, reason)

	if len(ret) == 0 {
		panic("no return value specified for CancelRun")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) (string, error)); ok {
		return rf(algorithmName, requestId, initiator, reason)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string) string); ok {
		r0 = rf(algorithmName, requestId, initiator, reason)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(algorithmName, requestId, initiator, reason)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRun provides a mock function with given fields: algorithmName, input, tag
func (_m *Service) CreateRun(algorithmName string, input algorithm.Payload, tag string) (string, error) {
	ret := _m.Called(algorithmName, input, tag)

	if len(ret) == 0 {
		panic("no return value specified for CreateRun")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, algorithm.Payload, string) (string, error)); ok {
		return rf(algorithmName, input, tag)
	}
	if rf, ok := ret.Get(0).(func(string, algorithm.Payload, string) string); ok {
		r0 = rf(algorithmName, input, tag)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, algorithm.Payload, string) error); ok {
		r1 = rf(algorithmName, input, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrievePayloadUri provides a mock function with given fields: runID, algorithmName
func (_m *Service) RetrievePayloadUri(runID string, algorithmName string) (*algorithm.PayloadResponse, error) {
	ret := _m.Called(runID, algorithmName)

	if len(ret) == 0 {
		panic("no return value specified for RetrievePayloadUri")
	}

	var r0 *algorithm.PayloadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*algorithm.PayloadResponse, error)); ok {
		return rf(runID, algorithmName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *algorithm.PayloadResponse); ok {
		r0 = rf(runID, algorithmName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*algorithm.PayloadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(runID, algorithmName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveRun provides a mock function with given fields: runID, algorithmName
func (_m *Service) RetrieveRun(runID string, algorithmName string) (string, error) {
	ret := _m.Called(runID, algorithmName)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveRun")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(runID, algorithmName)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(runID, algorithmName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(runID, algorithmName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
