package api

import (
	"fmt"
	"net/http"
	"team-management/members/utils"

	"github.com/google/jsonapi"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

type ErrorHandler struct {
	logger utils.Logger
}

func NewErrorHandler(logger utils.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (e *ErrorHandler) Handle(next CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			e.handleError(w, err)
		}
	}
}

func (e *ErrorHandler) handleError(w http.ResponseWriter, err error) {
	code, errorMessage := e.getCodeAndErrorMessage(err)
	if code == http.StatusInternalServerError {
		e.logger.Error(err)
	}

	e.setHeaders(w, code)
	e.write(w, code, errorMessage)
}

func (e *ErrorHandler) getCodeAndErrorMessage(err error) (int, string) {
	if e, is := err.(*utils.CustomError); is {
		return e.Code(), e.Error()
	}
	return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
}

func (e *ErrorHandler) setHeaders(w http.ResponseWriter, code int) {
	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(code)
}

func (e *ErrorHandler) write(w http.ResponseWriter, code int, errorMessage string) {
	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Detail: errorMessage,
		Status: fmt.Sprintf("%d", code),
	}})
}
