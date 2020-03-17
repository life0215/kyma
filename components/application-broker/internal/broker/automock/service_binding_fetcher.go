// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

// ServiceBindingFetcher is an autogenerated mock type for the ServiceBindingFetcher type
type ServiceBindingFetcher struct {
	mock.Mock
}

// GetServiceBindingSecretName provides a mock function with given fields: ns, externalID
func (_m *ServiceBindingFetcher) GetServiceBindingSecretName(ns string, externalID string) (string, error) {
	ret := _m.Called(ns, externalID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(ns, externalID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(ns, externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}