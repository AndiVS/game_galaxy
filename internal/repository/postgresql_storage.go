package repository

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UpdateStorage function for deleting item from a table
func (repos *Postgres) UpdateStorage(c context.Context, id uuid.UUID) error {
	_, err := repos.Pool.Exec(c,
		"UPDATE storages SET level = $2 WHERE id = $1", id, id)
	if err != nil {
		log.Errorf("Failed updating data in db: %s\n", err)
		return err
	}

	return nil
}
