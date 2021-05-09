package token

import (
	"strings"

	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
)

type TokenStore interface {
	GetById(string) (*AccessToken, *errors.Response)
}

type Service struct {
	Store TokenStore
}

func NewService(ts TokenStore) *Service {
	return &Service{
		Store: ts,
	}
}

func (s *Service) GetById(id string) (*AccessToken, *errors.Response) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.BadRequest("invalid access token id")
	}
	tok, err := s.Store.GetById(id)
	if err != nil {
		return nil, err
	}
	return tok, nil
}
