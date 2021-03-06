// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "gitlab.com/wallet-gpay/service-promotion/service_promotion/pkg/models/entities"
)

// ILogVoucher is an autogenerated mock type for the ILogVoucher type
type ILogVoucher struct {
	mock.Mock
}

// CountLogAll provides a mock function with given fields: user_id, voucher_id, action
func (_m *ILogVoucher) CountLogAll(user_id string, voucher_id string, action string) int64 {
	ret := _m.Called(user_id, voucher_id, action)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string) int64); ok {
		r0 = rf(user_id, voucher_id, action)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// CountLogPerDay provides a mock function with given fields: user_id, voucher_id, action
func (_m *ILogVoucher) CountLogPerDay(user_id string, voucher_id string, action string) int64 {
	ret := _m.Called(user_id, voucher_id, action)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string) int64); ok {
		r0 = rf(user_id, voucher_id, action)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Create provides a mock function with given fields: voucher
func (_m *ILogVoucher) Create(voucher entities.LogUserVoucher) error {
	ret := _m.Called(voucher)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.LogUserVoucher) error); ok {
		r0 = rf(voucher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindLogBy provides a mock function with given fields: trace_id, check_reverse
func (_m *ILogVoucher) FindLogBy(trace_id string, check_reverse bool) ([]*entities.LogUserVoucher, error) {
	ret := _m.Called(trace_id, check_reverse)

	var r0 []*entities.LogUserVoucher
	if rf, ok := ret.Get(0).(func(string, bool) []*entities.LogUserVoucher); ok {
		r0 = rf(trace_id, check_reverse)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.LogUserVoucher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(trace_id, check_reverse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogAll provides a mock function with given fields: user_id, voucher_id, action, check_total
func (_m *ILogVoucher) LogAll(user_id string, voucher_id string, action string, check_total int64) error {
	ret := _m.Called(user_id, voucher_id, action, check_total)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int64) error); ok {
		r0 = rf(user_id, voucher_id, action, check_total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogPerDay provides a mock function with given fields: user_id, voucher_id, action, check_total
func (_m *ILogVoucher) LogPerDay(user_id string, voucher_id string, action string, check_total int64) error {
	ret := _m.Called(user_id, voucher_id, action, check_total)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int64) error); ok {
		r0 = rf(user_id, voucher_id, action, check_total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogPerMonth provides a mock function with given fields: user_id, voucher_id, action, check_total
func (_m *ILogVoucher) LogPerMonth(user_id string, voucher_id string, action string, check_total int64) error {
	ret := _m.Called(user_id, voucher_id, action, check_total)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int64) error); ok {
		r0 = rf(user_id, voucher_id, action, check_total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogPerWeek provides a mock function with given fields: user_id, voucher_id, action, check_total
func (_m *ILogVoucher) LogPerWeek(user_id string, voucher_id string, action string, check_total int64) error {
	ret := _m.Called(user_id, voucher_id, action, check_total)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int64) error); ok {
		r0 = rf(user_id, voucher_id, action, check_total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogPerYear provides a mock function with given fields: user_id, voucher_id, action, check_total
func (_m *ILogVoucher) LogPerYear(user_id string, voucher_id string, action string, check_total int64) error {
	ret := _m.Called(user_id, voucher_id, action, check_total)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int64) error); ok {
		r0 = rf(user_id, voucher_id, action, check_total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: voucher
func (_m *ILogVoucher) Update(voucher entities.LogUserVoucher) error {
	ret := _m.Called(voucher)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.LogUserVoucher) error); ok {
		r0 = rf(voucher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
