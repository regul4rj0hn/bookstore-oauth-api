package cassandra

import (
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/utils/errors"
)

type TokenStore struct{}

func New() *TokenStore {
	return &TokenStore{}
}

func (ts *TokenStore) GetById(string) (*token.AccessToken, *errors.Response) {
	return nil, errors.InternalServerError("database connection not implemented")
}
