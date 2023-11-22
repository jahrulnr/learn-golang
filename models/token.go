package models

type AccessToken struct {
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
