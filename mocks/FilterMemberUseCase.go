// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	usecase "team-management/members/usecase"

	mock "github.com/stretchr/testify/mock"
)

// FilterMemberUseCase is an autogenerated mock type for the FilterMemberUseCase type
type FilterMemberUseCase struct {
	mock.Mock
}

// Handle provides a mock function with given fields: input
func (_m *FilterMemberUseCase) Handle(input *usecase.FilterMemberInput) ([]usecase.FilterMemberOutput, error) {
	ret := _m.Called(input)

	var r0 []usecase.FilterMemberOutput
	if rf, ok := ret.Get(0).(func(*usecase.FilterMemberInput) []usecase.FilterMemberOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]usecase.FilterMemberOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usecase.FilterMemberInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFilterMemberUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilterMemberUseCase creates a new instance of FilterMemberUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilterMemberUseCase(t mockConstructorTestingTNewFilterMemberUseCase) *FilterMemberUseCase {
	mock := &FilterMemberUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
