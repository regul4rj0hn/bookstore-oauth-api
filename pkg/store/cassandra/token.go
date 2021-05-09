package cassandra

import (
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
)

type TokenStore struct{}

func NewTokenStore() *TokenStore {
	return &TokenStore{}
}

func (ts *TokenStore) GetById(string) (*token.AccessToken, *errors.Response) {
	return nil, errors.InternalServerError("database connection not implemented")
}
