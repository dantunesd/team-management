package api

import (
	"encoding/json"
	"net/http"
	"team-management/members/usecase"
	"team-management/members/utils"

	"github.com/google/jsonapi"
)

type CreateMemberUseCase interface {
	Handle(input *usecase.CreateMemberInput) (*usecase.CreateMemberOutput, error)
}

type CreateMember struct {
	usecase   CreateMemberUseCase
	validator utils.Validator
}

func NewCreateMember(usecase CreateMemberUseCase, validator utils.Validator) *CreateMember {
	return &CreateMember{
		usecase:   usecase,
		validator: validator,
	}
}

type CreateMemberRequest struct {
	Name     string          `json:"name" validate:"required"`
	Type     string          `json:"type" validate:"required,oneof='employee' 'contractor'"`
	TypeData json.RawMessage `json:"type_data" validate:"required"`
	Tags     []string        `json:"tags"`
}

type CreateMemberResponse struct {
	ID        string `jsonapi:"primary,members"`
	CreatedAt string `jsonapi:"attr,created_at"`
}

func (c *CreateMember) Handle(w http.ResponseWriter, r *http.Request) error {
	var request CreateMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return utils.NewBadRequest("Invalid json payload")
	}

	if err := c.validator.Validate(request); err != nil {
		return utils.NewBadRequest(err.Error())
	}

	input := &usecase.CreateMemberInput{
		Name:     request.Name,
		Type:     request.Type,
		TypeData: request.TypeData,
		Tags:     request.Tags,
	}

	output, err := c.usecase.Handle(input)
	if err != nil {
		return err
	}

	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(201)
	return jsonapi.MarshalPayload(w, &CreateMemberResponse{
		ID:        output.ID,
		CreatedAt: output.CreatedAt,
	})
}
