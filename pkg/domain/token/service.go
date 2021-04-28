package token

import "github.com/regul4rj0hn/bookstore-oauth-api/pkg/utils/errors"

type TokenStore interface {
	GetById(string) (*AccessToken, *errors.Response)
}

type Service struct {
	Store TokenStore
}

func New(ts TokenStore) *Service {
	return &Service{
		Store: ts,
	}
}

func (s *Service) GetById(string) (*AccessToken, *errors.Response) {
	return nil, nil
}
