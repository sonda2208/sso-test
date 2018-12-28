package store

import "github.com/sonda2208/sso-test/model"

type TokenStore interface {
	Save(token *model.Token) error
	Delete(value string) error
	GetByValue(value string) (*model.Token, error)
	CleanUp()
}
