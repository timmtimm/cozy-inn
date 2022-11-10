// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	transactions "cozy-inn/businesses/transactions"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: email, transactionInput
func (_m *Repository) Create(email string, transactionInput transactions.Domain) (transactions.Domain, error) {
	ret := _m.Called(email, transactionInput)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(string, transactions.Domain) transactions.Domain); ok {
		r0 = rf(email, transactionInput)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, transactions.Domain) error); ok {
		r1 = rf(email, transactionInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: transactionID
func (_m *Repository) Delete(transactionID string) error {
	ret := _m.Called(transactionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(transactionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPaymentNotVerified provides a mock function with given fields:
func (_m *Repository) GetAllPaymentNotVerified() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReadyCheckIn provides a mock function with given fields:
func (_m *Repository) GetAllReadyCheckIn() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReadyCheckOut provides a mock function with given fields:
func (_m *Repository) GetAllReadyCheckOut() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTransaction provides a mock function with given fields:
func (_m *Repository) GetAllTransaction() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTransactionByEmail provides a mock function with given fields: email
func (_m *Repository) GetAllTransactionByEmail(email string) ([]transactions.Domain, error) {
	ret := _m.Called(email)

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func(string) []transactions.Domain); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByID provides a mock function with given fields: transactionID
func (_m *Repository) GetTransactionByID(transactionID string) (transactions.Domain, error) {
	ret := _m.Called(transactionID)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(string) transactions.Domain); ok {
		r0 = rf(transactionID)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByRoomAndEndDate provides a mock function with given fields: roomType, startDate, roomNumber
func (_m *Repository) GetTransactionByRoomAndEndDate(roomType string, startDate time.Time, roomNumber int) ([]transactions.Domain, error) {
	ret := _m.Called(roomType, startDate, roomNumber)

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func(string, time.Time, int) []transactions.Domain); ok {
		r0 = rf(roomType, startDate, roomNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Time, int) error); ok {
		r1 = rf(roomType, startDate, roomNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionOngoing provides a mock function with given fields:
func (_m *Repository) GetTransactionOngoing() ([]transactions.Domain, error) {
	ret := _m.Called()

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func() []transactions.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: transcationID, transactionDomain
func (_m *Repository) Update(transcationID string, transactionDomain transactions.Domain) error {
	ret := _m.Called(transcationID, transactionDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, transactions.Domain) error); ok {
		r0 = rf(transcationID, transactionDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}