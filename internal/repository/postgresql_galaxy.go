package repository

import (
	"context"
	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// InsertGalaxy function for inserting item from a table
func (repos *Postgres) InsertGalaxy(c context.Context, galaxy *model.Galaxy) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO galaxies(id, universe_name) VALUES ($1, $2) RETURNING id", galaxy.ID)

	err := row.Scan(&galaxy.ID)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectGalaxy function for selecting item from a table
func (repos *Postgres) SelectGalaxy(c context.Context, id uuid.UUID) (*model.Galaxy, error) {
	var galaxy model.Galaxy
	row := repos.Pool.QueryRow(c,
		"SELECT id, universe_name FROM galaxies WHERE id = $1", id)

	err := row.Scan(&galaxy.ID, &galaxy.UniverseName)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &galaxy, err
	}

	log.Printf("sec")

	return &galaxy, err
}

// DeleteGalaxy function for deleting item from a table
func (repos *Postgres) DeleteGalaxy(c context.Context, id uuid.UUID) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM galaxies WHERE id = $1", id)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
