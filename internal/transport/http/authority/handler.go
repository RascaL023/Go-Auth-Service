package authority

import (
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
