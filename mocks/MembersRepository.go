// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	domain "team-management/members/domain"

	mock "github.com/stretchr/testify/mock"
)

// MembersRepository is an autogenerated mock type for the MembersRepository type
type MembersRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: member
func (_m *MembersRepository) Create(member *domain.Member) (string, error) {
	ret := _m.Called(member)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.Member) string); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Member) error); ok {
		r1 = rf(member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *MembersRepository) Delete(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Filter provides a mock function with given fields: filters
func (_m *MembersRepository) Filter(filters map[string]string) ([]*domain.Member, error) {
	ret := _m.Called(filters)

	var r0 []*domain.Member
	if rf, ok := ret.Get(0).(func(map[string]string) []*domain.Member); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *MembersRepository) Get(id string) (*domain.Member, error) {
	ret := _m.Called(id)

	var r0 *domain.Member
	if rf, ok := ret.Get(0).(func(string) *domain.Member); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Member)
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

// Update provides a mock function with given fields: member
func (_m *MembersRepository) Update(member *domain.Member) (bool, error) {
	ret := _m.Called(member)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.Member) bool); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Member) error); ok {
		r1 = rf(member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMembersRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMembersRepository creates a new instance of MembersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMembersRepository(t mockConstructorTestingTNewMembersRepository) *MembersRepository {
	mock := &MembersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
