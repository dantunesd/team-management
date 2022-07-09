package api

import (
	"errors"
	"net/http"
	"strings"
	"team-management/members/usecase"
	"team-management/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateMember_ShouldReturnAnErrorIfIDIsInvalid(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")

	usecaseMock := mocks.NewUpdateMemberUseCase(t)
	validatorMock := mocks.NewValidator(t)

	handler := NewUpdateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestUpdateMember_ShouldReturnAnErrorIfDecodeFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")
	req = setIdParameter(req)

	usecaseMock := mocks.NewUpdateMemberUseCase(t)
	validatorMock := mocks.NewValidator(t)

	handler := NewUpdateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestUpdateMember_ShouldReturnAnErrorIfValidateFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", payloadExample)
	req = setIdParameter(req)

	usecaseMock := mocks.NewUpdateMemberUseCase(t)

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(errors.New("failed"))

	handler := NewUpdateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestUpdateMember_ShouldReturnAnErrorIfUseCaseFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", payloadExample)
	req = setIdParameter(req)

	usecaseMock := mocks.NewUpdateMemberUseCase(t)
	usecaseMock.On("Handle", "1", mock.Anything).Return(&usecase.UpdateMemberOutput{}, errors.New("failed"))

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(nil)

	handler := NewUpdateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestUpdateMember_ShouldSetResponseWritterCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", payloadExample)
	req = setIdParameter(req)

	usecaseMock := mocks.NewUpdateMemberUseCase(t)
	usecaseMock.On("Handle", "1", mock.Anything).Return(&usecase.UpdateMemberOutput{UpdatedAt: "now"}, nil)

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(nil)

	handler := NewUpdateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"data":{"type":"members","id":"1","attributes":{"updated_at":"now"}}}`, strings.Trim(res.Body.String(), "\n"))
}
