package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/clients/cassandra"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/token"
)

const (
	queryGetAccessToken    = "SELECT access_token, client_id, sub, expires FROM tokens WHERE access_token=?;"
	queryInsertAccessToken = "INSERT INTO tokens(access_token, client_id, sub, expires) VALUES (?,?,?,?);"
	queryUpdateExpires     = "UPDATE tokens SET expires=? WHERE access_token=?;"
)

type TokenStore struct{}

func NewTokenStore() *TokenStore {
	return &TokenStore{}
}

func (ts *TokenStore) Create(at token.AccessToken) *errors.Response {
	if err := cassandra.GetSession().Query(queryInsertAccessToken, at.AccessToken, at.ClientId, at.Subject, at.Expires).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (ts *TokenStore) GetById(id string) (*token.AccessToken, *errors.Response) {
	var token token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&token.AccessToken, &token.ClientId, &token.Subject, &token.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFound("token not found")
		}
		return nil, errors.InternalServerError(err.Error())
	}

	return &token, nil
}

func (ts *TokenStore) UpdateExpiration(at token.AccessToken) *errors.Response {
	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
