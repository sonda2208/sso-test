package model

import (
	"encoding/json"
	"io"
)

type FacebookUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func FacebookUserFromJSON(data io.Reader) *FacebookUser {
	decoder := json.NewDecoder(data)
	var fbu FacebookUser
	err := decoder.Decode(&fbu)
	if err == nil {
		return &fbu
	}

	return nil
}
