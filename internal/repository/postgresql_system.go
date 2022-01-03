package repository

import (
	"context"
	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// InsertSystem function for inserting item from a table
func (repos *Postgres) InsertSystem(c context.Context, system *model.System) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO systems(id, galaxy_id) VALUES ($1, $2) RETURNING id", system.ID, system.GalaxyID)

	err := row.Scan(&system.ID, &system.GalaxyID)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectSystem function for selecting item from a table
func (repos *Postgres) SelectSystem(c context.Context, id uuid.UUID) (*model.System, error) {
	var system model.System
	row := repos.Pool.QueryRow(c,
		"SELECT id, galaxy_id FROM systems WHERE id = $1", id)

	err := row.Scan(&system.ID, &system.GalaxyID)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &system, err
	}

	log.Printf("sec")

	return &system, err
}

// DeleteSystem function for deleting item from a table
func (repos *Postgres) DeleteSystem(c context.Context, id uuid.UUID) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM systems WHERE id = $1", id)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
