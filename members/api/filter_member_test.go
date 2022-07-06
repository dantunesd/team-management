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

func TestFilterMember_ShouldReturnAnErrorIfUseCaseFails(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	usecaseMock := mocks.NewFilterMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return([]usecase.FilterMemberOutput{}, errors.New("failed"))

	handler := NewFilterMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.NotNil(t, err)
}

func TestFilterMember_ShouldSetResponseWritterWithEmptyItemsCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	usecaseMock := mocks.NewFilterMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return([]usecase.FilterMemberOutput{}, nil)

	handler := NewFilterMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"data":[]}`, strings.Trim(res.Body.String(), "\n"))
}
func TestFilterMember_ShouldSetResponseWritterWithManyItemsCorrectly(t *testing.T) {
	req, res := getHttpRequestAndResponse("/?tags=backend", "")

	usecaseMock := mocks.NewFilterMemberUseCase(t)
	usecaseMock.On("Handle", mock.Anything).Return([]usecase.FilterMemberOutput{
		{ID: "1"},
		{ID: "2"},
	}, nil)

	handler := NewFilterMember(usecaseMock)
	err := handler.Handle(res, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"data":[{"type":"members","id":"1","attributes":{"created_at":"","name":"","tags":null,"type":"","type_data":null,"updated_at":""}},{"type":"members","id":"2","attributes":{"created_at":"","name":"","tags":null,"type":"","type_data":null,"updated_at":""}}]}`, strings.Trim(res.Body.String(), "\n"))
}
