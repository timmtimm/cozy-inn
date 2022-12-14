// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	users "cozy-inn/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: email
func (_m *Repository) Delete(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *Repository) GetUserByEmail(email string) (users.Domain, error) {
	ret := _m.Called(email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string) users.Domain); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserList provides a mock function with given fields:
func (_m *Repository) GetUserList() ([]users.Domain, error) {
	ret := _m.Called()

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func() []users.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
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

// Login provides a mock function with given fields: user
func (_m *Repository) Login(user users.Domain) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Domain) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: user
func (_m *Repository) Register(user users.Domain) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Domain) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: email, user
func (_m *Repository) Update(email string, user users.Domain) error {
	ret := _m.Called(email, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, users.Domain) error); ok {
		r0 = rf(email, user)
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
