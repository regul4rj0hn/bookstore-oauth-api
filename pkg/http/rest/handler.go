package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
)

func NewHandler(service TokenService) http.Handler {
	router := chi.NewRouter()

	router.NotFound(notFoundHandler)

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.Logger)

	accessTokenHandler := NewTokenHandler(service)
	router.Route("/oauth/token", accessTokenHandler.TokenRoutes)

	return router
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	render.Render(w, r, errors.NotFound("page not found"))
}
