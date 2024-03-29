// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	usecase "team-management/members/usecase"

	mock "github.com/stretchr/testify/mock"
)

// CreateMemberUseCase is an autogenerated mock type for the CreateMemberUseCase type
type CreateMemberUseCase struct {
	mock.Mock
}

// Handle provides a mock function with given fields: input
func (_m *CreateMemberUseCase) Handle(input *usecase.CreateMemberInput) (*usecase.CreateMemberOutput, error) {
	ret := _m.Called(input)

	var r0 *usecase.CreateMemberOutput
	if rf, ok := ret.Get(0).(func(*usecase.CreateMemberInput) *usecase.CreateMemberOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.CreateMemberOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usecase.CreateMemberInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCreateMemberUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateMemberUseCase creates a new instance of CreateMemberUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateMemberUseCase(t mockConstructorTestingTNewCreateMemberUseCase) *CreateMemberUseCase {
	mock := &CreateMemberUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
