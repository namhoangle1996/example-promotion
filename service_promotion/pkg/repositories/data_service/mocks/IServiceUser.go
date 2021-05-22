// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	aggregates "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/aggregates"

	entities "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"

	mock "github.com/stretchr/testify/mock"
)

// IServiceUser is an autogenerated mock type for the IServiceUser type
type IServiceUser struct {
	mock.Mock
}

// FindById provides a mock function with given fields: id
func (_m *IServiceUser) FindById(id string) (*entities.User, error) {
	ret := _m.Called(id)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(string) *entities.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDetailById provides a mock function with given fields: id
func (_m *IServiceUser) FindDetailById(id string) (*aggregates.UserDetail, error) {
	ret := _m.Called(id)

	var r0 *aggregates.UserDetail
	if rf, ok := ret.Get(0).(func(string) *aggregates.UserDetail); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregates.UserDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMerchantByCode provides a mock function with given fields: code
func (_m *IServiceUser) FindMerchantByCode(code string) (*entities.Merchant, error) {
	ret := _m.Called(code)

	var r0 *entities.Merchant
	if rf, ok := ret.Get(0).(func(string) *entities.Merchant); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Merchant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMerchantById provides a mock function with given fields: id
func (_m *IServiceUser) FindMerchantById(id string) (*entities.Merchant, error) {
	ret := _m.Called(id)

	var r0 *entities.Merchant
	if rf, ok := ret.Get(0).(func(string) *entities.Merchant); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Merchant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
