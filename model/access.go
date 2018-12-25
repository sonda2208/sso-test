package model

import (
	"encoding/json"
	"io"
)

const (
	AccessTokenGrantType = "authorization_code"
	AccessTokenType      = "bearer"
)

type AccessData struct {
	ClientId     string `json:"client_id"`
	UserId       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	RedirectUri  string `json:"redirect_uri"`
	ExpiresAt    int64  `json:"expires_at"`
	Scope        string `json:"scope"`
}

type AccessResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func (ad *AccessData) ToJson() string {
	b, _ := json.Marshal(ad)
	return string(b)
}

func AccessDataFromJson(data io.Reader) *AccessData {
	var ad *AccessData
	json.NewDecoder(data).Decode(&ad)
	return ad
}

func (ar *AccessResponse) ToJson() string {
	b, _ := json.Marshal(ar)
	return string(b)
}

func AccessResponseFromJson(data io.Reader) *AccessResponse {
	var ar *AccessResponse
	json.NewDecoder(data).Decode(&ar)
	return ar
}
