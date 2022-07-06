package api

import (
	"errors"
	"net/http"
	"strings"
	"team-management/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteMember_ShouldReturnAnErrorIfIDIsInvalid(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")

	usecaseMock := mocks.NewDeleteUseCase(t)

	handler := NewDeleteMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestDeleteMember_ShouldReturnAnErrorIfUseCaseFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")
	req = setIdParameter(req)

	usecaseMock := mocks.NewDeleteUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(errors.New("failed"))

	handler := NewDeleteMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestDeleteMember_ShouldSetResponseWritterCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")
	req = setIdParameter(req)

	usecaseMock := mocks.NewDeleteUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(nil)

	handler := NewDeleteMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"data":{"type":"members","id":"1"}}`, strings.Trim(res.Body.String(), "\n"))
}
