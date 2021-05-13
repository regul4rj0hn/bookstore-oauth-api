package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
)

type TokenService interface {
	Create(token.AccessToken) *errors.Response
	GetById(string) (*token.AccessToken, *errors.Response)
	UpdateExpiration(token.AccessToken) *errors.Response
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
	router.Post("/", h.Create)
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.GetById)
		router.Put("/", h.UpdateExpiration)
	})
}

func (h *tokenHandler) Create(w http.ResponseWriter, r *http.Request) {
	token := &token.AccessToken{}
	if err := render.Bind(r, token); err != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}
	if err := h.service.Create(*token); err != nil {
		render.Render(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, token)
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

func (h *tokenHandler) UpdateExpiration(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	toUpdate, err := h.service.GetById(id)
	if err != nil {
		render.Render(w, r, err)
		return
	}
	newExp := &token.AccessToken{}
	bindErr := render.Bind(r, newExp)
	if bindErr != nil {
		render.Render(w, r, errors.ErrBadRequest)
		return
	}
	toUpdate.Expires = newExp.Expires
	if err := h.service.UpdateExpiration(*toUpdate); err != nil {
		render.Render(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, toUpdate)
}
