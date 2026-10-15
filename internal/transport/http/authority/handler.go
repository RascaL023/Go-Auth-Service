package authority

import (
	"encoding/json"
	"go-auth-service/internal/domain/authority"
	"go-auth-service/internal/transport/http/helper"
	"net/http"
)

type AuthorityHandler struct {
	service	*authority.Service
}

func NewAuthorityHandler(service *authority.Service) (*AuthorityHandler) {
	return &AuthorityHandler{
		service: service,
	}
}

// =========== READ ===========
func (h *AuthorityHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r.PathValue("id"))
	if err != nil {
		helper.Respond(w, 400, err.Error(), nil, err)
		return
	}

	entity, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		helper.Respond(w, 400, err.Error(), nil, err)
		return
	}

	result := ToAuthorityPublicResponse(entity)
	helper.Respond(w, 200, "Authority successfully retrivied", result, nil)
}

// =========== WRITE ===========
func (h *AuthorityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request AuthorityAdminRequest
	if err := decode(r, &request); err != nil {
		helper.Respond(w, 400, err.Error(), nil, err)
		return
	}

	entity := authority.Authority{
		Name: request.Name,
		Description: request.Description,
	}

	saved, err := h.service.Create(r.Context(), entity)
	if err != nil {
		helper.Respond(w, 400, err.Error(), nil, err)
		return
	}

	helper.Respond(w, 201, "Authority successfully created", saved, nil)
}


// =========== UTILS ===========
func decode(r *http.Request, target any) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return helper.NewError(
			helper.ErrBadJSONBody,
			"JSON Body is invalid",
		)
	}

	return nil
}
