package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/http/rest"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/store/cassandra"
)

func main() {
	accessTokenService := token.New(cassandra.New())
	accessTokenHandler := rest.New(accessTokenService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/oauth/token", accessTokenHandler.TokenRouter())

	http.ListenAndServe(":8080", r)
}
