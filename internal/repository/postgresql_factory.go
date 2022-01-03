package repository

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UpdateFactory function for deleting item from a table
func (repos *Postgres) UpdateFactory(c context.Context, id uuid.UUID) error {
	_, err := repos.Pool.Exec(c,
		"UPDATE factories SET level = $2 WHERE id = $1", id, id)
	if err != nil {
		log.Errorf("Failed updating data in db: %s\n", err)
		return err
	}

	return nil
}
