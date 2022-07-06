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

func TestGetMember_ShouldReturnAnErrorIfIDIsInvalid(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")

	usecaseMock := mocks.NewGetMemberUseCase(t)

	handler := NewGetMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestGetMember_ShouldReturnAnErrorIfUseCaseFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")
	req = setIdParameter(req)

	usecaseMock := mocks.NewGetMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(&usecase.GetMemberOutput{}, errors.New("failed"))

	handler := NewGetMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestGetMember_ShouldSetResponseWritterCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/{id}", "")
	req = setIdParameter(req)

	usecaseMock := mocks.NewGetMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return(&usecase.GetMemberOutput{
		ID:        "id",
		Name:      "name",
		Type:      "employee",
		TypeData:  []byte(`{"role":"software engineer"}`),
		Tags:      []string{"backend"},
		CreatedAt: "now",
		UpdatedAt: "now",
	}, nil)

	handler := NewGetMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"data":{"type":"members","id":"id","attributes":{"created_at":"now","name":"name","tags":["backend"],"type":"employee","type_data":{"role":"software engineer"},"updated_at":"now"}}}`, strings.Trim(res.Body.String(), "\n"))
}
