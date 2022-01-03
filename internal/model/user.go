package model

import "github.com/google/uuid"

type User struct {
	ID         *uuid.UUID `json:"id"`
	UniverseID *uuid.UUID `json:"universe_id"`
	UserName   string     `json:"user_name"`
	//login
	//password
}
