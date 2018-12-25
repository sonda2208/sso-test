package model

import (
	"encoding/json"
	"io"
)

type GithubUser struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GithubUserFromJson(data io.Reader) *GithubUser {
	decoder := json.NewDecoder(data)
	var ghu GithubUser
	err := decoder.Decode(&ghu)
	if err == nil {
		return &ghu
	}

	return nil
}
