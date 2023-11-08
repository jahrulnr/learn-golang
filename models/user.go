package models

import "database/sql"

type User struct {
	Id          int
	Name        string
	Username    string
	Password    string
	Accesstoken AccessToken
	CreatedAt   string
	UpdatedAt   sql.NullString
}
