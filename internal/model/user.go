package model

import "github.com/google/uuid"

type User struct {
	ID           *uuid.UUID `json:"id"`
	AccountLogin string     `json:"account_login"`
	UniverseName string     `json:"universe_name"`
	Username     string     `json:"username"`
}
