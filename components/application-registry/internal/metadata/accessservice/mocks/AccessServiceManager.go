// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"
import mock "github.com/stretchr/testify/mock"

// AccessServiceManager is an autogenerated mock type for the AccessServiceManager type
type AccessServiceManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: application, serviceId, serviceName
func (_m *AccessServiceManager) Create(application string, serviceId string, serviceName string) apperrors.AppError {
	ret := _m.Called(application, serviceId, serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string) apperrors.AppError); ok {
		r0 = rf(application, serviceId, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: serviceName
func (_m *AccessServiceManager) Delete(serviceName string) apperrors.AppError {
	ret := _m.Called(serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Upsert provides a mock function with given fields: application, serviceId, serviceName
func (_m *AccessServiceManager) Upsert(application string, serviceId string, serviceName string) apperrors.AppError {
	ret := _m.Called(application, serviceId, serviceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string) apperrors.AppError); ok {
		r0 = rf(application, serviceId, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
