package repository

import (
	"context"
	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// InsertPlanet function for inserting item from a table
func (repos *Postgres) InsertPlanet(c context.Context, planet *model.Planet) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO planets(id, system_id, user_id, name) VALUES ($1, $2, $3, $4) RETURNING name",
		planet.ID, planet.SystemID, planet.UserID, planet.Name)

	err := row.Scan(&planet.Name)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectPlanet function for selecting item from a table
func (repos *Postgres) SelectPlanet(c context.Context, id uuid.UUID) (*model.Planet, error) {
	var planet model.Planet
	row := repos.Pool.QueryRow(c,
		"SELECT id, system_id, user_id, name FROM planets WHERE id = $1", id)

	err := row.Scan(&planet.ID, &planet.SystemID, &planet.UserID, &planet.Name)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &planet, err
	}

	log.Printf("sec")

	return &planet, err
}

// DeletePlanet function for deleting item from a table
func (repos *Postgres) DeletePlanet(c context.Context, id uuid.UUID) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM planets WHERE id = $1", id)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
