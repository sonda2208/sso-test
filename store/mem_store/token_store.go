package mem_store

import (
	"errors"
	"time"

	"github.com/sonda2208/sso-test/model"
)

const (
	MaxTokenExpiryTime = 24 * time.Hour
)

type OnMemoryTokenStore struct {
	tokens map[string]*model.Token
}

func NewOnMemoryTokenStore() *OnMemoryTokenStore {
	s := &OnMemoryTokenStore{}
	s.tokens = make(map[string]*model.Token)
	return s
}

func (s *OnMemoryTokenStore) Save(token *model.Token) error {
	s.tokens[token.Value] = token
	return nil
}

func (s *OnMemoryTokenStore) Delete(value string) error {
	delete(s.tokens, value)
	return nil
}

func (s *OnMemoryTokenStore) GetByValue(value string) (*model.Token, error) {
	token, ok := s.tokens[value]
	if !ok {
		return nil, errors.New("not found")
	}

	return token, nil
}

func (s *OnMemoryTokenStore) CleanUp() {
	for k, v := range s.tokens {
		elapsed := time.Since(v.CreatedAt)
		if elapsed >= MaxTokenExpiryTime {
			s.Delete(k)
		}
	}
}
