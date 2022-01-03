package model

import "github.com/google/uuid"

type System struct {
	ID       *uuid.UUID `json:"id"`
	GalaxyID *uuid.UUID `json:"galaxy_id"`
}
