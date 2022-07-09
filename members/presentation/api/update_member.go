package api

import (
	"encoding/json"
	"net/http"
	"team-management/members/usecase"
	"team-management/members/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/jsonapi"
)

type UpdateMemberUseCase interface {
	Handle(id string, input *usecase.UpdateMemberInput) (*usecase.UpdateMemberOutput, error)
}

type UpdateMember struct {
	usecase   UpdateMemberUseCase
	validator Validator
}

func NewUpdateMember(usecase UpdateMemberUseCase, validator Validator) *UpdateMember {
	return &UpdateMember{
		usecase:   usecase,
		validator: validator,
	}
}

type UpdateMemberRequest struct {
	Name     string          `json:"name" validate:"required"`
	Type     string          `json:"type" validate:"required,oneof='employee' 'contractor'"`
	TypeData json.RawMessage `json:"type_data" validate:"required"`
	Tags     []string        `json:"tags"`
}

type UpdateMemberResponse struct {
	ID        string `jsonapi:"primary,members"`
	UpdatedAt string `jsonapi:"attr,updated_at"`
}

func (c *UpdateMember) Handle(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		return utils.NewBadRequest("The id is missing")
	}

	var request UpdateMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return utils.NewBadRequest("Invalid json payload")
	}

	if err := c.validator.Validate(request); err != nil {
		return utils.NewBadRequest(err.Error())
	}

	input := &usecase.UpdateMemberInput{
		Name:     request.Name,
		Type:     request.Type,
		TypeData: request.TypeData,
		Tags:     request.Tags,
	}

	output, err := c.usecase.Handle(id, input)
	if err != nil {
		return err
	}

	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(200)
	return jsonapi.MarshalPayload(w, &UpdateMemberResponse{
		ID:        id,
		UpdatedAt: output.UpdatedAt,
	})
}
