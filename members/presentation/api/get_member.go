package api

import (
	"encoding/json"
	"net/http"
	"team-management/members/usecase"
	"team-management/members/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/jsonapi"
)

type GetMemberUseCase interface {
	Handle(id string) (*usecase.GetMemberOutput, error)
}

type GetMember struct {
	useCase GetMemberUseCase
}

func NewGetMember(useCase GetMemberUseCase) *GetMember {
	return &GetMember{
		useCase: useCase,
	}
}

type GetMemberResponse struct {
	ID        string          `jsonapi:"primary,members"`
	Name      string          `jsonapi:"attr,name"`
	Type      string          `jsonapi:"attr,type"`
	TypeData  json.RawMessage `jsonapi:"attr,type_data"`
	Tags      []string        `jsonapi:"attr,tags"`
	CreatedAt string          `jsonapi:"attr,created_at"`
	UpdatedAt string          `jsonapi:"attr,updated_at"`
}

func (h *GetMember) Handle(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		return utils.NewBadRequest("The id is missing")
	}

	output, err := h.useCase.Handle(id)
	if err != nil {
		return err
	}

	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(200)
	return jsonapi.MarshalPayload(w, &GetMemberResponse{
		ID:        output.ID,
		Name:      output.Name,
		Type:      output.Type,
		TypeData:  output.TypeData,
		Tags:      output.Tags,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	})
}
