package model

import "github.com/google/uuid"

type Galaxy struct {
	ID         *uuid.UUID `json:"id"`
	UniverseID *uuid.UUID `json:"universe_id"`
}
