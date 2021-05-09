package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
)

type TokenService interface {
	GetById(string) (*token.AccessToken, *errors.Response)
}

type tokenHandler struct {
	service TokenService
}

func NewTokenHandler(svc TokenService) *tokenHandler {
	return &tokenHandler{
		service: svc,
	}
}

func (h *tokenHandler) TokenRoutes(router chi.Router) {
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.GetById)
	})
}

func (h *tokenHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	tok, err := h.service.GetById(id)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, tok)
}
