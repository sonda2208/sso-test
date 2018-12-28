package model

import "time"

type Token struct {
	Value     string
	CreatedAt time.Time
	Extra     string
}

func NewToken(extra string) *Token {
	return &Token{
		Value:     NewUUID(),
		CreatedAt: time.Now(),
		Extra:     extra,
	}
}
