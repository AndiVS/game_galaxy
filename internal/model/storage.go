package model

import "github.com/google/uuid"

type Storage struct {
	ID              *uuid.UUID `json:"id"`
	PlanetID        *uuid.UUID `json:"planet_id"`
	Resource        string     `json:"resource"`
	Level           *int       `json:"level"`
	BaseCapacity    *int       `json:"base_capacity"`
	CurrentCapacity *int       `json:"current_capacity"`
	BaseUpdateCost  *int       `json:"base_update_cost"`
}
