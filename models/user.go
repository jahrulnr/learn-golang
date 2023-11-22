package models

import "database/sql"

type User struct {
	Id          int `json:"id,omitempty"`
	Name        string
	Username    string
	Password    string `json:"password,omitempty"`
	Accesstoken AccessToken
	CreatedAt   string         `json:"created_at,omitempty"`
	UpdatedAt   sql.NullString `json:"updated_at,omitempty"`
}
