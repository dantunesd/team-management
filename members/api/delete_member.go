package api

import (
	"net/http"
	"team-management/members/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/jsonapi"
)

type DeleteUseCase interface {
	Handle(id string) error
}

type DeleteMember struct {
	useCase DeleteUseCase
}

func NewDeleteMember(useCase DeleteUseCase) *DeleteMember {
	return &DeleteMember{
		useCase: useCase,
	}
}

type DeleteMemberResponse struct {
	ID string `jsonapi:"primary,members"`
}

func (h *DeleteMember) Handle(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		return utils.NewBadRequest("The id is missing")
	}

	if err := h.useCase.Handle(id); err != nil {
		return err
	}

	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(200)
	return jsonapi.MarshalPayload(w, &DeleteMemberResponse{ID: id})
}
