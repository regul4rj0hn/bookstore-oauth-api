package main

import (
	"log"
	"net"
	"net/http"

	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/http/rest"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/store/cassandra"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	tokenService := token.NewService(cassandra.NewTokenStore())
	httpHandler := rest.NewHandler(tokenService)
	server := http.Server{
		Handler: httpHandler,
	}

	server.Serve(listener)
}
