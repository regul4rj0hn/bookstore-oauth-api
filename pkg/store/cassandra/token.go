package cassandra

import (
	database "github.com/regul4rj0hn/bookstore-oauth-api/pkg/clients/cassandra"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
)

type TokenStore struct{}

func NewTokenStore() *TokenStore {
	return &TokenStore{}
}

func (ts *TokenStore) GetById(string) (*token.AccessToken, *errors.Response) {
	session, err := database.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.InternalServerError("database connection not implemented")
}
