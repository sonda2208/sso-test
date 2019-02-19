package model

import (
	"encoding/json"
	"io"
)

type GoogleEmail struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type GoogleName struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

type GoogleUser struct {
	ID          string        `json:"id"`
	Name        GoogleName    `json:"name"`
	Emails      []GoogleEmail `json:"emails"`
	DisplayName string        `json:"displayName"`
}

func GoogleUserFromJSON(data io.Reader) *GoogleUser {
	decoder := json.NewDecoder(data)
	var ggu GoogleUser
	err := decoder.Decode(&ggu)
	if err == nil {
		return &ggu
	}

	return nil
}
