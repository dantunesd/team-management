package api

import (
	"encoding/json"
	"net/http"
	"team-management/members/usecase"

	"github.com/google/jsonapi"
)

type FilterMemberUseCase interface {
	Handle(input *usecase.FilterMemberInput) ([]usecase.FilterMemberOutput, error)
}

type FilterMember struct {
	useCase FilterMemberUseCase
}

func NewFilterMember(useCase FilterMemberUseCase) *FilterMember {
	return &FilterMember{
		useCase: useCase,
	}
}

type FilterMemberResponse struct {
	ID        string          `jsonapi:"primary,members"`
	Name      string          `jsonapi:"attr,name"`
	Type      string          `jsonapi:"attr,type"`
	TypeData  json.RawMessage `jsonapi:"attr,type_data"`
	Tags      []string        `jsonapi:"attr,tags"`
	CreatedAt string          `jsonapi:"attr,created_at"`
	UpdatedAt string          `jsonapi:"attr,updated_at"`
}

func (h *FilterMember) Handle(w http.ResponseWriter, r *http.Request) error {
	filters := map[string]string{}
	for k, v := range r.URL.Query() {
		filters[k] = v[0]
	}

	input := &usecase.FilterMemberInput{Filters: filters}
	outputs, err := h.useCase.Handle(input)
	if err != nil {
		return err
	}

	var filterMemberResponse []*FilterMemberResponse
	for _, output := range outputs {
		filterMemberResponse = append(filterMemberResponse, &FilterMemberResponse{
			ID:        output.ID,
			Name:      output.Name,
			Type:      output.Type,
			TypeData:  output.TypeData,
			Tags:      output.Tags,
			CreatedAt: output.CreatedAt,
			UpdatedAt: output.UpdatedAt,
		})
	}

	w.Header().Add("content-type", jsonapi.MediaType)
	w.WriteHeader(200)
	return jsonapi.MarshalPayload(w, filterMemberResponse)
}
