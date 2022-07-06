package api

import (
	"errors"
	"net/http"
	"strings"
	"team-management/members/utils"
	"testing"

	"team-management/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestErrorHandler_ShouldReturnAnInternalServerErrorIfNextReturnsAGenericError(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	next := func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("failed")
	}

	loggerMock := mocks.NewLogger(t)
	loggerMock.On("Error", mock.Anything).Return()

	handler := NewErrorHandler(loggerMock)
	handlerFunc := handler.Handle(next)
	handlerFunc.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, `{"errors":[{"detail":"Internal Server Error","status":"500"}]}`, strings.Trim(res.Body.String(), "\n"))
}

func TestErrorHandler_ShouldReturnACustomErrorIfNextReturnsACustomError(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	next := func(w http.ResponseWriter, r *http.Request) error {
		return utils.NewBadRequest("invalid payload")
	}

	loggerMock := mocks.NewLogger(t)

	handler := NewErrorHandler(loggerMock)
	handlerFunc := handler.Handle(next)
	handlerFunc.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, `{"errors":[{"detail":"invalid payload","status":"400"}]}`, strings.Trim(res.Body.String(), "\n"))
}

func TestErrorHandler_ShouldReturnsSuccessIfNextDoesntReturnAnyError(t *testing.T) {
	req, res := getHttpRequestAndResponse("/", "")

	next := func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}

	loggerMock := mocks.NewLogger(t)

	handler := NewErrorHandler(loggerMock)
	handlerFunc := handler.Handle(next)
	handlerFunc.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, ``, res.Body.String())
}
