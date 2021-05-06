package rest

import (
	"encoding/json"
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
	id := chi.URLParam(r, "id")

	tok, err := h.service.GetById(id)
	if err != nil {
		http.Error(w, err.Data, err.Status)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tok)
}
