package authority

import (
	"net/http"
)

type AuthorityRouter struct {
	mux		*http.ServeMux
	handler *AuthorityHandler
}

const baseEndpoint = "/api/v1/authorities"

func NewAuthorityRouter(mux *http.ServeMux, handler *AuthorityHandler) *AuthorityRouter {
	return &AuthorityRouter{
		mux: mux,
		handler: handler,
	}
}

func (r *AuthorityRouter) Handle() {
	// r.mux.HandleFunc("GET " + baseEndpoint, r.handler.GetByID,)
	r.mux.HandleFunc("GET " + baseEndpoint + "/{id}", r.handler.GetByID,)
}
