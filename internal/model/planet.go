package model

import "github.com/google/uuid"

type Planet struct {
	ID       *uuid.UUID `json:"id"`
	SystemID *uuid.UUID `json:"system_id"`
	UserID   *uuid.UUID `json:"user_id"`
}
