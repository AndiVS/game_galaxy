package model

import "github.com/google/uuid"

type Factory struct {
	ID              *uuid.UUID `json:"id"`
	PlanetID        *uuid.UUID `json:"planet_id"`
	Resource        string     `json:"resource"`
	Level           *int       `json:"level"`
	BasePerformance *int       `json:"base_performance"`
	BaseUpdateCost  *int       `json:"base_update_cost"`
	BaseUpdateTime  *int       `json:"base_update_time"`
}
