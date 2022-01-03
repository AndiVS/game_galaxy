package model

import "github.com/google/uuid"

type Galaxy struct {
	ID           *uuid.UUID `json:"id"`
	UniverseName *uuid.UUID `json:"universe_name"`
}
