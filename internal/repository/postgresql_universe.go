package repository

import (
	"context"
	"github.com/AndiVS/game_galaxy/internal/model"
	log "github.com/sirupsen/logrus"
)

// InsertUniverse function for inserting item from a table
func (repos *Postgres) InsertUniverse(c context.Context, universe *model.Universe) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO universes(name) VALUES ($1) RETURNING name", universe.Name)

	err := row.Scan(&universe.Name)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectUniverse function for selecting item from a table
func (repos *Postgres) SelectUniverse(c context.Context, name string) (*model.Universe, error) {
	var universe model.Universe
	row := repos.Pool.QueryRow(c,
		"SELECT name FROM universes WHERE name = $1", name)

	err := row.Scan(&universe.Name)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &universe, err
	}

	log.Printf("sec")

	return &universe, err
}

// DeleteUniverse function for deleting item from a table
func (repos *Postgres) DeleteUniverse(c context.Context, name string) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM universes WHERE name = $1", name)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
