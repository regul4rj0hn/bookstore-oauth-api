package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/utils/errors"
)

type TokenService interface {
	GetById(string) (*token.AccessToken, *errors.Response)
}

type tokenHandler struct {
	service TokenService
}

func New(svc TokenService) *tokenHandler {
	return &tokenHandler{
		service: svc,
	}
}

func (h *tokenHandler) TokenRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{id}", h.GetById)
	return r
}

func (h *tokenHandler) GetById(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
