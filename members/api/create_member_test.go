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

func TestCreateMember_ShouldReturnAnErrorIfDecodeFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	usecaseMock := mocks.NewCreateMemberUseCase(t)
	validatorMock := mocks.NewValidator(t)

	handler := NewCreateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestCreateMember_ShouldReturnAnErrorIfValidatorFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", payloadExample)

	usecaseMock := mocks.NewCreateMemberUseCase(t)

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(errors.New("failed"))

	handler := NewCreateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestCreateMember_ShouldReturnAnErrorIfUseCaseFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", payloadExample)

	usecaseMock := mocks.NewCreateMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(&usecase.CreateMemberOutput{}, errors.New("failed"))

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(nil)

	handler := NewCreateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestCreateMember_ShouldSetResponseWritterCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", payloadExample)

	usecaseMock := mocks.NewCreateMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(&usecase.CreateMemberOutput{
		ID:        "id",
		CreatedAt: "now",
	}, nil)

	validatorMock := mocks.NewValidator(t)
	validatorMock.On("Validate", mock.Anything).Return(nil)

	handler := NewCreateMember(usecaseMock, validatorMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, `{"data":{"type":"members","id":"id","attributes":{"created_at":"now"}}}`, strings.Trim(res.Body.String(), "\n"))
}
